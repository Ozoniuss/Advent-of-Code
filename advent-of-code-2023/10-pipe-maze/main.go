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

type Board map[twod.Location]byte

func readlines() (twod.Location, Board) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	board := make(Board)

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

// Not defining consts for the tiles because it's easier to visualize this
// way.

// next returns a tile that could potentially continue the current road.
func next(b Board, current twod.Location) []twod.Location {
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
		uptile := twod.Move(current, twod.UP)
		downtile := twod.Move(current, twod.DOWN)
		lefttile := twod.Move(current, twod.LEFT)
		righttile := twod.Move(current, twod.RIGHT)
		if b[uptile] == '|' || b[uptile] == 'F' || b[uptile] == '7' {
			nextLocations = append(nextLocations, uptile)
		}
		if b[downtile] == '|' || b[downtile] == 'J' || b[downtile] == 'L' {
			nextLocations = append(nextLocations, downtile)
		}
		if b[lefttile] == '-' || b[lefttile] == 'F' || b[lefttile] == 'L' {
			nextLocations = append(nextLocations, lefttile)
		}
		if b[righttile] == '-' || b[righttile] == 'J' || b[righttile] == '7' {
			nextLocations = append(nextLocations, righttile)
		}
		return nextLocations
	}
	return nil
}

func findCyclicRoad(b Board, o twod.Location) []twod.Location {
	road := make([]twod.Location, 0, 100)
	road = append(road, o)
	visited := make(map[twod.Location]struct{})
	current := o
	visited[o] = struct{}{}
	for _, nextLocation := range next(b, current) {
		road = append(road, nextLocation)
		finalRoad := dfsRoad(b, nextLocation, &road, visited)
		if !finalRoad {
			road = road[:len(road)-1]
			continue
		} else {
			return road
		}
	}
	panic("no cycle detected from start")
}

func dfsRoad(b Board, current twod.Location, road *[]twod.Location, visited map[twod.Location]struct{}) bool {
	if b[current] == 'S' {
		return len(*road) != 3
	}
	visited[current] = struct{}{}
	for _, nextLocation := range next(b, current) {

		// The next vertex we add is already in the road, which means we found
		// a cycle.
		if slices.Contains(*road, nextLocation) && b[nextLocation] != 'S' {
			continue
		}

		// The next vertex we add is already visited, so there is no point in
		// going back.
		_, alreadyVisited := visited[nextLocation]
		if alreadyVisited && b[nextLocation] != 'S' {

			continue
		}
		(*road) = append((*road), nextLocation)
		foundCycle := dfsRoad(b, nextLocation, road, visited)
		// check next neighbours
		if !foundCycle {
			(*road) = (*road)[:len(*road)-1]
			continue
		} else {
			return true
		}
		// remove last element
	}
	return false
}

func part1() int {
	origin, brd := readlines()

	road := findCyclicRoad(brd, origin)
	return len(road) / 2
}

const (
	OUT     = "OUT"
	INDOWN  = "UPWARDS"
	INUP    = "DOWNWARDS"
	OUTDOWN = "OUTDOWN"
	OUTUP   = "OUTUP"
	IN      = "IN"
)

func part2(missingPiece byte) int {
	origin, brd := readlines()

	road := findCyclicRoad(brd, origin)
	roadmap := make(map[twod.Location]struct{})
	for _, tile := range road {
		roadmap[tile] = struct{}{}
	}

	// hack. Honestly figuring out which part is missing is a pain in the ass
	// and is not really part of the core algorithm. I'd rather "cheat" by
	// specifying the missing piece than trying to find a way to compute it
	// programatically.
	brd[origin] = missingPiece

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
	return count
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

	fmt.Println(part1())
	fmt.Println(part2('J'))
}
