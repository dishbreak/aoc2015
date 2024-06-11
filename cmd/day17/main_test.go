package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	assert.Equal(t, 4, part1(25, []int{20, 15, 10, 5, 5}))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 3, part2(25, []int{20, 15, 10, 5, 5}))
}
