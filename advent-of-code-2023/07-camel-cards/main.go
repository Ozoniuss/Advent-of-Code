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

type handWithScore struct {
	hand  string
	score int
}

//	var orderMap = map[byte]int{
//		'A': 14,
//		'K': 13,
//		'Q': 12,
//		'J': 11,
//		'T': 10,
//		'9': 9,
//		'8': 8,
//		'7': 7,
//		'6': 6,
//		'5': 5,
//		'4': 4,
//		'3': 3,
//		'2': 2,
//	}
var orderMap = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

func getOne(hands map[byte]int) (byte, int) {
	for k, v := range hands {
		return k, v
	}
	return 0, 0
}

// in order: five, four, full, three, two, one, high
func processHand(hand string) [7]int {
	ret := [7]int{0: 0}
	cards := make(map[byte]int, 5)
	for _, c := range hand {
		cards[byte(c)]++
	}

	// five of a kind
	if len(cards) == 1 {
		ret[0] = 5
	}
	if len(cards) == 2 {
		_, v := getOne(cards)
		// either 4 or 1 of a value means four of a kind
		if v == 4 || v == 1 {
			ret[1] = 1
		} else {
			// full house otherwise
			ret[2] = 1
		}
	}
	if len(cards) == 3 {
		values := maps.Values(cards)
		slices.SortFunc(values, func(a, b int) int {
			return b - a
		})
		if values[0] == 3 {
			// three of a kind
			ret[3]++
		}
		// two pairs
		if values[0] == 2 {
			ret[4]++
		}
	}
	if len(cards) == 4 {
		// one pair
		ret[5]++
	}
	if len(cards) == 5 {
		ret[6]++
	}
	// highest card, set to biggest value for comparisons
	// max := 0
	// for k := range cards {
	// 	if orderMap[k] > max {
	// 		max = orderMap[k]
	// 	}
	// }
	// ret[6] = max
	return ret
}

func completeToBestPair(card string) string {

	if card == "JJJJJ" {
		return "AAAAA"
	}

	cards := make(map[byte]int, 5)
	for _, c := range card {
		cards[byte(c)]++
	}

	if cards['J'] == 0 {
		return card
	}
	maxkey := byte(0)
	maxval := 0
	for k, v := range cards {
		if k == 'J' {
			continue
		}
		if v > maxval {
			maxval = v
			maxkey = k
		}
	}

	out := &strings.Builder{}
	for c := 0; c < len(card); c++ {
		if card[c] == 'J' {
			out.WriteByte(maxkey)
		} else {
			out.WriteByte(card[c])
		}
	}
	return out.String()
}

func compareCards(card1, card2 string) int {

	card1changed := completeToBestPair(card1)
	card2changed := completeToBestPair(card2)
	ret1 := processHand(card1changed)
	ret2 := processHand(card2changed)

	for i := 0; i < len(ret1); i++ {
		if ret1[i] > ret2[i] {
			return 1
		}
		if ret1[i] < ret2[i] {
			return -1
		}
		if ret1[i] == ret2[i] && ret1[i] == 1 {
			break
		}
	}
	fmt.Println(ret1, ret2, card1, card2)
	for i := 0; i < len(card1); i++ {
		if orderMap[card1[i]] > orderMap[card2[i]] {
			return 1
		}
		if orderMap[card1[i]] < orderMap[card2[i]] {
			return -1
		}
	}
	return 0
}

// func compareCards(card1, card2 string) int {
// 	ret1 := processHand(card1)
// 	ret2 := processHand(card2)

// 	for i := 0; i < len(ret1); i++ {
// 		if ret1[i] > ret2[i] {
// 			return 1
// 		}
// 		if ret1[i] < ret2[i] {
// 			return -1
// 		}
// 		if ret1[i] == ret2[i] && ret1[i] == 1 {
// 			break
// 		}
// 	}
// 	fmt.Println(ret1, ret2, card1, card2)
// 	for i := 0; i < len(card1); i++ {
// 		if orderMap[card1[i]] > orderMap[card2[i]] {
// 			return 1
// 		}
// 		if orderMap[card1[i]] < orderMap[card2[i]] {
// 			return -1
// 		}
// 	}
// 	return 0
// }

func part1() {
	hands := make([]handWithScore, 0, 100)
	for _, line := range inputLines {
		parts := strings.Split(line, " ")
		hand := parts[0]
		score, _ := strconv.Atoi(parts[1])
		hands = append(hands, handWithScore{
			hand:  hand,
			score: score,
		})
	}
	slices.SortFunc[[]handWithScore](hands, func(a, b handWithScore) int {
		return compareCards(a.hand, b.hand)
	})
	ttl := 0
	for i := 1; i <= len(hands); i++ {
		ttl += i * hands[i-1].score
	}
	fmt.Println(ttl)
}

// func part1() {
// 	hands := make([]handWithScore, 0, 100)
// 	for _, line := range inputLines {
// 		parts := strings.Split(line, " ")
// 		hand := parts[0]
// 		score, _ := strconv.Atoi(parts[1])
// 		hands = append(hands, handWithScore{
// 			hand:  hand,
// 			score: score,
// 		})
// 	}
// 	slices.SortFunc[[]handWithScore](hands, func(a, b handWithScore) int {
// 		return compareCards(a.hand, b.hand)
// 	})
// 	ttl := 0
// 	for i := 1; i <= len(hands); i++ {
// 		ttl += i * hands[i-1].score
// 	}
// 	fmt.Println(ttl)
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

	part1()
	// part2()
}
