package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
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
	fmt.Printf("Part 2: %d\n", part2(150, input))
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

type frame struct {
	amountLeft int
	containers int
	count      int
}

func part1(eggNogLiters int, containerSizes []int) int {
	return len(solve(eggNogLiters, containerSizes))
}

func part2(eggNogLiters int, containerSizes []int) int {
	solutions := solve(eggNogLiters, containerSizes)

	frames := make([]frame, 0)

	for _, v := range solutions {
		frames = append(frames, v)
	}

	sort.Slice(frames, func(i, j int) bool {
		return frames[i].count < frames[j].count
	})

	return func() int {
		for i := 0; i < len(frames)-1; i++ {
			if frames[i+1].count != frames[i].count {
				return i + 1
			}
		}
		return -1
	}()
}

func solve(eggNogLiters int, containerSizes []int) map[int]frame {
	acc := make(map[int]frame)

	q := make([]frame, 0)

	q = append(q, frame{eggNogLiters, 0, 0})

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
			acc[f.containers] = f
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
			nf.count = f.count + 1
			q = append(q, nf)
		}
	}

	return acc
}
