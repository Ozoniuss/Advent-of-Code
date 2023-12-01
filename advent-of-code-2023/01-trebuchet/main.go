package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/pkg/profile"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function).
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

func part1() int {
	sum := 0
	for _, line := range inputLines {

		first := -1
		last := 0
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				if first == -1 {
					first = int(line[i] - '0')
				}
				last = int(line[i] - '0')
			}
		}
		total := first*10 + last
		sum += total
	}
	return sum
}

func part2WithSubstr() int {

	substrings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	sum := 0
	for _, line := range inputLines {
		first := -1
		last := 0
		valleft, valright := findKeywordsInString(line, &substrings)
		if valleft != -1 && valright != -1 {
			if first == -1 {
				first = valleft
			}
			last = valright
		}
		total := first*10 + last
		sum += total
	}
	return sum
}

func findKeywordsInString(str string, keywords *map[string]int) (int, int) {
	if len(str) == 0 {
		return -1, -1
	}
	leftidx := 1000
	rightidx := -1000
	valleft := -1
	valright := -1
	for keyword, v := range *keywords {
		idx := strings.Index(str, keyword)
		if idx != -1 && idx < leftidx {
			leftidx = idx
			valleft = v
		}
		idx = strings.LastIndex(str, keyword)
		if idx != -1 && idx > rightidx {
			rightidx = idx
			valright = v
		}
	}
	return valleft, valright
}

func computeDigit(input string, substrings *map[string]int) int {
	if len(input) != 1 && len(input) != 3 && len(input) != 4 && len(input) != 5 {
		return -1
	}
	if val, ok := (*substrings)[input]; ok {
		return val
	}
	return -1
}

func part2NoSubstr() int {

	// should not escape to heap
	substrings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	sum := 0
	possibleLenghts := [5]int{5, 4, 3, 1}
	for _, line := range inputLines {
		first := -1
		last := 0

		for i := 1; i <= len(line); i++ {
			var digitLeft = -1
			var digitRight = -1
			for _, lenght := range possibleLenghts {
				digit := -1
				if i-lenght >= 0 {
					digit = computeDigit(line[i-lenght:i], &substrings)
				}
				// found a digit
				if digit != -1 && digitLeft == -1 {
					digitLeft = digit
				}
				if digit != -1 {
					digitRight = digit
				}
			}
			// If one is different, both are
			if digitRight != -1 {
				if first == -1 {
					first = digitLeft
				}
				last = digitRight
			}

		}
		total := first*10 + last
		sum += total
	}
	return sum
}

func main() {
	defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()
	part1()
	part2WithSubstr()
	part2NoSubstr()
}
