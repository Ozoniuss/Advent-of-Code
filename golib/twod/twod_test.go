package twod

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	type test struct {
		name     string
		loc      Location
		dir      Direction
		expected Location
	}
	tests := []test{
		{
			name:     "random directions 1",
			loc:      Location{1, 1},
			dir:      Direction{-1, 1},
			expected: Location{0, 2},
		},
		{
			name:     "random directions 2",
			loc:      Location{3, 4},
			dir:      Direction{1, -1},
			expected: Location{4, 3},
		},
		{
			name:     "random directions 3",
			loc:      Location{-1, -2},
			dir:      Direction{4, 5},
			expected: Location{3, 3},
		},
	}
	predefinedDirs := []Direction{
		LEFT, RIGHT, UP, DOWN, N, S, E, W, NE, NW, SE, SW,
	}
	for _, predefinedDir := range predefinedDirs {
		tests = append(tests, test{
			loc:      ORIGIN,
			dir:      predefinedDir,
			expected: Location{predefinedDir[0], predefinedDir[1]},
		})
	}
	for _, tt := range tests {
		if tt.expected != Move(tt.loc, tt.dir) {
			t.Errorf("got %v, want %v", Move(tt.loc, tt.dir), tt.expected)
		}
	}
}

func TestAddDirs(t *testing.T) {
	type test struct {
		dir1     Direction
		dir2     Direction
		expected Direction
	}
	tests := []test{
		{
			dir1:     Direction{1, 1},
			dir2:     Direction{-1, 1},
			expected: Direction{0, 2},
		},
		{
			dir1:     Direction{3, 4},
			dir2:     Direction{1, -1},
			expected: Direction{4, -3},
		},
		{
			dir1:     Direction{-1, -2},
			dir2:     Direction{4, 5},
			expected: Direction{3, 3},
		},
	}

	for _, tt := range tests {
		if tt.expected != AddDirs(tt.dir1, tt.dir2) {
			fmt.Printf("got %v, want %v", AddDirs(tt.dir1, tt.dir2), tt.expected)
		}
	}
}
