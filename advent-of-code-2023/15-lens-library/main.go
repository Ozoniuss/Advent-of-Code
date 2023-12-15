package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func readline() string {
	return readlines()[0]
}

var inputline = readline()

func computeHash(s string) int {
	c := 0
	for i := 0; i < len(s); i++ {
		c += int(s[i])
		c *= 17
		c = c % 256
	}
	return c
}

type headerWithVal struct {
	header string
	val    int
}

func part2() {
	boxes := [256][]headerWithVal{}
	for i := 0; i < 256; i++ {
		boxes[i] = []headerWithVal{}
	}
	// total := 0
	inputlineTrimmed := strings.Trim(inputline, "\n")
	parts := strings.Split(inputlineTrimmed, ",")

	for _, part := range parts {
		eqinx := strings.Index(part, "=")
		minidx := strings.Index(part, "-")

		if eqinx != -1 {
			header := part[:eqinx]
			valstr := part[eqinx+1:]

			val, _ := strconv.Atoi(valstr)

			boxNumber := computeHash(header)

			set := false
			for idx, len := range boxes[boxNumber] {
				if len.header == header {
					boxes[boxNumber][idx].val = val
					set = true
					break
				}
			}
			if !set {
				fmt.Println("here")
				boxes[boxNumber] = append(boxes[boxNumber], headerWithVal{header: header, val: val})
			}
		}

		if minidx != -1 {
			header := part[:minidx]
			boxNumber := computeHash(header)

			idx := -1
			for i := 0; i < len(boxes[boxNumber]); i++ {
				if boxes[boxNumber][i].header == header {
					idx = i
					break
				}
			}
			if idx != -1 {
				boxes[boxNumber] = append(slices.Clone(boxes[boxNumber][:idx]), slices.Clone(boxes[boxNumber][idx+1:])...)
			}
		}

	}
	total := 0
	for i := 0; i < 256; i++ {
		for j, hwv := range boxes[i] {
			total += (i + 1) * (j + 1) * hwv.val
		}
	}
	fmt.Println(total)
}

// func part1() {
// 	total := 0
// 	inputlineTrimmed := strings.Trim(inputline, "\n")
// 	parts := strings.Split(inputlineTrimmed, ",")
// 	for _, part := range parts {
// 		total += computeHash(part)
// 	}
// 	fmt.Println(total)
// }

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
