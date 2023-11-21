package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func findDistance(path []string, distances map[[2]string]int) int {
	total := 0
	for i := 1; i < len(path); i++ {
		total += distances[[2]string{path[i], path[i-1]}]
	}
	return total
}

// graph is complete
func findAllPaths(cities map[string]struct{}) [][]string {
	all := [][]string{}
	var q = make([][]string, 0, len(cities))
	for c := range cities {
		q = append(q, []string{c})
	}
	for len(q) != 0 {
		current := q[0]
		q = q[1:]
		for nextCity := range cities {
			if !slices.Contains(current, nextCity) {
				// wtf?
				next := slices.Clone(append(current, nextCity))
				if len(next) == len(cities) {
					all = append(all, next)
				} else {
					q = append(q, next)
				}
			}
		}
	}
	return all
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cities := make(map[string]struct{})
	distances := make(map[[2]string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		content := strings.Split(line, " ")
		city1, city2, distanceStr := content[0], content[2], content[4]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}
		cities[city1] = struct{}{}
		cities[city2] = struct{}{}
		distances[[2]string{city1, city2}] = distance
		distances[[2]string{city2, city1}] = distance
	}

	smallest := 999999999999
	largest := -1
	allPaths := findAllPaths(cities)
	for _, path := range allPaths {
		length := findDistance(path, distances)
		if length < smallest {
			// fmt.Println(path, length, i)
			smallest = length
		}
		if length > largest {
			largest = length
		}
	}
	fmt.Println(smallest)
	fmt.Println(largest)

}
