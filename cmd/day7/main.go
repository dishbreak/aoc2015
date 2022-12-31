package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day7.txt")
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	pt1 := part1(input)
	fmt.Printf("Part 1: %d\n", pt1)

	pt2 := part2(input, pt1)
	fmt.Printf("Part 2: %d\n", pt2)
}

func part1(input []string) uint16 {
	return runCircuit(input, "a")
}

func part2(input []string, pt1 uint16) uint16 {
	// ok this is a little sneaky
	// we override the value for b by literally replacing the instruction.
	for i, v := range input {
		if strings.HasSuffix(v, " -> b") {
			input[i] = fmt.Sprintf("%d -> b", pt1)
			break
		}
	}

	return runCircuit(input, "a")
}

var (
	wireLabel  = regexp.MustCompile(`([a-z]+)`)
	passThru   = regexp.MustCompile(`^(\d+|\w+)$`)
	twoArgExpr = regexp.MustCompile(`(\d+|\w+) (AND|OR|LSHIFT|RSHIFT) (\d+|\w+)`)
	notExpr    = regexp.MustCompile(`NOT (\w+)`)
)

type expression struct {
	raw, input, output string
}

func intOrWire(wires map[string]uint16, s string) uint16 {
	parsed, err := strconv.Atoi(s)
	if err != nil {
		return wires[s]
	}
	return uint16(parsed)
}

func (e expression) evaluate(wires map[string]uint16) uint16 {
	switch {
	case passThru.MatchString(e.input):
		return intOrWire(wires, e.input)
	case notExpr.MatchString(e.input):
		// NOT never takes a literal value, even though it could.
		// so we look for a wire label, look it up, and return the bitwise NOT
		labels := wireLabel.FindAllString(e.input, -1)
		return ^wires[labels[0]]
	default:
		m := twoArgExpr.FindStringSubmatch(e.input)
		lArg, rArg := intOrWire(wires, m[1]), intOrWire(wires, m[3])
		switch m[2] {
		case "AND":
			return lArg & rArg
		case "OR":
			return lArg | rArg
		case "LSHIFT":
			return lArg << rArg
		case "RSHIFT":
			return lArg >> rArg
		}
	}
	panic(errors.New("unexpected expression"))
}

func runCircuit(input []string, outputWire string) uint16 {

	exprs := make([]expression, len(input))
	visited := make(map[string]bool)

	for i, l := range input {
		parts := strings.Split(l, " -> ")
		exprs[i].input = parts[0]
		exprs[i].output = parts[1]
		exprs[i].raw = l
	}

	wires := make(map[string]uint16)
	depsMet := func(deps []string) bool {
		for _, dep := range deps {
			if !visited[dep] {
				return false
			}
		}
		return true
	}

	for len(exprs) > 0 {
		nExprs := make([]expression, 0)

		for _, expr := range exprs {
			deps := wireLabel.FindAllString(expr.input, -1)
			if !depsMet(deps) {
				nExprs = append(nExprs, expr)
				continue
			}
			result := expr.evaluate(wires)
			wires[expr.output] = result
			visited[expr.output] = true
		}

		if len(exprs) == len(nExprs) {
			panic(errors.New("failed to clear expressions on last round"))
		}
		exprs = nExprs
	}

	return wires[outputWire]
}
