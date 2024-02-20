package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTally(t *testing.T) {
	type testCase struct {
		input  string
		result int
	}

	testCases := []testCase{
		{
			`[1,2,3]`, 6,
		},
		{
			`{"a":2,"b":4}`, 6,
		},
		{
			`[[[3]]]`, 3,
		},
		{
			`{"a":{"b":4},"c":-1}`, 3,
		},
		{
			`[]`, 0,
		},
		{
			`{}`, 0,
		},
		{
			`false`, 0,
		},
		{
			`"foo"`, 0,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			var n interface{}
			if err := json.Unmarshal([]byte(tc.input), &n); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.result, tally(n))
		})
	}
}

func TestTallyNoRed(t *testing.T) {
	type testCase struct {
		input  string
		result int
	}

	testCases := []testCase{
		{
			`[1,2,3]`, 6,
		},
		{
			`{"a":2,"b":4}`, 6,
		},
		{
			`[[[3]]]`, 3,
		},
		{
			`{"a":{"b":4},"c":-1}`, 3,
		},
		{
			`[]`, 0,
		},
		{
			`{}`, 0,
		},
		{
			`false`, 0,
		},
		{
			`"foo"`, 0,
		},
		{
			`[1,{"c":"red","b":2},3]`, 4,
		},
		{
			`{"d":"red","e":[1,2,3,4],"f":5}`, 0,
		},
		{
			`[1,"red",3]`, 4,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprint("test case ", i), func(t *testing.T) {
			var n interface{}
			if err := json.Unmarshal([]byte(tc.input), &n); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.result, tallyNoRed(n))
		})
	}
}
