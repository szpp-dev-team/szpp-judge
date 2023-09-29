package sandboxtest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertCompilationNoOutput(t *testing.T, r *Result) {
	t.Helper()
	assert.Empty(t, r.Compile.Stdout)
	assert.Empty(t, r.Compile.Stderr)
	assert.False(t, r.Compile.StdoutOverflowed)
	assert.False(t, r.Compile.StderrOverflowed)
}

func assertCompilationTimeOK(t *testing.T, r *Result) {
	t.Helper()
	assert.NotZero(t, r.Compile.ExecTime)
	assert.Less(t, r.Compile.ExecTime, CompileLimit.TimeLimit)
}

func assertCompilationMemoryOK(t *testing.T, r *Result) {
	t.Helper()
	assert.NotZero(t, r.Compile.ExecMemory)
	assert.Less(t, r.Compile.ExecMemory, CompileLimit.MemoryLimit)
}

func assertCompilationTimeAndMemoryOK(t *testing.T, r *Result) {
	t.Helper()
	assertCompilationTimeOK(t, r)
	assertCompilationMemoryOK(t, r)
}

func assertExecutionTimeOK(t *testing.T, r *Result) {
	t.Helper()
	assert.NotZero(t, r.Exec.ExecTime)
	assert.Less(t, r.Exec.ExecTime, ExecLimit.TimeLimit)
}

func assertExecutionMemoryOK(t *testing.T, r *Result) {
	t.Helper()
	assert.NotZero(t, r.Exec.ExecMemory)
	assert.Less(t, r.Exec.ExecMemory, ExecLimit.MemoryLimit)
}

func assertExecutionTimeAndMemoryOK(t *testing.T, r *Result) {
	t.Helper()
	assertExecutionTimeOK(t, r)
	assertExecutionMemoryOK(t, r)
}

func assertExecutionTLE(t *testing.T, r *Result) {
	t.Helper()
	assert.Greater(t, r.Exec.ExecTime, ExecLimit.TimeLimit)
	assert.Less(t, r.Exec.ExecTime, ExecLimit.TimeLimit+TimeLimitMargin)
}

func assertExecutionMLE(t *testing.T, r *Result) {
	t.Helper()
	assert.Greater(t, r.Exec.ExecMemory, ExecLimit.MemoryLimit)
}
