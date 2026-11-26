package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	lineb, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	line := strings.TrimSpace(string(lineb))

	fmt.Println(string(line))
	origline := line

	minlen := math.MaxInt
	for cc := 'a'; cc <= 'z'; cc++ {
		fmt.Println("cc", cc)
		b := strings.Builder{}
		line := origline
		for _, r := range line {
			if r != cc && r != cc+'A'-'a' {
				b.WriteRune(r)
			}
		}
		line = b.String()
		shouldExit := true
		for {
			shouldExit = true
			for i := 1; i < len(line); i++ {
				if line[i]-line[i-1] == 'a'-'A' || line[i-1]-line[i] == 'a'-'A' {
					line = string(line[:i-1]) + string(line[i+1:])
					shouldExit = false
					break
				}
			}
			if shouldExit {
				break
			}
		}
		fmt.Println(len(line))
		minlen = min(minlen, len(line))
	}
	fmt.Println(minlen)
}

func part1(line string) {
	shouldExit := true
	for {
		shouldExit = true
		for i := 1; i < len(line); i++ {

			if line[i]-line[i-1] == 'a'-'A' || line[i-1]-line[i] == 'a'-'A' {
				line = string(line[:i-1]) + string(line[i+1:])
				shouldExit = false
				break
			}
		}
		if shouldExit {
			break
		}
	}

	fmt.Println(line, len(line))

}
