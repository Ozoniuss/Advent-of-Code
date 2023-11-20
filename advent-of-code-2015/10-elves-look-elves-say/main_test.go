package main

import "testing"

var out string

func BenchmarkLookAndSay(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 1; i < 30; i++ {
			val := lookAndSay("1321131112")
			out = val
		}
	}
}

func BenchmarkLookAndSayWithoutBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 1; i < 30; i++ {
			val := lookAndSayWithoutBuilder("1321131112")
			out = val
		}
	}
}
