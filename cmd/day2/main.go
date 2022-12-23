package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day2.txt")
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
		edges := toLWH(l)
		acc += getWrappingArea(edges[0], edges[1], edges[2])
	}
	return acc
}

func toLWH(line string) [3]int {
	parts := strings.Split(line, "x")

	ret := [3]int{}

	for i, p := range parts {
		ret[i], _ = strconv.Atoi(p)
	}

	return ret
}

func getWrappingArea(l, w, h int) int {
	faces := []int{l * w, l * h, w * h}
	sort.Ints(faces)

	acc := 0
	for i, a := range faces {
		if i == 0 {
			acc += a
		}
		acc += 2 * a
	}
	return acc
}

func getRibbonLength(l, w, h int) int {
	perims := []int{l + w, l + h, w + h}
	sort.Ints(perims)

	return 2*perims[0] + l*w*h
}

func part2(input []string) int {
	acc := 0
	for _, l := range input {
		edges := toLWH(l)
		acc += getRibbonLength(edges[0], edges[1], edges[2])
	}
	return acc
}
