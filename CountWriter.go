package kagus

import (
	"io"
)

// CountWriter wraps around an io.Writer and counts the amount of bytes written.
type CountWriter struct {
	W io.Writer
	N int64 // Total number of bytes written
}

func NewCountWriter(w io.Writer) *CountWriter {
	return &CountWriter{W: w}
}

func (cw *CountWriter) Write(p []byte) (int, error) {
	n, err := cw.W.Write(p)
	cw.N += int64(n)
	return n, err
}
