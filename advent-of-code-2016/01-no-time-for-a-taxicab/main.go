package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	. "aoc/golib/twod"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	current := ORIGIN
	currentDir := UP // doesn't matter

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	changes := strings.Split(line, ", ")
	for _, change := range changes {
		if change[0] == 'L' {
			currentDir = TurnLeft(currentDir)
		} else {
			currentDir = TurnRight(currentDir)
		}
		moves, _ := strconv.Atoi(change[1:])
		current = Move(current, TranslateByInteger(currentDir, moves))
	}
	fmt.Println(ManhattanDistance(current, ORIGIN))
}
func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	current := ORIGIN
	currentDir := UP // doesn't matter

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	changes := strings.Split(line, ", ")

	visited := make(map[Location]struct{})
	visited[ORIGIN] = struct{}{}
	for _, change := range changes {
		if change[0] == 'L' {
			currentDir = TurnLeft(currentDir)
		} else {
			currentDir = TurnRight(currentDir)
		}
		moves, _ := strconv.Atoi(change[1:])
		// current = Move(current, TranslateByInteger(currentDir, moves))
		for i := 1; i <= moves; i++ {
			current = Move(current, currentDir)
			if _, ok := visited[current]; ok {
				goto END
			}
			visited[current] = struct{}{}
		}
	}
END:
	fmt.Println(ManhattanDistance(current, ORIGIN))
}

func main() {
	part1()
	part2()
}
