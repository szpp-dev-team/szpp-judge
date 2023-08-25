package runner

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	docker "github.com/docker/docker/client"
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
		Resources:  container.Resources{},
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
