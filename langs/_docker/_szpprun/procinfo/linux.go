//go:build linux

package procinfo

import (
	"log/slog"
)

func FetchRSSKib(log *slog.Logger, pid int) uint {
	return FetchRSSKibFromProcStatusFile(log, pid)
}
