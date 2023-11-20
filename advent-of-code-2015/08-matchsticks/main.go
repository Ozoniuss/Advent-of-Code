package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// readPart reads the part number from the standard input.
func readPart() int {
	var part int
	fmt.Print("part: ")
	fmt.Scanf("%d\n", &part)
	return part
}

func isHexa(c byte) bool {
	return (c >= 'a' && c <= 'f') || (c >= '0' && c <= '9')
}

func processLine(line string) (total int, inmemory int) {

	// chars stored in file are all actual literals
	total = len(line)

	// works because they are ASCII encoded
	i := 1
	for i < len(line)-1 {
		if line[i] == '\x5c' && line[i-1] == '\x5c' {
			// inmemory-- because we're skipping
			i++
		}
		if line[i] == '"' && line[i-1] == '\x5c' {
			inmemory--
		}

		// we need entire condition to avoid shit like 'macac' or
		// a//x or mxcac
		if (i >= 3) && isHexa(line[i]) && isHexa(line[i-1]) && line[i-2] == 'x' && line[i-3] == '\x5c' {
			inmemory -= 3
		}
		inmemory++
		i++
	}
	return
}

func processLineStrconv(line string) (total, inmemory int) {
	// chars stored in file are all actual literals
	total = len(line)

	unquoted, err := strconv.Unquote(line)
	if err != nil {
		panic(err)
	}
	inmemory = len(unquoted)
	return
}

func processLinePart2(line string) int {
	total := 6
	for i := 1; i < len(line)-1; i++ {
		if line[i] == '\x5c' {
			total++
		}
		if line[i] == '"' {
			total++
		}
		total++
	}
	fmt.Println(line, total)
	return total - len(line)
}

// This also works
func processLinePart2Std(line string) int {
	return len(fmt.Sprintf("%#v", line)) - len(line)
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	part := readPart()

	if part == 1 {
		count := 0
		for scanner.Scan() {
			line := scanner.Text()
			total, inmemory := processLineStrconv(line)
			count += total - inmemory
		}
		fmt.Println(count)
	} else if part == 2 {
		count := 0
		for scanner.Scan() {
			line := scanner.Text()
			count += processLinePart2Std(line)
		}
		fmt.Println(count)
	} else {
		panic("wtf")
	}
}
