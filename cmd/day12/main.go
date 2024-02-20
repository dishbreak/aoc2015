package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day12.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

func part1(r io.Reader) int {
	dec := json.NewDecoder(r)
	var n interface{}

	dec.Decode(&n)

	return tally(n)
}

func part2(r io.Reader) int {
	dec := json.NewDecoder(r)
	var n interface{}

	dec.Decode(&n)

	return tallyNoRed(n)
}

func tally(n interface{}) int {
	if nFlt, ok := n.(float64); ok {
		return int(nFlt)
	}

	acc := 0
	if nMap, ok := n.(map[string]interface{}); ok {
		for _, v := range nMap {
			acc += tally(v)
		}
		return acc
	}

	if nSlc, ok := n.([]interface{}); ok {
		for _, v := range nSlc {
			acc += tally(v)
		}
		return acc
	}

	// if it's not an int, a list, or an object, it's something like a string or bool.
	return 0
}

func tallyNoRed(n interface{}) int {
	if nFlt, ok := n.(float64); ok {
		return int(nFlt)
	}

	acc := 0
	if nMap, ok := n.(map[string]interface{}); ok {
		for _, v := range nMap {
			if vStr, ok := v.(string); ok {
				if vStr == "red" {
					return 0
				}
			}
			acc += tallyNoRed(v)
		}
		return acc
	}

	if nSlc, ok := n.([]interface{}); ok {
		for _, v := range nSlc {
			acc += tallyNoRed(v)
		}
		return acc
	}

	// if it's not an int, a list, or an object, it's something like a string or bool.
	return 0
}
