package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day9.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

type road struct {
	from, to string
}

func part1(r io.Reader) int {
	min := math.MaxInt32
	travel(r, func(d int) {
		if d < min {
			min = d
		}
	})
	return min
}

func part2(r io.Reader) int {
	max := -1
	travel(r, func(d int) {
		if d > max {
			max = d
		}
	})
	return max
}

func travel(r io.Reader, check func(int)) {
	s := bufio.NewScanner(r)
	dists := make(map[road]int)
	cities := make(map[string]bool)
	for s.Scan() {
		l := s.Text()
		if l == "" {
			continue
		}

		parts := strings.Fields(l)
		a, b := parts[0], parts[2]
		d, err := strconv.Atoi(parts[4])
		if err != nil {
			panic(err)
		}
		cities[a], cities[b] = true, true
		dists[road{a, b}], dists[road{b, a}] = d, d
	}

	c := make([]string, 0)
	for city := range cities {
		c = append(c, city)
	}

	sequences := Permute(c)

	for _, route := range sequences {
		d := 0
		for i := 1; i < len(route); i++ {
			d += dists[road{route[i-1], route[i]}]
		}
		check(d)
	}

	return
}
