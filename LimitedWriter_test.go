package kagus

import (
	"bytes"
	"testing"
)

func TestLimitedWriting(t *testing.T) {
	buf := new(bytes.Buffer)
	lw := &LimitedWriter{buf, 10}

	if _, err := lw.Write([]byte{1, 2}); err != nil {
		t.Fatalf("lw.Write failed: %s", err)
	}

	if lw.N != 8 {
		t.Fatalf("lw.N was supposed to be 8, is %d", lw.N)
	}

	if !bytes.Equal(buf.Bytes(), []byte{1, 2}) {
		t.Errorf("Unexpected content of buf: %v", buf.Bytes())
	}

	n, err := lw.Write([]byte{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
	if (n != 8) || (err != WriteLimitReached) {
		t.Fatalf("lw.Write returned with unexpected n, err: %d, %s", n, err)
	}

	n, err = lw.Write([]byte("trololo"))
	if (n != 0) || (err != WriteLimitReached) {
		t.Fatalf("lw.Write returned with unexpected n, err: %d, %s", n, err)
	}
}
