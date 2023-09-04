package procinfo

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
)

var vmRSSLiteral []byte = []byte("VmRSS:")

func FetchRSSKibFromProcStatusFile(log *slog.Logger, pid int) uint {
	fpath := "/proc/" + strconv.Itoa(pid) + "/status"

	f, err := os.Open(fpath)
	if err != nil {
		log.Error("Cannot open file:", slog.String("path", fpath), slog.Any("err", err))
		return 0
	}

	buf := [1024]byte{}
	_, err = f.Read(buf[:])
	if err != nil {
		slog.Error("Cannot read content:", slog.String("path", fpath), slog.Any("err", err))
		return 0
	}

	// "VMRSS:" で始まる箇所を見つけ、さらに後続の空白をスキップ
	p := bytes.Index(buf[:], vmRSSLiteral) + len(vmRSSLiteral)
	for buf[p] == ' ' || buf[p] == '\t' {
		p++
	}

	// 十進数としてパース
	val := uint(0)
	for '0' <= buf[p] && buf[p] <= '9' {
		val = val*10 + uint(buf[p]-'0')
		p++
	}

	return val
}

func FetchRSSKibFromPsCommand(log *slog.Logger, pid int) uint {
	cmd := exec.Command("ps", "-o", "rss=", "-p", strconv.Itoa(pid))
	bs, err := cmd.Output()
	if err != nil {
		log.Error("Failed to spawn ps command:", "err", err)
		return 0
	}

	s := string(bytes.Trim(bs, " \t\n\r"))
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Error("Failed to parse output from ps command:", slog.String("output", s), slog.Any("err", err))
		return 0
	}
	return uint(n)
}
