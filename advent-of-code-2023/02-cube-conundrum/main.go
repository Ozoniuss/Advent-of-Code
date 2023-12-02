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

func processOneGame(game string) (int, int, int) {
	red := 0
	green := 0
	blue := 0

	var eoc int
	for eoc != -1 {
		var currentGameWithColor string
		eoc = strings.Index(game, ", ")
		if eoc == -1 {
			currentGameWithColor = game
		} else {
			currentGameWithColor = game[:eoc]
			game = game[eoc+2:]
		}

		space := strings.Index(currentGameWithColor, " ")
		vals := currentGameWithColor[:space]
		color := currentGameWithColor[space+1:]

		if color == "red" {
			val, _ := strconv.Atoi(vals)
			red += val
		}
		if color == "green" {
			val, _ := strconv.Atoi(vals)
			blue += val
		}
		if color == "blue" {
			val, _ := strconv.Atoi(vals)
			green += val
		}
	}
	return red, blue, green
}

func part1() int {
	sum := 0
	for idx, line := range inputLines {
		firstSpace := strings.Index(line, ": ")
		gamesPart := line[firstSpace+2:]

		possible := false
		days := gamesPart
		eod := 0
		for eod != -1 {
			eod = strings.Index(days, "; ")
			var currentDay string
			if eod == -1 {
				currentDay = days
			} else {
				currentDay = days[:eod]
				days = days[eod+2:]
			}
			red, green, blue := processOneGame(currentDay)
			if red <= 12 && green <= 13 && blue <= 14 {
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
	return sum
}
func part2() int {
	sum := 0
	for _, line := range inputLines {
		firstSpace := strings.Index(line, ": ")
		gamesPart := line[firstSpace+2:]

		days := gamesPart
		eod := 0
		var red, green, blue int
		for eod != -1 {
			eod = strings.Index(days, "; ")
			var currentDay string
			if eod == -1 {
				currentDay = days
			} else {
				currentDay = days[:eod]
				days = days[eod+2:]
			}
			redg, greeng, blueg := processOneGame(currentDay)
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
