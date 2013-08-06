package kagus

import (
	"errors"
	"io"
)

var (
	WriteLimitReached = errors.New("Write limit reached")
)

// LimitWriter wraps around a io.Writer and limits the amount of bytes that can be written to it.
type LimitedWriter struct {
	W io.Writer
	N int64
}

// Write writes at most lw.N bytes to LimitWriter. It updates lw.N on sucessful write operations.
// If the limit was reached, the error WriteLimitReached will be returned.
func (lw *LimitedWriter) Write(p []byte) (int, error) {
	towrite := int64(len(p))
	if lw.N < towrite {
		towrite = lw.N
	}

	n, err := lw.W.Write(p[:towrite])
	if err != nil {
		lw.N = 0
		return n, err
	}

	if n == len(p) {
		lw.N -= int64(n)
		return n, nil
	}

	lw.N = 0
	return n, WriteLimitReached
}
