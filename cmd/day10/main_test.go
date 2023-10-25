package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSequence(t *testing.T) {
	type testCase struct {
		input    string
		expected []seqRecord
	}

	testCases := []testCase{
		{"", nil},
		{"1", []seqRecord{{1, 1}}},
		{"11", []seqRecord{{1, 2}}},
		{"1211", []seqRecord{{1, 1}, {2, 1}, {1, 2}}},
		{"111221", []seqRecord{{1, 3}, {2, 2}, {1, 1}}},
		{"3113322113", []seqRecord{{3, 1}, {1, 2}, {3, 2}, {2, 2}, {1, 2}, {3, 1}}},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			assert.Equal(t, tc.expected, readSequence(tc.input))
		})
	}
}

func TestToString(t *testing.T) {
	r := []seqRecord{{1, 3}, {2, 2}, {1, 1}}
	assert.Equal(t, "312211", toString(r))
}

func TestIterate(t *testing.T) {
	assert.Equal(t, "312211", iterate("1", 5))
}
