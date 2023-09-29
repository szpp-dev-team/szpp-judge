package sandbox

import (
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

type option struct {
	devMode                 bool
	workingDir              string
	hostBindDirAbsPath      string
	containerBindDirAbsPath string
	containerMemoryLimit    unit.Byte
	procNumLimit            int64
	allowNetwork            bool
}

func defaultOption() option {
	return option{
		containerMemoryLimit: unit.GiB * 1,
		procNumLimit:         1024,
		devMode:              true,
	}
}

type optionFunc func(*option)

// ネットワーク接続を許可する
func WithAllowNetwork(o *option) {
	o.allowNetwork = true
}

// 開発モードを有効にするか否か (デフォルト：有効)
func WithDevMode(enabled bool) optionFunc {
	return func(o *option) {
		o.devMode = enabled
	}
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
func WithProcNumLimit(limit int64) optionFunc {
	return func(o *option) {
		o.procNumLimit = limit
	}
}

// コンテナ内で使えるメモリの総量。スワップメモリは0byte固定。
// 1プロセス当たりのメモリ使用量ではないので注意。
func WithContainerMemoryLimit(limit unit.Byte) optionFunc {
	return func(o *option) {
		o.containerMemoryLimit = limit
	}
}
