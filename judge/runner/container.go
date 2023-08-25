package runner

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	docker "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/google/uuid"
	"github.com/szpp-dev-team/szpp-judge/judge/runner/lang"
)

type DockerContaineredRunner struct {
	cli                 *docker.Client
	ID                  string
	Lang                lang.Lang
	ContainerWorkingDir string
	HostWorkingDir      string
}

type ExecResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

const (
	RUNNER_STOP_TIMEOUT int = 60 * 10 // 10 minutes
)

func ResolveDockerImageName(l lang.Lang) string {
	if l == lang.GCC {
		return "szpp-judge-images-gcc"
	}

	slog.Error("unknown language", "l", l)
	panic(l)
}

func NewDockerContaineredRunner(
	ctx context.Context,
	dc *docker.Client,
	l lang.Lang,
	hostWorkingDirAbsRoot string,
) (*DockerContaineredRunner, error) {
	iname := ResolveDockerImageName(l)

	u, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	hostWorkingDir := path.Join(hostWorkingDirAbsRoot, now.Format("0102_150405_")+u.String())

	containerWorkingDir := "/work"

	if err = os.MkdirAll(hostWorkingDir, 0750); err != nil {
		return nil, fmt.Errorf("failed to create temporary dir '%s' to bind to the container: %w", hostWorkingDir, err)
	}

	stopTimeout := RUNNER_STOP_TIMEOUT // const はアドレス取得できないので一旦変数へ
	runInit := true
	containerCfg := &container.Config{
		Tty:             true,
		Env:             []string{"SZPP_JUDGE=1"},
		Cmd:             []string{"/bin/sh"},
		Image:           iname,
		WorkingDir:      containerWorkingDir,
		NetworkDisabled: true,
		StopTimeout:     &stopTimeout,
	}
	hostCfg := &container.HostConfig{
		LogConfig:  container.LogConfig{},
		AutoRemove: true,
		Resources:  container.Resources{}, // TODO: ulimit などの設定
		Init:       &runInit,
		Mounts: []mount.Mount{
			{
				Type:        mount.TypeBind,
				Source:      hostWorkingDir,
				Target:      containerWorkingDir,
				Consistency: mount.ConsistencyFull,
			},
		},
	}

	resp, err := dc.ContainerCreate(ctx, containerCfg, hostCfg, nil, nil, "")
	if err != nil {
		return nil, err
	}

	if err = dc.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	return &DockerContaineredRunner{
		cli:                 dc,
		ID:                  resp.ID,
		Lang:                l,
		ContainerWorkingDir: containerWorkingDir,
		HostWorkingDir:      hostWorkingDir,
	}, nil
}

func (r *DockerContaineredRunner) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err := r.cli.ContainerRemove(ctx, r.ID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		slog.Error("Failed to remove container:", "ID", r.ID, "err", err)
	}

	slog.Info("Successfully removed the container:", "ID", r.ID)

	if err = os.RemoveAll(r.HostWorkingDir); err != nil {
		slog.Error("Failed to remove temporary dir on the host:", "path", r.HostWorkingDir)
	}
	cancel()
}

func (r *DockerContaineredRunner) Exec(ctx context.Context, cmd []string) (ExecResult, error) {
	execID, err := r.cli.ContainerExecCreate(ctx, r.ID, types.ExecConfig{
		User:         "1234:1234",
		Tty:          false,
		AttachStdin:  false,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          cmd,
	})
	if err != nil {
		return ExecResult{}, err
	}

	resp, err := r.cli.ContainerExecAttach(ctx, execID.ID, types.ExecStartCheck{})
	if err != nil {
		return ExecResult{}, err
	}
	defer resp.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		_, err := stdcopy.StdCopy(&stdoutBuf, &stderrBuf, resp.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return ExecResult{}, err
		}
		break

	case <-ctx.Done():
		return ExecResult{}, ctx.Err()
	}

	stdout, err := io.ReadAll(&stdoutBuf)
	if err != nil {
		return ExecResult{}, err
	}
	stderr, err := io.ReadAll(&stderrBuf)
	if err != nil {
		return ExecResult{}, err
	}

	res, err := r.cli.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{
		Stdout:   string(stdout),
		Stderr:   string(stderr),
		ExitCode: res.ExitCode,
	}, nil
}

func (r *DockerContaineredRunner) PrintLogs(ctx context.Context) error {
	out, err := r.cli.ContainerLogs(ctx, r.ID, types.ContainerLogsOptions{
		ShowStdout: true,
	})
	if err != nil {
		return err
	}

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	if err != nil {
		return err
	}

	return nil
}
