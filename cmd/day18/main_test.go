package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := `.#.#.#
...##.
#....#
..#...
#.#..#
####..`
	r := strings.NewReader(input)

	assert.Equal(t, 4, part1(r, 4))
}

func TestPart2(t *testing.T) {
	input := `##.#.#
...##.
#....#
..#...
#.#..#
####.#`
	r := strings.NewReader(input)

	assert.Equal(t, 17, part2(r, 5))
}
