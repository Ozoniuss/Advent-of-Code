package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)
	sum := 0
	nums := []int{}
	for s.Scan() {
		line := s.Text()
		no, err := strconv.Atoi(line)
		if err != nil {
			panic("parse")
		}

		nums = append(nums, no)
		sum += no
	}

	if s.Err() != nil {
		panic("lol")
	}

	fmt.Println(sum)
	sum = 0
	seen := map[int]struct{}{}

LOOP:
	for {
		for _, n := range nums {
			sum += n

			if _, ok := seen[sum]; ok {
				break LOOP
			}

			seen[sum] = struct{}{}
		}
	}

	fmt.Println(sum)
}
