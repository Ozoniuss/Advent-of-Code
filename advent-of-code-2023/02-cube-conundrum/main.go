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

func processGame(game string) (int, int, int) {
	parts := strings.Split(game, ", ")
	red := 0
	green := 0
	blue := 0
	for _, part := range parts {
		gameWithColor := strings.Split(part, " ")
		game := gameWithColor[0]
		color := gameWithColor[1]
		if color == "red" {
			val, _ := strconv.Atoi(game)
			red += val
		}
		if color == "green" {
			val, _ := strconv.Atoi(game)
			blue += val
		}
		if color == "blue" {
			val, _ := strconv.Atoi(game)
			green += val
		}
	}
	return red, blue, green
}

func part1() {
	sum := 0
	for idx, line := range inputLines {
		parts := strings.Split(line, ": ")
		games := parts[1]
		days := strings.Split(games, "; ")
		// red := 0
		// green := 0
		// blue := 0
		possible := false
		for _, day := range days {
			redg, greeng, blueg := processGame(day)
			// red += redg
			// green += greeng
			// blue += blueg
			if redg <= 12 && greeng <= 13 && blueg <= 14 {
				possible = true
			} else {
				possible = false
				break
			}
		}
		if possible {
			sum += idx + 1
		}
	}
	fmt.Println(sum)
}
func part2() {
	sum := 0
	for _, line := range inputLines {
		parts := strings.Split(line, ": ")
		games := parts[1]
		days := strings.Split(games, "; ")
		red := 0
		green := 0
		blue := 0
		for _, day := range days {
			redg, greeng, blueg := processGame(day)
			if redg > red {
				red = redg
			}
			if blueg > blue {
				blue = blueg
			}
			if greeng > green {
				green = greeng
			}
		}
		sum += red * blue * green
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

	// part1()
	part2()
}
