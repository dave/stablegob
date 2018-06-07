package stablegob

import (
	"bytes"
	"fmt"
	"testing"
)

func TestStable(t *testing.T) {
	s := struct{ S map[string]int }{
		S: map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9},
	}
	var last string
	for i := 0; i < 100; i++ {
		buf := &bytes.Buffer{}
		if err := NewEncoder(buf).Encode(s); err != nil {
			t.Fatal(err)
		}
		this := fmt.Sprintf("%x", buf.Bytes())
		if i > 0 && this != last {
			t.Fatalf("\nlast: %s, \nthis: %s", last, this)
			return
		}
		last = this
	}
}
