package sandboxtest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloC(t *testing.T) {
	ctx := context.Background()

	res := CompileAndExec(t, ctx, sbGcc, "hello.c", nil)

	assert.Equal(t, 0, res.Compile.ExitCode)
	assert.Empty(t, res.Compile.Stdout)
	assert.Empty(t, res.Compile.Stderr)
	assert.False(t, res.Compile.StdoutOverflowed)
	assert.False(t, res.Compile.StderrOverflowed)
	assert.NotZero(t, res.Compile.ExecTime)
	assert.NotZero(t, res.Compile.ExecMemory)

	assert.Equal(t, 0, res.Exec.ExitCode)
	assert.Equal(t, "hello\n", res.Exec.Stdout)
	assert.Empty(t, res.Exec.Stderr)
	assert.False(t, res.Exec.StdoutOverflowed)
	assert.False(t, res.Exec.StderrOverflowed)
	assert.NotZero(t, res.Exec.ExecTime)
	assert.NotZero(t, res.Exec.ExecMemory)
}
