package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextPassword(t *testing.T) {
	type testCase struct {
		before, after string
	}

	testCases := []testCase{
		{"aaaaaaa", "aaaaaab"},
		{"aaaaaab", "aaaaaac"},
		{"aaaaaaz", "aaaaaba"},
		{"aaaazzz", "aaabaaa"},
		{"zzzzzzz", "aaaaaaa"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.after, nextPassword(tc.before))
		})
	}
}

func TestIncreasingStraight(t *testing.T) {
	type testCase struct {
		password string
		result   bool
	}

	testCases := []testCase{
		{"hijklmmn", true},
		{"abbceffg", false},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
		{"resrdabc", true},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.result, oneIncreasingStraight(tc.password))
		})
	}
}

func TestTwoNonOverlappingPairs(t *testing.T) {
	type testCase struct {
		password string
		result   bool
	}

	testCases := []testCase{
		{"hijklmmn", false},
		{"abbceffg", true},
		{"abbcegjk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
		{"resrdabc", false},
		{"aabaafjk", false},
		{"aaabcdef", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.result, twoNonOverlappingPairs(tc.password))
		})
	}
}

func TestNoConfusingCharacters(t *testing.T) {
	type testCase struct {
		password string
		result   bool
	}

	testCases := []testCase{
		{"hijklmmn", false},
		{"abbceffg", true},
		{"abbceijk", false},
		{"abcdffaa", true},
		{"ghjaabcc", true},
		{"resroabc", false},
		{"aabaafik", false},
		{"aaalcdef", false},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.result, noConfusingCharacters(tc.password))
		})
	}
}

func TestNextValidPassword(t *testing.T) {
	type testCase struct {
		prev, next string
	}

	testCases := []testCase{
		{"abcdefgh", "abcdffaa"},
		{"ghijklmn", "ghjaabcc"},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.next, nextValidPassword(
				tc.prev,
				oneIncreasingStraight,
				twoNonOverlappingPairs,
				noConfusingCharacters,
			))
		})
	}

}
