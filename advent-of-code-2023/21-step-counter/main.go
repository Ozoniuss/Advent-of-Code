package main

import (
	"aoc/golib/rectboard"
	"aoc/golib/twod"
	"fmt"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
// var inputLines = readlines()

// func readlines() []string {
// 	f, err := os.Open("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

//		var lines []string
//		s := bufio.NewScanner(f)
//		for s.Scan() {
//			lines = append(lines, s.Text())
//		}
//		return lines
//	}
var board = rectboard.ReadBoardFromFile("input.txt")

func findStartingPosition(board *rectboard.RectBoard) twod.Location {
	bx := *board
	for i := 0; i < len(bx); i++ {
		for j := 0; j < len(bx[0]); j++ {
			if bx[i][j] == 'S' {
				return twod.Location{i, j}
			}
		}
	}
	panic("lol")
}

// stores how many points each location generates
var cacheGenerates = make(map[twod.Location]int)

func computeBfs(s twod.Location, board *rectboard.RectBoard, steps int) {

	odds := []int{}
	evens := []int{}

	bx := *board
	count := 0
	top := []twod.Location{s}
	for i := 0; i < steps; i++ {

		// the places reached this round to avoid duplicates
		visitedRound := make(map[twod.Location]struct{})
		nextq := make([]twod.Location, 0, 64)
		count = 0
		for _, pos := range top {
			for _, step := range []twod.Direction{twod.UP, twod.LEFT, twod.DOWN, twod.RIGHT} {
				nextPos := twod.Move(pos, step)
				// we got here already this round
				if _, ok := visitedRound[nextPos]; ok {
					continue
				}
				if !rectboard.IsInBoard(nextPos, board) {
					continue
				}
				// only use adjusted for in board criteria
				if bx[nextPos[0]][nextPos[1]] != '#' {
					nextq = append(nextq, nextPos)
					visitedRound[nextPos] = struct{}{}
					count++
				}
			}
		}

		if i%2 == 0 {
			odds = append(odds, count)
		} else {
			evens = append(evens, count)
		}

		// q = append(q, nextq)
		top = nextq
		fmt.Println(len(top))
	}
	fmt.Println("odds", odds)
	fmt.Println("evens", evens)
	// fmt.Println(len(top))
	fmt.Println(count)
}

func computeBfsInfiniteDumb(s twod.Location, board *rectboard.RectBoard, steps int) {

	nums := []int{}
	bx := *board
	count := 0
	top := []twod.Location{s}
	for i := 0; i < steps; i++ {

		fmt.Println("step", i)

		// the places reached this round to avoid duplicates
		visitedRound := make(map[twod.Location]struct{})
		nextq := make([]twod.Location, 0, 64)
		count = 0
		for _, pos := range top {
			for _, step := range []twod.Direction{twod.UP, twod.LEFT, twod.DOWN, twod.RIGHT} {
				nextPos := twod.Move(pos, step)
				// we got here already this round
				if _, ok := visitedRound[nextPos]; ok {
					continue
				}
				var nextPosAdjusted twod.Location
				nextPosAdjusted[0] = mod(nextPos[0], len(bx))
				nextPosAdjusted[1] = mod(nextPos[1], len(bx[0]))
				// only use adjusted for in board criteria
				if bx[nextPosAdjusted[0]][nextPosAdjusted[1]] != '#' {
					nextq = append(nextq, nextPos)
					visitedRound[nextPos] = struct{}{}
					count++
				}
			}
		}
		nums = append(nums, count)
		// q = append(q, nextq)
		top = nextq
		// fmt.Println(len(top))
	}
	// fmt.Println(len(top))
	fmt.Println(nums)
	fmt.Println(count)
}

// Note: some experimenting shows that this board becomes cyclic once it
// reaches 7796 steps. And the sequence is 7796, 7819
//
// For the input board the sequence is 39, 42, once it reaches 39.
//
// Note that while I believe it could be possible to reach one of the recurring
// sequence's size (e.g. 39 or 42) in a round without actually starting the
// recurring sequence, it's not the case for this problem and since I'm too lazy
// to generalize I will use these somehow.
//
// I also think it's possible that some cycles could be longer but whatever man

// hasreachedend keeps track of the "board" that reached the end.
// The board index of a position is x / row, y / col
//
// However, other board that start being filled at the same time from the sides
// could have a different cyclic pattern. Oh well...
// So let's just cache the boards.
//
// Hopefully the cache length is equal to 2...
//
// We can naively assume that the sequence will repeat itself to the same number
// of steps. Just by looking at the board it seems the starting point wouldn't
// really matter. We can print the sequences nonetheless

//
// Also it's important to know that 2 cached boards will not interfere with
// each other. If one goes to the other both will generate
//
//  O    OXO
//  O    OXO

func addVal(current [4]int, val int) [4]int {
	var next [4]int
	next[0] = current[1]
	next[1] = current[2]
	next[2] = current[3]
	next[3] = val
	return next
}

func isCyclic(vals [4]int) bool {
	return (vals[0] == vals[2]) && (vals[1] == vals[3])
}

type cacheVal struct {
	// latest values. If we have 4 consecutive values that repeat
	// a cycle
	val [4]int
}

var cacheRoundVals = make(map[[2]int]cacheVal)

// on the left value for even steps, on the right value for odd steps.
var cyclicCache = make(map[[2]int][2]int)

func computeBfsInfinite(s twod.Location, board *rectboard.RectBoard, steps int, startc, endc int) {

	bx := *board
	count := 0
	top := []twod.Location{s}
	for i := 0; i < steps; i++ {
		// fmt.Println(i)

		// the places reached this round to avoid duplicates
		visitedRound := make(map[twod.Location]struct{})
		nextq := make([]twod.Location, 0, 64)
		count = 0

		// need to figure out the total for each board
		var currentRoundTotal = make(map[[2]int]int)
		// fmt.Println(cyclicCache)
		// fmt.Println(len(top))

		// when the map had become cyclic, all the outer positions were
		// generated. So we can ignore every position from that map,
		// including those that are not cyclic, because we have the other
		// positions as well.
		//
		// we should no longer add positions to cached maps to generators.

		// to add the cache only once, for all the cached maps, add the
		// steps.
		// fmt.Println(cyclicCache)
		for _, v := range cyclicCache {
			count += v[steps%2]
		}
		// fmt.Println(count)

		for _, pos := range top {

			boardOfTop := computeBoardOfLocation(pos, board)
			// this position takes part of something that has been cached,
			// and by the time it had been cached it already generated
			// everything it could so ignore it.
			if _, ok := cyclicCache[boardOfTop]; ok {
				continue
			}

			// These come from a top that does not take part of a cached board

			for _, step := range []twod.Direction{twod.UP, twod.LEFT, twod.DOWN, twod.RIGHT} {
				nextPos := twod.Move(pos, step)

				// the "board" in which next is
				boardOfNext := computeBoardOfLocation(nextPos, board)
				// This next takes part of a board that had been cached. There
				// is no need to keep track of it, since it's generated also
				// by the cached board
				if _, ok := cyclicCache[boardOfNext]; ok {
					continue
				}

				// // for this board we're already cyclic

				// this board had not been encountered before, initialize
				if _, ok := cacheRoundVals[boardOfNext]; !ok {
					cacheRoundVals[boardOfNext] = cacheVal{}
				}

				var nextPosAdjusted twod.Location
				nextPosAdjusted[0] = mod(nextPos[0], len(bx))
				nextPosAdjusted[1] = mod(nextPos[1], len(bx[0]))
				// we got here already this round
				if _, ok := visitedRound[nextPos]; ok {
					continue
				}
				// count those that do not belong to a generated board at all.
				if bx[nextPosAdjusted[0]][nextPosAdjusted[1]] != '#' {
					nextq = append(nextq, nextPos)
					visitedRound[nextPos] = struct{}{}
					count++
					currentRoundTotal[boardOfNext]++
				}
			}
		}
		// fmt.Println(currentRoundTotal)
		for k, v := range currentRoundTotal {
			last4 := cacheRoundVals[k].val
			last4 = addVal(last4, v)
			cacheRoundVals[k] = cacheVal{val: last4}
		}
		// fmt.Println(cacheRoundVals)

		for k, v := range cacheRoundVals {
			if isCyclic(v.val) {
				delete(currentRoundTotal, k)
				// this is an even step which means odd steps will take the
				// start of the cycle since this is the end of the cycle
				if i%2 == 0 {
					cyclicCache[k] = [2]int{v.val[1], v.val[0]}
				} else {
					// even steps will take the start of the cycle
					cyclicCache[k] = [2]int{v.val[0], v.val[1]}
				}
			}
		}

		top = nextq
	}
	fmt.Println(len(top))
	fmt.Println(count)
}

func computeBoardOfLocation(loc twod.Location, board *rectboard.RectBoard) [2]int {
	bx := *board
	return [2]int{div(loc[0], len(bx)), div(loc[1], len(bx[0]))}
}

func isOnBorder(pos twod.Location, board *rectboard.RectBoard) bool {
	bx := *board
	if pos[0] == 0 || pos[0] == len(bx) {
		return true
	}
	if pos[1] == 0 || pos[1] == len(bx[0]) {
		return true
	}
	return false
}

func mod(x, d int) int {
	if x > 0 {
		return x % d
	} else {
		return (x%d + d) % d
	}
}

func div(x, d int) int {
	if x >= 0 {
		return x / d
	} else {
		return (x+1)/d - 1
	}
}

func part1() {

	s := findStartingPosition(board)
	fmt.Println(s)
	steps := 1000
	computeBfsInfiniteDumb(s, board, steps)
	// computeBfsInfiniteWithPerimeter(s, board, steps)
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

// func computeBfsInfiniteWithPerimeter(s twod.Location, board *rectboard.RectBoard, steps int) {

// 	bx := *board
// 	count := 0

// 	even := make(map[twod.Location]struct{})
// 	odd := make(map[twod.Location]struct{})

// 	// Always starts on an odd place by problem design
// 	odd[s] = struct{}{}

// 	// even i generates even positions.
// 	// odd i generates odd positions.
// 	for i := 0; i < steps; i++ {

// 		// this will be the new other
// 		visitedRound := make(map[twod.Location]struct{})

// 		var current map[twod.Location]struct{}
// 		var other map[twod.Location]struct{}

// 		if i%2 == 0 {
// 			// generating even positions, from the former odd positions
// 			current = odd
// 			other = even
// 		} else {
// 			// reverse
// 			current = even
// 			other = odd
// 		}

// 		// fmt.Println(i, odd, even, current)/

// 		// so the idea is that the positions we generated at step i will also
// 		// surely get generated at step i+2. We don't want to generate those
// 		// positions as the will contribute with nothing to the exploration,
// 		// since they already generated i+1
// 		for pos := range current {
// 			for _, step := range []twod.Direction{twod.UP, twod.LEFT, twod.DOWN, twod.RIGHT} {
// 				nextPos := twod.Move(pos, step)
// 				// we got here already this round
// 				if _, ok := visitedRound[nextPos]; ok {
// 					continue
// 				}
// 				// if the position is one from the previous rounds we already
// 				// used it to generate
// 				if _, ok := other[nextPos]; ok {
// 					continue
// 				}

// 				var nextPosAdjusted twod.Location
// 				nextPosAdjusted[0] = mod(nextPos[0], len(bx))
// 				nextPosAdjusted[1] = mod(nextPos[1], len(bx[0]))

// 				// only use adjusted for in board criteria
// 				if bx[nextPosAdjusted[0]][nextPosAdjusted[1]] != '#' {
// 					visitedRound[nextPos] = struct{}{}
// 					count++
// 				}
// 			}
// 		}
// 		other = visitedRound

// 		if i%2 == 0 {
// 			even = visitedRound
// 		} else {
// 			odd = visitedRound
// 		}

// 		// q = append(q, nextq)
// 		// fmt.Println(len(top))
// 	}
// 	// fmt.Println(len(top))
// 	fmt.Println(count)
// }
