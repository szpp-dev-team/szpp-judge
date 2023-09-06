package sandboxtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloC(t *testing.T) {
	ctx := context.Background()
	res := CompileAndExec(t, ctx, sbGcc, "hello.c", nil)

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Zero(t, res.Exec.ExitCode)
	assert.Equal(t, "hello\n", res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assert.False(t, res.Exec.StdoutOverflowed)
	assert.False(t, res.Exec.StderrOverflowed)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestHelloPy(t *testing.T) {
	ctx := context.Background()
	res := CompileAndExec(t, ctx, sbPython, "hello.py", nil)

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Zero(t, res.Exec.ExitCode)
	assert.Equal(t, "hello\n", res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assert.False(t, res.Exec.StdoutOverflowed)
	assert.False(t, res.Exec.StderrOverflowed)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestNetworkShouldBeDisabled(t *testing.T) {
	ctx := context.Background()
	res := CompileAndExec(t, ctx, sbPython, "get_example.com.py", nil)

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Greater(t, res.Exec.ExitCode, 0)
	assert.Empty(t, res.Exec.Stdout)
	assert.Contains(t, res.Exec.Stderr, "[Errno -3] Temporary failure in name resolution")
	assert.False(t, res.Exec.StdoutOverflowed)
	assert.False(t, res.Exec.StderrOverflowed)
	assertExecutionTimeAndMemoryOK(t, res)
}

func Test10MiBStdin(t *testing.T) {
	ctx := context.Background()

	v := []byte("0123456789ABCDEF")
	n := 10 * 1024 * 1024
	stdin := make([]byte, n)
	for i := 0; i < n; i += len(v) {
		copy(stdin[i:], v)
	}
	res := CompileAndExec(t, ctx, sbPython, "check_stdin.py", stdin)

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	wantStdout := fmt.Sprintf(
		`len(s)=%d
s[:5]='%s'
s[-5:]='%s'
`, n, "01234", "BCDEF")

	assert.Zero(t, res.Exec.ExitCode)
	assert.Equal(t, wantStdout, res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestStdinNonFullReadOK(t *testing.T) {
	ctx := context.Background()
	res := CompileAndExec(t, ctx, sbGcc, "only_read_1word.cpp", []byte("abc def xyz"))

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Zero(t, res.Exec.ExitCode)
	assert.Equal(t, "abc\n", res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestNoOutputOK(t *testing.T) {
	ctx := context.Background()
	res := CompileAndExec(t, ctx, sbGcc, "nop.c", nil)

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Zero(t, res.Exec.ExitCode)
	assert.Empty(t, res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assert.False(t, res.Exec.StdoutOverflowed)
	assert.False(t, res.Exec.StderrOverflowed)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestNoStdinReadOK(t *testing.T) {
	ctx := context.Background()
	res := CompileAndExec(t, ctx, sbGcc, "nop.c", []byte("helloooooooooooo"))

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Zero(t, res.Exec.ExitCode)
	assert.Empty(t, res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assert.False(t, res.Exec.StdoutOverflowed)
	assert.False(t, res.Exec.StderrOverflowed)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestExitCode(t *testing.T) {
	ctx := context.Background()
	res := CompileAndExec(t, ctx, sbGcc, "exit_42.c", nil)

	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Equal(t, 42, res.Exec.ExitCode)
	assert.Empty(t, res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assert.False(t, res.Exec.StdoutOverflowed)
	assert.False(t, res.Exec.StderrOverflowed)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestTimeLimit(t *testing.T) {
	ctx := context.Background()

	t.Run("InfiniteLoop", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbPython, "infinite_loop.py", nil)
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assertExecutionMemoryOK(t, res)
		assertExecutionTLE(t, res)
	})

	t.Run("Sleep", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbPython, "sleep_n_ms.py", []byte("2000"))
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assertExecutionMemoryOK(t, res)
		assertExecutionTLE(t, res)
	})

	t.Run("Sleep in DevMode sandbox", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbPythonDev, "sleep_n_ms.py", []byte("2000"))
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assert.Contains(t, res.Exec.Stderr, `level=INFO msg="Failed to wait a process:" err="signal: killed"`)
		assertExecutionMemoryOK(t, res)
		assertExecutionTLE(t, res)
	})
}

func TestMemoryLimit(t *testing.T) {
	ctx := context.Background()

	t.Run("alloc 64 MiB", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbGcc, "alloc_n_MiB.c", []byte("64"))
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assert.Empty(t, res.Exec.Stderr)
		assertExecutionTimeOK(t, res)
		assertExecutionMLE(t, res)
	})

	t.Run("alloc 64 MiB in DevMode sandbox", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbGccDev, "alloc_n_MiB.c", []byte("64"))
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assert.Contains(t, res.Exec.Stderr, `level=INFO`)
		assert.Contains(t, res.Exec.Stderr, `Memory limit exceeded`)
		assertExecutionTimeOK(t, res)
		assertExecutionMLE(t, res)
	})
}

func TestProcLimit(t *testing.T) {
	ctx := context.Background()

	res := CompileAndExec(t, ctx, sbPython, "spawn_n_subprocess.py", []byte("20"))
	assertCompilationNoOutput(t, res)
	assertCompilationTimeAndMemoryOK(t, res)

	assert.Greater(t, res.Exec.ExitCode, 0)
	assert.Contains(t, res.Exec.Stderr, `BlockingIOError: [Errno 11] Resource temporarily unavailable`)
	assertExecutionTimeAndMemoryOK(t, res)
}

func TestOutputReadLimit(t *testing.T) {
	ctx := context.Background()

	assert.NotEqual(t, ExecStdoutReadLimit, ExecStderrReadLimit)

	t.Run("InfiniteStdout", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbGcc, "infinite_stdout.c", nil)
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assert.Len(t, res.Exec.Stdout, int(ExecStdoutReadLimit))
		assert.Len(t, res.Exec.Stderr, 0)
		assert.True(t, res.Exec.StdoutOverflowed)
		assert.False(t, res.Exec.StderrOverflowed)
		assert.Equal(t, "####", res.Exec.Stdout[:4])
		assertExecutionTLE(t, res)
		assertExecutionMemoryOK(t, res)
	})

	t.Run("InfiniteStderr", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbGcc, "infinite_stderr.c", nil)
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assert.Len(t, res.Exec.Stdout, 0)
		assert.Len(t, res.Exec.Stderr, int(ExecStderrReadLimit))
		assert.False(t, res.Exec.StdoutOverflowed)
		assert.True(t, res.Exec.StderrOverflowed)
		assert.Equal(t, "XXXX", res.Exec.Stderr[:4])
		assertExecutionTLE(t, res)
		assertExecutionMemoryOK(t, res)
	})

	t.Run("InfiniteBothStdoutStderr", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbGcc, "infinite_stdout_stderr.c", nil)
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assert.Len(t, res.Exec.Stdout, int(ExecStdoutReadLimit))
		assert.Len(t, res.Exec.Stderr, int(ExecStderrReadLimit))
		assert.True(t, res.Exec.StdoutOverflowed)
		assert.True(t, res.Exec.StderrOverflowed)
		assert.Equal(t, "oooo", res.Exec.Stdout[:4])
		assert.Equal(t, "EEEE", res.Exec.Stderr[:4])
		assertExecutionTLE(t, res)
		assertExecutionMemoryOK(t, res)
	})
}

func TestFileWriteLimit(t *testing.T) {
	ctx := context.Background()

	t.Run("Write 3MiB to single file in DevMode sandbox", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbGccDev, "write_n_MiB_to_k_files.c", []byte("3 1"))
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Equal(t, -1, res.Exec.ExitCode)
		assert.Empty(t, res.Exec.Stdout)
		assert.Contains(t, res.Exec.Stderr, `level=INFO msg="Failed to wait a process:" err="signal: file size limit exceeded"`)
		assert.False(t, res.Exec.StdoutOverflowed)
		assert.False(t, res.Exec.StderrOverflowed)
		assertExecutionTimeAndMemoryOK(t, res)
	})

	// ulimit の fsize は「1ファイルあたり」の書き込みサイズ上限を指定するものである。
	// ulimit の ofiles は「同時に」開けるファイル数である。
	// そのため、1ファイルずつ fsize を超えない範囲で書き込みをすることを逐次繰り返せばディスクが食いつぶされる恐れがある
	// 以下のテストケースはそれを示すためのもの。
	t.Run("Write 1MiB to 5 files", func(t *testing.T) {
		res := CompileAndExec(t, ctx, sbGccDev, "write_n_MiB_to_k_files.c", []byte("1 5"))
		assertCompilationNoOutput(t, res)
		assertCompilationTimeAndMemoryOK(t, res)

		assert.Zero(t, res.Exec.ExitCode)
		assert.Empty(t, res.Exec.Stdout)
		assert.Empty(t, res.Exec.Stderr)
		assert.False(t, res.Exec.StdoutOverflowed)
		assert.False(t, res.Exec.StderrOverflowed)
		assertExecutionTimeAndMemoryOK(t, res)
	})
}
