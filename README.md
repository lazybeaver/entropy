Description
===========
Compute Shannon Entropy for a byte stream

See: http://mathworld.wolfram.com/Entropy.html

Example
=======
The command line tool can be invoked without any arguments.
It will emit the bits-per-byte value of entropy before it exits.

	$ dd if=/dev/random bs=128 count=1 | ./main
	1+0 records in
	1+0 records out
	128 bytes transferred in 0.000020 secs (6391320 bytes/sec)
	6.506879566572159

Benchmarks
==========
No particular optimizations have been done. This is a simple implementation.

	$ go test -bench=.
	PASS
	BenchmarkShannon	   50000	     72996 ns/op
	ok		github.com/lazybeaver/entropy	4.398s
