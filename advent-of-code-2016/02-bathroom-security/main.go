package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	. "aoc/golib/twod"
)

func dirToDir(dir rune) Direction {
	if dir == 'L' {
		return LEFT
	}
	if dir == 'R' {
		return RIGHT
	}
	if dir == 'D' {
		return DOWN
	}
	if dir == 'U' {
		return UP
	}
	return Direction{}
}

func isOutOfBound(position Location) bool {
	if position[0] < 0 || position[0] > 2 || position[1] < 0 || position[1] > 2 {
		return true
	}
	return false
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	current := Location{1, 1}

	scanner := bufio.NewScanner(f)
	code := ""
	for scanner.Scan() {
		line := scanner.Text()
		for _, dir := range line {
			next := Move(current, dirToDir(dir))
			if !isOutOfBound(next) {
				current = next
			}
		}
		code += strconv.Itoa(current[0]*3 + current[1] + 1)
	}
	fmt.Println(code)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isOutOfBoundRhombus(position Location) bool {
	if position[0] < -2 || position[0] > 2 {
		return true
	}
	if abs(position[0])+abs(position[1]) > 2 {
		return true
	}
	return false
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	current := Location{0, -2}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		for _, dir := range line {
			next := Move(current, dirToDir(dir))
			if !isOutOfBoundRhombus(next) {
				current = next
			}
		}
		fmt.Println(current) // convert manually from positions, too tired
		// to make this nice
	}
}

func main() {
	part2()
}
