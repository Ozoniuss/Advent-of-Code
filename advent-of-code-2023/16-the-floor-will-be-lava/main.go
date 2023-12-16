package main

import (
	"aoc/golib/rectboard"
	"aoc/golib/twod"
	"fmt"
)

// This setup is done not because I like global variables, but in order to avoid
// any potential system calls during benchmarks (and to also make it easier to
// just call partXXX function). From my benchmarks it also turns out that s.Text()
// allocates, even if reusing the variable for the line.
// var inputLines = readlines()

// func readlines() []string {
// 	f, err := os.Open("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	var lines []string
// 	s := bufio.NewScanner(f)
// 	for s.Scan() {
// 		lines = append(lines, s.Text())
// 	}
// 	return lines
// }

var board = rectboard.ReadBoardFromFile("input.txt")

type BeanWithDirection struct {
	pos twod.Location
	dir twod.Direction
}

func findNextBeans(bean BeanWithDirection, mirror byte) []BeanWithDirection {
	if mirror == '.' {
		return []BeanWithDirection{{dir: bean.dir, pos: twod.Move(bean.pos, bean.dir)}}
	}
	if bean.dir == twod.RIGHT {
		if mirror == '\\' {
			return []BeanWithDirection{
				{dir: twod.DOWN, pos: twod.Move(bean.pos, twod.DOWN)}}
		}
		if mirror == '/' {
			return []BeanWithDirection{
				{dir: twod.UP, pos: twod.Move(bean.pos, twod.UP)}}
		}
		if mirror == '|' {
			return []BeanWithDirection{
				{dir: twod.UP, pos: twod.Move(bean.pos, twod.UP)},
				{dir: twod.DOWN, pos: twod.Move(bean.pos, twod.DOWN)},
			}
		}
		if mirror == '-' {
			return []BeanWithDirection{
				{dir: bean.dir, pos: twod.Move(bean.pos, bean.dir)},
			}
		}
	}

	if bean.dir == twod.LEFT {
		if mirror == '/' {
			return []BeanWithDirection{
				{dir: twod.DOWN, pos: twod.Move(bean.pos, twod.DOWN)}}
		}
		if mirror == '\\' {
			return []BeanWithDirection{
				{dir: twod.UP, pos: twod.Move(bean.pos, twod.UP)}}
		}
		if mirror == '|' {
			return []BeanWithDirection{
				{dir: twod.UP, pos: twod.Move(bean.pos, twod.UP)},
				{dir: twod.DOWN, pos: twod.Move(bean.pos, twod.DOWN)},
			}
		}
		if mirror == '-' {
			return []BeanWithDirection{
				{dir: bean.dir, pos: twod.Move(bean.pos, bean.dir)},
			}
		}
	}
	if bean.dir == twod.UP {
		if mirror == '/' {
			return []BeanWithDirection{
				{dir: twod.RIGHT, pos: twod.Move(bean.pos, twod.RIGHT)}}
		}
		if mirror == '\\' {
			return []BeanWithDirection{
				{dir: twod.LEFT, pos: twod.Move(bean.pos, twod.LEFT)}}
		}
		if mirror == '-' {
			return []BeanWithDirection{
				{dir: twod.LEFT, pos: twod.Move(bean.pos, twod.LEFT)},
				{dir: twod.RIGHT, pos: twod.Move(bean.pos, twod.RIGHT)},
			}
		}
		if mirror == '|' {
			return []BeanWithDirection{
				{dir: bean.dir, pos: twod.Move(bean.pos, bean.dir)},
			}
		}
	}
	if bean.dir == twod.DOWN {
		if mirror == '\\' {
			return []BeanWithDirection{
				{dir: twod.RIGHT, pos: twod.Move(bean.pos, twod.RIGHT)}}
		}
		if mirror == '/' {
			return []BeanWithDirection{
				{dir: twod.LEFT, pos: twod.Move(bean.pos, twod.LEFT)}}
		}
		if mirror == '-' {
			return []BeanWithDirection{
				{dir: twod.LEFT, pos: twod.Move(bean.pos, twod.LEFT)},
				{dir: twod.RIGHT, pos: twod.Move(bean.pos, twod.RIGHT)},
			}
		}
		if mirror == '|' {
			return []BeanWithDirection{
				{dir: bean.dir, pos: twod.Move(bean.pos, bean.dir)},
			}
		}
	}
	return nil
}

func Move(currentBean BeanWithDirection, board *rectboard.RectBoard, visited map[twod.Location]struct{}, visitedWIthDir map[BeanWithDirection]struct{}) {
	bx := *board

	// fmt.Println(currentBean)

	if !rectboard.IsInBoard(currentBean.pos, board) {
		return
	}
	visited[currentBean.pos] = struct{}{}
	if _, ok := visitedWIthDir[currentBean]; ok {
		return
	}
	visitedWIthDir[currentBean] = struct{}{}

	nextBeans := findNextBeans(currentBean, bx[currentBean.pos[0]][currentBean.pos[1]])
	for _, next := range nextBeans {
		Move(next, board, visited, visitedWIthDir)
	}
}

func part1() {
	currentBean := BeanWithDirection{
		dir: twod.RIGHT,
		pos: twod.ORIGIN,
	}
	visited := make(map[twod.Location]struct{})
	visitedWithDir := make(map[BeanWithDirection]struct{})
	Move(currentBean, board, visited, visitedWithDir)
	fmt.Println(len(visited))
}

func part2() {
	bx := *board
	allPositions := make([]BeanWithDirection, 0, 4*len(bx)*len(bx[0]))

	// It doesn't matter these repeat
	for i := 0; i < len(bx); i++ {
		allPositions = append(
			allPositions,
			BeanWithDirection{pos: twod.Location{i, 0}, dir: twod.UP},
			BeanWithDirection{pos: twod.Location{i, 0}, dir: twod.DOWN},
			BeanWithDirection{pos: twod.Location{i, 0}, dir: twod.RIGHT},
			BeanWithDirection{pos: twod.Location{i, 0}, dir: twod.LEFT},
		)
		allPositions = append(
			allPositions,
			BeanWithDirection{pos: twod.Location{i, len(bx[0]) - 1}, dir: twod.UP},
			BeanWithDirection{pos: twod.Location{i, len(bx[0]) - 1}, dir: twod.DOWN},
			BeanWithDirection{pos: twod.Location{i, len(bx[0]) - 1}, dir: twod.RIGHT},
			BeanWithDirection{pos: twod.Location{i, len(bx[0]) - 1}, dir: twod.LEFT},
		)
	}

	for j := 0; j < len(bx); j++ {
		allPositions = append(
			allPositions,
			BeanWithDirection{pos: twod.Location{0, j}, dir: twod.UP},
			BeanWithDirection{pos: twod.Location{0, j}, dir: twod.DOWN},
			BeanWithDirection{pos: twod.Location{0, j}, dir: twod.RIGHT},
			BeanWithDirection{pos: twod.Location{0, j}, dir: twod.LEFT},
		)
		allPositions = append(
			allPositions,
			BeanWithDirection{pos: twod.Location{len(bx) - 1, j}, dir: twod.UP},
			BeanWithDirection{pos: twod.Location{len(bx) - 1, j}, dir: twod.DOWN},
			BeanWithDirection{pos: twod.Location{len(bx) - 1, j}, dir: twod.RIGHT},
			BeanWithDirection{pos: twod.Location{len(bx) - 1, j}, dir: twod.LEFT},
		)
	}
	fmt.Println(len(allPositions))

	maxlen := 0
	maxbean := BeanWithDirection{}
	for _, cbean := range allPositions {
		visited := make(map[twod.Location]struct{})
		visitedWithDir := make(map[BeanWithDirection]struct{})
		Move(cbean, board, visited, visitedWithDir)
		if len(visited) > maxlen {
			maxlen = len(visited)
			maxbean = cbean
		}
	}
	fmt.Println(maxlen, maxbean)

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
	part2()
}
