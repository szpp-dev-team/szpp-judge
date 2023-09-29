package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/docker/docker/client"
	"github.com/szpp-dev-team/szpp-judge/judge/config"
	"github.com/szpp-dev-team/szpp-judge/judge/sandbox"
	"github.com/szpp-dev-team/szpp-judge/judge/util/fsutil"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
	"github.com/szpp-dev-team/szpp-judge/langs"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: sandbox SOURCE_FILE [STDIN_FILE | -]\n")
}

func main() {
	err := mainSub(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainSub(args []string) error {
	slog.Info("mainSub():", "len(args)", len(args), "args", args)
	if len(args) != 2 && len(args) != 3 {
		printUsage()
		os.Exit(1)
	}

	cfg, err := config.New()
	if err != nil {
		return err
	}

	hostWorkingDir := path.Join(fsutil.GetGoModAbsDir(), "_workdir", "sandbox-demo")
	_ = os.RemoveAll(hostWorkingDir)
	if err = os.MkdirAll(hostWorkingDir, 0750); err != nil {
		return err
	}

	sourceFile := args[1]
	langID := DetectLang(sourceFile)
	langMeta, ok := langs.Get(langID)
	if !ok {
		return fmt.Errorf("Unknown langID: %q", langID)
	}

	sourceFileInBindDir := path.Join(hostWorkingDir, langMeta.SourceFile)
	slog.Info("Copy file:", "src", sourceFile, "dst", sourceFileInBindDir)
	err = fsutil.CopyFile(sourceFile, sourceFileInBindDir)
	if err != nil {
		return err
	}

	var stdin []byte = nil
	if len(args) >= 3 {
		stdin, err = ReadFromFileOrStdin(args[2])
		if err != nil {
			return err
		}
	}

	dc, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return err
	}
	defer dc.Close()

	ctx := context.Background()
	r, err := sandbox.New(ctx, dc, langMeta.DockerImage,
		sandbox.WithDevMode(cfg.ModeDev),
		sandbox.WithWorkingDir("/work"),
		sandbox.WithBindDir(hostWorkingDir, "/work"),
		sandbox.WithContainerMemoryLimit(1*unit.GiB),
		sandbox.WithProcNumLimit(32),
	)
	if err != nil {
		return err
	}
	defer r.Close()
	slog.Info("Created sandbox:", "ContainerID", r.ContainerID[:6])

	err = run(ctx, r, langMeta, stdin)
	if err != nil {
		return err
	}

	return nil
}

func run(
	ctx context.Context,
	r *sandbox.Sandbox,
	lm *langs.Meta,
	stdin []byte,
) error {
	slog.Info("Executing compile cmd:", "cmd", lm.CompileCmd)
	res, err := r.Exec(ctx, sandbox.ExecOption{
		AsRootUser:          true,
		Stdin:               nil,
		Cmd:                 lm.CompileCmd,
		StdoutReadLimit:     16 * unit.KiB,
		MergeStderrToStdout: true,
	}, sandbox.SzpprunOption{
		TimeLimit:        time.Second * 1,
		MemoryLimit:      unit.GiB * 1,
		FileWriteLimit:   unit.GiB * 2,
		NumOpenFileLimit: 1024,
	})
	if err != nil {
		return err
	}
	slog.Info("Finished compilation")
	showExecResult(res)
	if res.ExitCode != 0 {
		return fmt.Errorf("compile error: exitCode=%d", res.ExitCode)
	}

	slog.Info("Executing exec cmd:", "cmd", lm.ExecCmd, "len(stdin)", len(stdin), "stdin", truncateStr(string(stdin), 30))
	res, err = r.Exec(ctx, sandbox.ExecOption{
		AsRootUser:      false,
		Stdin:           bytes.NewBuffer(stdin),
		Cmd:             lm.ExecCmd,
		StdoutReadLimit: 256 * unit.KiB,
		StderrReadLimit: 64 * unit.KiB,
	}, sandbox.SzpprunOption{
		TimeLimit:        time.Millisecond * 500,
		MemoryLimit:      unit.MiB * 64,
		FileWriteLimit:   unit.MiB * 8,
		NumOpenFileLimit: 128,
	})
	if err != nil {
		return err
	}
	slog.Info("Finished execution")
	showExecResult(res)

	return nil
}

func truncateStr(s string, limitLen int) string {
	if len(s) < limitLen {
		return s
	} else {
		return s[:limitLen] + "..."
	}
}

func showExecResult(r sandbox.ExecResult) {
	fmt.Fprintln(os.Stderr, "- - - ExecResult - - -")
	fmt.Fprintf(os.Stderr, "len(Stdout): %d\nlen(Stderr): %d\n", len(r.Stdout), len(r.Stderr))
	fmt.Fprintf(os.Stderr, "Time:     %d ms\n", r.ExecTime.Milliseconds())
	fmt.Fprintf(os.Stderr, "Memory:   %d KiB\n", r.ExecMemory/unit.KiB)
	fmt.Fprintf(os.Stderr, "ExitCode: %d\n", r.ExitCode)

	write := func(prefix string, s string, limit int) {
		if len(s) > limit {
			fmt.Fprintf(os.Stderr, "%s%q\n", prefix, truncateStr(s, limit))
		} else {
			fmt.Fprintf(os.Stderr, "%s\n%s\n", prefix, s)
		}
	}

	write("Stdout: ", r.Stdout, 4000)
	write("Stderr: ", r.Stderr, 4000)
	fmt.Fprintln(os.Stderr, "- - - - - - - - - - - -")
}

func DetectLang(filepath string) langs.LangID {
	e := path.Ext(filepath)
	switch e {
	case ".c":
		return langs.C_11_GCC13
	case ".cpp", ".cxx":
		return langs.CPP_20_GCC13
	case ".java":
		return langs.JAVA_21_OPENJDK
	case ".py":
		return langs.PYTHON_311_CPYTHON
	}
	log.Fatalf("Unsupported language: filename='%s'", path.Base(filepath))
	panic(filepath)
}

func ReadFromFileOrStdin(filepathOrHyphen string) ([]byte, error) {
	var src io.ReadCloser

	if filepathOrHyphen == "-" {
		src = os.Stdin
	} else {
		f, err := os.Open(filepathOrHyphen)
		if err != nil {
			return nil, err
		}
		src = f
	}
	defer src.Close()

	bs, err := io.ReadAll(src)
	if err != nil {
		return nil, err
	}
	return bs, nil
}
