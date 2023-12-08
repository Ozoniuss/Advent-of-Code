package main

import (
	"aoc/golib/maths"
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/pkg/profile"
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

// Doesn't really reduce allocations because maps are allocated on the heap
// anyways since they are big. But, it does reduce overall memory usage.
type node [3]byte

var AAA = node{'A', 'A', 'A'}
var ZZZ = node{'Z', 'Z', 'Z'}

func part1() int {
	directions := inputLines[0]

	locations := make(map[node][2]node, 1024)

	for _, line := range inputLines[2:] {
		var start, left, right [3]byte
		copy(start[:], line[0:3])
		copy(left[:], line[7:10])
		copy(right[:], line[12:15])

		locations[start] = [2]node{left, right}
	}
	current := [3]byte{'A', 'A', 'A'}

	start := -1
	count := 0
	for {
		start++
		start = start % len(directions)
		count++

		if directions[start] == 'L' {
			current = locations[current][0]
		} else {
			current = locations[current][1]
		}
		if current == ZZZ {
			break
		}
	}
	return count
}

func part2() int {
	directions := inputLines[0]

	locations := make(map[node][2]node, 1024)

	for _, line := range inputLines[2:] {
		var start, left, right [3]byte
		copy(start[:], line[0:3])
		copy(left[:], line[7:10])
		copy(right[:], line[12:15])

		locations[start] = [2]node{left, right}
	}
	currents := []node{}

	for k := range locations {
		if k[2] == 'A' {
			currents = append(currents, k)
		}
	}

	rotations := make([]int, len(currents))
	rotationsmtx := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(len(currents))
	for i := 0; i < len(currents); i++ {
		start := -1
		count := 0
		go func(i, start, count int) {
			for {

				start++
				start = start % len(directions)
				count++

				if directions[start] == 'L' {
					currents[i] = locations[currents[i]][0]
				} else {
					currents[i] = locations[currents[i]][1]
				}
				if currents[i][2] == 'Z' {
					break
				}
			}
			rotationsmtx.Lock()
			rotations[i] = count
			rotationsmtx.Unlock()
			wg.Done()
		}(i, start, count)
	}
	wg.Wait()
	return maths.LCM(rotations...)
}

func main() {
	// Run only 1 profile at a time!
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.MemProfile, profile.ProfilePath("."), profile.MemProfileRate(1)).Stop()

	// No return value improves speed during contest, but make these functions
	// return something for benchmarks once problem is solved.
	//
	// Part 2 is not written above and commented below so that template compiles
	// while solving part 1.

	fmt.Println(part1())
	fmt.Println(part2())
}
