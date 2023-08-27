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

	"github.com/docker/docker/client"
	"github.com/szpp-dev-team/szpp-judge/judge/config"
	"github.com/szpp-dev-team/szpp-judge/judge/lang"
	"github.com/szpp-dev-team/szpp-judge/judge/runner"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: runner SOURCE_FILE [STDIN_FILE|-]\n")
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

	hostWorkingDir := path.Join(cfg.WorkingDirAbsRoot, "runner")
	_ = os.RemoveAll(hostWorkingDir)
	if err = os.MkdirAll(hostWorkingDir, 0750); err != nil {
		return err
	}

	sourceFile := args[1]
	langID := DetectLang(sourceFile)
	langMeta, err := lang.GetLangMeta(langID)
	if err != nil {
		return err
	}

	sourceFileInBindDir := path.Join(hostWorkingDir, langMeta.SourceFileName)
	slog.Info("Copy file:", "src", sourceFile, "dst", sourceFileInBindDir)
	err = CopyFile(sourceFile, sourceFileInBindDir)
	if err != nil {
		return err
	}

	var stdin []byte = nil
	if len(args) >= 3 {
		stdin, err = ReadFromStdinOrFile(args[2])
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
	r, err := runner.New(ctx, dc, langMeta.ImageName, hostWorkingDir)
	if err != nil {
		return err
	}
	defer r.Close()
	slog.Info("Created runner:", " ContainerID", r.ContainerID[:6])

	err = run(ctx, r, langMeta, stdin)
	if err != nil {
		return err
	}

	return nil
}

func run(
	ctx context.Context,
	r *runner.Runner,
	lm *lang.LangMeta,
	stdin []byte,
) error {
	slog.Info("Executing compile cmd:", "cmd", lm.CompileCmd)
	res, err := r.Exec(ctx, runner.ExecOption{
		AsRootUser: false,
		Stdin:      nil,
		Cmd:        lm.CompileCmd,
	})
	if err != nil {
		return err
	}
	slog.Info("Finished compilation")
	showExecResult(res)
	if res.ExitCode != 0 {
		return fmt.Errorf("compile error: exitCode=%d, msg=%s", res.ExitCode, res.Stderr)
	}

	slog.Info("Executing exec cmd:", "cmd", lm.ExecCmd, "len(stdin)", len(stdin), "stdin", truncateStr(string(stdin), 30))
	res, err = r.Exec(ctx, runner.ExecOption{
		AsRootUser: false,
		Stdin:      bytes.NewBuffer(stdin),
		Cmd:        lm.ExecCmd,
	})
	if err != nil {
		return err
	}
	slog.Info("Finished execution")
	showExecResult(res)

	fmt.Fprintln(os.Stderr, "----------------- logs ------------------")
	_ = r.PrintLogs(ctx)
	fmt.Fprintln(os.Stderr, "-----------------------------------------")

	return nil
}

func truncateStr(s string, limitLen int) string {
	if len(s) < limitLen {
		return s
	} else {
		return s[:limitLen] + "..."
	}
}

func showExecResult(r runner.ExecResult) {
	fmt.Fprintln(os.Stderr, "- - - ExecResult - - -")
	fmt.Fprintf(os.Stderr, "len(Stdout): %d\nlen(Stderr): %d\nExitCode: %d\n", len(r.Stdout), len(r.Stderr), r.ExitCode)

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

func DetectLang(filepath string) lang.LangID {
	e := path.Ext(filepath)
	switch e {
	case ".c":
		return lang.C
	case ".cpp", ".cxx":
		return lang.CPP
	case ".java":
		return lang.JAVA
	case ".py":
		return lang.PYTHON
	}
	log.Fatalf("Unsupported language: filename='%s'", path.Base(filepath))
	panic(filepath)
}

func ReadFromStdinOrFile(filepathOrHyphen string) ([]byte, error) {
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

func CopyFile(srcPath string, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}
