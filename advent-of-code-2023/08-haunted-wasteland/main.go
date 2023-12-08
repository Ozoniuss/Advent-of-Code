package main

import (
	"bufio"
	"fmt"
	"os"
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
func part1() {
	directions := inputLines[0]
	nodesstr := inputLines[2:]

	locations := make(map[string][2]string)

	for _, line := range nodesstr {
		start := line[0:3]
		left := line[7:10]
		right := line[12:15]

		locations[start] = [2]string{left, right}
	}
	current := "AAA"

	start := -1
	count := 0
	for {
		start++
		start = start % len(directions)
		count++

		if directions[start] == 'L' {
			current = locations[current][0]
		} else {
			current = locations[current][1]
		}
		if current == "ZZZ" {
			break
		}
	}
	fmt.Println(count)
}

func part2() {
	directions := inputLines[0]
	nodesstr := inputLines[2:]

	locations := make(map[string][2]string)

	for _, line := range nodesstr {
		start := line[0:3]
		left := line[7:10]
		right := line[12:15]

		locations[start] = [2]string{left, right}
	}
	currents := []string{}

	for k := range locations {
		if k[2] == 'A' {
			currents = append(currents, k)
		}
	}

	start := -1
	count := 0
	rotations := make([]int, len(currents))
	for i := 0; i < len(currents); i++ {
		for {

			start++
			start = start % len(directions)
			count++

			if directions[start] == 'L' {
				currents[i] = locations[currents[i]][0]
			} else {
				currents[i] = locations[currents[i]][1]
			}
			if currents[i][2] == 'Z' {
				break
			}
		}
		rotations[i] = count
		count = 0
	}

	// compute LCM of this
	fmt.Println(rotations)
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
