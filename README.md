# stablegob

Package stablegob is a fork of the Go standard library encoding/gob package at go version go1.10.1.

I've made some minor modifications that mean the binary output is deterministic. Maps are ordered by 
key. This will have a performance impact, because the map keys are encoded and buffered before being sorted, 
then written to the output.  

The binary output is 100% compatible with the standard library gob package.

# WARNING
My original changes had some issues which meant the output was not deterministic. The I've subsequently 
fixed the problems, but the number of changes has been significantly greater this time round, and will 
impact performance significantly more than previously. I wouldn't recommend using this package in it's 
current state. 