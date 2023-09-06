package sandboxtest

import (
	"time"

	"github.com/szpp-dev-team/szpp-judge/judge/sandbox"
	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

var (
	CompileLimit = sandbox.SzpprunOption{
		TimeLimit:        time.Millisecond * 500,
		MemoryLimit:      unit.MiB * 256,
		FileWriteLimit:   unit.MiB * 256,
		NumOpenFileLimit: 256,
	}
	ExecLimit = sandbox.SzpprunOption{
		TimeLimit:        time.Millisecond * 300,
		MemoryLimit:      unit.MiB * 64,
		FileWriteLimit:   unit.MiB * 2,
		NumOpenFileLimit: 32,
	}
)

const (
	ContainerProcNumLimit = 20
	ContainerMemoryLimit  = unit.MiB * 256
	// コンパイラの stdout と stderr をマージした結果の読み取り上限
	CompileMessageReadLimit = unit.KiB * 8
	// 実行時の stdout の読み取り上限
	ExecStdoutReadLimit = unit.KiB * 4
	// 実行時の stderr の読み取り上限 (テストのため ExecStdoutReadLimit と異なる値にすること)
	ExecStderrReadLimit = unit.KiB * 8
	// 時間制限に加算する猶予 (実際の値は szpprun のソースコードで定義されている)
	TimeLimitMargin = time.Millisecond * 105
)
