package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
var inputLines = readlines()

func readlines() [][][]byte {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var ret = make([][][]byte, 0, 64)

	var lines [][]byte
	s := bufio.NewScanner(f)
	for s.Scan() {
		if len(s.Text()) == 0 {
			ret = append(ret, slices.Clone(lines))
			lines = [][]byte{}
		} else {
			lines = append(lines, slices.Clone(s.Bytes()))
		}
	}
	ret = append(ret, slices.Clone(lines))
	return ret
}

func isReflectionColumn(col int, board [][]byte) bool {
	rightmost := 2*col - 1
	for j := col - 1; j >= 0; j-- {
		for i := 0; i < len(board); i++ {
			// ignore reflections that are larger
			if rightmost-j < len(board[0]) {
				if board[i][j] != board[i][rightmost-j] {
					return false
				}
			}
		}
	}
	return true
}

func isReflectionRow(row int, board [][]byte) bool {
	rightmost := 2*row - 1
	for i := row - 1; i >= 0; i-- {
		for j := 0; j < len(board[0]); j++ {
			// ignore reflections that are larger
			if rightmost-i < len(board) {
				if board[i][j] != board[rightmost-i][j] {
					return false
				}
			}
		}
	}
	return true
}

func isReflectionColumnWithSmudge(col int, board [][]byte) bool {
	rightmost := 2*col - 1
	cnt := 0
	total := 0
	for j := col - 1; j >= 0; j-- {
		for i := 0; i < len(board); i++ {
			// ignore reflections that are larger
			if rightmost-j < len(board[0]) {
				if board[i][j] == board[i][rightmost-j] {
					cnt++
				}
				total++
			}
		}
	}
	return cnt == total-1
}

func isReflectionRowWithSmudge(row int, board [][]byte) bool {
	rightmost := 2*row - 1
	cnt := 0
	total := 0
	for i := row - 1; i >= 0; i-- {
		for j := 0; j < len(board[0]); j++ {
			// ignore reflections that are larger
			if rightmost-i < len(board) {
				if board[i][j] == board[rightmost-i][j] {
					cnt++
				}
				total++
			}
		}
	}
	return cnt == total-1
}

func FindReflection(board [][]byte) (int, int) {
	for row := 1; row < len(board); row++ {
		if isReflectionRow(row, board) {
			return row, 0
		}
	}

	for col := 1; col < len(board[0]); col++ {
		if isReflectionColumn(col, board) {
			return 0, col
		}
	}
	return 0, 0
}

func FindReflectionSmudge(board [][]byte) (int, int) {
	for row := 1; row < len(board); row++ {
		if isReflectionRowWithSmudge(row, board) {
			return row, 0
		}
	}

	for col := 1; col < len(board[0]); col++ {
		if isReflectionColumnWithSmudge(col, board) {
			return 0, col
		}
	}
	return 0, 0
}

func part1() {
	count := 0
	for _, puzzle := range inputLines {
		// for i := 0; i < len(puzzle); i++ {
		// 	for j := 0; j < len(puzzle[0]); j++ {
		// 		fmt.Print(string(puzzle[i][j]))
		// 	}
		// 	fmt.Println()
		// }
		r, c := FindReflectionSmudge(puzzle)
		// fmt.Println(r, c)
		// fmt.Println(FindReflection(puzzle))
		count += 100*r + c
	}
	fmt.Println(count)
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
