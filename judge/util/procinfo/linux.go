//go:build linux

package procinfo

import (
	"bytes"
	"log"
	"os"
	"strconv"

	"github.com/szpp-dev-team/szpp-judge/judge/util/unit"
)

var vmRSSLiteral []byte = []byte("VmRSS:")

func FetchRSS(pid int) unit.ByteSize {
	fpath := "/proc/" + strconv.Itoa(pid) + "/status"

	f, err := os.Open(fpath)
	if err != nil {
		log.Printf("Cannot open %s: %v", fpath, err)
		return 0
	}

	buf := [1024]byte{}
	_, err = f.Read(buf[:])
	if err != nil {
		log.Printf("Cannot read content of %s: %v", fpath, err)
		return 0
	}

	// "VMRSS:" で始まる箇所を見つけ、さらに後続の空白をスキップ
	p := bytes.Index(buf[:], vmRSSLiteral) + len(vmRSSLiteral)
	for buf[p] == ' ' {
		p++
	}

	// 十進数としてパース
	val := 0
	for '0' <= buf[p] && buf[p] <= '9' {
		val = val*10 + int(buf[p]-'0')
		p++
	}

	return unit.ByteSize(val) * unit.KiB
}
