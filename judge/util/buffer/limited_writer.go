package buffer

// Writer that stores string at most N bytes
type LimitedWriter struct {
	N        int
	Buf      []byte
	overflow bool
}

// n バイトまでしか書きこめない Buffer。
// n バイト目以降の書き込みはエラーにならず無視される。
// 引数 `buf` は nil でも OK。
func NewLimitedWriter(n int, buf []byte) LimitedWriter {
	return LimitedWriter{
		N:   n,
		Buf: buf,
	}
}

func (w *LimitedWriter) Write(b []byte) (n int, err error) {
	avail := w.N - len(w.Buf)

	add := len(b)
	if add > avail {
		add = avail
		w.overflow = true
	}
	w.Buf = append(w.Buf, b[:add]...)
	return len(b), nil
}

func (w *LimitedWriter) HasOverflowed() bool {
	return w.overflow
}
