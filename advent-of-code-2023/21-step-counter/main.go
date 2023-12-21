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

func computeBfs(s twod.Location, board *rectboard.RectBoard) {
	q := make([][]twod.Location, 0)
	q = append(q, []twod.Location{s})

	bx := *board
	for i := 0; i < 64; i++ {
		top := q[0]
		q = q[1:]

		// the places reached this round to avoid duplicates
		visitedRound := make(map[twod.Location]struct{})
		nextq := make([]twod.Location, 0, 64)
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
				if bx[nextPos[0]][nextPos[1]] != '#' {
					nextq = append(nextq, nextPos)
					visitedRound[nextPos] = struct{}{}
				}
			}
		}
		q = append(q, nextq)
	}
	fmt.Println(len(q[0]))
}

func part1() {
	s := findStartingPosition(board)
	fmt.Println(s)
	computeBfs(s, board)
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
