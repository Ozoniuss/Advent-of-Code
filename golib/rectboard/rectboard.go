package rectboard

import (
	"aoc/golib/twod"
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

type RectBoard = [][]byte

func ReadBoardFromFile(inputfile string) *RectBoard {
	f, err := os.Open(inputfile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines [][]byte
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, slices.Clone(s.Bytes()))
	}
	return &lines
}

func ReadBoard(r io.Reader) *RectBoard {
	var lines [][]byte
	s := bufio.NewScanner(r)
	for s.Scan() {
		lines = append(lines, slices.Clone(s.Bytes()))
	}
	return &lines
}

func PrintBoard(r *RectBoard) {
	fmt.Println(AsString(r))
}

func AsString(r *RectBoard) string {
	rb := (*r)
	sb := &strings.Builder{}
	for i := 0; i < len(rb); i++ {
		for j := 0; j < len(rb[0]); j++ {
			sb.WriteByte(rb[i][j])
		}
		if i != len(rb)-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}

func IsInBoard(loc twod.Location, r *RectBoard) bool {
	rx := *r
	if loc[0] < 0 || loc[0] >= len(rx) {
		return false
	}
	if loc[1] < 0 || loc[1] >= len(rx[0]) {
		return false
	}
	return true
}
