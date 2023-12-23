package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
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

type Location = [3]int
type Brick struct {
	id  int
	loc Location
}

func fillBrickPositions(left, right Location, id int, brickMap map[Location]int, brickPos map[int][]Location) {
	locations := []Location{}
	// just one block
	if left == right {
		brickMap[left] = id
		brickPos[id] = []Location{left}
		return
	}
	if left[0] != right[0] {
		if left[0] < right[0] {
			for i := left[0]; i <= right[0]; i++ {
				loc := Location{i, left[1], left[2]}
				brickMap[loc] = id
				locations = append(locations, Location{i, left[1], left[2]})
			}
		}
		if left[0] > right[0] {
			for i := left[0]; i >= right[0]; i-- {
				loc := Location{i, left[1], left[2]}
				brickMap[loc] = id
				locations = append(locations, Location{i, left[1], left[2]})
			}
		}
	}
	if left[1] != right[1] {
		if left[1] < right[1] {
			for i := left[1]; i <= right[1]; i++ {
				loc := Location{left[0], i, left[2]}
				brickMap[loc] = id
				locations = append(locations, Location{left[0], i, left[2]})

			}
		}
		if left[1] > right[1] {
			for i := left[1]; i >= right[1]; i-- {
				loc := Location{left[0], i, left[2]}
				brickMap[loc] = id
				locations = append(locations, Location{left[0], i, left[2]})
			}
		}
	}
	if left[2] != right[2] {
		if left[2] < right[2] {
			for i := left[2]; i <= right[2]; i++ {
				loc := Location{left[0], left[1], i}
				brickMap[loc] = id
				locations = append(locations, Location{left[0], left[1], i})
			}
		}
		if left[2] > right[2] {
			for i := left[2]; i >= right[2]; i-- {
				loc := Location{left[0], left[1], i}
				brickMap[loc] = id
				locations = append(locations, Location{left[0], left[1], i})
			}
		}
	}
	brickPos[id] = locations
}

func readBrick(line string, id int, brickMap map[Location]int, brickPos map[int][]Location) {

	var left, right Location

	parts := strings.Split(line, "~")
	leftCoordsStr := strings.Split(parts[0], ",")
	for i, c := range leftCoordsStr {
		left[i], _ = strconv.Atoi(c)
	}

	rightCoordsStr := strings.Split(parts[1], ",")
	for i, c := range rightCoordsStr {
		right[i], _ = strconv.Atoi(c)
	}

	fillBrickPositions(left, right, id, brickMap, brickPos)
}

func findOrientation(id int, brickPos map[int][]Location) string {
	if len(brickPos[id]) == 0 {
		panic("lol")
	}
	if len(brickPos[id]) == 1 {
		return "x"
	}

	b1 := brickPos[id][0]
	b2 := brickPos[id][1]

	if b1[0] != b2[0] {
		return "x"
	}
	if b1[1] != b2[1] {
		return "y"
	}
	if b1[2] != b2[2] {
		return "z"
	}
	panic("orientation")
}

func findMinimumZ(loc []Location) int {
	if len(loc) == 0 {
		panic("loc should not be empty")
	}
	minz := 999999
	for _, l := range loc {
		if l[2] < minz {
			minz = l[2]
		}
	}
	return minz
}

func inStructureOrGround(loc Location, brickMap map[Location]int) bool {
	if loc[2] <= 0 {
		return true
	}
	if _, ok := brickMap[loc]; ok {
		return true
	}
	return false
}

func lowerCoordinates(id int, brickMap map[Location]int, brickPos map[int][]Location) {
	// This is not safe unless checked!!!!

	// Nothing below was found, move the brick down.
	toAppend := []Location{}
	for _, loc := range brickPos[id] {
		// Remove every location from the original map, and append the lower
		// one
		delete(brickMap, loc)

		// Lower location
		locc := loc
		locc[2] = loc[2] - 1
		toAppend = append(toAppend, locc)
	}
	// Add the new locations
	for _, loc := range toAppend {
		brickMap[loc] = id
	}

	// Do the same for positions. Remove the old brick reference, and
	// add the new one.
	delete(brickPos, id)
	brickPos[id] = toAppend
}

func moveBrickDownOnce(id int, brickMap map[Location]int, brickPos map[int][]Location) bool {
	orientation := findOrientation(id, brickPos)
	if orientation == "z" {
		minz := findMinimumZ(brickPos[id])
		// can get any x or y coordinate from the location since they are
		// stacked up
		xcoord := brickPos[id][0][0]
		ycoord := brickPos[id][0][1]
		nextDown := Location{xcoord, ycoord, minz - 1}

		// There is something below
		if inStructureOrGround(nextDown, brickMap) {
			return false
		}

		// Move the brick down
		lowerCoordinates(id, brickMap, brickPos)
		return true
	} else {

		// Check if there is anything below all places constituting the brick
		for _, loc := range brickPos[id] {
			nextDown := loc
			nextDown[2] = loc[2] - 1
			// There is something down in this case
			if inStructureOrGround(nextDown, brickMap) {
				return false
			}
		}

		// Move the brick down
		lowerCoordinates(id, brickMap, brickPos)
		return true
	}
	panic("wtf")
}

func findSupportingMap(brickMap map[Location]int, brickPos map[int][]Location) (map[int][]int, map[int][]int) {
	supporting := make(map[int][]int)
	supportedBy := make(map[int][]int)
	for id := range brickPos {
		supporting[id] = make([]int, 0)
		supportedBy[id] = make([]int, 0)
	}

	for id, locations := range brickPos {
		// Get the upper location of each loc, and add the bricks that the
		// current one is supporting
		alreadyAddedUp := make(map[int]struct{})
		for _, loc := range locations {
			locc := loc
			locc[2] += 1
			upperId, ok := brickMap[locc]
			if !ok {
				continue
			}
			// found a different brick, add it and remember it.
			if upperId != id {
				if _, ok := alreadyAddedUp[upperId]; !ok {
					supporting[id] = append(supporting[id], upperId)
					alreadyAddedUp[upperId] = struct{}{}
				}
			}
		}

		alreadyAddedDown := make(map[int]struct{})
		for _, loc := range locations {
			locc := loc
			locc[2] -= 1
			lowerId, ok := brickMap[locc]
			if !ok {
				continue
			}
			// found a different brick, add it and remember it.
			if lowerId != id {
				if _, ok := alreadyAddedDown[lowerId]; !ok {
					supportedBy[id] = append(supportedBy[id], lowerId)
					alreadyAddedDown[lowerId] = struct{}{}
				}
			}
		}
	}
	return supporting, supportedBy

}

func part1() {
	brickMap := make(map[Location]int)
	brickPos := make(map[int][]Location)

	for idx, line := range inputLines {
		readBrick(line, idx, brickMap, brickPos)
	}
	// fmt.Println(brickPos)

	for bid, l := range brickPos {
		if len(l) == 0 {
			panic(bid)
		}
	}

	allBrickIds := maps.Keys(brickPos)
	slices.SortFunc(allBrickIds, func(a, b int) int {
		locs1, ok := brickPos[allBrickIds[a]]
		if !ok {
			panic("sort")
		}
		locs2, ok := brickPos[allBrickIds[b]]
		if !ok {
			panic("sort")
		}
		minz1 := findMinimumZ(locs1)
		minz2 := findMinimumZ(locs2)
		return minz1 - minz2
	})

	for {
		hasMovedAnything := false
		// go through all bricks
		for _, bid := range allBrickIds {
			// move the brick as down as you can
			for {
				hasFallen := moveBrickDownOnce(bid, brickMap, brickPos)
				if !hasFallen {
					break
				} else {
					// tell that the brick has moved
					hasMovedAnything = true
				}
			}
		}
		if !hasMovedAnything {
			break
		}
	}
	// fmt.Println(brickPos)

	supporting, supportedBy := findSupportingMap(brickMap, brickPos)
	// fmt.Println("supporting", supporting)
	// fmt.Println("supported by", supportedBy)
	canDisintegrade := make(map[int]struct{})
	for bid := range brickPos {
		// The bricks that it is supporting
		over := supporting[bid]

		// The idea is that each of the bricks that this brick is supporting,
		// must be supported by at least another brick.
		canBeRemoved := true
		for _, ov := range over {
			// the brick over is supported by at least two bricks
			if len(supportedBy[ov]) == 1 {
				canBeRemoved = false
			}
		}
		if canBeRemoved {
			canDisintegrade[bid] = struct{}{}
		}
	}

	fmt.Println(len(brickMap))
	fmt.Println(len(brickPos))

	count := 0
	for disintegradedId := range brickPos {

		// Copy the fucking maps
		brickMapsCopy := maps.Clone(brickMap)
		brickPosCopy := make(map[int][]Location)
		for id, loc := range brickPos {
			brickPosCopy[id] = slices.Clone(loc)
		}

		locs := brickPosCopy[disintegradedId]
		// Delete the brick
		delete(brickPosCopy, disintegradedId)
		for _, loc := range locs {
			delete(brickMapsCopy, loc)
		}

		bricksMovedDown := make(map[int]struct{})
		for {
			hasMovedAnything := false
			// go through all bricks
			for _, bid := range allBrickIds {
				// Just ignore the one that got disintegrated.
				if bid == disintegradedId {
					continue
				}
				// move the brick as down as you can
				for {
					hasFallen := moveBrickDownOnce(bid, brickMapsCopy, brickPosCopy)
					if !hasFallen {
						break
					} else {
						// tell that the brick has moved
						hasMovedAnything = true
						bricksMovedDown[bid] = struct{}{}
					}
				}
			}
			if !hasMovedAnything {
				break
			}
		}
		count += len(bricksMovedDown)
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

	part1()
	// part2()
}
