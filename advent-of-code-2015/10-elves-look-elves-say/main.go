package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// lookAndSay applies the look and say rules ones.
func lookAndSay(input string) string {
	inputWithEnd := input + "x"

	var output string
	b := &strings.Builder{}
	length := 1
	for i := 1; i < len(inputWithEnd); i++ {
		if inputWithEnd[i] != inputWithEnd[i-1] {
			toWrite := inputWithEnd[i-1]
			b.WriteString(strconv.Itoa(length))
			b.WriteByte(toWrite)
			length = 1
		} else {
			length++
		}
	}
	output = b.String()
	return output
}

// lookAndSay applies the look and say rules ones.
func lookAndSayWithoutBuilder(input string) string {
	inputWithEnd := input + "x"

	var output string
	length := 1
	for i := 1; i < len(inputWithEnd); i++ {
		if inputWithEnd[i] != inputWithEnd[i-1] {
			toWrite := inputWithEnd[i-1]
			output += strconv.Itoa(length)
			output += string(toWrite)
			length = 1
		} else {
			length++
		}
	}
	return output
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	input := scanner.Text()
	out := input
	for i := 0; i < 40; i++ {
		out = lookAndSayWithoutBuilder(out)
	}
	fmt.Println(len(out))
}
