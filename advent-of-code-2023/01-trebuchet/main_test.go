package main

import "testing"

var ret = 0

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

// These each have to do an allocation because the contents of the map are
// strings. Everything is allocated at once and thus there is only one
// allocation.
func BenchmarkPart2WithSubstr(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = part2WithSubstr()
	}
	ret = r
}

// Substr from the standard library is a much faster algorithm.
func BenchmarkPart2NoSubstr(b *testing.B) {
	r := 0
	for n := 0; n < b.N; n++ {
		r = part2NoSubstr()
	}
	ret = r
}
