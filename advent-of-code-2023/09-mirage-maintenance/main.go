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

// isZero checks whether or not an array only has zero values.
func isZero(nums []int) bool {
	for _, num := range nums {
		if num != 0 {
			return false
		}
	}
	if len(nums) == 1 && nums[0] != 0 {
		panic("it happened")
	}
	return true
}

// getDifferenceArr returns a slice with the differences between the numbers
// in the original array.
func getDifferenceArr(original []int) []int {
	difference := make([]int, len(original)-1)
	for i := 1; i < len(original); i++ {
		difference[i-1] = original[i] - original[i-1]
	}
	return difference
}

// findRemainingNumber finds the last number that should be added to the array
// to eventually get to a "zero" difference.
func findRemainingNumber(original []int) int {
	lastNumbers := make([]int, 0, len(original))
	lastNumbers = append(lastNumbers, original[len(original)-1])
	difference := slices.Clone(original)
	for {
		difference = getDifferenceArr(difference)
		lastNumbers = append(lastNumbers, difference[len(difference)-1])
		if isZero(difference) {
			break
		}
	}
	remaining := 0
	for i := len(lastNumbers) - 1; i >= 0; i-- {
		remaining = remaining + lastNumbers[i]
	}
	return remaining
}

// findRemainingNumberReverse finds the leading number that should be added to
// the array to eventually get to a "zero" difference.
func findRemainingNumberReverse(original []int) int {
	firstNumbers := make([]int, 0, len(original))
	firstNumbers = append(firstNumbers, original[0])
	difference := slices.Clone(original)
	for {
		difference = getDifferenceArr(difference)
		firstNumbers = append(firstNumbers, difference[0])
		if isZero(difference) {
			break
		}
	}
	current := 0
	for i := len(firstNumbers) - 1; i >= 0; i-- {
		current = firstNumbers[i] - current
	}
	return current
}

func part1() {
	var historyNums = make([][]int, 0, len(inputLines))
	for _, line := range inputLines {
		numstrs := strings.Split(line, " ")
		numsArr := make([]int, 0, len(numstrs))
		for _, numstr := range numstrs {
			num, _ := strconv.Atoi(numstr)
			numsArr = append(numsArr, num)
		}
		historyNums = append(historyNums, numsArr)
	}
	c := 0
	for _, hist := range historyNums {
		c += findRemainingNumber(hist)
	}
	fmt.Println(c)
}

func part2() {
	var historyNums = make([][]int, 0, len(inputLines))
	for _, line := range inputLines {
		numstrs := strings.Split(line, " ")
		numsArr := make([]int, 0, len(numstrs))
		for _, numstr := range numstrs {
			num, _ := strconv.Atoi(numstr)
			numsArr = append(numsArr, num)
		}
		historyNums = append(historyNums, numsArr)
	}
	c := 0
	for _, hist := range historyNums {
		c += findRemainingNumberReverse(hist)
	}
	fmt.Println(c)
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
