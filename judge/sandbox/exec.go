package sandbox

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"strconv"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/szpp-dev-team/szpp-judge/judge/util/buffer"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

type ExecOption struct {
	// nil の場合は stdin への入力をしない
	Stdin *bytes.Buffer
	// root ユーザで実行するなら true
	AsRootUser bool
	// コマンド
	Cmd []string
	// stdout を読み取るサイズ上限 (この上限を超えてもプロセスを強制終了することはない)
	StdoutReadLimit unit.Byte
	// stderr を読み取るサイズ上限 (この上限を超えてもプロセスを強制終了することはない)
	StderrReadLimit unit.Byte
	// stderr の出力を stdout にマージする (この場合 ExecResult.Stderr は空文字列になる)
	MergeStderrToStdout bool
	// 環境変数 (e.g. "HOME=/home/hoge")
	Env []string
}

type ExecResult struct {
	// -1 の場合は ExitCode が得られなかったことを示す (TLEやMLEで強制終了させたときなど)
	ExitCode int
	Stdout   string
	Stderr   string
	// ExecOption.StdoutReadLimit を超えた場合は true
	StdoutOverflowed bool
	// ExecOption.StderrReadLimit を超えた場合は true
	StderrOverflowed bool
}

type ExecWithSzppRunResult struct {
	ExecResult
	ExecTime   time.Duration
	ExecMemory unit.Byte
}

func (sb *Sandbox) Exec(ctx context.Context, opt ExecOption) (ExecResult, error) {
	user := ""
	if !opt.AsRootUser {
		user = "1234:1234"
	}

	execID, err := sb.docker.ContainerExecCreate(ctx, sb.ContainerID, types.ExecConfig{
		User:         user,
		AttachStdin:  opt.Stdin != nil,
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          opt.Cmd,
		Env:          opt.Env,
	})
	if err != nil {
		return ExecResult{}, err
	}

	resp, err := sb.docker.ContainerExecAttach(ctx, execID.ID, types.ExecStartCheck{})
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
	stderr := stdout
	if !opt.MergeStderrToStdout {
		stderr = buffer.NewLimitedWriter(int(opt.StderrReadLimit), nil)
	}

	outputDone := make(chan error)
	go func() {
		_, err := stdcopy.StdCopy(stdout, stderr, resp.Reader)
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

	res, err := sb.docker.ContainerExecInspect(ctx, execID.ID)
	if err != nil {
		return ExecResult{}, err
	}

	stdoutStr := string(stdout.Buf)
	stderrStr := ""
	if !opt.MergeStderrToStdout {
		stderrStr = string(stderr.Buf)
	}

	return ExecResult{
		ExitCode:         res.ExitCode,
		Stdout:           stdoutStr,
		Stderr:           stderrStr,
		StdoutOverflowed: stdout.HasOverflowed(),
		StderrOverflowed: stderr.HasOverflowed(),
	}, nil
}

// szpprun 配下で実行する。
// szpprun については {GitRepoRoot}/lang/_docker/_szpprun/ を参照。
func (sb *Sandbox) ExecWithSzppRun(
	ctx context.Context,
	execTimeLimit time.Duration,
	execMemoryLimit unit.Byte,
	opt ExecOption,
) (ExecWithSzppRunResult, error) {
	cmd := []string{
		"szpprun",
		strconv.FormatInt(int64(execTimeLimit.Milliseconds()), 10),
		strconv.FormatUint(uint64(execMemoryLimit/unit.MiB), 10),
	}
	cmd = append(cmd, opt.Cmd...)
	opt.Cmd = cmd

	if !sb.DevMode {
		opt.Env = append(opt.Env, "SZPPRUN_LOG_LEVEL=warn")
	}

	res, err := sb.Exec(ctx, opt)
	if err != nil {
		return ExecWithSzppRunResult{}, err
	}

	t, mem, exitCode, err := parseExecResultFile(sb.HostBindDir)
	if err != nil {
		return ExecWithSzppRunResult{}, err
	}

	res.ExitCode = exitCode

	return ExecWithSzppRunResult{
		ExecResult: res,
		ExecTime:   t,
		ExecMemory: mem,
	}, nil
}
