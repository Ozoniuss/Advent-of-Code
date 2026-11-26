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

		stk := []byte{}
		for i := 0; i < len(line); i++ {
			stk = append(stk, line[i])
			if len(stk) >= 2 && (stk[len(stk)-1]-stk[len(stk)-1-1] == 'a'-'A' || stk[len(stk)-1-1]-stk[len(stk)-1] == 'a'-'A') {
				stk = stk[:len(stk)-2]
			}
		}

		fmt.Println(len(stk))
		minlen = min(minlen, len(stk))
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
