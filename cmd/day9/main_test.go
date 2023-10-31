package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput string = `
London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141
`

func TestPart1(t *testing.T) {
	r := strings.NewReader(testInput)
	assert.Equal(t, 605, part1(r))
}

func TestPart2(t *testing.T) {
	r := strings.NewReader(testInput)
	assert.Equal(t, 982, part2(r))
}
