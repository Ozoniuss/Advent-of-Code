package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
var inputLines = readlines()

func readlines() [][]byte {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines [][]byte
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, slices.Clone(s.Bytes()))
	}
	return lines
}

func calculateScore(rocks *[][]byte) int {
	s := 0
	// for j := 0; j < len((*rocks)[0]); j++ {
	// 	amountRocks := 0
	// 	for i := len((*rocks)) - 1; i >= 0; i-- {
	// 		if (*rocks)[i][j] == 'O' {
	// 			amountRocks++
	// 			s += amountRocks
	// 		}
	// 		if (*rocks)[i][j] == '#' {
	// 			amountRocks = 0
	// 		}
	// 	}
	// }
	for j := 0; j < len((*rocks)[0]); j++ {
		for i := 0; i < len((*rocks)); i++ {
			if (*rocks)[i][j] == 'O' {
				s += len(*rocks) - i
			}
		}
	}
	return s
}

func slideSingleRowNorth(row int, rocks *[][]byte) {
	for j := 0; j < len((*rocks)[0]); j++ {
		for i := row - 1; i >= 0; i-- {
			if (*rocks)[i][j] == '.' && (*rocks)[i+1][j] == 'O' {
				(*rocks)[i][j] = ((*rocks)[i+1][j])
				(*rocks)[i+1][j] = '.'
			} else {
				break
			}
		}
	}
}
func slideSingleRowSouth(row int, rocks *[][]byte) {
	for j := 0; j < len((*rocks)[0]); j++ {
		for i := row + 1; i < len((*rocks)); i++ {
			if (*rocks)[i][j] == '.' && (*rocks)[i-1][j] == 'O' {
				(*rocks)[i][j] = ((*rocks)[i-1][j])
				(*rocks)[i-1][j] = '.'
			} else {
				break
			}
		}
	}
}
func slideSingleRowEast(col int, rocks *[][]byte) {
	for i := 0; i < len((*rocks)); i++ {
		for j := col + 1; j < len((*rocks)[0]); j++ {
			if (*rocks)[i][j] == '.' && (*rocks)[i][j-1] == 'O' {
				(*rocks)[i][j] = ((*rocks)[i][j-1])
				(*rocks)[i][j-1] = '.'
			} else {
				break
			}
		}
	}
}

func slideSingleRowWest(col int, rocks *[][]byte) {
	for i := 0; i < len((*rocks)); i++ {
		for j := col - 1; j >= 0; j-- {
			if (*rocks)[i][j] == '.' && (*rocks)[i][j+1] == 'O' {

				(*rocks)[i][j] = ((*rocks)[i][j+1])
				(*rocks)[i][j+1] = '.'
			} else {
				break
			}
		}
	}
}

func SlideNorth(rocks *[][]byte) {
	for i := 1; i < len(*rocks); i++ {
		slideSingleRowNorth(i, rocks)
	}
}
func SlideSouth(rocks *[][]byte) {
	for i := len(*rocks) - 2; i >= 0; i-- {
		slideSingleRowSouth(i, rocks)
	}
}
func SlideWest(rocks *[][]byte) {
	for j := 1; j < len((*rocks)[0]); j++ {
		slideSingleRowWest(j, rocks)
	}
}
func SlideEast(rocks *[][]byte) {
	for j := len((*rocks)[0]) - 2; j >= 0; j-- {
		slideSingleRowEast(j, rocks)
	}
}

func SliceOneCycle(rocks *[][]byte, moves []func(*[][]byte)) {
	for _, move := range moves {
		move(rocks)
	}
}

func computeRocksCache(rocks *[][]byte) string {
	sb := &strings.Builder{}
	for i := 0; i < len(*rocks); i++ {
		for j := 0; j < len((*rocks)[0]); j++ {
			sb.WriteByte((*rocks)[i][j])
		}
	}
	return sb.String()
}

func reconstructFromCache(cache string, rows int, cols int) [][]byte {
	rocks := make([][]byte, rows)
	for i := 0; i < rows; i++ {
		rocks[i] = make([]byte, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rocks[i][j] = cache[i*rows+j]
		}
	}
	return rocks
}

func part2() {
	rocks := inputLines

	moves := []func(*[][]byte){SlideNorth, SlideWest, SlideSouth, SlideEast}

	cache := make(map[string]int)
	reverseCache := make(map[int]string)
	noMoves := 1000000000

	firstCacheHit := 0
	cacheval := ""
	secondCacheHit := 0

	for i := 0; i < noMoves; i++ {
		for mvidx, move := range moves {
			actualMoveIdx := i*4 + mvidx
			move(&rocks)

			rocksCache := computeRocksCache(&rocks)
			if _, ok := cache[rocksCache]; ok {
				if firstCacheHit == 0 {
					firstCacheHit = actualMoveIdx
					cacheval = rocksCache
					fmt.Println("firsthit", actualMoveIdx)
				} else {

					if rocksCache == cacheval {
						secondCacheHit = actualMoveIdx
						fmt.Println("secondhit", actualMoveIdx)
						goto END
					}
				}
			} else {
				cache[rocksCache] = actualMoveIdx
			}
			if firstCacheHit != 0 {
				reverseCache[actualMoveIdx] = rocksCache
			}
		}
	}
END:

	fmt.Println(firstCacheHit, secondCacheHit)
	cycleLenght := secondCacheHit - firstCacheHit

	current := firstCacheHit
	for current < 1000000000*4 {
		current += cycleLenght
	}
	fromEnd := current - 1000000000*4
	posInCycle := secondCacheHit - fromEnd - firstCacheHit

	currentRocksState := reconstructFromCache(reverseCache[firstCacheHit+posInCycle-1], len(rocks), len(rocks[0]))
	rocks = currentRocksState

	s := ""
	for i := 0; i < len(rocks); i++ {
		for j := 0; j < len(rocks[0]); j++ {
			s += (string((rocks)[i][j]))
		}
		s += "\n"
	}
	fmt.Println(s)
	fmt.Println(calculateScore(&rocks))
}

// func part1() {
// 	rocks := inputLines

// 	SlideNorth(&rocks)
// 	// slideSingleRow(1, &rocks)
// 	// slideSingleRow(3, &rocks)

// 	s := ""
// 	for i := 0; i < len(rocks); i++ {
// 		for j := 0; j < len(rocks[0]); j++ {
// 			s += (string((rocks)[i][j]))
// 		}
// 		s += "\n"
// 	}
// 	fmt.Println(s)
// 	fmt.Println(calculateScore(&rocks))
// }

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
