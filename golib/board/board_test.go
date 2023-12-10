package board

import (
	"aoc/golib/twod"
	"strings"
	"testing"
)

var input1string = ""
var input1lines = []string{}

var input2Untrimmed = `
abcd
1234
xyzt
`
var input2string = strings.Trim(input2Untrimmed, "\n")
var input2lines = []string{
	"abcd",
	"1234",
	"xyzt",
}

func TestNewBoard(t *testing.T) {
	t.Run("new empty board", func(t *testing.T) {
		r := strings.NewReader(input1string)
		testBoards := []struct {
			name string
			b    *Board
		}{
			{
				name: "from reader",
				b:    NewBoard(WithReader(r)),
			},
			{
				name: "from lines",
				b:    NewBoard(WithLines(&input1lines)),
			},
		}
		for _, tb := range testBoards {
			t.Run(tb.name, func(t *testing.T) {
				board := tb.b
				if s := board.Size(); s != 0 {
					t.Errorf("Expected size %d, got %d", 12, s)
				}
			})
		}

	})

	t.Run("new non-empty board", func(t *testing.T) {
		r := strings.NewReader(input2string)
		testBoards := []struct {
			name string
			b    *Board
		}{
			{
				name: "from reader",
				b:    NewBoard(WithReader(r)),
			},
			{
				name: "from lines",
				b:    NewBoard(WithLines(&input2lines)),
			},
		}

		for _, tb := range testBoards {
			t.Run(tb.name, func(t *testing.T) {
				board := tb.b
				if s := board.Size(); s != 12 {
					t.Errorf("Expected size %d, got %d", 12, s)
				}

				for i := 0; i < len(input2lines); i++ {
					for j := 0; j < len(input2lines[0]); j++ {
						if input2lines[i][j] != board.MustGet(twod.Location{i, j}) {
							t.Errorf("location (%d, %d): expected %d, got %d", i, j, input2lines[i][j], board.MustGet(twod.Location{i, j}))
						}
					}
				}
			})
		}

	})

	t.Run("with rect", func(t *testing.T) {
		r := strings.NewReader(input2string)
		testBoards := []struct {
			name string
			b    *Board
		}{
			{
				name: "from reader",
				b:    NewBoard(WithReader(r), WithRect()),
			},
			{
				name: "from lines",
				b:    NewBoard(WithLines(&input2lines), WithRect()),
			},
		}

		for _, tb := range testBoards {
			t.Run(tb.name, func(t *testing.T) {
				board := tb.b
				if board.GetUpperLeftBoundary() != twod.ORIGIN {
					t.Errorf("expected upper boundary %v, got %v", twod.ORIGIN, board.GetUpperLeftBoundary())
				}
				if board.GetLowerRightBoundary() != (twod.Location{2, 3}) {
					t.Errorf("expected upper boundary %v, got %v", twod.ORIGIN, board.GetUpperLeftBoundary())
				}
			})
		}

	})
}
