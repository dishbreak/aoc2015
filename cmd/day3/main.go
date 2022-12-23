package main

import (
	"fmt"
	"image"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day3.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
	fmt.Printf("Part 2: %d\n", part2(input[0]))
}

var directions = map[rune]image.Point{
	'^': image.Pt(0, 1),
	'v': image.Pt(0, -1),
	'>': image.Pt(1, 0),
	'<': image.Pt(-1, 0),
}

func part1(input string) int {
	houses := make(map[image.Point]bool)
	cursor := image.Pt(0, 0)
	houses[cursor] = true

	for _, c := range input {
		cursor = cursor.Add(directions[c])
		houses[cursor] = true
	}

	return len(houses)
}

func part2(input string) int {
	houses := make(map[image.Point]bool)
	cursors := [2]image.Point{}
	houses[cursors[0]] = true

	for i, c := range input {
		cursors[i%2] = cursors[i%2].Add(directions[c])
		houses[cursors[i%2]] = true
	}

	return len(houses)
}
