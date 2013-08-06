package kagus

import (
	"io"
)

// ReadByte reads a single byte from a reader
func ReadByte(r io.Reader) (byte, error) {
	buf := make([]byte, 1)

	for {
		n, err := r.Read(buf)
		if n == 1 {
			break
		}

		if err != nil {
			return 0, err
		}
	}

	return buf[0], nil
}
