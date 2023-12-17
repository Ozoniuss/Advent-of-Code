package main

import (
	"aoc/golib/rectboard"
	"aoc/golib/twod"
	"container/heap"
	"fmt"
	"slices"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
var board = rectboard.ReadBoardFromFile("input.txt")

type state struct {
	loc      []twod.Location
	dir      []twod.Direction
	sum      int
	forwards int
}

type exploredStateKey struct {
	loc      twod.Location
	dir      twod.Direction
	forwards int
}

func findForwardsLeft(dir *[]twod.Direction) int {
	if len((*dir)) == 0 {
		return 0
	}
	remaining := 3
	last := (*dir)[len((*dir))-1]
	for i := len((*dir)) - 2; i >= 0; i-- {
		if i < 0 {
			return remaining
		}
		if (*dir)[i] != last {
			return remaining
		}
		remaining--
	}
	return remaining
}

// nextVertices returs the neighbours from a vertex.
func nextVertices(current exploredStateKey, board *rectboard.RectBoard) []exploredStateKey {
	nextVertices := []exploredStateKey{}
	nextDirs := []twod.Direction{
		twod.TurnLeft(current.dir),
		twod.TurnRight(current.dir),
	}
	for _, nd := range nextDirs {
		nextLoc := twod.Move(current.loc, nd)
		if rectboard.IsInBoard(nextLoc, board) {
			nextVertices = append(nextVertices, exploredStateKey{
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
			nextVertices = append(nextVertices, exploredStateKey{
				loc:      nextLoc,
				dir:      current.dir,
				forwards: current.forwards + 1,
			})
		}
	}
	return nextVertices
}

func nextVerticesPart2(current exploredStateKey, board *rectboard.RectBoard) []exploredStateKey {

	nextVertices := []exploredStateKey{}
	// nextDirs := []twod.Direction{}
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
				nextVertices = append(nextVertices, exploredStateKey{
					loc:      nextLoc,
					dir:      nd,
					forwards: current.forwards + 1,
				})
			} else {
				nextVertices = append(nextVertices, exploredStateKey{
					loc:      nextLoc,
					dir:      nd,
					forwards: 0,
				})
			}
		}
	}
	return nextVertices
}

func minimumDistanceNextAdjacent(current exploredStateKey, board *rectboard.RectBoard) exploredStateKey {
	bx := *board
	nextv := nextVertices(current, board)
	if len(nextv) == 0 {
		panic("wtf")
	}
	minDistance := 10
	minVertex := nextv[0]
	for _, nv := range nextv {
		distance := int(bx[nv.loc[0]][nv.loc[1]] - '0')
		if distance < minDistance {
			minVertex = nv
		}
	}
	return minVertex
}

func findMinimumDistanceValueVertex(board *rectboard.RectBoard, sptset map[exploredStateKey]struct{}) (exploredStateKey, bool) {
	bx := *board
	mdist := 99999
	mvertex := exploredStateKey{}
	found := false

	// go through all possible vertices. That is, every vertex along with
	// every direction and every forwards combination.
	for i := 0; i < len(bx); i++ {
		for j := 0; j < len(bx[0]); j++ {
			for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
				for forwards := 0; forwards < 3; forwards++ {
					vertex := exploredStateKey{
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
		return exploredStateKey{}, false
	}
	return mvertex, true
}

// getAllVertices returns all the vertices of the board
func getAllVertices(board *rectboard.RectBoard) []exploredStateKey {
	bx := *board
	ret := []exploredStateKey{}
	for i := 0; i < len(bx); i++ {
		for j := 0; j < len(bx[0]); j++ {
			for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
				for forwards := 0; forwards < 3; forwards++ {
					vertex := exploredStateKey{
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

func getAllVerticesPart2(board *rectboard.RectBoard) []exploredStateKey {
	bx := *board
	ret := []exploredStateKey{}
	for i := 0; i < len(bx); i++ {
		for j := 0; j < len(bx[0]); j++ {
			for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
				for forwards := 0; forwards < 10; forwards++ {
					vertex := exploredStateKey{
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

func findMinDistanceVertex(V []exploredStateKey, dist map[exploredStateKey]int, sptSet map[exploredStateKey]struct{}) exploredStateKey {
	// go through all vertices
	min := 1000000
	minv := exploredStateKey{}
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
	vertex exploredStateKey
}

func dijkstra(board *rectboard.RectBoard, startingVertex exploredStateKey) map[exploredStateKey]int {
	bx := *board

	dist := make(map[exploredStateKey]int)
	sptset := make(map[exploredStateKey]struct{})

	// V := getAllVertices(board)
	V := getAllVerticesPart2(board)

	// initialize distances to something ridiculous
	for _, v := range V {
		dist[v] = 999999
	}
	dist[startingVertex] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item{
		value:    startingVertex,
		priority: 0,
	})

	for len(pq) != 0 {
		// minDistanceVertex := findMinDistanceVertex(V, dist, sptset)

		topVertex := heap.Pop(&pq).(*Item)
		minDistanceVertex := topVertex.value

		sptset[minDistanceVertex] = struct{}{}

		for _, next := range nextVerticesPart2(minDistanceVertex, board) {
			if _, exists := sptset[next]; exists {
				continue
			}
			if _, exists := dist[next]; !exists {
				panic("wrong")
			}
			distToMin := dist[minDistanceVertex]
			if distToMin+int(bx[next.loc[0]][next.loc[1]]-'0') < dist[next] {
				dist[next] = distToMin + int(bx[next.loc[0]][next.loc[1]]-'0')
				heap.Push(&pq, &Item{
					value:    next,
					priority: -dist[next],
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

func getEndingStates(board *rectboard.RectBoard) []exploredStateKey {
	bx := *board
	ending := []exploredStateKey{}
	for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
		for forwards := 0; forwards < 3; forwards++ {
			ending = append(ending, exploredStateKey{
				loc:      twod.Location{len(bx) - 1, len(bx[0]) - 1},
				dir:      d,
				forwards: forwards,
			})
		}
	}
	return ending
}

func getEndingStatesPart2(board *rectboard.RectBoard) []exploredStateKey {
	bx := *board
	ending := []exploredStateKey{}
	for _, d := range []twod.Direction{twod.DOWN, twod.UP, twod.LEFT, twod.RIGHT} {
		for forwards := 3; forwards < 10; forwards++ {
			ending = append(ending, exploredStateKey{
				loc:      twod.Location{len(bx) - 1, len(bx[0]) - 1},
				dir:      d,
				forwards: forwards,
			})
		}
	}
	return ending
}

func part1() {
	// exploreDfs(board)

	startingVertex1 := exploredStateKey{
		loc:      twod.ORIGIN,
		dir:      twod.DOWN,
		forwards: 0,
	}
	startingVertex2 := exploredStateKey{
		loc:      twod.ORIGIN,
		dir:      twod.DOWN,
		forwards: 0,
	}
	startingVertices := []exploredStateKey{
		startingVertex1,
		startingVertex2,
	}

	mind := 1000000
	for _, src := range startingVertices {
		dist := dijkstra(board, src)
		for _, ends := range getEndingStates(board) {
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
	startingVertex1 := exploredStateKey{
		loc:      twod.ORIGIN,
		dir:      twod.DOWN,
		forwards: 0,
	}
	startingVertex2 := exploredStateKey{
		loc:      twod.ORIGIN,
		dir:      twod.RIGHT,
		forwards: 0,
	}
	startingVertices := []exploredStateKey{
		startingVertex1,
		startingVertex2,
	}

	mind := 1000000
	for _, src := range startingVertices {
		dist := dijkstra(board, src)
		for _, ends := range getEndingStatesPart2(board) {
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

	// part1()
	part2()
}

// Backtracking approaches that don't work

func explore(board *rectboard.RectBoard) {

	bx := *board
	visited := make(map[exploredStateKey]int)
	minsum := 3721

	q := make([]state, 0)

	// Don't add those to visited yet because you can revisit them.
	// Their heat doens't matetr
	q = append(q,
		state{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.RIGHT},
			sum:      0,
			forwards: 0,
		},
		state{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.DOWN},
			sum:      0,
			forwards: 0,
		},
	)

	for {
		fmt.Println(len(visited))
		if len(q) == 0 {
			fmt.Println("done")
			return
		}
		top := q[0]
		q = q[1:]

		if top.sum > minsum {
			continue
		}

		// fmt.Println(top.forwards)
		if top.loc[len(top.loc)-1] == (twod.Location{len(bx) - 1, len(bx[0]) - 1}) {

			if top.forwards < 3 {
				// fmt.Println(top.forwards)
				continue
			}

			if top.sum < minsum {
				minsum = top.sum
				fmt.Println(top.forwards)

				fmt.Println("minsum", minsum)
			}
			continue
		}

		currentKey := exploredStateKey{
			loc:      top.loc[len(top.loc)-1],
			dir:      top.dir[len(top.dir)-1],
			forwards: top.forwards,
		}

		// I can come from a different direction, allowing for multiple
		// exploration opportunities.
		// It's also not the same if I come from the same direction having
		// 2 more forwards of having no more forwards.
		if val, ok := visited[currentKey]; ok {
			// visited but the path is more costley
			if top.sum >= val {
				continue
			} else {
				visited[currentKey] = top.sum
			}
		} else {
			// not visited at all
			visited[currentKey] = top.sum
		}

		var nextDirs []twod.Direction
		if currentKey.forwards < 3 {
			nextDirs = []twod.Direction{top.dir[len(top.dir)-1]}
		} else if currentKey.forwards >= 9 {
			nextDirs = []twod.Direction{
				twod.TurnLeft(top.dir[len(top.dir)-1]),
				twod.TurnRight(top.dir[len(top.dir)-1]),
			}
		} else {
			nextDirs = []twod.Direction{
				twod.TurnLeft(top.dir[len(top.dir)-1]),
				twod.TurnRight(top.dir[len(top.dir)-1]),
				top.dir[len(top.dir)-1],
			}
		}
		for _, nd := range nextDirs {
			nextLoc := twod.Move(top.loc[len(top.loc)-1], nd)
			if rectboard.IsInBoard(nextLoc, board) {
				state := state{
					loc:      slices.Clone(append(top.loc, twod.Move(top.loc[len(top.loc)-1], nd))),
					dir:      slices.Clone(append(top.dir, nd)),
					sum:      top.sum + int(bx[nextLoc[0]][nextLoc[1]]-'0'),
					forwards: top.forwards,
				}
				// last direction didn't change
				if currentKey.dir == nd {
					state.forwards++
				} else {
					state.forwards = 0
				}
				q = append(q, state)
			}
		}

		// nextDirs := []twod.Direction{
		// 	twod.TurnLeft(top.dir[len(top.dir)-1]),
		// 	twod.TurnRight(top.dir[len(top.dir)-1]),
		// }
		// if len(top.loc) < 3 {
		// 	nextDirs = append(nextDirs, top.dir[len(top.dir)-1])
		// } else {
		// 	// previous two directions were the same
		// 	if !(top.dir[len(top.dir)-2] == top.dir[len(top.dir)-1] &&
		// 		top.dir[len(top.dir)-3] == top.dir[len(top.dir)-1]) {
		// 		nextDirs = append(nextDirs, top.dir[len(top.dir)-1])
		// 	}
		// }

	}
}

func exploreDfs(board *rectboard.RectBoard) {
	visited := make(map[exploredStateKey]int)
	minsum := 99999
	startingStates := []state{
		{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.RIGHT},
			sum:      0,
			forwards: 0,
		},
		{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.DOWN},
			sum:      0,
			forwards: 0,
		}}
	for _, state := range startingStates {
		dfs(board, state, visited, &minsum)
	}
}

func dfs(board *rectboard.RectBoard, current state, visited map[exploredStateKey]int, minsum *int) {
	bx := *board

	top := current
	if top.sum > *minsum {
		return
	}

	if top.loc[len(top.loc)-1] == (twod.Location{len(bx) - 1, len(bx[0]) - 1}) {

		if top.forwards < 3 {
			// fmt.Println(top.forwards)
			return
		}

		fmt.Println("reached the end")
		fmt.Println(top.sum)
		if top.sum < *minsum {
			*minsum = top.sum
		}
		return
	}

	currentKey := exploredStateKey{
		loc:      top.loc[len(top.loc)-1],
		dir:      top.dir[len(top.dir)-1],
		forwards: top.forwards,
	}

	if val, ok := visited[currentKey]; ok {
		// visited but the path is more costly
		if top.sum >= val {
			return
		} else {
			visited[currentKey] = top.sum
			if len(visited)%10000 == 0 {
				fmt.Println(len(visited))
			}
		}
	} else {
		// not visited at all

		visited[currentKey] = top.sum
		if len(visited)%10000 == 0 {
			fmt.Println(len(visited))
		}

	}

	var nextDirs []twod.Direction
	if currentKey.forwards < 3 {
		nextDirs = []twod.Direction{top.dir[len(top.dir)-1]}
	} else if currentKey.forwards >= 9 {
		nextDirs = []twod.Direction{
			twod.TurnLeft(top.dir[len(top.dir)-1]),
			twod.TurnRight(top.dir[len(top.dir)-1]),
		}
	} else {
		nextDirs = []twod.Direction{
			twod.TurnLeft(top.dir[len(top.dir)-1]),
			twod.TurnRight(top.dir[len(top.dir)-1]),
			top.dir[len(top.dir)-1],
		}
	}
	for _, nd := range nextDirs {
		nextLoc := twod.Move(top.loc[len(top.loc)-1], nd)
		if rectboard.IsInBoard(nextLoc, board) {
			state := state{
				loc:      slices.Clone(append(top.loc, twod.Move(top.loc[len(top.loc)-1], nd))),
				dir:      slices.Clone(append(top.dir, nd)),
				sum:      top.sum + int(bx[nextLoc[0]][nextLoc[1]]-'0'),
				forwards: top.forwards,
			}
			// last direction didn't change
			if currentKey.dir == nd {
				state.forwards++
			} else {
				state.forwards = 0
			}
			dfs(board, state, visited, minsum)
		}
	}

}
