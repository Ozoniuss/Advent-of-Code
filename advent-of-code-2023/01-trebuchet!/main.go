package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := -1
		last := 0
		for i := 0; i < len(line); i++ {
			if line[i] >= '0' && line[i] <= '9' {
				if first == -1 {
					first = int(line[i] - '0')
				}
				last = int(line[i] - '0')
			}
		}
		total, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		fmt.Println(first, last)
		sum += total
	}
	fmt.Println(sum)
}

func part2() {
	f, err := os.Open("processed.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	substrings := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	scanner := bufio.NewScanner(f)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		first := -1
		last := 0

		parts := strings.Split(line, " ")
		fmt.Println(parts)
		for _, part := range parts {
			if len(part) == 0 {
				continue
			}
			if len(part) == 1 && part[0] >= '0' && part[0] <= '9' {
				if first == -1 {
					first = int(part[0] - '0')
				}
				last = int(part[0] - '0')
			} else if len(part) == 1 {
				continue
			} else {
				valleft, valright := findKeywordsInString(substrings, part)
				if valleft != -1 && valright != -1 {
					if first == -1 {
						first = valleft
					}
					last = valright
				}
			}
		}
		total, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
		fmt.Println("final", first, last, total)
		sum += total

	}
	fmt.Println(sum)
}

func findKeywordsInString(keywords map[string]int, str string) (int, int) {
	if len(str) < 3 {
		return -1, -1
	}
	leftidx := 1000
	rightidx := -1000
	valleft := -1
	valright := -1
	// keymin := ""
	// keymax := ""
	for keyword, v := range keywords {
		idx := strings.Index(str, keyword)
		if idx != -1 && idx < leftidx {
			leftidx = idx
			valleft = v
			// keymin = keyword
		}
		idx = strings.LastIndex(str, keyword)
		if idx != -1 && idx > rightidx {
			rightidx = idx
			valright = v
			// keymax = keyword
		}
	}
	// fmt.Println(keymin, keymax)
	return valleft, valright
}

func main() {
	// part1()
	part2()
}
