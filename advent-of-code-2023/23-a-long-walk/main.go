package main

import (
	"aoc/golib/rectboard"
	"aoc/golib/twod"
	"fmt"
	"maps"
	"slices"
)

var board = rectboard.ReadBoardFromFile("input.txt")

// traverseBfs does a bfs while caching distances to other vertices.
func traverseBfs(s twod.Location, e twod.Location, board *rectboard.RectBoard) {
	bx := (*board)

	forcedDirs := make(map[byte]twod.Direction)
	forcedDirs['>'] = twod.RIGHT
	forcedDirs['<'] = twod.LEFT
	forcedDirs['^'] = twod.UP
	forcedDirs['v'] = twod.DOWN

	q := make([][]twod.Location, 0)
	q = append(q, []twod.Location{s})

	maxlen := 0

	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		prev := twod.Location{-1, -1}
		if len(top) >= 2 {
			prev = top[len(top)-2]
		}

		loc := top[len(top)-1]
		// fmt.Println(loc)

		nextDirs := []twod.Direction{twod.DOWN, twod.RIGHT, twod.UP, twod.LEFT}
		if nx, ok := forcedDirs[bx[loc[0]][loc[1]]]; ok {
			nextDirs = []twod.Direction{nx}
			// fmt.Println(nextDirs)
		}

		for _, nextDir := range nextDirs {
			next := twod.Move(loc, nextDir)
			// fmt.Println(next)

			if next == prev {
				continue
			}
			if slices.Contains(top, next) {
				continue
			}

			if next == e {
				fmt.Println("ending", len(top)+1, top)
				maxlen = len(top)
				continue
			}

			if !rectboard.IsInBoard(next, board) {
				continue
			}

			if (bx[next[0]][next[1]]) == '#' {
				continue
			}

			topc := slices.Clone(top)
			topc = append(topc, next)
			q = append(q, topc)
		}
	}

	fmt.Println("maxlen", maxlen)
}

func containsFromReverse(s []twod.Location, e twod.Location) bool {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == e {
			return true
		}
	}
	return false
}

func traverseBfsNoForcedDirsSlices(s twod.Location, e twod.Location, board *rectboard.RectBoard) {
	bx := (*board)
	visited := make(map[twod.Location]int)

	q := make([][]twod.Location, 0)
	q = append(q, []twod.Location{s})

	visited[s] = 0
	maxlen := 0

	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		prev := twod.Location{-1, -1}
		if len(top) >= 2 {
			prev = top[len(top)-2]
		}

		shouldGenerate := true

		loc := top[len(top)-1]
		if prevDist, ok := visited[loc]; !ok {
			visited[loc] = len(top) - 1
		} else {
			// The previous distance to loc is smaller than the current distance
			// we have to top.
			if prevDist < len(top)-1 {
				visited[loc] = len(top) - 1
			}
			goto LOOP
		}

		// check the past 10 distances. If found one with a shorter
		// distance to loc, it means we don't need it
		// shouldGenerate := true
		if len(top) > 10 {
			for i := len(top) - 1; i >= len(top)-10; i-- {
				// In this top array, the distance to loc is smaller than the
				// biggest distance we know to loc. It should not generate
				// any more solutions.
				loc := top[i]
				if i < visited[loc] {
					shouldGenerate = false
					break
				}
			}
		}

		if !shouldGenerate {
			continue
		}

	LOOP:
		nextDirs := []twod.Direction{twod.DOWN, twod.RIGHT, twod.UP, twod.LEFT}

		for _, nextDir := range nextDirs {
			next := twod.Move(loc, nextDir)
			// fmt.Println(next)

			if next == prev {
				continue
			}

			if containsFromReverse(top, next) {
				continue
			}

			if next == e {
				fmt.Println("ending", len(top)+1)
				maxlen = len(top)
				continue
			}

			if !rectboard.IsInBoard(next, board) {
				continue
			}

			if (bx[next[0]][next[1]]) == '#' {
				continue
			}

			topc := slices.Clone(top)
			topc = append(topc, next)
			q = append(q, topc)
		}
	}

	fmt.Println("maxlen", maxlen)
}

type qstate struct {
	current twod.Location
	// Dists are paths that were visited in this state that are NOT part of the
	// global visited.
	dists map[twod.Location]int
}

// traverseBfs does a bfs while caching distances to other vertices.
func traverseBfsNoForcedDirs(s twod.Location, e twod.Location, board *rectboard.RectBoard) {
	bx := (*board)
	visited := make(map[twod.Location]int, 4096)

	dists := make(map[twod.Location]int, 4096)
	dists[s] = 0

	// Have the initial states
	initialState := qstate{
		current: s,
		dists:   dists,
	}
	q := make([]qstate, 0)
	q = append(q, initialState)

	// global visited
	visited[s] = 0
	maxlen := 0

	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		loc := top.current
		// fmt.Println(loc)

		// if otherPathDist, ok := visited[loc]; !ok {
		// 	visited[loc] = top.dists[top.current]
		// } else {

		// }

		nextDirs := []twod.Direction{twod.DOWN, twod.RIGHT, twod.UP, twod.LEFT}
		for _, nextDir := range nextDirs {
			next := twod.Move(loc, nextDir)

			// This is invalid because it's in the current path.
			if _, hasVisited := top.dists[next]; hasVisited {
				continue
			}

			if next == e {
				fmt.Println("ending", len(top.dists)+1)
				maxlen = len(top.dists)
				continue
			}

			if !rectboard.IsInBoard(next, board) {
				continue
			}

			if (bx[next[0]][next[1]]) == '#' {
				continue
			}

			topdistsc := maps.Clone(top.dists)
			topdistsc[next] = topdistsc[top.current] + 1
			nextq := qstate{
				current: next,
				dists:   topdistsc,
			}
			q = append(q, nextq)
		}
	}

	fmt.Println("maxlen", maxlen)
}

type SimplifiedGraph struct {
	edges    map[[2]twod.Location]int
	vertices map[twod.Location]struct{}
}

type simplifiedState struct {
	path        []twod.Location
	fromNode    twod.Location
	currentDist int
}

// It's too fucking expensive to explore the whole thing. However, note that
// if there is a straight path from A to B, we don't really care about the
// full path. We only care about the distance.
//
// The idea is that there aren't really that many places where you can actually
// go in multiple directions. We should create a graph where those are the
// vertices and there are only edges between those. The idea is that an edge
// should be associated to the only path from A to B, and its cost is the
// length of that path.
//
// Then, we can BFS the fuck out of that graph.
func generateSimplifiedGraph(s twod.Location, e twod.Location, board *rectboard.RectBoard) SimplifiedGraph {
	bx := (*board)
	visited := make(map[twod.Location]struct{})
	g := SimplifiedGraph{
		edges: make(map[[2]twod.Location]int),
	}

	i := simplifiedState{
		path:        []twod.Location{s},
		currentDist: 0,
		fromNode:    s,
	}

	q := make([]simplifiedState, 0)
	q = append(q, i)

	for len(q) != 0 {
		state := q[0]
		q = q[1:]

		// There is a catch when generating this graph, see below. We can't
		// do it with by keeping track of visited vertices.

		loc := state.path[len(state.path)-1]
		prev := twod.Location{-69, -69}
		if len(state.path) >= 2 {
			prev = state.path[len(state.path)-2]
		}
		nextDirs := []twod.Direction{twod.DOWN, twod.RIGHT, twod.UP, twod.LEFT}

		availableDirs := []twod.Direction{}
		for _, nextDir := range nextDirs {
			next := twod.Move(loc, nextDir)

			// Not in board or bullshit, not available
			if !rectboard.IsInBoard(next, board) {
				continue
			}
			if (bx[next[0]][next[1]]) == '#' {
				continue
			}
			// we just went back
			if next == prev {
				continue
			}

			// edge case, got back to the original node
			if next == i.fromNode {
				continue
			}

			availableDirs = append(availableDirs, nextDir)
		}

		// Just go further down this path, absolutely nothing changes.
		if len(availableDirs) == 1 {
			nextState := simplifiedState{
				path:        slices.Clone(state.path),
				fromNode:    state.fromNode,
				currentDist: state.currentDist + 1,
			}
			nextState.path = append(nextState.path, twod.Move(loc, availableDirs[0]))
			q = append(q, nextState)
		} else {

			e1 := [2]twod.Location{state.fromNode, state.path[len(state.path)-1]}
			e2 := [2]twod.Location{state.path[len(state.path)-1], state.fromNode}
			_, hase1 := g.edges[e1]
			_, hase2 := g.edges[e2]

			// we generated an existing edge, we are retarded.
			//
			// The only way to generate this graph is by keeping track of the
			// edges we generated, not the places we visited. Consider this
			// scenario. There is an edge between X and Y. We got to X via A,
			// and now we go to Y. But, we also got to Y via B, and now we go
			// to X. These explorations will meet in between, and if we decide
			// to abort exploration when we encounter a visited map location,
			// we won't be adding XY.
			if hase1 || hase2 {
				continue
			}

			// Time to add a fucking edge.
			g.edges[[2]twod.Location{state.fromNode, state.path[len(state.path)-1]}] = state.currentDist

			for _, nx := range availableDirs {
				nextState := simplifiedState{
					path: slices.Clone(state.path),

					// This new path starts from the last node
					fromNode:    state.path[len(state.path)-1],
					currentDist: 1,
				}
				nextState.path = append(nextState.path, twod.Move(loc, nx))
				q = append(q, nextState)
			}
		}

		// after the shenanigans are done, mark this as visited
		visited[state.path[len(state.path)-1]] = struct{}{}
	}

	return g
}

func getNeighbours(v twod.Location, g *SimplifiedGraph) []twod.Location {
	n := make([]twod.Location, 0)
	for e := range g.edges {
		if e[0] == v {
			n = append(n, e[1])
		}
		if e[1] == v {
			n = append(n, e[0])
		}
	}
	return n
}

func calculatePathDistance(path []twod.Location, g *SimplifiedGraph) int {
	sum := 0
	for i := 1; i < len(path); i++ {
		e1 := [2]twod.Location{path[i-1], path[i]}
		e2 := [2]twod.Location{path[i], path[i-1]}

		if dist, ok := g.edges[e1]; ok {
			sum += dist
		} else if dist, ok := g.edges[e2]; ok {
			sum += dist
		} else {
			panic("sth went wrong")
		}
	}
	return sum
}

func traverseBfsSimplifiedGraph(s twod.Location, e twod.Location, g *SimplifiedGraph) {
	bx := (*board)

	q := make([][]twod.Location, 0)
	q = append(q, []twod.Location{s})

	maxlen := 0
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		prev := twod.Location{-1, -1}
		if len(top) >= 2 {
			prev = top[len(top)-2]
		}

		loc := top[len(top)-1]

		neighbours := getNeighbours(loc, g)
		for _, next := range neighbours {

			if next == prev {
				continue
			}
			if slices.Contains(top, next) {
				continue
			}

			if next == e {
				// fmt.Println("ending", top)
				topc := slices.Clone(top)
				len := calculatePathDistance(append(topc, next), g)
				if len > maxlen {
					maxlen = len
					fmt.Println("maxlen", maxlen)
				}
				continue
			}

			if !rectboard.IsInBoard(next, board) {
				continue
			}

			if (bx[next[0]][next[1]]) == '#' {
				continue
			}

			topc := slices.Clone(top)
			topc = append(topc, next)
			q = append(q, topc)
		}
	}

	fmt.Println("maxlen", maxlen)
}

func part1() {
	// rectboard.PrintBoard(board)
	bx := *board
	e := twod.Location{len(bx) - 1, len(bx[0]) - 2}

	traverseBfs(twod.Location{0, 1}, e, board)
}

func part2() {
	// rectboard.PrintBoard(board)
	bx := *board
	s := twod.Location{0, 1}
	e := twod.Location{len(bx) - 1, len(bx[0]) - 2}

	g := generateSimplifiedGraph(s, e, board)
	fmt.Println(g)
	traverseBfsSimplifiedGraph(s, e, &g)
	// traverseBfsNoForcedDirsSlices(twod.Location{0, 1}, e, board)
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

	// part1()
	part2()
}
