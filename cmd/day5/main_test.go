package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsNiceString(t *testing.T) {
	type testCase struct {
		input    string
		expected bool
	}

	testCases := []testCase{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, isNiceString(tc.input))
		})
	}

}

func TestIsNicerString(t *testing.T) {
	type testCase struct {
		input    string
		expected bool
	}

	testCases := []testCase{
		{"qjhvhtzxzqqjkmpb", true},
		{"aaa", false},
		{"aaaa", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, isNicerString(tc.input))
		})
	}
}
