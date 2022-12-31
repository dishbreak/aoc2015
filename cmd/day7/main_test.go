package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	input := []string{
		"NOT x -> h",
		"NOT y -> i",
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
	}

	outputs := map[string]uint16{
		"d": 72,
		"e": 507,
		"f": 492,
		"g": 114,
		"h": 65412,
		"i": 65079,
		"x": 123,
		"y": 456,
	}

	for k, v := range outputs {
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, v, runCircuit(input, k))
		})
	}
}
