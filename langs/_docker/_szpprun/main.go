package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/szpp-dev-team/szpp-judge/szpprun/procinfo"
	"golang.org/x/sys/unix"
)

const (
	MemUsageCheckInterval        = 3 * time.Millisecond
	LogLevelEnvKey               = "SZPPRUN_LOG_LEVEL"
	RlimitStackSize       uint64 = unix.RLIM_INFINITY // スタックサイズ制限なし
)

func printUsage() {
	fmt.Fprintf(os.Stderr,
		`
  Usage: szpprun timeLimit_ms memLimit_MiB fileSizeLimit_MiB numOpenFileLimit CMD...

  ENVIRONMENT:
		%s  error|warn|info|debug  (default: info)
`, LogLevelEnvKey)
}

func main() {
	logger := newLoggerFromEnv()
	if err := subMain(os.Args, logger); err != nil {
		log.Fatal(err)
	}
}

func parseSlogLevel(name string, fallback slog.Level) slog.Level {
	switch strings.ToLower(name) {
	case "error":
		return slog.LevelError
	case "warn":
		return slog.LevelWarn
	case "info":
		return slog.LevelInfo
	case "debug":
		return slog.LevelDebug
	case "":
		return fallback
	default:
		log.Printf("Invalid log level %q: using fallback %v", name, fallback)
		return fallback
	}
}

func newLoggerFromEnv() *slog.Logger {
	level := parseSlogLevel(os.Getenv(LogLevelEnvKey), slog.LevelInfo)
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	}))
}

type parseArgResult struct {
	timeLimit          time.Duration
	memLimitKib        uint
	rlimitFileSizeByte uint64
	rlimitNumOpenFiles uint64
	cmd                []string
}

func parseArg(args []string) (parseArgResult, error) {
	if len(args) <= 5 {
		return parseArgResult{}, fmt.Errorf("Too few arguments")
	}

	// parse timeLimit as [ms]
	n, err := strconv.ParseUint(args[1], 10, 32)
	if err != nil {
		return parseArgResult{}, fmt.Errorf("Expected time limit [ms] in args[1], but got %v", args[1])
	}
	timeLimit := time.Millisecond * time.Duration(n)

	// parse memLimit as [MiB]
	n, err = strconv.ParseUint(args[2], 10, 32)
	if err != nil {
		return parseArgResult{}, fmt.Errorf("Expected memory limit [MiB] in args[2], but got %v", args[2])
	}
	memLimitKib := uint(n) << 10

	// parse fileSizeLimit as [MiB]
	n, err = strconv.ParseUint(args[3], 10, 32)
	if err != nil {
		return parseArgResult{}, fmt.Errorf("Expected file size limit [MiB] in args[3], but got %v", args[3])
	}
	fileSizeLimitByte := n << 20

	// parse numOpenFileLimit
	n, err = strconv.ParseUint(args[4], 10, 32)
	if err != nil {
		return parseArgResult{}, fmt.Errorf("Expected num of open files limit in args[4], but got %v", args[4])
	}
	numOpenFileLimit := n

	return parseArgResult{
		timeLimit:          timeLimit,
		memLimitKib:        memLimitKib,
		rlimitFileSizeByte: fileSizeLimitByte,
		rlimitNumOpenFiles: numOpenFileLimit,
		cmd:                args[5:],
	}, nil
}

// log パッケージを誤って使わないようあえて変数名を log にしている
func subMain(args []string, log *slog.Logger) error {
	a, err := parseArg(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		printUsage()
		os.Exit(1)
	}

	// set ulimit (この親プロセス自身にも ulimit 制約が課せられるので注意)
	unix.Setrlimit(unix.RLIMIT_STACK, &unix.Rlimit{Cur: RlimitStackSize, Max: RlimitStackSize})
	unix.Setrlimit(unix.RLIMIT_FSIZE, &unix.Rlimit{Cur: a.rlimitFileSizeByte, Max: a.rlimitFileSizeByte})
	unix.Setrlimit(unix.RLIMIT_NOFILE, &unix.Rlimit{Cur: a.rlimitNumOpenFiles, Max: a.rlimitNumOpenFiles})

	ctx, cancel := context.WithTimeout(context.Background(), a.timeLimit+100*time.Millisecond)

	cmd := exec.CommandContext(ctx, a.cmd[0], a.cmd[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	maxRSSKib := uint(0)

	if err := cmd.Start(); err != nil {
		cancel()
		return err
	}
	pid := cmd.Process.Pid
	startAt := time.Now()

	// 定期的に RSS (Resident Set Size) を調べる
	go func() {
		ticker := time.NewTicker(MemUsageCheckInterval)
		defer ticker.Stop()

		// 最初の数ミリ秒間は短い間隔で取得
		maxRSSKib = procinfo.FetchRSSKib(log, pid)

		time.Sleep(time.Microsecond * 500)
		rss := procinfo.FetchRSSKib(log, pid)
		if rss > maxRSSKib {
			maxRSSKib = rss
		}

		time.Sleep(time.Microsecond * 500)
		rss = procinfo.FetchRSSKib(log, pid)
		if rss > maxRSSKib {
			maxRSSKib = rss
		}

		for {
			select {
			case <-ctx.Done():
				return

			case <-ticker.C:
				if rss = procinfo.FetchRSSKib(log, pid); rss > maxRSSKib {
					maxRSSKib = rss
					if maxRSSKib > a.memLimitKib {
						log.Info("Memory limit exceeded:", slog.Uint64("RSS[kiB]", uint64(maxRSSKib)))
						cancel()
						return
					}
				}
			}
		}
	}()

	err = cmd.Wait()
	elapsed := time.Since(startAt)
	cancel()
	if err != nil {
		// context の WithTimeout で既に kill された場合もここに引っ掛かるので無視
		log.Info("Failed to wait a process:", "err", err)
	}

	os.Stdout.Sync()
	os.Stderr.Sync()
	if log.Enabled(context.Background(), slog.LevelDebug) {
		fmt.Println("--------------------")
		fmt.Printf("elapsed: %d ms\n", elapsed.Milliseconds())
		fmt.Printf("max RSS: %d kiB (=%d MiB)\n", maxRSSKib, maxRSSKib>>10)
		fmt.Printf("exit code: %d\n", cmd.ProcessState.ExitCode())
	}

	if err := saveResult(elapsed, maxRSSKib, cmd.ProcessState.ExitCode()); err != nil {
		return err
	}

	return nil
}

func saveResult(elapsed time.Duration, maxRSSKib uint, exitCode int) error {
	const NAME string = ".szpprun-result.txt"
	f, err := os.Create(NAME)
	if err != nil {
		return fmt.Errorf("cannot create file named %q: %w", NAME, err)
	}
	defer f.Close()
	fmt.Fprintf(f, "%d %d %d", elapsed.Milliseconds(), maxRSSKib, exitCode)
	return nil
}
