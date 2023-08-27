package runner

import (
	"math"
	"time"

	docker "github.com/docker/go-units"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

type option struct {
	workingDir              string
	hostBindDirAbsPath      string
	containerBindDirAbsPath string
	maxContainerMemoryBytes int64
	maxProcNum              *int64
	maxStackKiB             *int64
	maxWriteFileBytes       *int64
	maxOpenFiles            *int64
	maxCPUTimeSecs          *int64
	allowNetwork            bool
}

func defaultOption() option {
	return option{}
}

type optionFunc func(*option)

// ネットワーク接続を許可する
func WithAllowNetwork(o *option) {
	o.allowNetwork = true
}

func WithWorkingDir(containerAbsPath string) optionFunc {
	return func(o *option) {
		o.workingDir = containerAbsPath
	}
}

// ホストとコンテナのディレクトリをバインドする。複数回の指定は不可能。
func WithBindDir(hostAbsPath, containerAbsPath string) optionFunc {
	return func(o *option) {
		o.hostBindDirAbsPath = hostAbsPath
		o.containerBindDirAbsPath = containerAbsPath
	}
}

// コンテナ内で同時に存在できるプロセス数
func WithMaxProcNum(limit int64) optionFunc {
	return func(o *option) {
		o.maxProcNum = &limit
	}
}

// コンテナ内で使えるメモリの総量。スワップメモリは0byte固定。
// 1プロセス当たりのメモリ使用量ではないので注意。
func WithMaxContainerMemory(limit unit.ByteSize) optionFunc {
	return func(o *option) {
		o.maxContainerMemoryBytes = int64(limit)
	}
}

// 1プロセスあたりのスタックサイズ
func WithMaxStack(limit unit.ByteSize) optionFunc {
	x := int64(-1)
	if limit > 0 {
		x = int64(limit / unit.KiB)
	}
	return func(o *option) {
		o.maxStackKiB = &x
	}
}

// 1プロセスあたりの書き込み可能なファイルサイズ
func WithMaxWriteFile(limit unit.ByteSize) optionFunc {
	x := int64(limit)
	return func(o *option) {
		o.maxWriteFileBytes = &x
	}
}

// 1プロセスあたりのファイルオープン上限数
func WithMaxOpenFiles(limit int64) optionFunc {
	return func(o *option) {
		o.maxOpenFiles = &limit
	}
}

// 1プロセスあたりの CPU Time の上限
func WithMaxCPUTime(limit time.Duration) optionFunc {
	x := int64(-1)
	if limit > 0 {
		x = int64(math.Ceil(limit.Seconds()))
	}
	return func(o *option) {
		o.maxCPUTimeSecs = &x
	}
}

func (o *option) createUlimits() []*docker.Ulimit {
	// NOTE:
	// ulimit の nproc による指定は docker デーモンに対しての制限であり、
	// コンテナに対しての制限ではないので注意。
	// ulimit の代わりに docker の HostConfig.PidsLimit を使うとよい。
	//  https://docs.docker.jp/engine/reference/commandline/run.html#ulimits-ulimit

	ulimits := make([]*docker.Ulimit, 0, 6)
	if o.maxOpenFiles != nil {
		ulimits = append(ulimits, &docker.Ulimit{
			Name: "nofile",
			Hard: *o.maxOpenFiles,
			Soft: *o.maxOpenFiles,
		})
	}
	if o.maxStackKiB != nil {
		ulimits = append(ulimits, &docker.Ulimit{
			Name: "stack",
			Hard: *o.maxStackKiB,
			Soft: *o.maxStackKiB,
		})
	}
	if o.maxWriteFileBytes != nil {
		ulimits = append(ulimits, &docker.Ulimit{
			Name: "fsize",
			Hard: *o.maxWriteFileBytes,
			Soft: *o.maxWriteFileBytes,
		})
	}
	if o.maxCPUTimeSecs != nil {
		ulimits = append(ulimits, &docker.Ulimit{
			Name: "cpu",
			Hard: *o.maxCPUTimeSecs,
			Soft: *o.maxCPUTimeSecs,
		})
	}
	return ulimits
}
