package main

import "testing"

// fill type
var ret int

// Input is parsed into an array of strings in the main file to avoid any
// system calls or allocations from calling scanner.Text(). It's also much
// more convenient to write the benchmark code.

// go test -run=XXX -bench=. -benchmem -count=10 main*.go > benchmarks/official.txt

func BenchmarkPartX(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = partX()
	}
	ret = r
}
