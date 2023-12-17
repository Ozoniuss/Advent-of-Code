package main

import (
	"aoc/golib/rectboard"
	"aoc/golib/twod"
	"fmt"
	"slices"
)

// Backtracking approaches that don't finish executing.

type state struct {
	loc      []twod.Location
	dir      []twod.Direction
	sum      int
	forwards int
}

func explore(board *rectboard.RectBoard) {

	bx := *board
	visited := make(map[vertex]int)
	minsum := 3721

	q := make([]state, 0)

	// Don't add those to visited yet because you can revisit them.
	// Their heat doens't matetr
	q = append(q,
		state{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.RIGHT},
			sum:      0,
			forwards: 0,
		},
		state{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.DOWN},
			sum:      0,
			forwards: 0,
		},
	)

	for {
		fmt.Println(len(visited))
		if len(q) == 0 {
			fmt.Println("done")
			return
		}
		top := q[0]
		q = q[1:]

		if top.sum > minsum {
			continue
		}

		// fmt.Println(top.forwards)
		if top.loc[len(top.loc)-1] == (twod.Location{len(bx) - 1, len(bx[0]) - 1}) {

			if top.forwards < 3 {
				// fmt.Println(top.forwards)
				continue
			}

			if top.sum < minsum {
				minsum = top.sum
				fmt.Println(top.forwards)

				fmt.Println("minsum", minsum)
			}
			continue
		}

		currentKey := vertex{
			loc:      top.loc[len(top.loc)-1],
			dir:      top.dir[len(top.dir)-1],
			forwards: top.forwards,
		}

		// I can come from a different direction, allowing for multiple
		// exploration opportunities.
		// It's also not the same if I come from the same direction having
		// 2 more forwards of having no more forwards.
		if val, ok := visited[currentKey]; ok {
			// visited but the path is more costley
			if top.sum >= val {
				continue
			} else {
				visited[currentKey] = top.sum
			}
		} else {
			// not visited at all
			visited[currentKey] = top.sum
		}

		var nextDirs []twod.Direction
		if currentKey.forwards < 3 {
			nextDirs = []twod.Direction{top.dir[len(top.dir)-1]}
		} else if currentKey.forwards >= 9 {
			nextDirs = []twod.Direction{
				twod.TurnLeft(top.dir[len(top.dir)-1]),
				twod.TurnRight(top.dir[len(top.dir)-1]),
			}
		} else {
			nextDirs = []twod.Direction{
				twod.TurnLeft(top.dir[len(top.dir)-1]),
				twod.TurnRight(top.dir[len(top.dir)-1]),
				top.dir[len(top.dir)-1],
			}
		}
		for _, nd := range nextDirs {
			nextLoc := twod.Move(top.loc[len(top.loc)-1], nd)
			if rectboard.IsInBoard(nextLoc, board) {
				state := state{
					loc:      slices.Clone(append(top.loc, twod.Move(top.loc[len(top.loc)-1], nd))),
					dir:      slices.Clone(append(top.dir, nd)),
					sum:      top.sum + int(bx[nextLoc[0]][nextLoc[1]]-'0'),
					forwards: top.forwards,
				}
				// last direction didn't change
				if currentKey.dir == nd {
					state.forwards++
				} else {
					state.forwards = 0
				}
				q = append(q, state)
			}
		}

		// nextDirs := []twod.Direction{
		// 	twod.TurnLeft(top.dir[len(top.dir)-1]),
		// 	twod.TurnRight(top.dir[len(top.dir)-1]),
		// }
		// if len(top.loc) < 3 {
		// 	nextDirs = append(nextDirs, top.dir[len(top.dir)-1])
		// } else {
		// 	// previous two directions were the same
		// 	if !(top.dir[len(top.dir)-2] == top.dir[len(top.dir)-1] &&
		// 		top.dir[len(top.dir)-3] == top.dir[len(top.dir)-1]) {
		// 		nextDirs = append(nextDirs, top.dir[len(top.dir)-1])
		// 	}
		// }

	}
}

func exploreDfs(board *rectboard.RectBoard) {
	visited := make(map[vertex]int)
	minsum := 99999
	startingStates := []state{
		{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.RIGHT},
			sum:      0,
			forwards: 0,
		},
		{
			loc:      []twod.Location{twod.ORIGIN},
			dir:      []twod.Direction{twod.DOWN},
			sum:      0,
			forwards: 0,
		}}
	for _, state := range startingStates {
		dfs(board, state, visited, &minsum)
	}
}

func dfs(board *rectboard.RectBoard, current state, visited map[vertex]int, minsum *int) {
	bx := *board

	top := current
	if top.sum > *minsum {
		return
	}

	if top.loc[len(top.loc)-1] == (twod.Location{len(bx) - 1, len(bx[0]) - 1}) {

		if top.forwards < 3 {
			// fmt.Println(top.forwards)
			return
		}

		fmt.Println("reached the end")
		fmt.Println(top.sum)
		if top.sum < *minsum {
			*minsum = top.sum
		}
		return
	}

	currentKey := vertex{
		loc:      top.loc[len(top.loc)-1],
		dir:      top.dir[len(top.dir)-1],
		forwards: top.forwards,
	}

	if val, ok := visited[currentKey]; ok {
		// visited but the path is more costly
		if top.sum >= val {
			return
		} else {
			visited[currentKey] = top.sum
			if len(visited)%10000 == 0 {
				fmt.Println(len(visited))
			}
		}
	} else {
		// not visited at all

		visited[currentKey] = top.sum
		if len(visited)%10000 == 0 {
			fmt.Println(len(visited))
		}

	}

	var nextDirs []twod.Direction
	if currentKey.forwards < 3 {
		nextDirs = []twod.Direction{top.dir[len(top.dir)-1]}
	} else if currentKey.forwards >= 9 {
		nextDirs = []twod.Direction{
			twod.TurnLeft(top.dir[len(top.dir)-1]),
			twod.TurnRight(top.dir[len(top.dir)-1]),
		}
	} else {
		nextDirs = []twod.Direction{
			twod.TurnLeft(top.dir[len(top.dir)-1]),
			twod.TurnRight(top.dir[len(top.dir)-1]),
			top.dir[len(top.dir)-1],
		}
	}
	for _, nd := range nextDirs {
		nextLoc := twod.Move(top.loc[len(top.loc)-1], nd)
		if rectboard.IsInBoard(nextLoc, board) {
			state := state{
				loc:      slices.Clone(append(top.loc, twod.Move(top.loc[len(top.loc)-1], nd))),
				dir:      slices.Clone(append(top.dir, nd)),
				sum:      top.sum + int(bx[nextLoc[0]][nextLoc[1]]-'0'),
				forwards: top.forwards,
			}
			// last direction didn't change
			if currentKey.dir == nd {
				state.forwards++
			} else {
				state.forwards = 0
			}
			dfs(board, state, visited, minsum)
		}
	}

}
