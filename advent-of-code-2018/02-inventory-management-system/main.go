package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	s := bufio.NewScanner(f)

	two := 0
	three := 0

	lines := []string{}

	for s.Scan() {
		cnt := make(map[byte]int)
		line := s.Text()
		lines = append(lines, line)

		for i := range len(line) {
			cnt[line[i]] += 1
		}

		seentwo := false
		seenthree := false
		for _, v := range cnt {
			if v == 2 {
				seentwo = true
			}
			if v == 3 {
				seenthree = true
			}
		}

		if seentwo {
			two += 1
		}
		if seenthree {
			three += 1
		}
	}

	if s.Err() != nil {
		panic("lol")
	}

	fmt.Println(two * three)

	for i := range len(lines) - 1 {
		for j := i + 1; j < len(lines); j++ {
			line1 := lines[i]
			line2 := lines[j]

			if len(line1) != len(line2) {
				panic("lol")
			}

			N := len(line1)
			b := strings.Builder{}

			for x := range N {
				if line1[x] == line2[x] {
					b.WriteByte(line1[x])
				}
			}

			s := b.String()
			if len(s) == N-1 {
				fmt.Println(s)
				break
			}

			b.Reset()
		}
	}
}

// abcde
// abcfg
// abdde
