package board

import (
	"aoc/golib/twod"
	"bufio"
	"io"
	"maps"
	"strings"
)

// For speed sake, code here will panic if something goes wrong, since I
// won't care about handling errors during the contest.

// Board stores all the tiles (locations) of a board.
type Board struct {
	locations     map[twod.Location]byte
	initialCap    int
	fromLines     *[]string
	fromReader    io.Reader
	LeftBoundary  twod.Location
	RightBoundary twod.Location
	isRect        bool
}

type boardOption func(u *Board)

// WithCap sets an initial capacity to the locations map. If used more than
// once, the final capacity will apply.
func WithCap(cap int) boardOption {
	return func(u *Board) {
		u.initialCap = cap
	}
}

// WithReader sets an io.Reader to read the board from. It is assumed that each
// byte from the reader represents a tile, and consecutive rows are separated
// by newlines.
func WithReader(r io.Reader) boardOption {
	return func(u *Board) {
		u.fromReader = r
	}
}

// WithLines creates a board from an input which was parsed into an array of
// lines.
func WithLines(lines *[]string) boardOption {
	return func(u *Board) {
		u.fromLines = lines
	}
}

// WithLeftBoundary sets an upper left boundary for the board. Note that if the
// board is not rectangular, or the boundary is set incorrectly, this option
// results in undefined behavior.
func WithLeftBoundary(left twod.Location) boardOption {
	return func(u *Board) {
		u.LeftBoundary = left
	}
}

// / WithRightBoundary sets an lower right boundary for the board. Note that if
// the board is not rectangular, or the boundary is set incorrectly, this option
// results in undefined behavior.
func WithRightBoundary(right twod.Location) boardOption {
	return func(u *Board) {
		u.RightBoundary = right
	}
}

// WithRect announces the parses that the input board is rectangular, and the
// boundaries will automatically be set. If this and WithXXXBoundary are
// provided, the behaviour will be undefined.
func WithRect() boardOption {
	return func(u *Board) {
		u.isRect = true
	}
}

// NewBoard generates a new board.
func NewBoard(opts ...boardOption) *Board {
	// many aoc boards were 140 * 140. Setting this higher just in case.
	b := Board{
		locations:  nil,
		initialCap: 256 * 256,
		fromLines:  nil,
		fromReader: nil,
	}
	for _, o := range opts {
		o(&b)
	}
	if b.fromLines != nil && b.fromReader != nil {
		panic("cannot parse both lines and reader")
	}

	locations := make(map[twod.Location]byte, b.initialCap)
	b.locations = locations

	if b.fromLines != nil {
		for i := 0; i < len(*b.fromLines); i++ {
			for j := 0; j < len((*b.fromLines)[i]); j++ {
				b.locations[twod.Location{i, j}] = (*b.fromLines)[i][j]
			}
		}
		if b.isRect {
			b.LeftBoundary = twod.Location{0, 0}
			b.RightBoundary = twod.Location{len(*b.fromLines) - 1, len((*b.fromLines)[0]) - 1}
		}
	}

	if b.fromReader != nil {
		s := bufio.NewScanner(b.fromReader)
		i := 0
		var line []byte
		for s.Scan() {
			line = s.Bytes()
			for j, c := range line {
				b.locations[twod.Location{i, j}] = c
			}
			i++
		}
		if b.isRect {
			b.LeftBoundary = twod.ORIGIN
			b.RightBoundary = twod.Location{i - 1, len(line) - 1}
		}
	}

	return &b
}

func (u *Board) Size() int {
	return len(u.locations)
}

func (u *Board) GetLocations() map[twod.Location]byte {
	return u.locations
}

func (u *Board) GetLocationsCopy() map[twod.Location]byte {
	return maps.Clone(u.locations)
}

func (u *Board) GetStringRepresentation() string {
	b := &strings.Builder{}
	if u.isRect {
		for i := u.LeftBoundary[0]; i <= u.RightBoundary[0]; i++ {
			for j := u.LeftBoundary[1]; j <= u.RightBoundary[1]; j++ {
				b.WriteByte(u.locations[twod.Location{i, j}])
			}
			b.WriteByte('\n')
		}
		return b.String()
	} else {
		panic("unimplemented")
	}
}

func (u *Board) Set(loc twod.Location, val byte) {
	u.locations[loc] = val
}

func (u *Board) UpdateIfExists(loc twod.Location, val byte) {
	if _, ok := u.locations[loc]; ok {
		u.locations[loc] = val
	}
}

func (u *Board) Get(loc twod.Location) (byte, bool) {
	v, ok := u.locations[loc]
	return v, ok
}

func (u *Board) MustGet(loc twod.Location) byte {
	return u.locations[loc]
}

func (u *Board) GetUpperLeftBoundary() twod.Location {
	return u.LeftBoundary
}

func (u *Board) GetLowerRightBoundary() twod.Location {
	return u.RightBoundary
}

// GoAround goes around the provided location and calls f, only for the
// neighbours that are on the board.
func (u *Board) GoAround(loc twod.Location, f func(l twod.Location, v byte) bool) {
	for i := loc[0] - 1; i <= loc[0]+1; i++ {
		for j := loc[1] - 1; j <= loc[1]+1; j++ {
			if i == loc[0] && j == loc[1] {
				continue
			}

			cloc := twod.Location{i, j}
			if val, ok := u.Get(cloc); ok {
				shouldQuit := f(cloc, val)
				if shouldQuit {
					return
				}
			}

		}
	}
}
