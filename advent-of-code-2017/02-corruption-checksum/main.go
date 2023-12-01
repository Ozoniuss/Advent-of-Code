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

func part1() {
	sum := 0
	for _, line := range inputLines {
		nums := strings.Split(line, "\t")
		max := -9999
		min := 9999
		for _, num := range nums {
			numv, _ := strconv.Atoi(num)
			if numv > max {
				max = numv
			}
			if numv < min {
				min = numv
			}
		}
		sum += max - min
	}
	fmt.Println(sum)
}

func part2() {
	sum := 0
	for _, line := range inputLines {
		nums := strings.Split(line, "\t")
		numarr := make([]int, 0, len(nums))
		for _, num := range nums {
			numv, _ := strconv.Atoi(num)
			numarr = append(numarr, numv)
		}
		for i := 0; i < len(numarr)-1; i++ {
			for j := i + 1; j < len(numarr); j++ {
				a := numarr[i]
				b := numarr[j]
				if a < b {
					a, b = b, a
				}
				if a%b == 0 {
					sum += a / b
				}
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	// Run only 1 profile at a time!
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()

	// This improves speed, but make these functions return something for
	// benchmarks once problem is solved.
	part1()
	part2()
}
