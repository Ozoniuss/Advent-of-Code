package main

import (
	"aoc/golib/twod"
	"bufio"
	"fmt"
	"os"
	"slices"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
// var inputLines = readlines()

type gameboard map[twod.Location]byte

func readlines() (twod.Location, gameboard) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	board := make(gameboard)

	s := bufio.NewScanner(f)
	i := 0
	origin := twod.ORIGIN
	for s.Scan() {
		line := s.Bytes()
		for j, c := range line {
			board[twod.Location{i, j}] = c
			if c == 'S' {
				origin = twod.Location{i, j}
			}
		}
		i++
	}
	return origin, board
}

func next(b gameboard, current twod.Location) []twod.Location {
	nextLocations := make([]twod.Location, 0, 8)
	if b[current] == '.' {
		return nil
	}
	if b[current] == '|' {
		nextLocations = append(nextLocations, twod.Move(current, twod.UP), twod.Move(current, twod.DOWN))
		return nextLocations
	}
	if b[current] == '-' {
		nextLocations = append(nextLocations, twod.Move(current, twod.LEFT), twod.Move(current, twod.RIGHT))
		return nextLocations
	}
	if b[current] == 'L' {
		nextLocations = append(nextLocations, twod.Move(current, twod.UP), twod.Move(current, twod.RIGHT))
		return nextLocations
	}
	if b[current] == 'J' {
		nextLocations = append(nextLocations, twod.Move(current, twod.UP), twod.Move(current, twod.LEFT))
		return nextLocations
	}
	if b[current] == '7' {
		nextLocations = append(nextLocations, twod.Move(current, twod.DOWN), twod.Move(current, twod.LEFT))
		return nextLocations
	}
	if b[current] == 'F' {
		nextLocations = append(nextLocations, twod.Move(current, twod.DOWN), twod.Move(current, twod.RIGHT))
		return nextLocations
	}
	if b[current] == 'S' {
		// CHEATING
		// nextLocations = append(nextLocations, twod.Move(current, twod.DOWN), twod.Move(current, twod.RIGHT))
		// return nextLocations
		nextLocations = append(nextLocations, twod.Move(current, twod.UP), twod.Move(current, twod.LEFT))
		return nextLocations
		// for _, d := range []twod.Direction{twod.UP, twod.DOWN, twod.LEFT, twod.RIGHT} {
		// 	if p := twod.Move(current, d); b[p] != '.' {
		// 		nextLocations = append(nextLocations, p)
		// 	}
		// }
		// return nextLocations
	}
	return nil
}

func findCycle(b gameboard, o twod.Location) []twod.Location {
	road := make([]twod.Location, 0, 100)
	road = append(road, o)
	visited := make(map[twod.Location]struct{})
	current := o
	visited[o] = struct{}{}
	for _, nextLocation := range next(b, current) {
		roadCopy := slices.Clone(road)
		roadCopy = append(roadCopy, nextLocation)
		finalRoad := dfs(b, nextLocation, roadCopy, visited)
		if finalRoad == nil {
			continue
		} else {
			return finalRoad
		}
	}
	panic("shouldntve happended")
}

func dfs(b gameboard, current twod.Location, road []twod.Location, visited map[twod.Location]struct{}) []twod.Location {
	if b[current] == 'S' {
		if len(road) == 3 {
			fmt.Println("here")
			return nil
		}
		return road
	}
	visited[current] = struct{}{}
	for _, nextLocation := range next(b, current) {

		// The next vertex we add is already in the road, which means we found
		// a cycle.
		if slices.Contains(road, nextLocation) && b[nextLocation] != 'S' {
			continue
		}

		// The next vertex we add is already visited, so there is no point in
		// going back.
		_, alreadyVisited := visited[nextLocation]
		if alreadyVisited && b[nextLocation] != 'S' {

			continue
		}
		roadCopy := slices.Clone(road)
		roadCopy = append(roadCopy, nextLocation)
		wholeRoad := dfs(b, nextLocation, roadCopy, visited)
		// check next neighbours
		if wholeRoad == nil {
			continue
		} else {
			return wholeRoad
		}
		// return dfs(b, nextLocation, roadCopy, visited)
	}
	return nil
}

func part1() {
	origin, brd := readlines()
	fmt.Println(origin)

	road := findCycle(brd, origin)
	fmt.Println(len(road) / 2)

	// fmt.Println(maxsteps)
}

const (
	OUT     = "OUT"
	INDOWN  = "UPWARDS"
	INUP    = "DOWNWARDS"
	OUTDOWN = "OUTDOWN"
	OUTUP   = "OUTUP"
	IN      = "IN"
)

func part2() {
	origin, brd := readlines()

	road := findCycle(brd, origin)
	roadmap := make(map[twod.Location]struct{})
	for _, tile := range road {
		roadmap[tile] = struct{}{}
	}

	// hacks
	brd[origin] = 'J'

	// we are not on the roadmap yet
	currentState := OUT

	count := 0
	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			currentLoc := twod.Location{i, j}
			// we are not on the roadmap
			if _, ok := roadmap[currentLoc]; !ok {
				if currentState == IN {
					count++
				}
				// we are on the roadmap, probably need to make changes
			} else {
				if currentState == OUT {
					if brd[currentLoc] == '|' {
						currentState = IN
						// These are the only tiles I can encounter if I'm out
					} else if brd[currentLoc] == 'L' {
						currentState = OUTDOWN
					} else if brd[currentLoc] == 'F' {
						currentState = OUTUP
					} else {
						panic("out")
					}
				} else if currentState == IN {
					if brd[currentLoc] == '|' {
						currentState = OUT
					} else if brd[currentLoc] == 'L' {
						currentState = INDOWN
					} else if brd[currentLoc] == 'F' {
						currentState = INUP
					} else {
						panic("in")
					}
				} else if currentState == INDOWN {
					if brd[currentLoc] == 'J' {
						currentState = IN
					} else if brd[currentLoc] == '7' {
						currentState = OUT
					} else if brd[currentLoc] != '-' {
						panic("indown")
					}
				} else if currentState == INUP {
					if brd[currentLoc] == 'J' {
						currentState = OUT
					} else if brd[currentLoc] == '7' {
						currentState = IN
					} else if brd[currentLoc] != '-' {
						panic("indown")
					}
				} else if currentState == OUTDOWN {
					if brd[currentLoc] == 'J' {
						currentState = OUT
					} else if brd[currentLoc] == '7' {
						currentState = IN
					} else if brd[currentLoc] != '-' {
						panic("indown")
					}
				} else if currentState == OUTUP {
					if brd[currentLoc] == 'J' {
						currentState = IN
					} else if brd[currentLoc] == '7' {
						currentState = OUT
					} else if brd[currentLoc] != '-' {
						panic("indown")
					}
				}
			}
		}
	}
	fmt.Println(count)
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

	// part1()
	part2()
}
