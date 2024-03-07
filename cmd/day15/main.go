package main

import (
	"regexp"
	"strconv"
)

var regexIngredients = regexp.MustCompile(`(\w+) (-?\d+)`)

func parseIngredients(l string) map[string]int {
	m := make(map[string]int)

	matches := regexIngredients.FindAllStringSubmatch(l, -1)
	for _, match := range matches {
		val, _ := strconv.Atoi(match[2])
		m[match[1]] = val
	}

	return m
}
