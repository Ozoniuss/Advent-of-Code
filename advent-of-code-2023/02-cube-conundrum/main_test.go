package main

import "testing"

// fill type
var ret int

// Input is parsed into an array of strings in the main file to avoid any
// system calls or allocations from calling scanner.Text(). It's also much
// more convenient to write the benchmark code.

func BenchmarkPart1(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = part1()
	}
	ret = r
}

func BenchmarkPart2(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = part2()
	}
	ret = r
}
