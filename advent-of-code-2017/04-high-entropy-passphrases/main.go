package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
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
	count := 0
	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		words := make(map[string]struct{})
		count++
		for _, part := range parts {
			if _, ok := words[part]; ok {
				count--
				break
			}
			words[part] = struct{}{}
		}
	}
	fmt.Println(count)
}

func part2() {
	count := 0
	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		words := make(map[string]struct{})
		count++
		for _, part := range parts {
			partBytes := []byte(part)
			sort.Slice(partBytes, func(i, j int) bool {
				return partBytes[i] < partBytes[j]
			})
			part = string(partBytes)
			if _, ok := words[part]; ok {
				count--
				break
			}
			words[part] = struct{}{}
		}
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
	part2()
}
