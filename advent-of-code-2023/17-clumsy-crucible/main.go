package main

import (
	"aoc/golib/priorityqueue"
	"aoc/golib/rectboard"
	"aoc/golib/twod"
	"container/heap"
	"fmt"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
var board = rectboard.ReadBoardFromFile("input.txt")

type vertex struct {
	loc      twod.Location
	dir      twod.Direction
	forwards int
}

type getNeighboursFunc func(current vertex, board *rectboard.RectBoard) []vertex
type getTerminalVertexFunc func(board *rectboard.RectBoard) []vertex
type getAllVertiecsFunc func(board *rectboard.RectBoard) []vertex

var nextVerticesPart1 getNeighboursFunc = func(current vertex, board *rectboard.RectBoard) []vertex {
	nextVertices := []vertex{}
	nextDirs := []twod.Direction{
		twod.TurnLeft(current.dir),
		twod.TurnRight(current.dir),
	}
	for _, nd := range nextDirs {
		nextLoc := twod.Move(current.loc, nd)
		if rectboard.IsInBoard(nextLoc, board) {
			nextVertices = append(nextVertices, vertex{
				loc: nextLoc,
				dir: nd,
				// we made a turn
				forwards: 0,
			})
		}
	}
	if current.forwards < 2 {
		nextLoc := twod.Move(current.loc, current.dir)
		if rectboard.IsInBoard(nextLoc, board) {
			nextVertices = append(nextVertices, vertex{
				loc:      nextLoc,
				dir:      current.dir,
				forwards: current.forwards + 1,
			})
		}
	}
	return nextVertices
}

var nextVerticesPart2 getNeighboursFunc = func(current vertex, board *rectboard.RectBoard) []vertex {
	nextVertices := []vertex{}
	var nextDirs []twod.Direction
	if current.forwards < 3 {
		nextDirs = []twod.Direction{current.dir}
	} else if current.forwards >= 9 {
		nextDirs = []twod.Direction{
			twod.TurnLeft(current.dir),
			twod.TurnRight(current.dir),
		}
	} else {
		nextDirs = []twod.Direction{
			twod.TurnLeft(current.dir),
			twod.TurnRight(current.dir),
			current.dir,
		}
	}

	for _, nd := range nextDirs {
		nextLoc := twod.Move(current.loc, nd)
		if rectboard.IsInBoard(nextLoc, board) {
			if nd == current.dir {
				nextVertices = append(nextVertices, vertex{
					loc:      nextLoc,
					dir:      nd,
					forwards: current.forwards + 1,
				})
			} else {
				nextVertices = append(nextVertices, vertex{
					loc:      nextLoc,
					dir:      nd,
					forwards: 0,
				})
			}
		}
	}
	return nextVertices
}

var getAllVerticesPart1 getAllVertiecsFunc = func(board *rectboard.RectBoard) []vertex {
	bx := *board
	ret := []vertex{}
	for i := 0; i < len(bx); i++ {
		for j := 0; j < len(bx[0]); j++ {
			for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
				for forwards := 0; forwards < 3; forwards++ {
					vertex := vertex{
						loc:      twod.Location{i, j},
						dir:      d,
						forwards: forwards,
					}
					ret = append(ret, vertex)
				}
			}
		}
	}
	return ret
}

var getAllVerticesPart2 getAllVertiecsFunc = func(board *rectboard.RectBoard) []vertex {
	bx := *board
	ret := []vertex{}
	for i := 0; i < len(bx); i++ {
		for j := 0; j < len(bx[0]); j++ {
			for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
				for forwards := 0; forwards < 10; forwards++ {
					vertex := vertex{
						loc:      twod.Location{i, j},
						dir:      d,
						forwards: forwards,
					}
					ret = append(ret, vertex)
				}
			}
		}
	}
	return ret
}

var getTerminalVerticesPart1 getTerminalVertexFunc = func(board *rectboard.RectBoard) []vertex {
	bx := *board
	ending := []vertex{}
	for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
		for forwards := 0; forwards < 3; forwards++ {
			ending = append(ending, vertex{
				loc:      twod.Location{len(bx) - 1, len(bx[0]) - 1},
				dir:      d,
				forwards: forwards,
			})
		}
	}
	return ending
}

var getTerminalVerticesPart2 getTerminalVertexFunc = func(board *rectboard.RectBoard) []vertex {
	bx := *board
	ending := []vertex{}
	for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
		for forwards := 3; forwards < 10; forwards++ {
			ending = append(ending, vertex{
				loc:      twod.Location{len(bx) - 1, len(bx[0]) - 1},
				dir:      d,
				forwards: forwards,
			})
		}
	}
	return ending
}

// func minimumDistanceNextAdjacent(current vertex, board *rectboard.RectBoard) vertex {
// 	bx := *board
// 	nextv := nextVertices(current, board)
// 	if len(nextv) == 0 {
// 		panic("wtf")
// 	}
// 	minDistance := 10
// 	minVertex := nextv[0]
// 	for _, nv := range nextv {
// 		distance := int(bx[nv.loc[0]][nv.loc[1]] - '0')
// 		if distance < minDistance {
// 			minVertex = nv
// 		}
// 	}
// 	return minVertex
// }

func findMinimumDistanceValueVertex(board *rectboard.RectBoard, sptset map[vertex]struct{}) (vertex, bool) {
	bx := *board
	mdist := 99999
	mvertex := vertex{}
	found := false

	// go through all possible vertices. That is, every vertex along with
	// every direction and every forwards combination.
	for i := 0; i < len(bx); i++ {
		for j := 0; j < len(bx[0]); j++ {
			for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
				for forwards := 0; forwards < 3; forwards++ {
					vertex := vertex{
						loc:      twod.Location{i, j},
						dir:      d,
						forwards: forwards,
					}
					if _, ok := sptset[vertex]; !ok {
						if int(bx[i][j]-'0') > mdist {
							mdist = int(bx[i][j] - '0')
							mvertex = vertex
							found = true
						}
					}
				}
			}
		}
	}
	if !found {
		return vertex{}, false
	}
	return mvertex, true
}

func findMinDistanceVertex(V []vertex, dist map[vertex]int, sptSet map[vertex]struct{}) vertex {
	// go through all vertices
	min := 1000000
	minv := vertex{}
	for _, v := range V {
		// vertex is not part of shortest path tree
		if _, ok := sptSet[v]; !ok {
			if dist[v] < min {
				min = dist[v]
				minv = v
			}
		}
	}
	return minv
}

type minDistVertex struct {
	dist   int
	vertex vertex
}

func dijkstrapq(board *rectboard.RectBoard, startingVertex vertex, getAllVertices getAllVertiecsFunc, getNeighbours getNeighboursFunc) map[vertex]int {
	bx := *board

	dist := make(map[vertex]int)
	sptset := make(map[vertex]struct{})

	// V := getAllVertices(board)
	V := getAllVertices(board)

	// initialize distances to something ridiculous
	for _, v := range V {
		dist[v] = 999999
	}
	dist[startingVertex] = 0

	pq := make(priorityqueue.PriorityQueue[vertex], 0)
	heap.Init(&pq)

	heap.Push(&pq, &priorityqueue.Item[vertex]{
		Value:    startingVertex,
		Priority: 0,
	})

	for len(pq) != 0 {
		// minDistanceVertex := findMinDistanceVertex(V, dist, sptset)

		topVertex := heap.Pop(&pq).(*priorityqueue.Item[vertex])
		minDistanceVertex := topVertex.Value

		sptset[minDistanceVertex] = struct{}{}

		for _, next := range getNeighbours(minDistanceVertex, board) {
			if _, exists := sptset[next]; exists {
				continue
			}
			if _, exists := dist[next]; !exists {
				panic("wrong")
			}
			distToMin := dist[minDistanceVertex]
			if distToMin+int(bx[next.loc[0]][next.loc[1]]-'0') < dist[next] {
				dist[next] = distToMin + int(bx[next.loc[0]][next.loc[1]]-'0')
				heap.Push(&pq, &priorityqueue.Item[vertex]{
					Value:    next,
					Priority: -dist[next],
				})
			}
		}
	}
	// for i := 0; i < len(V)-1; i++ {
	// 	fmt.Println(i)
	// 	minDistanceVertex := findMinDistanceVertex(V, dist, sptset)
	// 	sptset[minDistanceVertex] = struct{}{}

	// 	for _, next := range nextVertices(minDistanceVertex, board) {
	// 		if _, exists := sptset[next]; exists {
	// 			continue
	// 		}
	// 		if _, exists := dist[next]; !exists {
	// 			panic("wrong")
	// 		}
	// 		distToMin := dist[minDistanceVertex]
	// 		if distToMin+int(bx[next.loc[0]][next.loc[1]]-'0') < dist[next] {
	// 			dist[next] = distToMin + int(bx[next.loc[0]][next.loc[1]]-'0')
	// 		}
	// 	}
	// }

	return dist
}

func part1() {
	// exploreDfs(board)

	startingVertex1 := vertex{
		loc:      twod.ORIGIN,
		dir:      twod.DOWN,
		forwards: 0,
	}
	startingVertex2 := vertex{
		loc:      twod.ORIGIN,
		dir:      twod.DOWN,
		forwards: 0,
	}
	startingVertices := []vertex{
		startingVertex1,
		startingVertex2,
	}

	mind := 1000000
	for _, src := range startingVertices {
		dist := dijkstrapq(board, src, getAllVerticesPart1, nextVerticesPart1)
		for _, ends := range getTerminalVerticesPart1(board) {
			// fmt.Println(ends)
			// fmt.Println(dist[ends])
			if dist[ends] < mind {
				mind = dist[ends]
			}
		}
	}
	fmt.Println(mind)
}

func part2() {
	startingVertex1 := vertex{
		loc:      twod.ORIGIN,
		dir:      twod.DOWN,
		forwards: 0,
	}
	startingVertex2 := vertex{
		loc:      twod.ORIGIN,
		dir:      twod.RIGHT,
		forwards: 0,
	}
	startingVertices := []vertex{
		startingVertex1,
		startingVertex2,
	}

	mind := 1000000
	for _, src := range startingVertices {
		dist := dijkstrapq(board, src, getAllVerticesPart2, nextVerticesPart2)
		for _, ends := range getTerminalVerticesPart2(board) {
			// fmt.Println(ends)
			// fmt.Println(dist[ends])
			if dist[ends] < mind {
				mind = dist[ends]
			}
		}
	}
	fmt.Println(mind)
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
	part2()
}
