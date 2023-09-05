package sandbox

import (
	"context"
	"log/slog"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	docker "github.com/docker/docker/client"
	"github.com/docker/go-units"
)

type Sandbox struct {
	docker              *docker.Client
	DevMode             bool
	ContainerID         string
	ImageName           string
	ContainerWorkingDir string
	ContainerBindDir    string
	HostBindDir         string
}

const (
	CONTAINER_STOP_TIMEOUT int = 60 * 10 // 10 minutes
)

func New(
	ctx context.Context,
	dc *docker.Client,
	imageName string,
	opts ...optionFunc,
) (*Sandbox, error) {
	opt := defaultOption()
	for _, f := range opts {
		f(&opt)
	}

	mounts := make([]mount.Mount, 0, 1)
	if opt.hostBindDirAbsPath != "" {
		mounts = append(mounts, mount.Mount{
			Type:        mount.TypeBind,
			Source:      opt.hostBindDirAbsPath,
			Target:      opt.containerBindDirAbsPath,
			Consistency: mount.ConsistencyFull,
		})
	}

	stopTimeout := CONTAINER_STOP_TIMEOUT // const はアドレス取得できないので一旦変数へ
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
			Memory:     int64(opt.containerMemoryLimit),
			MemorySwap: int64(opt.containerMemoryLimit),
			PidsLimit:  &opt.procNumLimit,
			Ulimits: []*units.Ulimit{
				{
					// core ファイルを作らない
					Name: "core",
					Hard: 0,
					Soft: 0,
				},
				{
					// スタックサイズを無制限にする
					Name: "stack",
					Hard: -1,
					Soft: -1,
				},
				{
					// 1プロセスあたりの CPU 時間の上限 (15sec 固定)
					// szpprun で実行時間は制限できるが、万が一のためにここに最後の砦を築く
					// ※ CPU 時間のみを制限するので sleep 等による時間や IO 時間は制限できない
					Name: "cpu",
					Hard: 15,
					Soft: 15,
				},
			},
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

	return &Sandbox{
		docker:              dc,
		DevMode:             opt.devMode,
		ContainerID:         resp.ID,
		ImageName:           imageName,
		ContainerWorkingDir: opt.workingDir,
		ContainerBindDir:    opt.containerBindDirAbsPath,
		HostBindDir:         opt.hostBindDirAbsPath,
	}, nil
}

func (sb *Sandbox) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := sb.docker.ContainerRemove(ctx, sb.ContainerID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		slog.Error("Failed to remove container:", "ID", sb.ContainerID, "err", err)
	} else {
		slog.Info("Successfully removed the container:", "ID", sb.ContainerID)
	}
}
