package main

import (
	"aoc/golib/twod"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	OUT     = "OUT"
	INDOWN  = "UPWARDS"
	INUP    = "DOWNWARDS"
	OUTDOWN = "OUTDOWN"
	OUTUP   = "OUTUP"
	IN      = "IN"
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

func processInstruction(instruction string) (string, int) {
	distancestrhex := instruction[2:7]
	direncoded := instruction[7]

	// distancestrhex = "0x" + distancestrhex
	steps64, _ := strconv.ParseInt(distancestrhex, 16, 64)

	var dirstr string
	var steps = int(steps64)
	if direncoded == '0' {
		dirstr = "R"
	}
	if direncoded == '1' {
		dirstr = "D"
	}
	if direncoded == '2' {
		dirstr = "L"
	}
	if direncoded == '3' {
		dirstr = "U"
	}

	return dirstr, steps
}

const (
	TOP   = "TOP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

func isOutsideBorder(direction twod.Direction, wherePolygon string) bool {
	if wherePolygon == "right" {
		if direction == twod.RIGHT {
			return true
		} else if direction == twod.DOWN {
			return true
		} else {
			return false
		}
	}
	if wherePolygon == "left" {
		if direction == twod.LEFT {
			return true
		} else if direction == twod.UP {
			return true
		} else {
			return false
		}
	}
	panic("lol")
}

func fillCoordinate(current twod.Location, state string, dirstr string, steps int, coordiantes *[]twod.Location, outercount *int, wherePolygon string) twod.Location {
	dir := twod.DirFromString(dirstr)
	curr := current
	for i := 0; i < steps; i++ {
		curr = twod.Move(curr, dir)
	}
	if isOutsideBorder(dir, wherePolygon) {
		*outercount += steps
	}
	(*coordiantes) = append((*coordiantes), curr)
	return curr
}

func advance(current twod.Location, dirstr string, steps int, board map[twod.Location]struct{}) twod.Location {
	dir := twod.DirFromString(dirstr)
	curr := current
	for i := 0; i < steps; i++ {
		curr = twod.Move(curr, dir)
		board[curr] = struct{}{}
	}
	return curr
}

func polygonArea(coords []twod.Location) int {
	S := 0
	n := len(coords)
	for i := 0; i < n; i++ {
		S += coords[i%n][0]*coords[(i+1)%n][1] - coords[(i+1)%n][0]*coords[i%n][1]
	}
	return S / 2
}

func polygonAreaExt(coords []twod.Location) int {
	S := 0
	n := len(coords)
	for i := 0; i < n; i++ {
		S += (coords[i%n][0]+1)*(coords[(i+1)%n][1]) - (coords[(i+1)%n][0]+1)*(coords[i%n][1])
	}
	return -S / 2
}

func part2() {
	coordinates := make([]twod.Location, 0)

	current := twod.ORIGIN
	// FILL STATE BASED ON NEEDS
	state := INDOWN
	coordinates = append(coordinates, twod.ORIGIN)
	outercount := 0
	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		dir, steps := processInstruction(parts[2])
		// dirstr := parts[0]
		// steps, _ := strconv.Atoi(parts[1])
		// fmt.Println(dir, steps)
		current = fillCoordinate(current, state, dir, steps, &coordinates, &outercount, "right")
	}
	// fmt.Println(coordinates)
	fmt.Println(polygonArea(coordinates), outercount, -polygonArea(coordinates)+outercount+1)
	//fmt.Println(polygonArea(coordinates) + outercount)

	// minr := -9990
	// minc := -9990
	// maxr := 9990
	// maxc := 9990

	// const (
	//  OUT     = "OUT"
	//  INDOWN  = "UPWARDS"
	//  INUP    = "DOWNWARDS"
	//  OUTDOWN = "OUTDOWN"
	//  OUTUP   = "OUTUP"
	//  IN      = "IN"
	// )
	// state := OUT
	// in := false

	// count := 0
	// fmt.Println(len(board))
	// for row := minr; row < maxr; row++ {
	//  for col := minc; col < maxc; col++ {

	//      _, exists := board[twod.Location{row, col}]
	//      _, existsup := board[twod.Location{row - 1, col}]
	//      _, existsdown := board[twod.Location{row + 1, col}]

	//      if !exists && state == IN {
	//          // board[twod.Location{row, col}] = struct{}{}
	//          count++
	//      }

	//      if state == OUT {
	//          if exists && existsup && existsdown {
	//              state = IN

	//          } else if exists && existsup && !existsdown {
	//              state = OUTDOWN
	//          } else if exists && !existsup && existsdown {
	//              state = OUTUP
	//          }
	//      } else if state == IN {
	//          if exists && existsup && existsdown {
	//              state = OUT
	//          } else if exists && existsup && !existsdown {
	//              state = INDOWN
	//          } else if exists && !existsup && existsdown {
	//              state = INUP
	//          }
	//      } else if state == INDOWN {
	//          if exists && existsup && !existsdown {
	//              state = IN
	//          } else if exists && !existsup && existsdown {
	//              state = OUT
	//          }
	//      } else if state == INUP {
	//          if exists && existsup && !existsdown {
	//              state = OUT
	//          } else if exists && !existsup && existsdown {
	//              state = IN
	//          }
	//      } else if state == OUTDOWN {
	//          if exists && existsup && !existsdown {
	//              state = OUT
	//          } else if exists && !existsup && existsdown {
	//              state = IN
	//          }
	//      } else if state == OUTUP {
	//          if exists && existsup && !existsdown {
	//              state = IN
	//          } else if exists && !existsup && existsdown {
	//              state = OUT
	//          }
	//      }
	//  }
	// }
	// fmt.Println(len(board) + count)
}

func part1() {
	board := make(map[twod.Location]struct{})

	current := twod.ORIGIN
	board[current] = struct{}{}
	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		dir := parts[0]
		steps, _ := strconv.Atoi(parts[1])
		current = advance(current, dir, steps, board)
	}
	// fmt.Println(board)

	minr := -9990
	minc := -9990
	maxr := 9990
	maxc := 9990

	const (
		OUT     = "OUT"
		INDOWN  = "UPWARDS"
		INUP    = "DOWNWARDS"
		OUTDOWN = "OUTDOWN"
		OUTUP   = "OUTUP"
		IN      = "IN"
	)
	state := OUT
	// in := false

	count := 0
	fmt.Println(len(board))
	for row := minr; row < maxr; row++ {
		for col := minc; col < maxc; col++ {

			_, exists := board[twod.Location{row, col}]
			_, existsup := board[twod.Location{row - 1, col}]
			_, existsdown := board[twod.Location{row + 1, col}]

			if !exists && state == IN {
				// board[twod.Location{row, col}] = struct{}{}
				count++
			}

			if state == OUT {
				if exists && existsup && existsdown {
					state = IN

				} else if exists && existsup && !existsdown {
					state = OUTDOWN
				} else if exists && !existsup && existsdown {
					state = OUTUP
				}
			} else if state == IN {
				if exists && existsup && existsdown {
					state = OUT
				} else if exists && existsup && !existsdown {
					state = INDOWN
				} else if exists && !existsup && existsdown {
					state = INUP
				}
			} else if state == INDOWN {
				if exists && existsup && !existsdown {
					state = IN
				} else if exists && !existsup && existsdown {
					state = OUT
				}
			} else if state == INUP {
				if exists && existsup && !existsdown {
					state = OUT
				} else if exists && !existsup && existsdown {
					state = IN
				}
			} else if state == OUTDOWN {
				if exists && existsup && !existsdown {
					state = OUT
				} else if exists && !existsup && existsdown {
					state = IN
				}
			} else if state == OUTUP {
				if exists && existsup && !existsdown {
					state = IN
				} else if exists && !existsup && existsdown {
					state = OUT
				}
			}
		}
	}
	fmt.Println(len(board) + count)
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
