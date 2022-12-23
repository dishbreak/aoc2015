package main

import (
	"crypto/md5"
	"fmt"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day4.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part 1: %d\n", part1(input[0]))
	fmt.Printf("Part 2: %d\n", part2(input[0]))
}

func part1(input string) int {
	for i := 0; ; i++ {
		h := generateHash(input, i)
		if strings.HasPrefix(h, "00000") {
			return i
		}
	}
}

func part2(input string) int {
	for i := 0; ; i++ {
		h := generateHash(input, i)
		if strings.HasPrefix(h, "000000") {
			return i
		}
	}
}

func generateHash(secret string, seed int) string {
	data := fmt.Sprint(secret, seed)

	h := md5.New()

	h.Write([]byte(data))
	r := h.Sum(nil)

	return fmt.Sprintf("%x", r)
}
