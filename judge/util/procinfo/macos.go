//go:build darwin

package procinfo

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"

	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

// WARN: ps コマンドのRSSはあまり信用できないのであくまでも仮実装。
// linux バージョンの FetchRSS() は /proc/{pid}/status を見る実装で、しっくりくる結果が得られる
func FetchRSS(pid int) unit.ByteSize {
	cmd := exec.Command("ps", "-o", "rss=", "-p", strconv.Itoa(pid))
	bs, err := cmd.Output()
	if err != nil {
		log.Printf("Failed to spawn ps command: %v", err)
		return 0
	}

	s := string(bytes.Trim(bs, " \t\n\r"))
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("Failed to parse output from ps command: %q: %v", s, err)
		return 0
	}
	return unit.ByteSize(n) * unit.KiB
}
