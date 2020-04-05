package prefixsuffix

import (
	"io"
	"testing"
)

func TestSaver(t *testing.T) {
	s := &Saver{N: 4}
	io.WriteString(s, "1234 ---- 5678")
	if string(s.Bytes()) != "1234\n... omitting 6 bytes ...\n5678" {
		t.Fatalf("got %q", s.Bytes())
	}
}
