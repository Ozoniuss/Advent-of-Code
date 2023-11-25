package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validTriangle(x1, x2, x3 int) bool {
	if x1+x2 <= x3 {
		return false
	}
	if x2+x3 <= x1 {
		return false
	}
	if x3+x1 <= x2 {
		return false
	}
	return true
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		num1 := line[:5]
		num2 := line[5:10]
		num3 := line[10:15]
		fmt.Println(num1, num2, num3)
		num1val, _ := strconv.Atoi(strings.TrimSpace(num1))
		num2val, _ := strconv.Atoi(strings.TrimSpace(num2))
		num3val, _ := strconv.Atoi(strings.TrimSpace(num3))
		if validTriangle(num1val, num2val, num3val) {
			count++
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

	var rows = [][3]string{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num1 := line[:5]
		num2 := line[5:10]
		num3 := line[10:15]
		rows = append(rows, [3]string{num1, num2, num3})
	}
	count := 0
	for i := 0; i < len(rows); i += 3 {
		for j := 0; j < 3; j++ {
			num1 := rows[i][j]
			num2 := rows[i+1][j]
			num3 := rows[i+2][j]
			num1val, _ := strconv.Atoi(strings.TrimSpace(num1))
			num2val, _ := strconv.Atoi(strings.TrimSpace(num2))
			num3val, _ := strconv.Atoi(strings.TrimSpace(num3))
			if validTriangle(num1val, num2val, num3val) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func main() {
	// part1()
	part2()
}
