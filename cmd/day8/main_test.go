package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testFile = `""
"abc"
"aaa\"aaa"
"\x27"`

func TestPart1(t *testing.T) {
	input := strings.Split(testFile, "\n")
	assert.Equal(t, 12, part1(input))
}
