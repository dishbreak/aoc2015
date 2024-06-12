package main

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"os"
)

func main() {
	f, err := os.Open("inputs/day18.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f, 100))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f, 100))
}

var neighbors = []image.Point{
	image.Pt(-1, -1),
	image.Pt(-1, 0),
	image.Pt(-1, 1),
	image.Pt(0, 1),
	image.Pt(1, 1),
	image.Pt(1, 0),
	image.Pt(1, -1),
	image.Pt(0, -1),
}

func parseSpace(r io.Reader) (space map[image.Point]bool, dim int) {
	s := bufio.NewScanner(r)

	space = make(map[image.Point]bool)

	for line := 0; s.Scan(); line++ {
		dim = len(s.Text())
		for col, char := range s.Text() {
			p := image.Pt(col, line)
			val := false
			if char == '#' {
				val = true
			}
			space[p] = val
		}
	}

	return
}

func tallyLitNeighbors(pt image.Point, space map[image.Point]bool) (acc int) {
	for _, n := range neighbors {
		o := pt.Add(n)
		val, ok := space[o]
		if !ok {
			continue
		}
		if val {
			acc++
		}
	}
	return
}

func part1(r io.Reader, steps int) int {

	space, _ := parseSpace(r)

	for step := 0; step < steps; step++ {

		newSpace := make(map[image.Point]bool, len(space))

		for pt, lightOn := range space {
			litNeigbors := tallyLitNeighbors(pt, space)
			litCell := false
			if !lightOn {
				litCell = litNeigbors == 3
			} else {
				litCell = (litNeigbors == 2 || litNeigbors == 3)
			}
			newSpace[pt] = litCell
		}
		space = newSpace
	}

	acc := 0
	for _, lightOn := range space {
		if lightOn {
			acc++
		}
	}

	return acc
}

func part2(r io.Reader, steps int) int {

	space, dim := parseSpace(r)

	corners := []image.Point{
		image.Pt(0, 0),
		image.Pt(0, dim-1),
		image.Pt(dim-1, 0),
		image.Pt(dim-1, dim-1),
	}

	for _, c := range corners {
		space[c] = true
	}

	isCorner := func(p image.Point) bool {
		for _, c := range corners {
			if p.Eq(c) {
				return true
			}
		}
		return false
	}

	for step := 0; step < steps; step++ {

		newSpace := make(map[image.Point]bool, len(space))
		for pt, lightOn := range space {
			if isCorner(pt) {
				newSpace[pt] = true
				continue
			}
			litNeigbors := tallyLitNeighbors(pt, space)
			litCell := false
			if !lightOn {
				litCell = litNeigbors == 3
			} else {
				litCell = (litNeigbors == 2 || litNeigbors == 3)
			}
			newSpace[pt] = litCell
		}
		space = newSpace
	}

	acc := 0
	for _, lightOn := range space {
		if lightOn {
			acc++
		}
	}

	return acc
}
