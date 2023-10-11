package sandbox

import (
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

// szpprun が生成する .szpprun-result.txt を読み取る。
// 実行時間、メモリ使用量、exitCode を返す
func parseSzpprunResultFile(dir string) (time.Duration, unit.Byte, int, error) {
	const NAME = ".szpprun-result.txt"
	fpath := path.Join(dir, NAME)
	f, err := os.Open(fpath)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to open %q: %w", fpath, err)
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to read content of %q: %w", fpath, err)
	}

	// .szpprun-result.txt は実行時間、メモリ使用量、exitCodeがこの順に空白区切りで書きこまれている
	// 実行時間、メモリ使用量は負数で記述されていることはないが、exitCode は負数で保存されている可能性があるので注意
	// 末尾の改行はない

	i := 0
	t := uint(0)
	for '0' <= bs[i] && bs[i] <= '9' {
		t = t*10 + uint(bs[i]-'0')
		i++
	}

	i++ // skip whitespace
	mem := uint(0)
	for '0' <= bs[i] && bs[i] <= '9' {
		mem = mem*10 + uint(bs[i]-'0')
		i++
	}

	i++ // skip whitespace
	rest := string(bs[i:])
	exitCode, err := strconv.Atoi(rest)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to parse as exitCode integer: str=%q: %w", rest, err)
	}

	return time.Duration(t) * time.Millisecond, unit.Byte(mem) * unit.KiB, exitCode, nil
}
