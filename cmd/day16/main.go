package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day16.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

var constraints = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

var reProperty = regexp.MustCompile(` (\w+): (\d+)`)

func parse(s string) map[string]int {
	result := make(map[string]int)
	matches := reProperty.FindAllStringSubmatch(s, -1)
	for _, match := range matches {
		result[match[1]], _ = strconv.Atoi(match[2])
	}
	return result
}

var reSueId = regexp.MustCompile(`^Sue (\d+):`)

func part1(r io.Reader) int {
	s := bufio.NewScanner(r)
	for s.Scan() {
		properties := parse(s.Text())
		matched := true
		for prop, val := range properties {
			if constraints[prop] != val {
				matched = false
				continue
			}
		}
		if !matched {
			continue
		}
		match := reSueId.FindStringSubmatch(s.Text())
		sueId, _ := strconv.Atoi(match[1])
		return sueId
	}
	return -1
}

type constraint func(int) bool

func mustEqual(n int) constraint {
	return func(i int) bool {
		return i == n
	}
}

func lessThan(n int) constraint {
	return func(i int) bool {
		return i < n
	}
}

func greaterThan(n int) constraint {
	return func(i int) bool {
		return i > n
	}
}

var limits = map[string]constraint{
	"children":    mustEqual(3),
	"cats":        greaterThan(7),
	"samoyeds":    mustEqual(2),
	"pomeranians": lessThan(3),
	"akitas":      mustEqual(0),
	"vizslas":     mustEqual(0),
	"goldfish":    lessThan(5),
	"trees":       greaterThan(3),
	"cars":        mustEqual(2),
	"perfumes":    mustEqual(1),
}

func part2(r io.Reader) int {
	s := bufio.NewScanner(r)
	for s.Scan() {
		properties := parse(s.Text())
		matched := true
		for prop, val := range properties {
			if !limits[prop](val) {
				matched = false
				continue
			}
		}
		if !matched {
			continue
		}
		match := reSueId.FindStringSubmatch(s.Text())
		sueId, _ := strconv.Atoi(match[1])
		return sueId
	}
	return -1
}
