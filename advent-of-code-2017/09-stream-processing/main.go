package main

import (
	"fmt"
	"os"
)

func main() {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input := string(b)

	gc := 0
	garbageCount := 0
	score := 0
	skip := false
	isgarbage := false
	for _, r := range input {
		if isgarbage {
			garbageCount += 1
		}
		if skip {
			skip = false
			garbageCount -= 1
			continue
		}
		switch r {
		case '<':
			if !isgarbage {
				isgarbage = true
			}
		case '>':
			if isgarbage {
				isgarbage = false
				garbageCount -= 1
			}
		case '!':
			if isgarbage {
				skip = true
				// ! only appears in garbage
				garbageCount -= 1
			}
		case '{':
			if !isgarbage {
				gc += 1
			}
		case '}':
			if !isgarbage {
				score += gc
				gc -= 1
			}
		default:
		}
	}
	fmt.Println(score)
	fmt.Println(garbageCount)

}
