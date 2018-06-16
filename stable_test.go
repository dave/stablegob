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
	t.Skip("TODO")
	Register(A{})
	Register(B{})
	Register(C{})
	Register(D{})
	Register(E{})
	Register(F{})
	Register(G{})
	Register(H{})
	Register(I{})
	s := map[interface{}]string{A{}: "a", B{}: "b", C{}: "c", D{}: "d", E{}: "e", F{}: "f", G{}: "g", H{}: "h", I{}: "i"}
	buf := &bytes.Buffer{}
	sha := sha1.New()
	w := io.MultiWriter(buf, sha)
	if err := NewEncoder(w).Encode(s); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprintf("%x", sha.Sum(nil)) != "50c9d1c7bfde3c4bdf1d511c416e068f20a92cc0" {
		t.Fatalf("hash not as expected, got %x", sha.Sum(nil))
	}

}
