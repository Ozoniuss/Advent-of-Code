package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 6:18, 7:28

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
// var inputNums = readlines()

func readlines() []int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Bytes()
		// does this alloc?
		num, _ := strconv.Atoi(string(line))
		lines = append(lines, num)
	}
	return lines
}

func part1() int {
	inputNums := readlines()
	pos := 0
	count := 0
	for {
		jump := inputNums[pos]
		inputNums[pos]++

		pos = pos + jump
		count++
		if pos < 0 || pos >= len(inputNums) {
			break
		}
	}
	return count
}

func part2() int {
	inputNums := readlines()
	pos := 0
	count := 0
	for {
		jump := inputNums[pos]
		if jump >= 3 {
			inputNums[pos]--
		} else {
			inputNums[pos]++
		}

		pos = pos + jump
		count++
		if pos < 0 || pos >= len(inputNums) {
			return count
		}
	}
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
