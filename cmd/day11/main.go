package main

import (
	"fmt"
)

const input string = "hepxcrrq"

func main() {
	fmt.Printf("Part 1: %s\n", part1())
	fmt.Printf("Part 2: %s\n", part2())
}

func part1() string {
	return nextValidPassword(
		input,
		oneIncreasingStraight,
		noConfusingCharacters,
		twoNonOverlappingPairs,
	)
}

func part2() string {
	pw := nextValidPassword(
		input,
		oneIncreasingStraight,
		noConfusingCharacters,
		twoNonOverlappingPairs,
	)
	return nextValidPassword(
		pw,
		oneIncreasingStraight,
		noConfusingCharacters,
		twoNonOverlappingPairs,
	)
}

type passwordRule func(string) bool

func nextValidPassword(s string, rules ...passwordRule) string {
	isValid := func(s string) bool {
		for _, r := range rules {
			if !r(s) {
				return false
			}
		}
		return true
	}

	for s = nextPassword(s); !isValid(s); s = nextPassword(s) {
	}

	return s
}

func increment(s []byte, pos int) []byte {
	if s[pos] != 'z' {
		s[pos]++
		return s
	}

	s[pos] = 'a'

	if pos == 0 {
		return s
	}

	return increment(s, pos-1)
}

func nextPassword(s string) string {
	return string(increment([]byte(s), len(s)-1))
}

// 01234567
func oneIncreasingStraight(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1]-1 && s[i] == s[i+2]-2 {
			return true
		}
	}
	return false
}

func twoNonOverlappingPairs(s string) bool {
	skip := false
	firstPairIdx := -1
	for i := 0; i < len(s)-1; i++ {
		if skip {
			skip = false
			continue
		}
		if s[i] == s[i+1] {
			skip = true
			if firstPairIdx == -1 {
				firstPairIdx = i
				continue
			}
			if s[i] == s[firstPairIdx] {
				continue
			}
			return true
		}
	}
	return false
}

func noConfusingCharacters(s string) bool {
	for _, c := range s {
		switch c {
		case 'i':
			fallthrough
		case 'o':
			fallthrough
		case 'l':
			return false
		default:
			continue
		}
	}
	return true
}
