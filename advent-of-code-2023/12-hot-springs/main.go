package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

func processLine(line string) (string, []int) {
	parts := strings.Split(line, " ")
	springs := parts[0]

	nums := make([]int, 0, 128)
	numbersStr := strings.Split(parts[1], ",")
	for _, n := range numbersStr {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	return springs, nums
}

func processLinePart2(line string) (string, []int) {
	parts := strings.Split(line, " ")
	springs := parts[0]

	nums := make([]int, 0, 128)
	numbersStr := strings.Split(parts[1], ",")
	for _, n := range numbersStr {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	return explode(springs, nums)
}

func CountDamaged(springs string) []int {
	dots := strings.Split(springs, ".")
	nums := make([]int, 0, 64)
	for _, dot := range dots {
		if len(dot) > 0 {
			nums = append(nums, len(dot))
		}
	}
	return nums
}

func existingNumberOfHashtags(springs string) int {
	cnt := 0
	for _, c := range springs {
		if c == '#' {
			cnt++
		}
	}
	return cnt
}

func findQuestionIndices(springs string) []int {
	inds := make([]int, 0, 32)
	for i := 0; i < len(springs); i++ {
		if springs[i] == '?' {
			inds = append(inds, i)
		}
	}
	return inds
}

func sum(n []int) int {
	s := 0
	for _, ni := range n {
		s += ni
	}
	return s
}

type ComboWithCount struct {
	combo      []byte
	startsWith []int
	chash      int
}

func StartsWithSlice(nums []int, startsWith []int) bool {
	if len(startsWith) == 0 {
		return true
	}
	if len(startsWith) > len(nums) {
		return false
	}
	for i := 0; i < len(startsWith); i++ {
		if startsWith[i] != nums[i] {
			return false
		}
	}
	return true
}

func StartsWithSprint(spring string, chars []byte) bool {
	if len(chars) == 0 {
		return true
	}
	if len(chars) > len(spring) {
		return false
	}
	for i := 0; i < len(chars); i++ {
		if spring[i] != '?' && (spring[i] != chars[i]) {
			return false
		}
	}
	return true
}

func NextInLine(nums []int, startsWith []int) int {
	if len(startsWith) >= len(nums) {
		return -1
	}

	return nums[len(startsWith)]
}

func computeCacheKey(nums []int) string {
	b := &strings.Builder{}
	for _, n := range nums {
		b.WriteString(strconv.Itoa(n))
	}
	return b.String()
}

type Cache struct {
	startsWith string
	length     int
}

func isGoodStart(val ComboWithCount, springs string, nums []int) bool {
	if !StartsWithSlice(nums, val.startsWith) {
		return false
	}
	if !StartsWithSprint(springs, val.combo) {
		return false
	}
	return true
}

func isGoodStart2(val ComboStartsWith, springs string, nums []int) bool {
	if !StartsWithSlice(nums, val.startsWith) {
		return false
	}
	if !StartsWithSprint(springs, val.combo) {
		return false
	}
	return true
}

func findNumberOfArrangmentsGenerative(springs string, nums []int) int {
	q := []ComboWithCount{
		{
			combo:      []byte{},
			chash:      0,
			startsWith: []int{},
		},
	}
	count := 0
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		if (len(top.combo) == len(springs)) && top.combo[len(top.combo)-1] == '#' {
			top.startsWith = append(top.startsWith, top.chash)
		}

		if !StartsWithSlice(nums, top.startsWith) {
			continue
		}
		if !StartsWithSprint(springs, top.combo) {
			continue
		}
		// fmt.Println("waaa", string(top.combo), top.startsWith)
		// not generating anything off this
		if len(top.combo) == len(springs) {

			// fmt.Println(string(top.combo), springs, top.startsWith, nums)
			if len(top.startsWith) == len(nums) {
				count++
			}
			continue
		}
		// fmt.Println(string(top.combo), top.startsWith)
		for _, c := range []byte{'.', '#'} {
			topCopyCombo := slices.Clone(top.combo)
			topCopyCombo = append(topCopyCombo, c)

			if c == '.' {
				if len(top.combo) == 0 {
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: []int{},
					})
					continue
				}
				if top.combo[len(top.combo)-1] == '.' {
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: slices.Clone(top.startsWith),
					})
					// breaking the hashtag
				} else if top.combo[len(top.combo)-1] == '#' {
					cnums := slices.Clone(top.startsWith)
					cnums = append(cnums, top.chash)
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: cnums,
					})
				}

			} else if c == '#' {
				if len(top.combo) > 0 && top.combo[len(top.combo)-1] == '#' {
					continue
				}
				// here get the next number of nums and fill this motherfucker
				nextInLine := NextInLine(nums, top.startsWith)

				for i := 0; i < nextInLine-1; i++ {
					topCopyCombo = append(topCopyCombo, '#')
				}
				next := ComboWithCount{
					combo:      topCopyCombo,
					chash:      nextInLine,
					startsWith: slices.Clone(top.startsWith),
				}
				q = append(q, next)
			}

		}
	}
	return count

}

func computeComboNums(combo string) []int {
	nums := make([]int, 0, 32)
	parts := strings.Split(combo, ".")
	for _, part := range parts {
		if len(part) != 0 {
			nums = append(nums, len(part))
		}
	}
	return nums
}

type ComboStartsWith struct {
	combo      []byte
	startsWith []int
}

func findNumberOfArrangmentsCompleteBacktracking(springs string, nums []int) int {
	q := make([][]ComboStartsWith, 100)

	cache := make(map[Cache]int)

	q[0] = []ComboStartsWith{
		{
			combo:      []byte{},
			startsWith: []int{},
		},
	}
	count := 0
	qidx := 0
	fmt.Println(len(springs))
	for qidx <= len(springs)+1 {
		if len(q[qidx]) == 0 {
			qidx++
			continue
		}
		top := q[qidx][0]
		q[qidx] = q[qidx][1:]

		// this will contain the latest ones as well
		prev, prevExists := cache[Cache{
			startsWith: computeCacheKey(top.startsWith),
			length:     len(top.combo),
		}]
		if !prevExists {
			prev = 1
		}
		if len(top.combo) >= 2 && top.combo[len(top.combo)-1] == '.' && top.combo[len(top.combo)-2] == '#' {
			if len(top.combo) == len(springs)+1 &&
				len(top.startsWith) == len(nums) {
				count += prev
				continue
			}
		}

		if !isGoodStart2(top, springs, nums) {
			continue
		}

		if len(top.combo) == len(springs) &&
			len(top.startsWith) == len(nums) {
			// fmt.Println("exists", prevExists)
			count += prev
			continue
		}

		for _, c := range []byte{'.', '#'} {

			if c == '.' {
				topCopyCombo := slices.Clone(top.combo)
				topCopyCombo = append(topCopyCombo, c)

				next := ComboStartsWith{
					combo:      topCopyCombo,
					startsWith: slices.Clone(top.startsWith),
				}

				cachedval := Cache{startsWith: computeCacheKey(next.startsWith), length: len(next.combo)}

				// adding a dot means a new cacheable value, increase by prev.
				// if we can get to (*). in k ways and (*).. is valid, we can
				// get to (*).. in k+1 ways
				if _, ok := cache[cachedval]; !ok {
					q[len(next.combo)] = append(q[len(next.combo)], next)
				}
				cache[cachedval] += prev

			} else if c == '#' {

				topCopyCombo := slices.Clone(top.combo)
				topCopyCombo = append(topCopyCombo, c)
				// here get the next number of nums and fill this motherfucker
				nextInLine := NextInLine(nums, top.startsWith)

				for i := 0; i < nextInLine-1; i++ {
					topCopyCombo = append(topCopyCombo, '#')
				}
				// add this for caching purposes
				topCopyCombo = append(topCopyCombo, '.')

				next := ComboStartsWith{
					combo:      topCopyCombo,
					startsWith: computeComboNums(string(topCopyCombo)),
				}

				cachedval := Cache{startsWith: computeCacheKey(next.startsWith), length: len(next.combo)}
				if _, ok := cache[cachedval]; !ok {
					// fmt.Println("not cached", string(next.combo), cachedval.length, cachedval.startsWith)
					q[len(next.combo)] = append(q[len(next.combo)], next)
				}
				// fmt.Println("cache increasing", string(next.combo), cachedval.length, cachedval.startsWith)
				cache[cachedval] += prev
			}
		}
	}
	return count

}

func findNumberOfArrangmentsWithCaching(springs string, nums []int) int {
	q := []ComboWithCount{
		{
			combo:      []byte{},
			chash:      0,
			startsWith: []int{},
		},
	}
	count := 0
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		if (len(top.combo) == len(springs)) && top.combo[len(top.combo)-1] == '#' {
			top.startsWith = append(top.startsWith, top.chash)
		}

		if !StartsWithSlice(nums, top.startsWith) {
			continue
		}
		if !StartsWithSprint(springs, top.combo) {
			continue
		}
		// fmt.Println("waaa", string(top.combo), top.startsWith)
		// not generating anything off this
		if len(top.combo) == len(springs) {

			// fmt.Println(string(top.combo), springs, top.startsWith, nums)
			if len(top.startsWith) == len(nums) {
				count++
			}
			continue
		}
		// fmt.Println(string(top.combo), top.startsWith)
		for _, c := range []byte{'.', '#'} {
			topCopyCombo := slices.Clone(top.combo)
			topCopyCombo = append(topCopyCombo, c)

			if c == '.' {
				if len(top.combo) == 0 {
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: []int{},
					})
					continue
				}
				if top.combo[len(top.combo)-1] == '.' {
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: slices.Clone(top.startsWith),
					})
					// breaking the hashtag
				} else if top.combo[len(top.combo)-1] == '#' {
					cnums := slices.Clone(top.startsWith)
					cnums = append(cnums, top.chash)
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: cnums,
					})
				}

			} else if c == '#' {
				if len(top.combo) > 0 && top.combo[len(top.combo)-1] == '#' {
					continue
				}
				// here get the next number of nums and fill this motherfucker
				nextInLine := NextInLine(nums, top.startsWith)

				for i := 0; i < nextInLine-1; i++ {
					topCopyCombo = append(topCopyCombo, '#')
				}
				next := ComboWithCount{
					combo:      topCopyCombo,
					chash:      nextInLine,
					startsWith: slices.Clone(top.startsWith),
				}
				q = append(q, next)
				// fmt.Println(string(next.combo), nextInLine, top.startsWith)
			}

		}
	}
	return count

}

func explode(springs string, nums []int) (string, []int) {
	b := &strings.Builder{}
	newNums := make([]int, 0, 256)
	for i := 0; i < 5; i++ {
		b.WriteString(springs)
		if i != 4 {
			b.WriteByte('?')
		}
		newNums = append(newNums, nums...)
	}
	return b.String(), newNums
}

type CacheKey struct {
	springs string
	nums    string
}

func countRecursive(springs string, nums []int, cache *map[CacheKey]uint64) uint64 {

	if val, ok := (*cache)[CacheKey{
		springs: springs,
		nums:    computeCacheKey(nums),
	}]; ok {
		return val
	}

	if len(springs) == 0 {
		// If there were nums left we couldn't fill them.
		if len(nums) == 0 {
			return 1
		}
		return 0
	}

	// No numbers left means we can only support . and ? which turns to .
	if len(nums) == 0 {
		if strings.Contains(springs, "#") {
			return 0
		}
		return 1
	}

	total := uint64(0)
	// First character acts as a dot. We don't care, this doesn't change nums
	// at all. We can basically just ignore this part.
	if springs[0] == '.' || springs[0] == '?' {
		total += countRecursive(springs[1:], nums, cache)
	}

	// In this case ? acts as a #. So basically, we need to fill the block
	// with the next number from nums. If we can't, we can ignore this bit.
	if springs[0] == '?' || springs[0] == '#' {
		if len(springs) >= nums[0] && !strings.Contains(springs[:nums[0]], ".") {

			// If the length exactly matches the first number, pass an empty
			// string.
			if len(springs) == nums[0] {
				total += countRecursive("", nums[1:], cache)
			} else {
				// This if is necessary because, in the next call to the
				// recursive function, we don't want to enter the last if.
				// If the last bit is a hashtag, it will enter the last if.
				// So it must either be a ? or a ., in order for us to consider
				// it a dot in this case and not send it further.
				if springs[nums[0]] != '#' {
					// If the length is greater, slice after the first dot as well
					total += countRecursive(springs[nums[0]+1:], nums[1:], cache)
				}
				// total += countRecursive(springs[nums[0]:], nums[1:])
			}
		}
	}
	(*cache)[CacheKey{
		springs: springs,
		nums:    computeCacheKey(nums),
	}] = total
	return total
}

func part1() {
	c := uint64(0)
	for _, line := range inputLines {
		cache := make(map[CacheKey]uint64, 4096)
		springs, nums := processLine(line)
		c += countRecursive(springs, nums, &cache)
	}
	fmt.Println(c)
}

func part2() {
	c := uint64(0)
	for _, line := range inputLines {
		cache := make(map[CacheKey]uint64, 4096)

		springs, nums := processLinePart2(line)
		fmt.Println(springs, nums)
		total := countRecursive(springs, nums, &cache)
		// fmt.Println("total", total)
		c += total
	}
	fmt.Println(c)
}

// func part1() {
// 	c := 0
// 	for idx, line := range inputLines {
// 		fmt.Println(idx)
// 		springs, nums := processLine(line)
// 		c += findNumberOfArrangments(springs, nums)
// 	}
// 	fmt.Println(c)
// }

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

// 	b := &strings.Builder{}
// 	carr := 0
// 	for _, qx := range q {
// 		b.Reset()
// 		qidx := 0
// 		for i := 0; i < len(springs); i++ {
// 			if springs[i] != '?' {
// 				b.WriteByte(springs[i])
// 			} else {
// 				b.WriteByte(qx.combo[qidx])
// 				qidx++
// 			}
// 		}
// 		numsActual := CountDamaged(b.String())
// 		if slices.Compare(nums, numsActual) == 0 {
// 			carr++
// 		}
// 	}
// 	return carr
// }

// part 1
// func findNumberOfArrangments(springs string, nums []int) int {
// 	inds := findQuestionIndices(springs)
// 	fmt.Println(inds)

// 	q := [][]byte{{'.'}, {'#'}}
// 	for {
// 		top := q[0]
// 		if len(top) >= len(inds) {
// 			break
// 		}
// 		q = q[1:]
// 		for _, c := range []byte{'.', '#'} {
// 			topCopy := slices.Clone(top)
// 			topCopy = append(topCopy, c)
// 			q = append(q, topCopy)
// 		}
// 	}

func findNumberOfArrangmentsWorking(springs string, nums []int) int {
	q := []ComboWithCount{
		{
			combo:      []byte{'.'},
			startsWith: []int{},
			chash:      0,
		},
		{
			combo:      []byte{'#'},
			startsWith: []int{},
			chash:      1,
		},
	}
	count := 0
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		if (len(top.combo) == len(springs)) && top.combo[len(top.combo)-1] == '#' {
			top.startsWith = append(top.startsWith, top.chash)
		}

		if !StartsWithSlice(nums, top.startsWith) {
			continue
		}
		if !StartsWithSprint(springs, top.combo) {
			continue
		}
		// fmt.Println(string(top.combo), top.startsWith)
		// not generating anything off this
		if len(top.combo) == len(springs) {

			// fmt.Println(string(top.combo), springs, top.startsWith, nums)
			if len(top.startsWith) == len(nums) {
				count++
			}
			continue
		}
		for _, c := range []byte{'.', '#'} {
			topCopyCombo := slices.Clone(top.combo)
			topCopyCombo = append(topCopyCombo, c)

			if c == '.' {
				if top.combo[len(top.combo)-1] == '.' {
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: slices.Clone(top.startsWith),
					})
					// breaking the hashtag
				} else if top.combo[len(top.combo)-1] == '#' {
					cnums := slices.Clone(top.startsWith)
					cnums = append(cnums, top.chash)
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      0,
						startsWith: cnums,
					})
				}

			} else if c == '#' {
				if top.combo[len(top.combo)-1] == '#' {
					q = append(q, ComboWithCount{
						combo:      topCopyCombo,
						chash:      top.chash + 1,
						startsWith: slices.Clone(top.startsWith),
					})
					continue
				}

				// here get the next number of nums and fill this motherfucker
				nextInLine := NextInLine(nums, top.startsWith)

				for i := 0; i < nextInLine-1; i++ {
					topCopyCombo = append(topCopyCombo, '#')
				}
				next := ComboWithCount{
					combo:      topCopyCombo,
					chash:      nextInLine,
					startsWith: slices.Clone(top.startsWith),
				}
				q = append(q, next)
				// fmt.Println(string(next.combo), nextInLine, top.startsWith)
			}

		}
	}
	return count

}
