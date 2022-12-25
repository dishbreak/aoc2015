package main

import (
	"fmt"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day5.txt")
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
		if isNiceString(l) {
			acc++
		}
	}

	return acc
}

func isNiceString(l string) bool {
	return !hasNaughtyString(l) && hasThreeVowels(l) && hasDoubleLetter(l)
}

func hasThreeVowels(input string) bool {
	letters := [26]int{}

	for _, c := range input {
		letters[c-'a']++
	}

	acc := 0
	for _, c := range "aeiou" {
		acc += letters[c-'a']
	}

	return acc >= 3
}

func hasNaughtyString(input string) bool {
	pairs := make(map[string]bool)
	for i := range input {
		if i == len(input)-1 {
			continue
		}
		pairs[input[i:i+2]] = true
	}

	for _, s := range []string{"ab", "cd", "pq", "xy"} {
		if pairs[s] {
			return true
		}
	}

	return false
}

func hasDoubleLetter(input string) bool {
	for i := range input {
		if i == 0 {
			continue
		}
		if input[i-1] == input[i] {
			return true
		}
	}
	return false
}

func part2(input []string) int {
	acc := 0
	for _, l := range input {
		if isNicerString(l) {
			acc++
		}
	}

	return acc
}

func isNicerString(input string) bool {
	return hasSandwichPair(input) && hasNonOverlappingDupePair(input)
}

func hasSandwichPair(input string) bool {
	for i := range input {
		if i < 2 {
			continue
		}

		if input[i-2] == input[i] {
			return true
		}
	}
	return false
}

func hasNonOverlappingDupePair(input string) bool {
	pairs := make(map[string][]int)

	for i := range input {
		if i == len(input)-1 {
			continue
		}
		pair := input[i : i+2]
		if _, ok := pairs[pair]; !ok {
			pairs[pair] = make([]int, 0)
		}
		pairs[pair] = append(pairs[pair], i)
	}

	for key, idxs := range pairs {
		if len(idxs) == 1 {
			continue
		}

		if key[0] != key[1] {
			return true
		}

		if len(idxs) >= 3 {
			return true
		}

		if idxs[0] != idxs[1]-1 {
			return true
		}
	}

	return false
}
