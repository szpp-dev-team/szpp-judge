package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	procinfo "github.com/szpp-dev-team/szpp-judge/judge/util/procinfo"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

func main() {
	if err := subMain(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: monitoring_exec TIME_LIMI_ms MEM_LIMIT_MiB CMD ARGS...")
}

type parseArgResult struct {
	timeLimit time.Duration
	memLimit  unit.ByteSize
	cmd       []string
}

func parseArg(args []string) (parseArgResult, error) {
	if len(args) < 4 {
		return parseArgResult{}, fmt.Errorf("Too few arguments")
	}
	n, err := strconv.Atoi(args[1])
	if err != nil {
		return parseArgResult{}, fmt.Errorf("Expected time limit[ms] in args[1], but got %v", args[1])
	}
	timeLimit := time.Millisecond * time.Duration(n)

	n, err = strconv.Atoi(args[2])
	if err != nil {
		return parseArgResult{}, fmt.Errorf("Expected memory limit[MiB] in args[2], but got %v", args[2])
	}
	memLimit := unit.MiB * unit.ByteSize(n)

	return parseArgResult{
		timeLimit: timeLimit,
		memLimit:  memLimit,
		cmd:       args[3:],
	}, nil
}

func subMain(args []string) error {
	a, err := parseArg(args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		printUsage()
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "timeLimit: %d ms\n", a.timeLimit.Milliseconds())
	fmt.Fprintf(os.Stderr, "meoryLimit: %d kiB\n", a.memLimit/unit.KiB)

	ctx, cancel := context.WithTimeout(context.Background(), a.timeLimit+500*time.Millisecond)

	cmd := exec.CommandContext(ctx, a.cmd[0], a.cmd[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	maxRSS := unit.ByteSize(0)

	if err := cmd.Start(); err != nil {
		cancel()
		return err
	}
	pid := cmd.Process.Pid
	startAt := time.Now()

	// 定期的に RSS を調べる
	go func() {
		ticker := time.NewTicker(3 * time.Millisecond)
		defer ticker.Stop()

		maxRSS = procinfo.FetchRSS(pid)

		for {
			select {
			case <-ctx.Done():
				log.Println("Finish ticking.")
				return

			case <-ticker.C:
				if rss := procinfo.FetchRSS(pid); rss > maxRSS {
					maxRSS = rss
					if maxRSS > a.memLimit {
						log.Println("Memory limit exceeded!")
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
		log.Printf("Failed to wait process: %v", err)
	}

	os.Stdout.Sync()
	os.Stderr.Sync()
	fmt.Println("--------------------")
	fmt.Printf("elapsed: %d ms\n", elapsed.Milliseconds())
	fmt.Printf("max RSS: %d kiB\n", maxRSS/unit.KiB)
	fmt.Printf("exit code: %d\n", cmd.ProcessState.ExitCode())

	return nil
}
