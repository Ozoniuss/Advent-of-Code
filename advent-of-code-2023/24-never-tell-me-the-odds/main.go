package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

// (x*t + x0, y*t + y0, z*t + z0), t>0
type Hailstone struct {
	x0 int
	y0 int
	z0 int
	x  int
	y  int
	z  int
}

/*
Equation is t*(x, y) + (x0, y0)

Slope is given by (y/x). If the slopes are equal, lines never intersect.
We're basically solving the equation
t1a + a0 = t2c + c0
t1b = b0 = t2d + d0
b/a != d/c

a*t1 + a0 = b*t2 + b0
c*t1 + c0 = d*t2 + d0

a*t1 - b*t2 = b0 - a0
c*t1 - d*t2 = d0 - c0

matrix form
(a - b) (t1) = (b0 - a0)
(c - d) (t2) = (d0 - c0)

Inverse matrix (with a factor of 1/(bc-ad) in front)
(-d b)
(-c a)

General solution

(t1) = ( -d(b0-a0)+b(d0-c0)/(bc-ad) )
(t2) = ( -c(b0-a0)+a(d0-c0)/(bc-ad) )

a0 = h1.x0, c0 = h1.y0
b0 = h2.x0, d0 = h2.y0
*/

func doesIntersectInArea2d(h1, h2 Hailstone, boundlow, boundhigh int) bool {
	if h1.x == 0 || h1.y == 0 || h2.x == 0 || h2.y == 0 {
		panic("zero")
	}

	// Never intersect
	if h1.x*h2.y == h1.y*h2.x {
		fmt.Println("never intersects")
		return false
	}

	a0 := h1.x0
	a := h1.x

	b0 := h2.x0
	b := h2.x

	c0 := h1.y0
	c := h1.y

	d0 := h2.y0
	d := h2.y

	det := -a*d + b*c

	t1 := float64(((-d)*(b0-a0) + b*(d0-c0))) / float64(det)
	t2 := float64(((-c)*(b0-a0) + a*(d0-c0))) / float64(det)

	// fmt.Println(t1, t2, "intersects at", float64(a)*t1+float64(a0), float64(c)*t1+float64(c0))
	if t1 < 0 || t2 < 0 {
		return false
	}
	intersectAt := [2]float64{float64(a)*t1 + float64(a0), float64(c)*t1 + float64(c0)}
	if intersectAt[0] < float64(boundlow) || intersectAt[0] > float64(boundhigh) {
		return false
	}

	if intersectAt[1] < float64(boundlow) || intersectAt[1] > float64(boundhigh) {
		return false
	}
	return true
}

func processHailstone(line string) Hailstone {
	idx := strings.Index(line, " @ ")
	initialLocs := line[:idx]
	velocitiies := line[idx+3:]

	partsi := strings.Split(initialLocs, ", ")
	x0, _ := strconv.Atoi(partsi[0])
	y0, _ := strconv.Atoi(partsi[1])
	z0, _ := strconv.Atoi(partsi[2])

	partsv := strings.Split(velocitiies, ", ")
	x, _ := strconv.Atoi(partsv[0])
	y, _ := strconv.Atoi(partsv[1])
	z, _ := strconv.Atoi(partsv[2])

	return Hailstone{x0: x0, y0: y0, z0: z0, x: x, y: y, z: z}
}

// Assume there is a solution to this motherfucker. Then, if we project all
// lines on the plane perpendicular to the line that intersects them all, all
// intersections will project to a single point.
//
// Also, notice the author asks for the coordinates of point 0, but there are
// infinitely many points 0. This means the values of the direction vector add
// up to 0.
//
// Perhaps the best trick to spot here is that if there are two parallel lines,
// the plane that is perpendicular to the line we need to compute is also
// perpendicular to this plane. Hopefully that is the case.
func areParallel(h1, h2 Hailstone) bool {
	if h1.x == 0 || h1.y == 0 || h2.x == 0 || h2.y == 0 || h1.z == 0 || h2.z == 0 {
		panic("zero")
	}

	if h1.x*h2.y == h1.y*h2.x && h1.x*h2.z == h1.z*h2.x {
		fmt.Println("never intersects")
		return true
	}
	return false
}

// func part1() {
// 	hailstones := []Hailstone{}
// 	for _, line := range inputLines {
// 		hailstones = append(hailstones, processHailstone(line))
// 	}
// 	fmt.Println(hailstones)

// 	bl := 200000000000000
// 	bh := 400000000000000

//		c := 0
//		for i := 0; i < len(hailstones)-1; i++ {
//			for j := i + 1; j < len(hailstones); j++ {
//				intersects := doesIntersectInArea2d(hailstones[i], hailstones[j], bl, bh)
//				if intersects {
//					c++
//				}
//			}
//		}
//		fmt.Println(c)
//	}
func part2() {
	hailstones := []Hailstone{}
	for _, line := range inputLines {
		hailstones = append(hailstones, processHailstone(line))
	}
	// fmt.Println(hailstones)

	c := 0
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if areParallel(hailstones[i], hailstones[j]) {
				fmt.Println(hailstones[i], hailstones[j])
			}
		}
	}
	fmt.Println(c)
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
