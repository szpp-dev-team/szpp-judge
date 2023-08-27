package runner

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	docker "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

type Runner struct {
	docker              *docker.Client
	ContainerID         string
	ImageName           string
	ContainerWorkingDir string
	HostWorkingDir      string
}

const (
	RUNNER_STOP_TIMEOUT int = 60 * 10 // 10 minutes
)

func New(
	ctx context.Context,
	dc *docker.Client,
	imageName string,
	hostWorkingDirAbsPath string,
) (*Runner, error) {
	containerWorkingDir := "/work"

	if err := os.MkdirAll(hostWorkingDirAbsPath, 0750); err != nil {
		return nil, fmt.Errorf("failed to create temporary dir '%s' to bind to the container: %w", hostWorkingDirAbsPath, err)
	}

	stopTimeout := RUNNER_STOP_TIMEOUT // const はアドレス取得できないので一旦変数へ
	runInit := true
	containerCfg := &container.Config{
		Tty:             true,
		Env:             []string{"SZPP_JUDGE=1"},
		Cmd:             []string{"/bin/sh"},
		Image:           imageName,
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
				Source:      hostWorkingDirAbsPath,
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

	return &Runner{
		docker:              dc,
		ImageName:           imageName,
		ContainerID:         resp.ID,
		ContainerWorkingDir: containerWorkingDir,
		HostWorkingDir:      hostWorkingDirAbsPath,
	}, nil
}

func (r *Runner) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err := r.docker.ContainerRemove(ctx, r.ContainerID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		slog.Error("Failed to remove container:", "ID", r.ContainerID, "err", err)
	}

	slog.Info("Successfully removed the container:", "ID", r.ContainerID)

	// if err = os.RemoveAll(r.HostWorkingDir); err != nil {
	// 	slog.Error("Failed to remove temporary dir on the host:", "path", r.HostWorkingDir)
	// }
	cancel()
}

type ExecOption struct {
	AsRootUser bool
	Stdin      *bytes.Buffer
	Cmd        []string
}

type ExecResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

func (r *Runner) Exec(ctx context.Context, opt ExecOption) (ExecResult, error) {
	user := ""
	if !opt.AsRootUser {
		user = "1234:1234"
	}

	execID, err := r.docker.ContainerExecCreate(ctx, r.ContainerID, types.ExecConfig{
		User:         user,
		AttachStdin:  opt.Stdin != nil,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          opt.Cmd,
	})
	if err != nil {
		return ExecResult{}, err
	}

	resp, err := r.docker.ContainerExecAttach(ctx, execID.ID, types.ExecStartCheck{})
	if err != nil {
		return ExecResult{}, err
	}
	defer resp.Close()

	inputDone := make(chan struct{})
	go func() {
		if opt.Stdin != nil {
			_, err = io.Copy(resp.Conn, opt.Stdin)
			if err != nil {
				slog.Error("Cannot send stdin:", "err", err)
			}
		}
		if err := resp.CloseWrite(); err != nil {
			slog.Error("Cannot close stdin:", "err", err)
		}
		close(inputDone)
	}()

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
	case <-inputDone:
		if err := <-outputDone; err != nil {
			return ExecResult{}, err
		}
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

	res, err := r.docker.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{
		Stdout:   string(stdout),
		Stderr:   string(stderr),
		ExitCode: res.ExitCode,
	}, nil
}

func (r *Runner) PrintLogs(ctx context.Context) error {
	out, err := r.docker.ContainerLogs(ctx, r.ContainerID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
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
