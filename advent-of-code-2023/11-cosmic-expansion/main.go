package main

import (
	brd "aoc/golib/board"
	"aoc/golib/twod"
	"bufio"
	"os"
	"slices"
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

func readBoard() *brd.Board {
	return brd.NewBoard(brd.WithLines(&inputLines), brd.WithRect())
}

func findGalaxies(board *brd.Board) []twod.Location {
	galaxies := make([]twod.Location, 0, 100)
	for k, v := range board.GetLocations() {
		if v != '.' {
			galaxies = append(galaxies, k)
		}
	}
	return galaxies
}

func findEmptyColumns(board *brd.Board) []int {
	columns := make([]int, 0, 256)
	for col := board.GetUpperLeftBoundary()[1]; col <= board.GetLowerRightBoundary()[1]; col++ {
		noGalaxies := true
		for i := board.GetUpperLeftBoundary()[0]; i <= board.GetLowerRightBoundary()[0]; i++ {
			if board.MustGet(twod.Location{i, col}) != '.' {
				noGalaxies = false
			}
		}
		if noGalaxies {
			columns = append(columns, col)
		}
	}
	return columns
}

func findEmptyRows(board *brd.Board) []int {
	rows := make([]int, 0, 256)
	for row := board.GetUpperLeftBoundary()[0]; row <= board.GetLowerRightBoundary()[0]; row++ {
		noGalaxies := true
		for j := board.GetUpperLeftBoundary()[1]; j <= board.GetLowerRightBoundary()[1]; j++ {
			if board.MustGet(twod.Location{row, j}) != '.' {
				noGalaxies = false
			}
		}

		if noGalaxies {
			rows = append(rows, row)
		}
	}
	return rows
}

func superManhattanDistance(l1 twod.Location, l2 twod.Location, emptyRows []int, emptyCols []int, padding int) int {
	var startRow, endRow, startCol, endCol int
	if l1[0] < l2[0] {
		startRow, endRow = l1[0], l2[0]
	} else {
		startRow, endRow = l2[0], l1[0]
	}

	if l1[1] < l2[1] {
		startCol, endCol = l1[1], l2[1]
	} else {
		startCol, endCol = l2[1], l1[1]
	}

	extradistance := 0
	for i := startRow; i < endRow; i++ {
		if slices.Contains(emptyRows, i) {
			extradistance += padding
		}
	}
	for j := startCol; j < endCol; j++ {
		if slices.Contains(emptyCols, j) {
			extradistance += padding
		}
	}
	return twod.ManhattanDistance(l1, l2) + extradistance
}

func partX() int {
	board := readBoard()
	emptyCols := findEmptyColumns(board)
	emptyRows := findEmptyRows(board)

	galaxies := findGalaxies(board)

	s := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			// s += superManhattanDistance(galaxies[i], galaxies[j], emptyRows, emptyCols, 1)
			s += superManhattanDistance(galaxies[i], galaxies[j], emptyRows, emptyCols, 999999)
		}
	}
	return s
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

	partX()
}
