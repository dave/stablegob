# stablegob

Package stablegob is a fork of the Go standard library encoding/gob package at go version go1.10.1.

I've made some minor modifications that mean the binary output is deterministic. Maps are ordered by 
key. This will have a performance impact, beause the map keys are encoded and buffered before being sorted, 
then written to the output.  