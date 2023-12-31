package twod

import "aoc/golib/maths"

// Location represents the two-dimensional integer coordinates of a point in
// the place.
type Location [2]int

// Direction consists of two values, representing the values added to each
// coordinate when moving to that direction. It can be thought of as a vector
// in the plane.
type Direction [2]int

// ManhattanDistance returns the Manhattan distance between two locations.
func ManhattanDistance(l1, l2 Location) int {
	return maths.Abs(l1[0]-l2[0]) + maths.Abs(l1[1]-l2[1])
}

// Move takes an initial location and a direction, and returns the new location
// after moving in that direction.
func Move(l Location, d Direction) Location {
	return Location{l[0] + d[0], l[1] + d[1]}
}

// AddDirs adds two directions by individually adding the direction's horizontal
// and vertical component. It works the same way as adding two vectors.
func AddDirs(d1, d2 Direction) Direction {
	return Direction{d1[0] + d2[0], d1[1] + d2[1]}
}

// DirFromRune returns a basic direction from a rune representation.
func DirFromRune(dir rune) Direction {
	if dir == 'L' {
		return LEFT
	}
	if dir == 'R' {
		return RIGHT
	}
	if dir == 'D' {
		return DOWN
	}
	if dir == 'U' {
		return UP
	}
	return NULL
}

// DirFromChar returns a basic direction from a char representation.
func DirFromChar(dir byte) Direction {
	if dir == 'L' {
		return LEFT
	}
	if dir == 'R' {
		return RIGHT
	}
	if dir == 'D' {
		return DOWN
	}
	if dir == 'U' {
		return UP
	}
	return NULL
}

// DirFromString returns a basic direction from a string representation.
func DirFromString(dir string) Direction {
	if dir == "L" {
		return LEFT
	}
	if dir == "R" {
		return RIGHT
	}
	if dir == "D" {
		return DOWN
	}
	if dir == "U" {
		return UP
	}
	return NULL
}

var (
	// Origin represents the center of the plane, or (0,0).
	ORIGIN Location = Location{0, 0}
)

// basic directions
var (
	// Moves one unit to the left of the x-axis.
	LEFT Direction = Direction{0, -1}
	// Moves one unit to the right of the x-axis.
	RIGHT Direction = Direction{0, 1}
	// Moves one unit upwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	UP Direction = Direction{-1, 0}
	// Moves one unit downwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	DOWN Direction = Direction{1, 0}
	// Doesn't move.
	NULL Direction = Direction{0, 0}
)

// TranslateByInteger translates the current direction by a given integer factor.
func TranslateByInteger(d Direction, factor int) Direction {
	return Direction{d[0] * factor, d[1] * factor}
}

// TurnLeft rotates the current direction by 90 degrees counter-clockwise
func TurnLeft(d Direction) Direction {
	return Direction{-d[1], d[0]}
}

// TurnRight rotates the current direction by 90 degrees clockwise
func TurnRight(d Direction) Direction {
	return Direction{d[1], -d[0]}
}

// TurnAround reverses the current direction, such that the original direction
// and the new one add up to 0.
func TurnAround(d Direction) Direction {
	return Direction{-d[0], -d[1]}
}

/* Directions, expressed as cardinal directions. */

var (
	// Moves one unit upwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	N Direction = UP
	// Moves one unit upwards the y-axis. Likely for parsing reasons, in all
	// Advent-of-Code problems the y-axis is reversed.
	S Direction = DOWN
	// Moves one unit to the left of the x-axis.
	W Direction = LEFT
	// Moves one unit to the right of the x-axis.
	E Direction = RIGHT
	// Moves one unit upwards the y-axis and and one unit to the right of the
	// y-axis.
	NE Direction = Direction{-1, 1}
	// Moves one unit upwards the y-axis and and one unit to the left of the
	// y-axis.
	NW Direction = Direction{-1, -1}
	// Moves one unit downwards the y-axis and and one unit to the right of the
	// y-axis.
	SE Direction = Direction{1, 1}
	// Moves one unit downwards the y-axis and and one unit to the right of the
	// y-axis.
	SW Direction = Direction{1, -1}
)
