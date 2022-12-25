package main

import (
	"fmt"
	"image"
	"regexp"
	"strconv"
	"strings"

	"github.com/dishbreak/aoc-common/lib"
)

func main() {
	input, err := lib.GetInput("inputs/day6.txt")
	if err != nil {
		panic(err)
	}
	input = input[:len(input)-1]

	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))
}

func part1(input []string) int {
	l := lightMatrix{
		lightOperator: &lightMatrixV1{
			bulb: make(map[image.Point]bool),
		},
	}

	for _, line := range input {
		l.Execute(line)
	}

	return l.count()
}

type lightMatrixV1 struct {
	bulb map[image.Point]bool
}

type lightMatrix struct {
	lightOperator
}

type lightOperator interface {
	on(image.Point)
	off(image.Point)
	toggle(image.Point)
	count() int
}

type operator func(image.Point)

func (l *lightMatrix) execute(op operator, min, max image.Point) {
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			op(image.Pt(x, y))
		}
	}
}

func (l *lightMatrix) Count() int {
	return l.count()
}

func (l *lightMatrixV1) toggle(p image.Point) {
	l.bulb[p] = !l.bulb[p]
}

func (l *lightMatrixV1) on(p image.Point) {
	l.bulb[p] = true
}

func (l *lightMatrixV1) off(p image.Point) {
	l.bulb[p] = false
}

var commandRegex = regexp.MustCompile(`(\d+),(\d+) through (\d+),(\d+)$`)

func (l *lightMatrix) Execute(command string) {
	match := commandRegex.FindStringSubmatch(command)
	var min, max image.Point
	min.X, _ = strconv.Atoi(match[1])
	min.Y, _ = strconv.Atoi(match[2])
	max.X, _ = strconv.Atoi(match[3])
	max.Y, _ = strconv.Atoi(match[4])

	var op operator = l.toggle
	switch {
	case strings.HasPrefix(command, "turn on"):
		op = l.on
	case strings.HasPrefix(command, "turn off"):
		op = l.off
	}

	l.execute(op, min, max)
}

func (l *lightMatrixV1) count() int {
	acc := 0
	for _, lit := range l.bulb {
		if lit {
			acc++
		}
	}
	return acc
}

type lightMatrixV2 struct {
	bulb map[image.Point]int
}

func (l *lightMatrixV2) on(p image.Point) {
	l.bulb[p]++
}

func (l *lightMatrixV2) off(p image.Point) {
	if l.bulb[p] == 0 {
		return
	}
	l.bulb[p]--
}

func (l *lightMatrixV2) toggle(p image.Point) {
	l.bulb[p] += 2
}

func (l *lightMatrixV2) count() int {
	acc := 0
	for _, brightness := range l.bulb {
		acc += brightness
	}

	return acc
}

func part2(input []string) int {
	l := &lightMatrix{
		lightOperator: &lightMatrixV2{
			bulb: make(map[image.Point]int),
		},
	}

	for _, line := range input {
		l.Execute(line)
	}

	return l.Count()
}
