package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("inputs/day19.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
}

type replacement struct {
	old, new string
}

func parseInput(r io.Reader) (molecule string, replacements []replacement) {
	s := bufio.NewScanner(r)

	for s.Scan() {
		if s.Text() == "" {
			break
		}
		pts := strings.Split(s.Text(), " => ")
		replacements = append(replacements, replacement{pts[0], pts[1]})
	}

	s.Scan()
	molecule = s.Text()
	return
}

func replaceMolecule(s string, idx int, old, new string) string {
	return s[:idx] + new + s[idx+len(old):]
}

func findIndexes(s string, m string) (idxs []int) {
	offset := 0
	for len(s) > len(m) {
		idx := strings.Index(s, m)
		if idx == -1 {
			break
		}
		s = s[idx+1:]
		idxs = append(idxs, offset+idx)
		offset += idx + 1
	}

	return
}

func part1(r io.Reader) int {
	molecule, replacements := parseInput(r)
	molecules := make(map[string]int)
	for _, rep := range replacements {
		idxs := findIndexes(molecule, rep.old)
		for _, idx := range idxs {
			molecules[replaceMolecule(molecule, idx, rep.old, rep.new)]++
		}
	}

	return len(molecules)
}
