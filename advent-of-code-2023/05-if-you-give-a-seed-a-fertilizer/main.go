package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
var inputLines = readlines()

func readlines() []string {
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
	return lines
}

func addRangeToArr(dest, source, length int, arr *[][3]int) {
	if arr == nil {
		return
	}
	(*arr) = append((*arr), [3]int{dest, source, length})
}

func findCorrespondentInArray(arr [][3]int, number int) int {
	pos := 0
	end := true
	for i := 0; i < len(arr); i++ {
		if arr[i][1] > number {
			pos = i - 1
			end = false
			break
		}
	}
	if end {
		pos = len(arr) - 1
	}
	if pos == -1 {
		return number
	}
	if number >= arr[pos][1]+arr[pos][2] {
		return number
	}
	return number - arr[pos][1] + arr[pos][0]
}

// Took a few minutes to run
func part2() {
	var seeds []int
	var seedtosoil = make([][3]int, 0, 100)
	var soiltofert = make([][3]int, 0, 100)
	var ferttowater = make([][3]int, 0, 100)
	var watertolight = make([][3]int, 0, 100)
	var lighttotemp = make([][3]int, 0, 100)
	var temptohum = make([][3]int, 0, 100)
	var humtoloc = make([][3]int, 0, 100)

	var currentArr *[][3]int
	inmap := false
	for idx, line := range inputLines {
		if idx == 0 {
			_, seedsstr, _ := strings.Cut(line, ": ")
			seedNumsstr := strings.Split(seedsstr, " ")
			for _, seedstr := range seedNumsstr {
				seed, _ := strconv.Atoi(seedstr)
				seeds = append(seeds, seed)
			}
		}

		if len(line) == 0 {
			currentArr = nil
			inmap = false
		}
		if inmap {
			ranges := strings.Split(line, " ")
			dest, _ := strconv.Atoi(ranges[0])
			source, _ := strconv.Atoi(ranges[1])
			length, _ := strconv.Atoi(ranges[2])
			addRangeToArr(dest, source, length, currentArr)
		}

		if line == "seed-to-soil map:" {
			fmt.Println("here")
			inmap = true
			currentArr = &seedtosoil
		}
		if line == "soil-to-fertilizer map:" {

			inmap = true
			currentArr = &soiltofert
		}
		if line == "fertilizer-to-water map:" {
			inmap = true
			currentArr = &ferttowater
		}
		if line == "water-to-light map:" {
			inmap = true
			currentArr = &watertolight
		}
		if line == "light-to-temperature map:" {
			inmap = true
			currentArr = &lighttotemp
		}
		if line == "temperature-to-humidity map:" {
			inmap = true
			currentArr = &temptohum
		}
		if line == "humidity-to-location map:" {
			inmap = true
			currentArr = &humtoloc
		}
	}
	sort.Slice(seedtosoil, func(i, j int) bool {
		return seedtosoil[i][1] < seedtosoil[j][1]
	})
	sort.Slice(soiltofert, func(i, j int) bool {
		return soiltofert[i][1] < soiltofert[j][1]
	})
	sort.Slice(ferttowater, func(i, j int) bool {
		return ferttowater[i][1] < ferttowater[j][1]
	})
	sort.Slice(watertolight, func(i, j int) bool {
		return watertolight[i][1] < watertolight[j][1]
	})
	sort.Slice(lighttotemp, func(i, j int) bool {
		return lighttotemp[i][1] < lighttotemp[j][1]
	})
	sort.Slice(temptohum, func(i, j int) bool {
		return temptohum[i][1] < temptohum[j][1]
	})
	sort.Slice(humtoloc, func(i, j int) bool {
		return humtoloc[i][1] < humtoloc[j][1]
	})
	fmt.Println(seeds)
	all := [][][3]int{
		seedtosoil,
		soiltofert,
		ferttowater,
		watertolight,
		lighttotemp,
		temptohum,
		humtoloc,
	}
	// fmt.Println(all)
	lowest := math.MaxInt
	for i := 0; i < len(seeds); i += 2 {
		seed := seeds[i]
		length := seeds[i+1]
		fmt.Println(i)
		for currentSeed := seed; currentSeed < seed+length; currentSeed++ {
			curr := currentSeed
			for _, a := range all {
				curr = findCorrespondentInArray(a, curr)
			}
			if curr < lowest {
				lowest = curr
			}
		}
	}

	fmt.Println(lowest)
}

func part1() {
	var seeds []int
	var seedtosoil = make([][3]int, 0, 100)
	var soiltofert = make([][3]int, 0, 100)
	var ferttowater = make([][3]int, 0, 100)
	var watertolight = make([][3]int, 0, 100)
	var lighttotemp = make([][3]int, 0, 100)
	var temptohum = make([][3]int, 0, 100)
	var humtoloc = make([][3]int, 0, 100)

	var currentArr *[][3]int
	inmap := false
	for idx, line := range inputLines {
		if idx == 0 {
			_, seedsstr, _ := strings.Cut(line, ": ")
			seedNumsstr := strings.Split(seedsstr, " ")
			for _, seedstr := range seedNumsstr {
				seed, _ := strconv.Atoi(seedstr)
				seeds = append(seeds, seed)
			}
		}

		if len(line) == 0 {
			currentArr = nil
			inmap = false
		}
		if inmap {
			ranges := strings.Split(line, " ")
			dest, _ := strconv.Atoi(ranges[0])
			source, _ := strconv.Atoi(ranges[1])
			length, _ := strconv.Atoi(ranges[2])
			addRangeToArr(dest, source, length, currentArr)
		}

		if line == "seed-to-soil map:" {
			fmt.Println("here")
			inmap = true
			currentArr = &seedtosoil
		}
		if line == "soil-to-fertilizer map:" {

			inmap = true
			currentArr = &soiltofert
		}
		if line == "fertilizer-to-water map:" {
			inmap = true
			currentArr = &ferttowater
		}
		if line == "water-to-light map:" {
			inmap = true
			currentArr = &watertolight
		}
		if line == "light-to-temperature map:" {
			inmap = true
			currentArr = &lighttotemp
		}
		if line == "temperature-to-humidity map:" {
			inmap = true
			currentArr = &temptohum
		}
		if line == "humidity-to-location map:" {
			inmap = true
			currentArr = &humtoloc
		}
	}
	fmt.Println(seeds)
	all := [][][3]int{
		seedtosoil,
		soiltofert,
		ferttowater,
		watertolight,
		lighttotemp,
		temptohum,
		humtoloc,
	}
	// fmt.Println(all)
	lowest := math.MaxInt
	for _, s := range seeds {
		curr := s
		for _, a := range all {
			curr = findCorrespondentInArray(a, curr)
		}
		if curr < lowest {
			lowest = curr
		}
	}
	fmt.Println(lowest)
}

func main() {
	// Run only 1 profile at a time!
	// defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()

	// No return value improves speed during contest, but make these functions
	// return something for benchmarks once problem is solved.
	//
	// Part 2 is not written above and commented below so that template compiles
	// while solving part 1.

	part1()
	part2()
}
