package main

import (
	"aoc/golib/twod"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
var inputLines = readlines()

func readlines() []string {
	f, err := os.Open("padded.txt")
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

func isNumber(r byte) bool {
	return r >= '0' && r <= '9'
}

func isSymbol(r byte) bool {
	return r != '.' && !isNumber(r)
}

func isOutBound(i, j int) bool {
	return (i < 0 || i >= len(inputLines)) || (j < 0 || j >= len(inputLines[0]))
}

func part1() int {
	sum := 0
	for idx, line := range inputLines {
		inNumber := false
		var startPos, endPos int
		var found = false
		for lineIdx := 0; lineIdx < len(line); lineIdx++ {
			// currently in number, don't care
			if isNumber(line[lineIdx]) && inNumber {
				continue
			}
			// currently out of number and not a number, don't care
			if !isNumber(line[lineIdx]) && !inNumber {
				continue
			}

			// reached the beginning of the number, care
			if isNumber(line[lineIdx]) && !inNumber {
				inNumber = true
				startPos = lineIdx
				continue
			}

			// reached the end of number, check it
			if !isNumber(line[lineIdx]) && inNumber {
				inNumber = false
				endPos = lineIdx
			}
			num, err := strconv.Atoi(line[startPos:endPos])
			if err != nil {
				panic("strconv")
			}

			// here work with line not current line
			// num is startPos->endPos-1
			for i := idx - 1; i <= idx+1; i++ {
				for j := startPos - 1; j <= endPos; j++ {
					if isOutBound(i, j) {
						continue
					}
					if i == idx && (j <= endPos-1 && j >= startPos) {
						continue
					}

					if isSymbol(inputLines[i][j]) {
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				sum += num
			}
			startPos, endPos = 0, 0
			found = false
		}

	}
	return sum
}

func part2() int {
	positions := make(map[twod.Location][2]int)
	numberIndex := 0
	for idx, line := range inputLines {
		current := line + "."
		inNumber := false
		var startPos, endPos int
		for lineIdx := 0; lineIdx < len(current); lineIdx++ {
			// currently in number, don't care
			if isNumber(current[lineIdx]) && inNumber {
				continue
			}
			// currently out of number and not a number, don't care
			if !isNumber(current[lineIdx]) && !inNumber {
				continue
			}

			// reached the beginning of the number, care
			if isNumber(current[lineIdx]) && !inNumber {
				inNumber = true
				startPos = lineIdx
				continue
			}

			// reached the end of number, check it
			if !isNumber(current[lineIdx]) && inNumber {
				inNumber = false
				endPos = lineIdx
			}
			num, err := strconv.Atoi(current[startPos:endPos])
			if err != nil {
				panic("strconv")
			}
			for i := startPos; i < endPos; i++ {
				positions[twod.Location{idx, i}] = [2]int{num, numberIndex}
			}
			numberIndex++

		}

	}
	sum := 0
	for i := 0; i < len(inputLines); i++ {
		for j := 0; j < len(inputLines[i]); j++ {
			prod := 1
			neighbours := make(map[int]int)
			if inputLines[i][j] == '*' {
				for stari := i - 1; stari <= i+1; stari++ {
					for starj := j - 1; starj <= j+1; starj++ {
						if val, ok := positions[twod.Location{stari, starj}]; ok {
							neighbours[val[1]] = val[0]
						}
					}
				}
				if len(neighbours) == 2 {
					for _, n := range neighbours {
						prod = prod * n
					}
					sum = sum + prod
				}
			}
		}
	}
	return sum
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

	fmt.Println(part1())
	fmt.Println(part2())
}
