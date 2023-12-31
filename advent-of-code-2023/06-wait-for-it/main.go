package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func getInput() ([]int, []int) {

	r := regexp.MustCompile("[0-9]+")
	allTimes := r.FindAllString(inputLines[0], -1)
	allDistances := r.FindAllString(inputLines[1], -1)

	times := make([]int, len(allTimes))
	distances := make([]int, len(allDistances))

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(allTimes[i])
		distance, _ := strconv.Atoi(allDistances[i])

		times[i] = time
		distances[i] = distance
	}

	return times, distances
}

func getInputMerged() (int, int) {

	r := regexp.MustCompile("[0-9]+")
	allTimes := r.FindAllString(inputLines[0], -1)
	allDistances := r.FindAllString(inputLines[1], -1)

	bigTime := strings.Join(allTimes, "")
	bigDistances := strings.Join(allDistances, "")

	time, _ := strconv.Atoi(bigTime)
	distance, _ := strconv.Atoi(bigDistances)

	return time, distance
}

var inputTimes, inputDistances = getInput()
var inputBigTime, inputBigDistance = getInputMerged()

// findFromLeftAndRight attempts to optimize the problem given the nature of
// a second-degree polynomial. The first root is time=0 and the second root
// is time=distance.
func findFromLeftAndRight(time int, distance int) int {
	var left, right int
	for i := 0; i < time; i++ {
		if i*(time-i) > distance {
			left = i
			break
		}
	}
	for j := time - 1; j >= 0; j-- {
		if j*(time-j) > distance {
			right = j
			break
		}
	}
	return right - left + 1
}

func part1(times, distances []int) int {
	count := 1
	for round := 0; round < len(inputTimes); round++ {
		roundCnt := findFromLeftAndRight(times[round], distances[round])
		count *= roundCnt
	}
	return count
}

func part2(time, distance int) int {
	return findFromLeftAndRight(time, distance)
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

	fmt.Println(part1(inputTimes, inputDistances))
	fmt.Println(part2(inputBigTime, inputBigDistance))
}
