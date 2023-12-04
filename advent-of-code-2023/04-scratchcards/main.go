package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// func part1() {

// 	sum := 0
// 	for _, line := range inputLines {
// 		start := strings.Index(line, ":")
// 		second := strings.Index(line, "|")

// 		firstSetStr := line[start+2 : second-1]
// 		secondSetStr := line[second+2:]

// 		firstSet := make(map[int]struct{})
// 		secondSet := make(map[int]struct{})

// 		fmt.Println("f", firstSetStr)
// 		fmt.Println("s", secondSetStr)

// 		partsFirst := strings.Split(firstSetStr, " ")
// 		partsSecond := strings.Split(secondSetStr, " ")

// 		for _, part := range partsFirst {
// 			num, err := strconv.Atoi(part)
// 			if err == nil {
// 				firstSet[num] = struct{}{}
// 			}
// 		}
// 		for _, part := range partsSecond {
// 			num, err := strconv.Atoi(part)
// 			if err == nil {
// 				secondSet[num] = struct{}{}
// 			}
// 		}

//			total := 0
//			for num := range firstSet {
//				if _, ok := secondSet[num]; ok {
//					if total == 0 {
//						total = 1
//					} else {
//						total *= 2
//					}
//				}
//			}
//			sum += total
//		}
//		fmt.Println(sum)
//	}
func part1() {

	sum := 0
	totalCopies := make(map[int]int)

	for i := 1; i <= len(inputLines); i++ {
		totalCopies[i] = 1
	}

	totalCards := len(inputLines)
	for idx, line := range inputLines {
		idx = idx + 1
		start := strings.Index(line, ":")
		second := strings.Index(line, "|")

		firstSetStr := line[start+2 : second-1]
		secondSetStr := line[second+2:]

		firstSet := make(map[int]struct{})
		secondSet := make(map[int]struct{})

		partsFirst := strings.Split(firstSetStr, " ")
		partsSecond := strings.Split(secondSetStr, " ")

		for _, part := range partsFirst {
			num, err := strconv.Atoi(part)
			if err == nil {
				firstSet[num] = struct{}{}
			}
		}
		for _, part := range partsSecond {
			num, err := strconv.Atoi(part)
			if err == nil {
				secondSet[num] = struct{}{}
			}
		}

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

	}
	for k := range totalCopies {
		sum += totalCopies[k]
	}
	fmt.Println(sum)
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
