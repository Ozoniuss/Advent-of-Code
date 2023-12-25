package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	dbgraph "github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
var inputLines = readlines()

func readlines() []string {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

type Graph struct {
	vertices   []string
	neighbours map[string][]string
	edges      map[[2]string]struct{}
}

func (g *Graph) hasVertex(v string) bool {
	return slices.Contains(g.vertices, v)
}

func (g *Graph) hasEdge(s, e string) bool {
	_, ok1 := g.edges[[2]string{s, e}]
	_, ok2 := g.edges[[2]string{e, s}]
	return ok1 || ok2
}

func addIfNotExists(s string, sl *[]string) {
	if slices.Contains(*sl, s) {
		return
	}
	(*sl) = append((*sl), s)
}

func readLine(line string, g *Graph, g2 dbgraph.Graph[string, string]) {
	parts := strings.Split(line, ": ")
	source := parts[0]

	if !g.hasVertex(source) {
		g.vertices = append(g.vertices, source)
		g.neighbours[source] = make([]string, 0)
	}

	neighbours := strings.Split(parts[1], " ")
	for _, n := range neighbours {
		if !g.hasEdge(source, n) {
			g.edges[[2]string{source, n}] = struct{}{}
		}

		if !g.hasVertex(n) {
			g.vertices = append(g.vertices, n)
			g.neighbours[n] = make([]string, 0)
		}

		neighs := g.neighbours[source]
		addIfNotExists(n, &neighs)
		g.neighbours[source] = neighs

		neighn := g.neighbours[n]
		addIfNotExists(source, &neighn)
		g.neighbours[n] = neighn

		_ = g2.AddVertex(source)
		_ = g2.AddVertex(n)
		_ = g2.AddEdge(n, source)
	}
}

func first3Neighbours(s string, g *Graph) []string {
	all := []string{}
	q := []string{s}

	visited := make(map[string]struct{})
	for i := 0; i < 3; i++ {
		nextq := []string{}
		// get all neighbours of all elements in the queue
		for _, el := range q {
			visited[el] = struct{}{}
			ns := g.neighbours[el]
			for _, n := range ns {
				addIfNotExists(n, &all)
				if _, ok := visited[n]; !ok {
					nextq = append(nextq, n)
				}
			}
		}
		q = nextq
	}
	return all
}

func part1() {
	g := Graph{
		vertices:   make([]string, 0),
		edges:      map[[2]string]struct{}{},
		neighbours: make(map[string][]string),
	}
	g2 := dbgraph.New(dbgraph.StringHash, dbgraph.Directed())
	for _, line := range inputLines {
		readLine(line, &g, g2)
	}
	// fmt.Println(g
	fmt.Println(len(g.edges))
	fmt.Println(len(g.vertices))
	return

	fmt.Println(g.neighbours["zqv"])
	fmt.Println(first3Neighbours("zqv", &g))

	for _, v := range g.vertices {
		fmt.Println(v, first3Neighbours(v, &g))
	}

	file, _ := os.Create("./mygraph.gv")
	_ = draw.DOT(g2, file)
}

func main() {
	// Run only 1 profile at a time!
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()

	// No return value improves speed during contest, but make these functions
	// return something for benchmarks once problem is solved.
	//
	// Part 2 is not written above and commented below so that template compiles
	// while solving part 1.

	part1()
	// part2()
}
