package main

import (
	"fmt"
	"strings"
)

const input string = "3113322113"

type seqRecord struct {
	digit, count int
}

func readSequence(input string) []seqRecord {
	if len(input) == 0 {
		return nil
	}

	r := seqRecord{
		digit: int(input[0] - '0'),
		count: 1,
	}

	if len(input) == 1 {
		return []seqRecord{r}
	}

	result := []seqRecord{}
	for i := 1; i < len(input); i++ {
		digit := int(input[i] - '0')
		if digit == r.digit {
			r.count++
			continue
		}
		result = append(result, r)
		r = seqRecord{digit, 1}
	}

	result = append(result, r)

	return result
}

func toString(input []seqRecord) string {
	var b strings.Builder
	for _, r := range input {
		b.WriteString(fmt.Sprintf("%d%d", r.count, r.digit))
	}

	return b.String()
}

func iterate(input string, count int) string {
	s := input

	for i := 0; i < count; i++ {
		s = toString(readSequence(s))
	}

	return s
}

func main() {
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input string) int {
	return len(iterate(input, 40))
}

func part2(input string) int {
	return len(iterate(input, 50))
}
