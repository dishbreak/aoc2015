package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day1.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
	fmt.Printf("Part 2: %d\n", part2(input[0]))
}

func part1(input string) int {
	floor := 0
	for _, c := range input {
		if c == '(' {
			floor++
			continue
		}
		floor--
	}

	return floor
}

func part2(input string) int {
	floor := 0
	for i, c := range input {
		switch c {
		case '(':
			floor++
		case ')':
			floor--
		}

		if floor == -1 {
			return i + 1
		}
	}
	return -1
}
