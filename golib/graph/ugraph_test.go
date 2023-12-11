package graph

import (
	"aoc/golib/graph/bytetype"
	"aoc/golib/graph/stringtype"
	"testing"
)

var ok string
var mymap map[string]struct{}

func BenchmarkReadFromFile1kStringType(b *testing.B) {
	g := stringtype.NewUnorderedGraph()
	for n := 0; n <= b.N; n++ {
		g.ReadFromFile("graph1k.txt")
		for val := range g.GetVerticesMap() {
			ok = val
		}
	}
	mymap = g.GetVerticesMap()
}
func BenchmarkReadFromFileWithCapacity1kStringType(b *testing.B) {
	g := stringtype.NewUnorderedGraphWithCapacity(1000, 4000)
	for n := 0; n <= b.N; n++ {
		g.ReadFromFile("graph1k.txt")
		for val := range g.GetVerticesMap() {
			ok = val
		}
	}
	mymap = g.GetVerticesMap()
}

func BenchmarkReadFromFile1kByteType(b *testing.B) {
	g := bytetype.NewUnorderedGraph[[4]byte]()
	for n := 0; n <= b.N; n++ {
		g.ReadFromFile("graph1k.txt")
	}
}
func BenchmarkReadFromFileWithCapacity1kByteType(b *testing.B) {
	g := bytetype.NewUnorderedGraphWithCapacity[[4]byte](1000, 4000)
	for n := 0; n <= b.N; n++ {
		g.ReadFromFile("graph1k.txt")
	}
}
