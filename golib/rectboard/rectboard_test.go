package rectboard

import (
	"strings"
	"testing"
)

var input4x4Untrimmed = `
abcd
xyzt
1234
5678
`

var input4x4 = strings.Trim(input4x4Untrimmed, "\n")

func TestReadBoard(t *testing.T) {
	t.Run("read 4 by 4", func(t *testing.T) {
		r := strings.NewReader(input4x4)
		brd := ReadBoard(r)
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				board := (*brd)
				// ignore all string newlines
				if board[i][j] != input4x4[5*i+j] {
					t.Errorf("expected %c, got %c", input4x4[5*i+j], board[i][j])
				}
			}
		}
	})
}

func TestAsString(t *testing.T) {
	t.Run("converts 4 by 4", func(t *testing.T) {
		r := strings.NewReader(input4x4)
		brd := ReadBoard(r)
		brds := AsString(brd)
		if input4x4 != brds {
			t.Errorf("expected %s, got %s", input4x4, brds)
		}
	})
}
