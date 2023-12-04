package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/pkg/profile"
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

func isDigit(d byte) bool {
	return d >= '0' && d <= '9'
}

// allNumbersFromLine executes a function for all numbers from that line.
// It is the generalized version of the function below, but adds a bit of
// overhead by calling a function instead of assigning to a map directly.
func allNumbersFromLine(line string, f func(n int)) {
	innumber := false
	start := 0
	for i := 0; i < len(line); i++ {
		if !innumber && isDigit(line[i]) {
			innumber = true
			start = i
		}
		if innumber && i == len(line)-1 {
			num, _ := strconv.Atoi(line[start : i+1])
			f(num)
			return
		} else if innumber && !isDigit(line[i]) {
			num, _ := strconv.Atoi(line[start:i])
			f(num)
			innumber = false
		}
	}
}

// allNumbersFromLine executes a function for all numbers from that line.
func allNumberFromLineToStruct(line string, all map[int]struct{}) {
	innumber := false
	start := 0
	for i := 0; i < len(line); i++ {
		if !innumber && isDigit(line[i]) {
			innumber = true
			start = i
		}
		if innumber && i == len(line)-1 {
			num, _ := strconv.Atoi(line[start : i+1])
			all[num] = struct{}{}
			return
		} else if innumber && !isDigit(line[i]) {
			num, _ := strconv.Atoi(line[start:i])
			all[num] = struct{}{}
			innumber = false
		}
	}
}

func part1() int {

	sum := 0

	firstSet := make(map[int]struct{}, 16)
	secondSet := make(map[int]struct{}, 32)
	for _, line := range inputLines {
		start := strings.Index(line, ":")
		second := strings.Index(line, "|")

		firstLine := line[start+2 : second-1]
		secondLine := line[second+2:]

		allNumbersFromLine(firstLine, func(n int) {
			firstSet[n] = struct{}{}
		})
		allNumbersFromLine(secondLine, func(n int) {
			secondSet[n] = struct{}{}
		})
		// allNumberFromLineToStruct(firstLine, firstSet)
		// allNumberFromLineToStruct(secondLine, secondSet)

		total := 0
		for num := range firstSet {
			if _, ok := secondSet[num]; ok {
				if total == 0 {
					total = 1
				} else {
					total *= 2
				}
			}
		}
		sum += total
		clear(firstSet)
		clear(secondSet)
	}
	return sum
}

func part2() int {

	sum := 0

	// Otherwise it's hard to figure out how many copies each
	// number has.
	totalCopies := make(map[int]int, 256)

	for i := 1; i <= len(inputLines); i++ {
		totalCopies[i] = 1
	}
	firstSet := make(map[int]struct{}, 16)
	secondSet := make(map[int]struct{}, 32)

	totalCards := len(inputLines)
	for idx, line := range inputLines {
		idx = idx + 1
		start := strings.Index(line, ":")
		second := strings.Index(line, "|")

		firstLine := line[start+2 : second-1]
		secondLine := line[second+2:]

		allNumbersFromLine(firstLine, func(n int) {
			firstSet[n] = struct{}{}
		})
		allNumbersFromLine(secondLine, func(n int) {
			secondSet[n] = struct{}{}
		})

		copies := 0
		for num := range firstSet {
			if _, ok := secondSet[num]; ok {
				copies++
			}
		}
		var nextCardCopies = 0
		if idx+copies > totalCards {
			nextCardCopies = totalCards - idx
		} else {
			nextCardCopies = copies
		}

		for i := 1; i <= nextCardCopies; i++ {
			totalCopies[idx+i] += totalCopies[idx]
		}
		clear(firstSet)
		clear(secondSet)

	}
	for k := range totalCopies {
		sum += totalCopies[k]
	}
	return sum
}

func main() {
	// Run only 1 profile at a time!
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()

	// No return value improves speed during contest, but make these functions
	// return something for benchmarks once problem is solved.
	//
	// Part 2 is not written above and commented below so that template compiles
	// while solving part 1.

	r := 0
	for i := 0; i < 10000; i++ {
		r += part1() % 53
	}
	fmt.Println(r)
	// fmt.Println(part2())
}
