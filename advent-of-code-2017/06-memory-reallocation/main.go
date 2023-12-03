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
// var inputNums = readlines()

func readlines() []int {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var nums []int
	s := bufio.NewScanner(f)
	s.Scan()
	line := s.Text()
	parts := strings.Split(line, "\t")
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		nums = append(nums, num)
	}
	return nums
}

func maxPos(nums []int) int {
	pos := -1
	max := -9999
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			pos = i
		}
	}
	return pos
}

func toKey(nums []int) [16]int {
	var key [16]int
	for idx, num := range nums {
		key[idx] = num
	}
	return key
}

func redistributeFromPos(pos int, inputNums *[]int) {
	blocks := (*inputNums)[pos]
	(*inputNums)[pos] = 0
	current := pos + 1
	for ; blocks > 0; blocks-- {
		(*inputNums)[current%len(*inputNums)]++
		current++
	}
}

func part1() {
	inputNums := readlines()
	var states = make(map[[16]int]struct{}, 16)
	states[toKey(inputNums)] = struct{}{}

	fmt.Println(inputNums)
	count := 0

	for {
		// fmt.Println(inputNums)
		max := maxPos(inputNums)
		redistributeFromPos(max, &inputNums)
		fmt.Println(inputNums)
		count++

		key := toKey(inputNums)
		if _, ok := states[key]; ok {
			fmt.Println("count", count)
			return
		}
		states[key] = struct{}{}
	}
}

func part2() {
	inputNums := readlines()
	var states = make(map[[16]int]struct{}, 16)
	states[toKey(inputNums)] = struct{}{}

	fmt.Println(inputNums)
	count := 0

	var start [16]int
	for {
		max := maxPos(inputNums)
		redistributeFromPos(max, &inputNums)

		key := toKey(inputNums)
		if _, ok := states[key]; ok {
			start = key
			break
		}
		states[key] = struct{}{}
	}

	count = 0
	for {
		max := maxPos(inputNums)
		redistributeFromPos(max, &inputNums)
		count++

		key := toKey(inputNums)
		if key == start {
			fmt.Println("count", count)
			break
		}
	}
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
