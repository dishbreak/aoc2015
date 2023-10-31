package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	A, B, C := "A", "B", "C"
	input := []string{A, B, C}
	expected := [][]string{
		{A, B, C},
		{C, B, A},
		{B, C, A},
		{A, C, B},
		{B, A, C},
		{C, A, B},
	}

	actual := Permute(input)
	assert.ElementsMatch(t, expected, actual)
}
