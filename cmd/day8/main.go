package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func calculateLength(s string) int {
	acc := 2

	slashMode, hexMode := false, false
	hexCtr := 0
	for _, c := range s[1 : len(s)-1] {
		if c == '\\' && !slashMode {
			slashMode = true
		} else if hexMode {
			hexCtr--
			if hexCtr == 0 {
				acc += 3
				hexMode = false
			}
		} else if slashMode {
			if c == '\\' || c == '"' {
				acc++
				slashMode = false
			} else if c == 'x' {
				hexMode = true
				hexCtr = 2
				slashMode = false
			}
		}

	}

	return acc
}

func getEncodedChars(s string) int {
	acc := 4

	for _, c := range s[1 : len(s)-1] {
		if c == '\\' || c == '"' {
			acc++
		}
	}

	return acc
}

func main() {
	input, err := lib.GetInput("inputs/day8.txt")
	if err != nil {
		panic(err)
	}

	input = input[:len(input)-1]

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	acc := 0

	for _, l := range input {
		acc += calculateLength(l)
	}

	return acc
}

func part2(input []string) int {
	acc := 0

	for _, l := range input {
		acc += getEncodedChars(l)
	}
	return acc
}
