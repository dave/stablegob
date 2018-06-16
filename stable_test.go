package stablegob

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"testing"
)

type A struct{}
type B struct{}
type C struct{}
type D struct{}
type E struct{}
type F struct{}
type G struct{}
type H struct{}
type I struct{}

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
		Debug(buf)
		return fmt.Sprintf("%x", sha.Sum(nil))
	}
	if run() != run() {
		t.Fatalf("hash different")
	}
}
