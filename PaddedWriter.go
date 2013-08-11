package kagus

import (
	"io"
)

// PaddedWriter is a writer for data that needs to be padded / aligned.
type PaddedWriter struct {
	cw      *CountWriter
	padding int
}

func NewPaddedWriter(w io.Writer, padding int) *PaddedWriter {
	return &PaddedWriter{NewCountWriter(w), padding}
}

func (pw *PaddedWriter) Write(p []byte) (int, error) {
	return pw.cw.Write(p)
}

// Pad needs to be called if the padding should be inserted. Data is padded with zero-bytes.
func (pw *PaddedWriter) Pad() error {
	mod := int(pw.cw.N % int64(pw.padding))
	if mod == 0 {
		return nil
	}

	paddingLen := pw.padding - mod
	padding := make([]byte, paddingLen)
	_, err := pw.cw.Write(padding)
	return err
}
