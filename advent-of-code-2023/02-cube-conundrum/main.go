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

func findScores(round string) (int, int, int) {
	red := 0
	green := 0
	blue := 0

	var eoc int
	for eoc != -1 {
		var currentScoreWithColor string
		eoc = strings.Index(round, ", ")
		if eoc == -1 {
			currentScoreWithColor = round
		} else {
			currentScoreWithColor = round[:eoc]
			round = round[eoc+2:]
		}

		space := strings.Index(currentScoreWithColor, " ")
		scoreStr := currentScoreWithColor[:space]
		color := currentScoreWithColor[space+1:]

		score, _ := strconv.Atoi(scoreStr)
		if color == "red" {
			red += score
		}
		if color == "green" {
			blue += score
		}
		if color == "blue" {
			green += score
		}
	}
	return red, blue, green
}

func isRoundPossible(red, green, blue int) bool {
	return red <= 12 && green <= 13 && blue <= 14
}

func part1() int {
	sum := 0
	for gameIdx, line := range inputLines {
		gameContentStartIdx := strings.Index(line, ": ")
		game := line[gameContentStartIdx+2:]

		possible := false
		eod := 0
		for eod != -1 {
			eod = strings.Index(game, "; ")
			var round string
			if eod == -1 {
				round = game
			} else {
				round = game[:eod]
				game = game[eod+2:]
			}
			red, green, blue := findScores(round)
			if isRoundPossible(red, green, blue) {
				possible = true
			} else {
				possible = false
				break
			}
		}
		if possible {
			sum += gameIdx + 1
		}
	}
	return sum
}
func part2() int {
	sum := 0
	for _, line := range inputLines {
		gameContentStartIdx := strings.Index(line, ": ")
		game := line[gameContentStartIdx+2:]

		eod := 0
		var red, green, blue int
		for eod != -1 {
			eod = strings.Index(game, "; ")
			var round string
			if eod == -1 {
				round = game
			} else {
				round = game[:eod]
				game = game[eod+2:]
			}
			redg, greeng, blueg := findScores(round)
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
