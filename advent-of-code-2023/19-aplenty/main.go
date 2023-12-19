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

func getConditionFunc(cond string) func(p Part) bool {
	if strings.Index(cond, "<") != -1 {
		parts := strings.Split(cond, "<")
		structMember := parts[0]
		val, _ := strconv.Atoi(parts[1])

		if structMember == "x" {
			return func(p Part) bool {
				return p.x < val
			}
		}
		if structMember == "a" {
			return func(p Part) bool {
				return p.a < val
			}
		}
		if structMember == "s" {
			return func(p Part) bool {
				return p.s < val
			}
		}
		if structMember == "m" {
			return func(p Part) bool {
				return p.m < val
			}
		}
	}
	if strings.Index(cond, ">") != -1 {
		parts := strings.Split(cond, ">")
		structMember := parts[0]
		val, _ := strconv.Atoi(parts[1])

		if structMember == "x" {
			return func(p Part) bool {
				return p.x > val
			}
		}
		if structMember == "a" {
			return func(p Part) bool {
				return p.a > val
			}
		}
		if structMember == "s" {
			return func(p Part) bool {
				return p.s > val
			}
		}
		if structMember == "m" {
			return func(p Part) bool {
				return p.m > val
			}
		}
	}
	panic("wrong")
}

var workflows = make(map[string][]rulefunc)
var workflowsPart2 = make(map[string][]Step)

func applyWorkflowRules(p Part, workflowName string) string {
	// fmt.Println(workflowName, len(workflowName))
	for _, rule := range workflows[workflowName] {
		val := rule(p)
		if val != "continue" {
			return val
		}
	}
	panic("all rules applied and didn't get accepted")
}

func processCondWithValString(condwithval string) CondWithVal {
	if strings.Index(condwithval, "<") != -1 {
		parts := strings.Split(condwithval, "<")
		structMember := parts[0]
		val, _ := strconv.Atoi(parts[1])

		return CondWithVal{
			cond:  "<",
			field: structMember,
			val:   val,
		}
	}
	if strings.Index(condwithval, ">") != -1 {
		parts := strings.Split(condwithval, ">")
		structMember := parts[0]
		val, _ := strconv.Atoi(parts[1])

		return CondWithVal{
			cond:  ">",
			field: structMember,
			val:   val,
		}
	}
	panic("wrong")
}

func processIndividualRule(rulestr string) rulefunc {
	// fmt.Println(rulestr)
	if rulestr == "A" {
		return func(p Part) string { return "accepted" }
	}
	if rulestr == "R" {
		return func(p Part) string { return "rejected" }
	}

	// In this case the rule is a workflow name
	if strings.Index(rulestr, ":") == -1 {
		return func(part Part) string {
			return applyWorkflowRules(part, rulestr)
		}
	}

	ruleParts := strings.Split(rulestr, ":")
	condition := ruleParts[0]

	condFunc := getConditionFunc(condition)
	result := ruleParts[1]
	return func(part Part) string {
		if !condFunc(part) {
			return "continue"
		} else {
			return processIndividualRule(result)(part)
		}
	}
}

// ar part b, we need cond, val and next
func processIndividualRulePartb(rulestr string) Step {
	// fmt.Println(rulestr)
	if rulestr == "A" || rulestr == "B" {
		return Step{
			cwv:  nil,
			next: rulestr,
		}
	}

	// In this case the rule is a workflow name
	if strings.Index(rulestr, ":") == -1 {
		// so return the result that we get by applying all the workflow rules.
		return Step{
			cwv:  nil,
			next: rulestr,
		}
	}

	parts := strings.Split(rulestr, ":")
	cwv := processCondWithValString(parts[0])

	return Step{
		cwv:  &cwv,
		next: parts[1],
	}

}

func processWorkflow(workflowstr string) {
	nameend := strings.Index(workflowstr, "{")
	name := workflowstr[:nameend]
	rulesstr := workflowstr[nameend+1 : len(workflowstr)-1]
	rulesParts := strings.Split(rulesstr, ",")

	workflowRules := make([]rulefunc, 0)
	for _, rp := range rulesParts {
		rulefunction := processIndividualRule(rp)
		workflowRules = append(workflowRules, rulefunction)
	}
	workflows[name] = workflowRules
}

type CondWithVal struct {
	cond  string
	val   int
	field string
}

type Step struct {
	cwv  *CondWithVal
	next string
}

// at part 2, instead of figuring out the function, store each step with cond
// and value.
func processWorkflowPart2(workflowstr string) {
	nameend := strings.Index(workflowstr, "{")
	name := workflowstr[:nameend]

	// define an arrat
	wintervals := make([]interval, 0)
	workflowIntervals[name] = wintervals

	rulesstr := workflowstr[nameend+1 : len(workflowstr)-1]
	rulesParts := strings.Split(rulesstr, ",")

	workflowSteps := make([]Step, 0)
	for _, rp := range rulesParts {
		step := processIndividualRulePartb(rp)
		workflowSteps = append(workflowSteps, step)
		// if step.cwv != nil {
		// 	fmt.Println(*step.cwv, step.next)
		// } else {
		// 	fmt.Println(step.cwv, step.next)
		// }
	}
	workflowsPart2[name] = workflowSteps
}

func processPart(part string) Part {
	actualpart := part[1 : len(part)-1]
	values := strings.Split(actualpart, ",")
	p := Part{}
	for _, valstr := range values {
		eqi := strings.Index(valstr, "=")
		structfield := valstr[:eqi]
		val, _ := strconv.Atoi(valstr[eqi+1:])
		if structfield == "x" {
			p.x = val
		}
		if structfield == "a" {
			p.a = val
		}
		if structfield == "s" {
			p.s = val
		}
		if structfield == "m" {
			p.m = val
		}
	}
	return p
}

type Part struct {
	x int
	m int
	a int
	s int
}

type rulefunc func(part Part) string

func part1() {
	processingParts := false
	parts := make([]Part, 0)

	for _, line := range inputLines {
		if len(line) == 0 {
			processingParts = true
			continue
		}
		if !processingParts {
			processWorkflow(line)
		} else {
			currentPart := processPart(line)
			parts = append(parts, currentPart)
		}
	}
	accepted := make([]Part, 0)
	for _, p := range parts {
		current := "continue"
		for current == "continue" {
			current = applyWorkflowRules(p, "in")
		}
		if current == "accepted" {
			accepted = append(accepted, p)
		}
	}
	s := 0
	for _, ap := range accepted {
		s += ap.x + ap.s + ap.m + ap.a
	}
	fmt.Println(s)
}

type interval struct {
	xranges [2]int
	mranges [2]int
	aranges [2]int
	sranges [2]int
}

func (i interval) length() int {
	return (i.xranges[1] - i.xranges[0]) *
		(i.mranges[1] - i.mranges[0]) *
		(i.aranges[1] - i.aranges[0]) *
		(i.sranges[1] - i.sranges[0])
}

// splitRange splits a range in two, given the condition and value.
// All numbers in the first interval satisgy the condition with the value.
func splitRange(cond string, val int, intrange [2]int) ([2]int, [2]int) {
	// no splitting because the range is empty
	if intrange[0] == intrange[1] {
		return intrange, intrange
	}
	// every number i in first interval satisfies i < val
	if cond == "<" {
		// v <= a; nil, [a,b). No i from [a,b) satisfies i < a ( >= v )
		if val <= intrange[0] {
			return [2]int{0, 0}, intrange
			// v >= b; [a,b), nil. Every i from [a,b) satisfies i < b <= v
		} else if val >= intrange[1] {
			return intrange, [2]int{0, 0}
		} else {
			// Every i from [a,v) satisfies i<v
			return [2]int{intrange[0], val}, [2]int{val, intrange[1]}
		}
	} else if cond == ">" {
		// Every i from [a,b) satisfies i > a-1 >= v
		if val <= intrange[0]-1 {
			return intrange, [2]int{0, 0}
			// No i from [a,b) satisfies i > b-1 <= v
		} else if val >= intrange[1]-1 {
			return [2]int{0, 0}, intrange
		} else {
			// Every i from [v+1, b) satisfies i > v
			return [2]int{val + 1, intrange[1]}, [2]int{intrange[0], val + 1}
		}
	}
	panic("lol")
}

func split(intrvl interval, axis string, cond string, val int) (interval, interval) {
	if axis == "x" {
		respecting, remaining := splitRange(cond, val, intrvl.xranges)
		respectingInterval := intrvl
		remainingInterval := intrvl
		respectingInterval.xranges = respecting
		remainingInterval.xranges = remaining
		return respectingInterval, remainingInterval
	}
	if axis == "m" {
		respecting, remaining := splitRange(cond, val, intrvl.mranges)
		respectingInterval := intrvl
		remainingInterval := intrvl
		respectingInterval.mranges = respecting
		remainingInterval.mranges = remaining
		return respectingInterval, remainingInterval
	}
	if axis == "a" {
		respecting, remaining := splitRange(cond, val, intrvl.aranges)
		respectingInterval := intrvl
		remainingInterval := intrvl
		respectingInterval.aranges = respecting
		remainingInterval.aranges = remaining
		return respectingInterval, remainingInterval
	}
	if axis == "s" {
		respecting, remaining := splitRange(cond, val, intrvl.sranges)
		respectingInterval := intrvl
		remainingInterval := intrvl
		respectingInterval.sranges = respecting
		remainingInterval.sranges = remaining
		return respectingInterval, remainingInterval
	}
	panic("lol intervals")
}

func applySplit(name string, intervals []interval) {

	// clone just to be sure
	intervalscp := slices.Clone(intervals)

	if name == "A" {
		workflowIntervals["A"] = append(workflowIntervals["A"], intervalscp...)
	}
	if name == "B" {
		workflowIntervals["B"] = append(workflowIntervals["B"], intervalscp...)
	}

	steps := workflowsPart2[name]
	for _, step := range steps {
		// no more conditions, just send those intervals to the next condition
		// at this point we may as well just return, it doesn't matter anymore
		if step.cwv == nil {
			applySplit(step.next, intervalscp)
			return
		} else {
			respectingCurrentStep := make([]interval, 0)
			remainingCurrentStep := make([]interval, 0)

			// get all the intervals for which we apply the current step
			for _, icp := range intervalscp {
				respectingInterval, remainingInterval := split(icp, step.cwv.field, step.cwv.cond, step.cwv.val)
				respectingCurrentStep = append(respectingCurrentStep, respectingInterval)
				remainingCurrentStep = append(remainingCurrentStep, remainingInterval)
			}
			applySplit(step.next, respectingCurrentStep)

			// those will go to the next step
			intervalscp = remainingCurrentStep
		}
	}
}

var workflowIntervals = make(map[string][]interval)

func part2() {
	processingParts := false

	for _, line := range inputLines {
		if len(line) == 0 {
			processingParts = true
			continue
		}
		if !processingParts {
			processWorkflowPart2(line)
		} else {
			break
		}
	}
	// fmt.Println(workflowsPart2)

	initialInterval := interval{
		xranges: [2]int{1, 4001},
		mranges: [2]int{1, 4001},
		aranges: [2]int{1, 4001},
		sranges: [2]int{1, 4001},
	}
	// workflowIntervals["in"] = []interval{initialInterval}
	workflowIntervals["A"] = []interval{}
	workflowIntervals["R"] = []interval{}
	fmt.Println(workflowIntervals)

	fmt.Println(workflowsPart2["in"][0].cwv)
	fmt.Println(split(initialInterval, "s", "<", 1351))

	applySplit("in", []interval{initialInterval})

	total := 0
	for _, intrv := range workflowIntervals["A"] {
		l := intrv.length()
		// fmt.Println(l)
		total += l
	}
	fmt.Println(total)
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
