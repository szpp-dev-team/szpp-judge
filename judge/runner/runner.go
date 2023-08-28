package runner

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	docker "github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/szpp-dev-team/szpp-judge/judge/util/buffer"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

type Runner struct {
	docker              *docker.Client
	ContainerID         string
	ImageName           string
	ContainerWorkingDir string
	HostBindDir         string
}

const (
	RUNNER_STOP_TIMEOUT int = 60 * 10 // 10 minutes
)

func New(
	ctx context.Context,
	dc *docker.Client,
	imageName string,
	opts ...optionFunc,
) (*Runner, error) {
	opt := defaultOption()
	for _, f := range opts {
		f(&opt)
	}

	mounts := make([]mount.Mount, 0, 1)
	if len(opt.hostBindDirAbsPath) > 0 {
		mounts = append(mounts, mount.Mount{
			Type:        mount.TypeBind,
			Source:      opt.hostBindDirAbsPath,
			Target:      opt.containerBindDirAbsPath,
			Consistency: mount.ConsistencyFull,
		})
	}

	stopTimeout := RUNNER_STOP_TIMEOUT // const はアドレス取得できないので一旦変数へ
	runInit := true
	containerCfg := &container.Config{
		Tty:             true,
		Env:             []string{"SZPP_JUDGE=1"},
		Cmd:             []string{"/bin/sh"},
		Image:           imageName,
		WorkingDir:      opt.workingDir,
		NetworkDisabled: !opt.allowNetwork,
		StopTimeout:     &stopTimeout,
	}

	hostCfg := &container.HostConfig{
		LogConfig:  container.LogConfig{},
		AutoRemove: true,
		IpcMode:    container.IPCModeNone,
		Init:       &runInit,
		Resources: container.Resources{
			Memory:     opt.maxContainerMemoryBytes,
			MemorySwap: opt.maxContainerMemoryBytes,
			PidsLimit:  opt.maxProcNum,
			Ulimits:    opt.createUlimits(),
		},
		Mounts: mounts,
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
		ContainerWorkingDir: opt.workingDir,
		HostBindDir:         opt.hostBindDirAbsPath,
	}, nil
}

func (r *Runner) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	err := r.docker.ContainerRemove(ctx, r.ContainerID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		slog.Error("Failed to remove container:", "ID", r.ContainerID, "err", err)
	} else {
		slog.Info("Successfully removed the container:", "ID", r.ContainerID)
	}

	// if err = os.RemoveAll(r.HostWorkingDir); err != nil {
	// 	slog.Error("Failed to remove temporary dir on the host:", "path", r.HostWorkingDir)
	// }
	cancel()
}

type ExecOption struct {
	AsRootUser      bool
	Stdin           *bytes.Buffer
	Cmd             []string
	StdoutReadLimit unit.ByteSize
	StderrReadLimit unit.ByteSize
}

type ExecResult struct {
	ExitCode          int
	Stdout            string
	Stderr            string
	StdoutIsTruncated bool
	StderrIsTruncated bool
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

	stdout := buffer.NewLimitedWriter(int(opt.StdoutReadLimit), nil)
	stderr := buffer.NewLimitedWriter(int(opt.StderrReadLimit), nil)
	outputDone := make(chan error)
	go func() {
		_, err := stdcopy.StdCopy(&stdout, &stderr, resp.Reader)
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

	res, err := r.docker.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{
		ExitCode:          res.ExitCode,
		Stdout:            string(stdout.Buf),
		Stderr:            string(stderr.Buf),
		StdoutIsTruncated: stdout.HasOverflowed(),
		StderrIsTruncated: stderr.HasOverflowed(),
	}, nil
}

func (r *Runner) PrintLogs(ctx context.Context) error {
	os.Stdout.Sync()
	os.Stderr.Sync()

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
