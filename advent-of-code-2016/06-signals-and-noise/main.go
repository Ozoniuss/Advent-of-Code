package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	all := make([]string, 0, 20)
	for scanner.Scan() {
		line := scanner.Text()
		all = append(all, line)
	}

	length := len(all[0])
	mostCommon := make([]byte, length, length)
	for i := 0; i < length; i++ {
		allChars := make(map[byte]int)
		for _, str := range all {
			allChars[str[i]]++
		}
		mostCommonChar := finxMax(allChars)
		mostCommon[i] = mostCommonChar
	}
	fmt.Println(string(mostCommon))
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	all := make([]string, 0, 20)
	for scanner.Scan() {
		line := scanner.Text()
		all = append(all, line)
	}

	length := len(all[0])
	mostCommon := make([]byte, length, length)
	for i := 0; i < length; i++ {
		allChars := make(map[byte]int)
		for _, str := range all {
			allChars[str[i]]++
		}
		mostCommonChar := finxMin(allChars)
		mostCommon[i] = mostCommonChar
	}
	fmt.Println(string(mostCommon))
}

func finxMax(chars map[byte]int) byte {
	max := -1
	var maxChar byte
	for k, v := range chars {
		if v > max {
			max = v
			maxChar = k
		}
	}
	return maxChar
}

func finxMin(chars map[byte]int) byte {
	min := 9999
	var minChar byte
	for k, v := range chars {
		if v < min {
			min = v
			minChar = k
		}
	}
	return minChar
}

func main() {
	part1()
	part2()
}
