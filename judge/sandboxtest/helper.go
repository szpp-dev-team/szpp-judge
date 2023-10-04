package sandboxtest

import (
	"bytes"
	"path"
	"testing"

	"github.com/szpp-dev-team/szpp-judge/judge/sandbox"
	"github.com/szpp-dev-team/szpp-judge/judge/util/fsutil"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
	"github.com/szpp-dev-team/szpp-judge/langs"
	"golang.org/x/net/context"
)

type Result struct {
	Compile sandbox.ExecResult
	Exec    sandbox.ExecResult
}

func CompileAndExec(t *testing.T, ctx context.Context, sb *sandbox.Sandbox, filename string, stdin []byte) *Result {
	t.Helper()

	langMeta := prepare(t, sb, filename)

	compileRes, err := sb.Exec(ctx, sandbox.ExecOption{
		Cmd:                 langMeta.CompileCmd,
		StdoutReadLimit:     CompileMessageReadLimit,
		MergeStderrToStdout: true,
	}, CompileLimit)
	if err != nil {
		t.Fatalf("Failed to exec compile: %v", err)
	}
	if compileRes.ExitCode != 0 {
		t.Fatalf("Compile error:\n%s", compileRes.Stdout)
	}
	t.Logf(
		"Execution   done: time=%3d[ms], mem=%5d[kiB] exitCode=%d, len(out)=%d, len(err)=%d",
		compileRes.ExecTime.Milliseconds(),
		compileRes.ExecMemory/unit.KiB,
		compileRes.ExitCode,
		len(compileRes.Stdout),
		len(compileRes.Stderr),
	)

	var stdinBuf *bytes.Buffer = nil
	if stdin != nil {
		stdinBuf = bytes.NewBuffer(stdin)
	}

	execRes, err := sb.Exec(ctx, sandbox.ExecOption{
		Cmd:             langMeta.ExecCmd,
		Stdin:           stdinBuf,
		StdoutReadLimit: ExecStdoutReadLimit,
		StderrReadLimit: ExecStderrReadLimit,
	}, ExecLimit)
	if err != nil {
		t.Fatalf("Failed to exec: %v", err)
	}
	t.Logf(
		"Compilation done: time=%3d[ms], mem=%4d[kiB] exitCode=%d, len(out)=%d, len(err)=%d",
		execRes.ExecTime.Milliseconds(),
		execRes.ExecMemory/unit.KiB,
		execRes.ExitCode,
		len(execRes.Stdout),
		len(execRes.Stderr),
	)

	return &Result{
		Compile: compileRes,
		Exec:    execRes,
	}
}

func prepare(t *testing.T, sb *sandbox.Sandbox, filename string) *langs.Meta {
	t.Helper()

	langID := DetectLang(t, filename)
	t.Logf("CompileAndExec(): filename=%s, langID=%s", filename, langID)

	langMeta := GetLangMeta(t, langID)

	{
		src := path.Join(fsutil.GetGoModAbsDir(), "_test_code", filename)
		dst := path.Join(sb.HostBindDir, langMeta.SourceFile)
		if err := fsutil.CopyFile(src, dst); err != nil {
			t.Fatalf("Cannot copy file: %v", err)
		}
	}

	return langMeta
}

func GetLangMeta(t *testing.T, id langs.LangID) *langs.Meta {
	m, ok := langs.Get(id)
	if !ok {
		t.Fatalf("Cannot get lang meta: langID=%s", id)
	}
	return m
}

func DetectLang(t *testing.T, filepath string) langs.LangID {
	t.Helper()

	switch path.Ext(filepath) {
	case ".c":
		return langs.C_11_GCC
	case ".cpp":
		return langs.CPP_20_GCC
	case ".java":
		return langs.JAVA_21_OPENJDK
	case ".py":
		return langs.PYTHON_311_CPYTHON
	}

	t.Fatalf("Unsupported language: filename='%s'", path.Base(filepath))
	panic(filepath)
}
