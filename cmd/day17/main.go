package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("inputs/day17.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	input := parseInput(f)
	fmt.Printf("Part 1: %d\n", part1(150, input))
}

func parseInput(r io.Reader) (result []int) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		amt, err := strconv.Atoi(s.Text())
		if err != nil {
			panic(fmt.Errorf("failed to parse input: %w", err))
		}
		result = append(result, amt)
	}

	return
}

func part1(eggNogLiters int, containerSizes []int) int {
	acc := make(map[int]bool)

	type frame struct {
		amountLeft int
		containers int
	}

	q := make([]frame, 0)

	q = append(q, frame{eggNogLiters, 0})

	for len(q) != 0 {
		f := q[0]
		q = q[1:]

		// if there's a negative amount left or all container bits are 1, skip this frame.
		if f.amountLeft < 0 || (f.containers != 0 && f.containers+1&f.containers == 0) {
			continue
		}

		// if we've got no nog left, we're set
		if f.amountLeft == 0 {
			//register the combination of containers as
			acc[f.containers] = true
			continue
		}

		// write a frame for each empty container left in the current frame.
		for i, c := range containerSizes {
			// if the bit for the container is set, it's in use already
			// skip it.
			if f.containers&(1<<i) > 0 {
				continue
			}

			// write a new frame to the queue with an updated amount remaining
			// and the bit for the corresponding container set.
			var nf frame
			nf.containers = f.containers | (1 << i)
			nf.amountLeft = f.amountLeft - c
			q = append(q, nf)
		}
	}

	return len(acc)
}
