package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func startsWithZeroes(s string) bool {
	if len(s) < 5 {
		return false
	}
	if s[0] == '0' && s[1] == '0' && s[2] == '0' && s[3] == '0' && s[4] == '0' {
		return true
	}
	return false
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	input := s.Text()

	pass := &strings.Builder{}
	length := 0
	for i := 0; ; i++ {
		sum := md5.Sum(append([]byte(input), strconv.Itoa(i)...))
		sumhex := fmt.Sprintf("%x", sum)
		if startsWithZeroes(sumhex) {
			pass.WriteByte(sumhex[5])
			length++
		}
		if length == 8 {
			break
		}
	}
	fmt.Println(pass.String())
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Scan()
	input := s.Text()

	pass := [8]byte{}
	length := 0
	for i := 0; ; i++ {
		sum := md5.Sum(append([]byte(input), strconv.Itoa(i)...))
		sumhex := fmt.Sprintf("%x", sum)
		if startsWithZeroes(sumhex) {
			poschar := sumhex[5]
			pos := sumhex[5] - '0'
			if poschar < '8' && poschar >= '0' && pass[pos] == 0 {
				pass[pos] = sumhex[6]
				fmt.Println(sumhex[5]-'0', sumhex[6])
				length++
			}
		}
		if length == 8 {
			break
		}
	}
	fmt.Println(string(pass[:]))
}

func main() {
	part1()
	part2()
}
