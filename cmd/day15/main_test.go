package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseIngredient(t *testing.T) {
	type testCase struct {
		input  string
		result map[string]int
	}

	testCases := []testCase{
		{
			"Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8",
			map[string]int{
				"capacity":   -1,
				"durability": -2,
				"flavor":     6,
				"texture":    3,
				"calories":   8,
			},
		},
		{
			"Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3",
			map[string]int{
				"capacity":   2,
				"durability": 3,
				"flavor":     -2,
				"texture":    -1,
				"calories":   3,
			},
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.result, parseIngredients(tc.input))
		})
	}
}
