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
	scanner.Scan()
	input := scanner.Text()

	count := 0
	for i := 0; i < len(input); i++ {
		if input[i%len(input)] == input[(i+1)%len(input)] {
			count += int((input[i] - '0'))
		}
	}
	fmt.Println(count)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	input := scanner.Text()

	count := 0
	for i := 0; i < len(input); i++ {
		if input[i%len(input)] == input[(i+len(input)/2)%len(input)] {
			count += int((input[i] - '0'))
		}
	}
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
