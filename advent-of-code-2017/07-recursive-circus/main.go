package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	parents := make(map[string]string, 0)
	allchildren := make(map[string][]string)
	allweights := make(map[string]int)

	for _, line := range lines {
		parts := strings.Split(line, "->")
		subparts := strings.Split(parts[0], " ")
		program := subparts[0]
		if _, ok := parents[program]; !ok {
			parents[program] = ""
		}

		weight := subparts[1][1 : len(subparts[1])-1]
		weightint, err := strconv.Atoi(weight)
		if err != nil {
			panic(err)
		}
		allweights[program] = weightint

		if len(parts) == 2 {
			children := strings.Split(parts[1], ",")
			for i := range children {
				children[i] = strings.Trim(children[i], " ")
				parents[children[i]] = program
			}
			allchildren[program] = append(allchildren[program], children...)
		}
	}

	root := ""
	for c, p := range parents {
		if p == "" {
			root = c
		}
	}
	fmt.Println(root, allweights[root])
	fmt.Println(allchildren[root])
	allstackweights := make(map[string]int)
	t := stackWeight(root, allstackweights, allweights, allchildren)
	fmt.Println(allstackweights, t)

	fmt.Println("root", root)
	unbalanced := findUnbalanced("", root, allstackweights, allweights, allchildren)
	parent := parents[unbalanced]
	diff := 0

	for _, c := range allchildren[parent] {
		if allstackweights[c] != allstackweights[unbalanced] {
			diff = allstackweights[c] - allstackweights[unbalanced]
		}
	}
	fmt.Println(allweights[unbalanced] + diff)
}

func stackWeight(program string, allstackweights map[string]int, allweights map[string]int, allchildren map[string][]string) int {
	if len(allchildren[program]) == 0 {
		allstackweights[program] = allweights[program]
		return allweights[program]
	}
	s := 0
	for _, c := range allchildren[program] {
		s += stackWeight(c, allstackweights, allweights, allchildren)
	}
	allstackweights[program] = s + allweights[program]
	return allstackweights[program]
}

func findUnbalanced(previous string, current string, allstackweights map[string]int, allweights map[string]int, allchildren map[string][]string) string {
	unbalanced := findUnbalancedChild(current, allstackweights, allchildren)
	fmt.Println("u", unbalanced, "c", current, "p", previous)
	if unbalanced == "" {
		return current
	} else {
		return findUnbalanced(current, unbalanced, allstackweights, allweights, allchildren)
	}
}

func findUnbalancedChild(current string, allstackweights map[string]int, allchildren map[string][]string) string {

	if len(allchildren[current]) == 0 {
		return ""
	}

	s := make(map[int]int)
	for _, c := range allchildren[current] {
		s[allstackweights[c]] += 1
	}
	if len(s) == 1 {
		return ""
	}

	for c := range allchildren {
		if s[allstackweights[c]] == 1 {
			return c
		}
	}

	return ""
}
