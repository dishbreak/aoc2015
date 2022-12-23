package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWrappingArea(t *testing.T) {
	type testCase struct {
		l, w, h  int
		expected int
	}

	testCases := []testCase{
		{
			l:        2,
			w:        3,
			h:        4,
			expected: 58,
		},
		{
			l:        1,
			w:        1,
			h:        10,
			expected: 43,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, getWrappingArea(tc.l, tc.w, tc.h))
		})
	}
}

func TestGetRibbonLength(t *testing.T) {
	type testCase struct {
		l, w, h  int
		expected int
	}

	testCases := []testCase{
		{
			l:        2,
			w:        3,
			h:        4,
			expected: 34,
		},
		{
			l:        1,
			w:        1,
			h:        10,
			expected: 14,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, getRibbonLength(tc.l, tc.w, tc.h))
		})
	}
}

func TestPart1(t *testing.T) {
	input := []string{
		"2x3x4",
		"1x1x10",
	}

	assert.Equal(t, 101, part1(input))
}
