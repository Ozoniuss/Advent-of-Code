package main

import (
	"bufio"
	"fmt"
	"os"
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

// deliverSignal marks the signal as having been sent to the other modules.
func deliverSignal(sig signal) []string {
	delivered := []string{}
	next := nextmodules[sig.sender]
	// go through all connected modules
	for _, nextm := range next {

		if nextm == "rx" {
			fmt.Println("here")
			// panic(step)
			// fmt.Println(modules["ps"].received)
			fmt.Println(modules["xc"].received)
			// fmt.Println(modules["ml"].received)
			fmt.Println(modules["th"].received)
			// fmt.Println(modules["zh"].received)
			fmt.Println(modules["pd"].received)
			// fmt.Println(modules["kh"].received)
			// fmt.Println(modules["mk"].received)
			fmt.Println(modules["bp"].received)
		}

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

// calculateSendingSignal finds out what signal to send further. It returns
// false as a second value if it should not send any signal further.
func calculateSendingSignal(mname string) (signal, bool) {
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
	if m.mtype == CONJUNCTION {
		for _, s := range m.received {
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

func sendAllSignals() (int, int) {
	q := []string{}

	highs := 0
	lows := 1 // signal sent to broadcasetr

	// start with BROADCASTER and bfs to all modules
	q = append(q, BROADCASTER)
	for len(q) != 0 {
		top := q[0]
		q = q[1:]

		sig, ok := calculateSendingSignal(top)
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
		delivered := deliverSignal(sig)
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
	for i := 0; i < 100000; i++ {
		highsx, lowsx := sendAllSignals()
		highs += highsx
		lows += lowsx
	}
	fmt.Println(highs, lows, highs*lows)
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
	// part2()
}
