package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Flip-flop modules (prefix %) are either on or off; they are initially off. If a flip-flop module receives a high pulse, it is ignored and nothing happens. However, if a flip-flop module receives a low pulse, it flips between on and off. If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.

// Conjunction modules (prefix &) remember the type of the most recent pulse received from each of their connected input modules; they initially default to remembering a low pulse for each input. When a pulse is received, the conjunction module first updates its memory for that input. Then, if it remembers high pulses for all inputs, it sends a low pulse; otherwise, it sends a high pulse.

// There is a single broadcast module (named broadcaster). When it receives a pulse, it sends the same pulse to all of its destination modules.

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

const (
	BROADCASTER = "broadcaster"
	CONJUNCTION = "conjunction"
	FLIPFLOP    = "flip-flop"
)

type signal struct {
	sender   string
	strenght int // 0 or 1
}

// also remember untyped testing modules

const (
	ON  = "on"
	OFF = "off"
)

type module struct {
	name  string
	mtype string
	// for flip-flop
	state         string
	receivedPulse []int
	// lastReceived signal
	received map[string]int
	// receivedFrom map[string]signal
	// prev       []*module
	// next       []*module
}

func appendOrCreate(key string, arr []string, m map[string][]string) {
	if _, ok := m[key]; !ok {
		m[key] = arr
	} else {
		m[key] = append(m[key], arr...)
	}
}

func parseLine(line string) {
	parts := strings.Split(line, "->")
	mname := strings.Trim(parts[0], " ")
	nextmodulesfull := strings.Trim(parts[1], " ")
	nextmodulesparts := strings.Split(nextmodulesfull, ", ")

	var m module

	if mname == BROADCASTER {
		m = module{
			name:     BROADCASTER,
			mtype:    BROADCASTER,
			received: nil,
		}
	} else if mname[0] == '%' {
		mname = mname[1:]
		m = module{
			name:          mname,
			mtype:         FLIPFLOP,
			state:         OFF,
			receivedPulse: make([]int, 0),
			// received: make(map[string]int),
		}
	} else if mname[0] == '&' {
		mname = mname[1:]
		m = module{
			name:  mname,
			mtype: CONJUNCTION,
			// received: make(map[string]int),
		}
	}

	for _, next := range nextmodulesparts {
		appendOrCreate(next, []string{mname}, prevmodules)
	}

	// Create the module
	modules[mname] = m
	nextm := make([]string, 0)
	// Create nexts for that module (assuming they are only filled once)
	nextm = append(nextm, nextmodulesparts...)
	nextmodules[mname] = nextm
}

var nextmodules = make(map[string][]string)
var prevmodules = make(map[string][]string)
var modules = make(map[string]module)

func flip(state string) string {
	if state == OFF {
		return ON
	}
	if state == ON {
		return OFF
	}
	panic("flip?")
}

func cacheMap(m map[string]int) string {
	b := &strings.Builder{}
	for _, v := range m {
		b.WriteString(strconv.Itoa(v))
	}
	return b.String()
}

var cache = make(map[string]int)
var original string = ""

// deliverSignal marks the signal as having been sent to the other modules.
func deliverSignal(step int, sig signal) []string {
	delivered := []string{}
	next := nextmodules[sig.sender]
	// go through all connected modules
	for _, nextm := range next {
		recvm, ok := modules[nextm]
		if !ok {
			continue
		}
		if recvm.mtype == BROADCASTER {
			panic("lol broadcaster")
		} else if recvm.mtype == FLIPFLOP {
			recvm.receivedPulse = append(recvm.receivedPulse, sig.strenght)
			// flip the state only when sending to process incoming pulses
			// accordingly
			// if sig.strenght == 0 {
			// 	recvm.state = flip(recvm.state)
			// }

			// remember to set the map
			modules[nextm] = recvm

			// mark the received pulse strenght
		} else if recvm.mtype == CONJUNCTION {
			recvm.received[sig.sender] = sig.strenght
			modules[nextm] = recvm
		}
		delivered = append(delivered, nextm)
	}
	return delivered
}

var globps int
var globml int
var globkh int
var globmk int

// calculateSendingSignal finds out what signal to send further. It returns
// false as a second value if it should not send any signal further.
func calculateSendingSignal(step int, mname string) (signal, bool) {
	if mname == BROADCASTER {
		return signal{
			sender:   BROADCASTER,
			strenght: 0,
		}, true
	}

	// "tester" module that doesn't output
	m, ok := modules[mname]
	if !ok {
		return signal{}, false
	}
	// This was a stupid hack I used during the contest. I figured out the
	// number of inputs these conjuction modules had and then I noticed these
	// sequences at some point become arithmetic sequences. So I figured I just
	// had to find a common number of these arithmetic sequences.
	//
	// These are the sequences I got (I stopped randomly the execution, it
	// doesn't matter what the added constant is because the sequences generate
	// the same numbers anyway)
	//
	// kh 4001a + 176043
	// ml 3823b + 175857
	// mk 3877c + 174464
	// ps 3847d + 173114
	//
	// Put these in wolfram alpha to find the following solution (here is the
	// exact query)
	// https://www.wolframalpha.com/input?i=solve+4001a+%2B+176043+%3D+3823b+%2B+175857+%3D3877c+%2B+174464+%3D+3847d+%2B+173114+over+the+integers
	//
	// An integer solution is for example a = 57019352993 + 57019353037n
	// So you could just replace a with 57019352993, compute kh, add 1 since
	// you start steps from 0 and you're done.
	if m.mtype == CONJUNCTION {
		if m.name == "ps" && cacheMap(m.received) == "1111111" {
			fmt.Println("ps", step-globps, step)
			globps = step
		}
		if m.name == "ml" && cacheMap(m.received) == "1111111111" {
			fmt.Println("ml", step-globml, step)
			globml = step
		}
		if m.name == "kh" && cacheMap(m.received) == "1111111" {
			fmt.Println("kh", step-globkh, step)
			globkh = step
		}
		if m.name == "mk" && cacheMap(m.received) == "1111111" {

			fmt.Println("mk", step-globmk, step)
			globmk = step
		}

		for _, s := range m.received {
			//
			// a low pulse exists, send a high pulse
			if s == 0 {
				return signal{sender: mname, strenght: 1}, true
			}
		}
		// all inputs sent high strenght (this was off by one...)
		return signal{sender: mname, strenght: 0}, true
	} else if m.mtype == FLIPFLOP {
		var ret signal
		var retb bool
		// received a low pulse, should flip state and send accordingly
		firstReceived := m.receivedPulse[0]
		m.receivedPulse = m.receivedPulse[1:]
		if firstReceived == 0 {
			// flip the state just before sending the signal, not when it is
			// delivered.
			m.state = flip(m.state)
			if m.state == OFF {
				ret, retb = signal{sender: mname, strenght: 0}, true
			} else {
				ret, retb = signal{sender: mname, strenght: 1}, true
			}
			// received high pulse, ignore
		} else if firstReceived == 1 {
			ret, retb = signal{}, false
		}
		// Since this step involved removing a signal from the queue and
		// switching states, update the module.
		modules[mname] = m
		return ret, retb
	}
	panic("lol")
	// return signal{}, false
}

func sendAllSignals(step int) (int, int) {
	q := []string{}

	highs := 0
	lows := 1 // signal sent to broadcasetr

	// start with BROADCASTER and bfs to all modules
	q = append(q, BROADCASTER)
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		sig, ok := calculateSendingSignal(step, top)
		if !ok {
			// fmt.Println("unforuante", top)
			continue
		}
		// fmt.Printf("debug %+v %+v %v\n", modules[top], sig, nextmodules[top])

		// Remember that you are sending the signal to multiple modules.
		if sig.strenght == 1 {
			highs += len(nextmodules[top])
		} else {
			lows += len(nextmodules[top])
		}
		delivered := deliverSignal(step, sig)
		q = append(q, delivered...)
	}
	return highs, lows
}

func part1() {
	for _, line := range inputLines {
		parseLine(line)
	}
	fmt.Println(modules)
	fmt.Println("next", nextmodules)
	fmt.Println("prev", prevmodules)
	// for each CONJUNCTION module fill the received map
	for mname := range modules {
		m := modules[mname]
		if m.mtype == CONJUNCTION {
			prev := make(map[string]int)
			for _, p := range prevmodules[m.name] {
				prev[p] = 0
			}
			fmt.Println("prevvvv", prev)
			// modules[mname].received = prev
			m.received = prev
			// wtf?
			modules[mname] = m
		}
	}
	fmt.Printf("modules, %+v\n", modules)
	fmt.Printf("next, %+v\n", nextmodules)
	highs, lows := 0, 0
	for i := 0; i < 10000000; i++ {
		highsx, lowsx := sendAllSignals(i)
		highs += highsx
		lows += lowsx
	}
	fmt.Println(highs, lows, highs*lows)
}

// Notes for part 2
// There is only the conj module pointing to rx
// there are 4 conj modules pointing to the last conj module
// fmt.Println(modules["th"].received) <- kh
// fmt.Println(modules["pd"].received) <- mk
// fmt.Println(modules["xc"].received) <- ps
// fmt.Println(modules["bp"].received) <- ml
// And there are 4 conj modules pointing only to the latter
// 4 conj modules (and their only parent)
// fmt.Println(modules["ps"].received)
// fmt.Println(modules["ml"].received)
// fmt.Println(modules["kh"].received)
// fmt.Println(modules["mk"].received)
// One of the latter conj modules needs to output 0
// rx is the last one in the chain.
// One of ps, ml, kh and mk needs to be full 1.

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
