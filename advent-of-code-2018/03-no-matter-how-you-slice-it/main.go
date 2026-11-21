package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	board := map[[2]int]int{}
	overlaps := map[[2]int]int{}
	overlapsnums := map[int]struct{}{}
	lno := 0
	for s.Scan() {
		lno += 1
		line := s.Text()
		parts := strings.Split(line, " ")
		coords := parts[2][:len(parts[2])-1]
		coordsp := strings.Split(coords, ",")

		c0, _ := strconv.Atoi(coordsp[0])
		c1, _ := strconv.Atoi(coordsp[1])

		size := parts[3]
		sizep := strings.Split(size, "x")
		s0, _ := strconv.Atoi(sizep[0])
		s1, _ := strconv.Atoi(sizep[1])

		for i := c0; i < c0+s0; i++ {
			for j := c1; j < c1+s1; j++ {
				board[[2]int{i, j}] += 1

				if v, ok := overlaps[[2]int{i, j}]; ok {
					if v == -1 {
						overlapsnums[lno] = struct{}{}
					} else {
						overlapsnums[lno] = struct{}{}
						overlapsnums[v] = struct{}{}
						overlaps[[2]int{i, j}] = -1
					}
				} else {
					overlaps[[2]int{i, j}] = lno
				}
			}
		}

	}

	if s.Err() != nil {
		panic("lol")
	}

	fmt.Println(board)
	cnt := 0
	for _, v := range board {
		if v > 1 {
			cnt += 1
		}
	}

	fmt.Println(cnt)
	fmt.Println(len(overlapsnums))

	expected := 1247 * 1248 / 2
	actual := 0
	for k := range overlapsnums {
		actual += k
	}
	fmt.Println(expected - actual)
}
