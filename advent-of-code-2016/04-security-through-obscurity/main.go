package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type letterWithFreq struct {
	letter byte
	freq   int
}

func processLine(line string) int {
	var letters string
	var number int
	var hash string
	fmt.Sscanf(line, "%s %d %s", &letters, &number, &hash)

	var allLetters = make(map[byte]int, 0)
	for i := 0; i < len(letters); i++ {
		allLetters[letters[i]]++
	}
	var lettersWithFreq = make([]letterWithFreq, 0)
	for l, f := range allLetters {
		lettersWithFreq = append(lettersWithFreq, letterWithFreq{
			letter: l,
			freq:   f,
		})
	}
	slices.SortFunc(lettersWithFreq, func(a, b letterWithFreq) int {
		if a.freq == b.freq {
			return int(a.letter) - int(b.letter)
		}
		return b.freq - a.freq
	})

	b := &strings.Builder{}
	for i := 0; i < 5; i++ {
		b.WriteByte(lettersWithFreq[i].letter)
	}

	if b.String() == hash {
		return number
	}

	return 0
}

func rotateLetter(l byte, i int) byte {
	return 'a' + byte((int(l)-'a'+i)%26)
}

func processLineRotation(line string) (string, int) {
	var letters string
	var number int
	var hash string
	fmt.Sscanf(line, "%s %d %s", &letters, &number, &hash)

	b := &strings.Builder{}
	for _, letter := range letters {
		b.WriteByte(rotateLetter(byte(letter), number))
	}

	return b.String(), number
}

func part1() {
	f, err := os.Open("processed.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		total += processLine(line)
	}
	fmt.Println(total)
}

func part2() {
	f, err := os.Open("processed.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// grep afterwards
		fmt.Println(processLineRotation(scanner.Text()))
	}
}

func main() {
	// part1()
	part2()
}
