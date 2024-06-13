package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplaceMolecule(t *testing.T) {
	assert.Equal(t, "HOOH", replaceMolecule("HOH", 0, "H", "HO"))
}

func TestParseInput(t *testing.T) {
	input := `H => HO
H => OH
O => HH

HOH`
	r := strings.NewReader(input)
	molecule, replacements := parseInput(r)
	assert.Equal(t, "HOH", molecule)
	assert.Equal(t, []replacement{{"H", "HO"}, {"H", "OH"}, {"O", "HH"}}, replacements)
}

func TestFindIndexes(t *testing.T) {
	assert.Equal(t, []int{2, 3}, findIndexes("HOHHHO", "HH"))
}

func TestPart1(t *testing.T) {
	type testCase struct {
		input  string
		result int
	}

	testCases := []testCase{
		{
			input: `H => HO
H => OH
O => HH

HOH`,
			result: 4,
		},
		{
			input: `H => HO
H => OH
O => HH

HOHOHO`,
			result: 7,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			r := strings.NewReader(tc.input)
			assert.Equal(t, tc.result, part1(r))
		})
	}
}
