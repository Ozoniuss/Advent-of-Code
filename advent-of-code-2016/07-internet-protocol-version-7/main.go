package main

import (
	"bufio"
	"fmt"
	"os"
)

func isMirroed(word [4]byte) bool {
	return word[0] == word[3] && word[1] == word[2] && word[0] != word[1]
}

func isPalindrome(word [3]byte) bool {
	return word[0] == word[2] && word[0] != word[1]
}

func supportsTLS(ip string) bool {
	insideSquares := false
	found := false
	for i := 3; i < len(ip); i++ {
		if ip[i] == '[' {
			insideSquares = true
			continue
		}
		if ip[i] == ']' {
			insideSquares = false
			continue
		}
		word := [4]byte{ip[i-3], ip[i-2], ip[i-1], ip[i]}
		if insideSquares && isMirroed(word) {
			return false
		} else if !insideSquares && isMirroed(word) {
			found = true
		}
	}
	return found
}

func supportsSSL(ip string) bool {
	insideSquares := false
	matchesInside := make(map[[3]byte]struct{}, 64)
	matchesOutside := make(map[[3]byte]struct{}, 64)
	for i := 2; i < len(ip); i++ {
		if ip[i] == '[' {
			insideSquares = true
			continue
		}
		if ip[i] == ']' {
			insideSquares = false
			continue
		}
		word := [3]byte{ip[i-2], ip[i-1], ip[i]}
		if insideSquares && isPalindrome(word) {
			matchesInside[word] = struct{}{}
		} else if !insideSquares && isPalindrome(word) {
			matchesOutside[word] = struct{}{}
		}
	}
	for k := range matchesInside {
		if _, ok := matchesOutside[[3]byte{k[1], k[0], k[1]}]; ok {
			return true
		}
	}
	return false
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
		if supportsTLS(line) {
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

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		if supportsSSL(line) {
			count++
		}
	}
	fmt.Println(count)
}

func main() {
	part1()
	part2()
}
