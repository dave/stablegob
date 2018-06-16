package stablegob

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

type A struct{ I int }
type B struct{ I int }
type C struct{ I int }
type D struct{ I int }
type E struct{ I int }
type F struct{ I int }
type G struct{ I int }
type H struct{ I int }
type I struct{ I int }

func TestStable(t *testing.T) {
	Register(A{})
	Register(B{})
	Register(C{})
	Register(D{})
	Register(E{})
	Register(F{})
	Register(G{})
	Register(H{})
	Register(I{})
	run := func() string {
		s := map[interface{}]string{A{}: "a", B{}: "b", C{}: "c", D{}: "d", E{}: "e", F{}: "f", G{}: "g", H{}: "h", I{}: "i"}
		buf := &bytes.Buffer{}
		sha := sha1.New()
		w := io.MultiWriter(buf, sha)
		if err := NewEncoder(w).Encode(s); err != nil {
			t.Fatal(err)
		}
		return fmt.Sprintf("%x", sha.Sum(nil))
	}
	if run() != run() {
		t.Fatalf("hash different")
	}
}
