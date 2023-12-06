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

	// input was processed manually
	f, err := os.Open("processed.txt")
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
	timesStr := inputLines[0]
	distancesStr := inputLines[1]

	times := make([]int, 0, 4)
	for _, timestr := range strings.Split(timesStr, " ") {
		time, _ := strconv.Atoi(timestr)
		times = append(times, time)
	}

	distances := make([]int, 0, 4)
	for _, distancestr := range strings.Split(distancesStr, " ") {
		distance, _ := strconv.Atoi(distancestr)
		distances = append(distances, distance)
	}

	fmt.Println(times, distances)
	count := 1
	for round := 0; round < len(times); round++ {
		roundCnt := 0
		for i := 0; i <= times[round]; i++ {
			totalDistance := (i) * (times[round] - i)
			if round == 0 {
				fmt.Println(totalDistance)
			}
			if totalDistance > distances[round] {
				roundCnt++
			}
		}
		count *= roundCnt
	}
	fmt.Println(count)
}
func part1() {
	timesStr := inputLines[0]
	distancesStr := inputLines[1]

	times := make([]int, 0, 4)
	for _, timestr := range strings.Split(timesStr, " ") {
		time, _ := strconv.Atoi(timestr)
		times = append(times, time)
	}

	distances := make([]int, 0, 4)
	for _, distancestr := range strings.Split(distancesStr, " ") {
		distance, _ := strconv.Atoi(distancestr)
		distances = append(distances, distance)
	}

	fmt.Println(times, distances)
	count := 1
	for round := 0; round < len(times); round++ {
		roundCnt := 0
		for i := 0; i <= times[round]; i++ {
			totalDistance := (i) * (times[round] - i)
			if totalDistance > distances[round] {
				roundCnt++
			}
		}
		count *= roundCnt
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
	// part2()
}
