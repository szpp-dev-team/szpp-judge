//go:build darwin

package procinfo

import (
	"log/slog"
)

// WARN: ps コマンドのRSSはあまり信用できないのであくまでも仮実装。
// linux バージョンの FetchRSS() は /proc/{pid}/status を見る実装で、しっくりくる結果が得られる
func FetchRSSKib(log *slog.Logger, pid int) uint {
	return FetchRSSKibFromPsCommand(log, pid)
}
