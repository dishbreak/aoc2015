package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"sync"
)

func main() {
	f, err := os.Open("inputs/day13.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("Part 1: %d\n", part1(f))
	f.Seek(0, 0)
	fmt.Printf("Part 2: %d\n", part2(f))
}

var regexpRule regexp.Regexp = *regexp.MustCompile(`^(\w+) would (gain|lose) (\d+) happiness units by sitting next to (\w+).$`)

func parseRuleSet(r io.Reader) map[string]map[string]int {
	s := bufio.NewScanner(r)
	rules := make(map[string]map[string]int)

	for s.Scan() {
		match := regexpRule.FindStringSubmatch(s.Text())
		if match == nil {
			continue
		}

		outerMap, ok := rules[match[1]]
		if !ok {
			outerMap = make(map[string]int)
		}

		delta, _ := strconv.Atoi(match[3])
		if match[2] == "lose" {
			delta = delta * -1
		}

		outerMap[match[4]] = delta
		rules[match[1]] = outerMap
	}

	return rules
}

func calculateHappiness(seats []string, rules map[string]map[string]int) int {
	acc := 0
	for i, seat := range seats {
		left := i - 1
		right := i + 1

		if i == 0 {
			left = len(seats) - 1
		}

		if i == len(seats)-1 {
			right = 0
		}

		acc += rules[seat][seats[left]]
		acc += rules[seat][seats[right]]

	}

	return acc
}

func Permute(values []string) [][]string {
	result := make([][]string, 0)
	result = permute(values, 0, len(values)-1, result)
	return result
}

func permute(values []string, j, n int, acc [][]string) [][]string {
	if j == n {
		result := make([]string, len(values))
		copy(result, values)
		return append(acc, result)

	}

	for i := j; i <= n; i++ {
		values[j], values[i] = values[i], values[j]
		acc = permute(values, j+1, n, acc)
		values[j], values[i] = values[i], values[j]
	}
	return acc
}

func part1(r io.Reader) int {
	reports := make(chan int)

	rules := parseRuleSet(r)

	var seats []string
	for seat, _ := range rules {
		seats = append(seats, seat)
	}

	permutations := Permute(seats)

	var wg sync.WaitGroup
	wg.Add(len(permutations))

	for _, combo := range permutations {
		go func(seats []string) {
			defer wg.Done()
			reports <- calculateHappiness(seats, rules)
		}(combo)
	}

	go func() {
		wg.Wait()
		close(reports)
	}()

	max := math.MinInt
	for result := range reports {
		if result > max {
			max = result
		}
	}

	return max

}

func part2(r io.Reader) int {
	reports := make(chan int)

	rules := parseRuleSet(r)

	var seats []string
	for seat, _ := range rules {
		seats = append(seats, seat)
	}

	innerMap := make(map[string]int)
	for seat, _ := range rules {
		rules[seat]["Vishal"] = 0
		innerMap[seat] = 0
	}
	rules["Vishal"] = innerMap
	seats = append(seats, "Vishal")

	permutations := Permute(seats)

	var wg sync.WaitGroup
	wg.Add(len(permutations))

	for _, combo := range permutations {
		go func(seats []string) {
			defer wg.Done()
			reports <- calculateHappiness(seats, rules)
		}(combo)
	}

	go func() {
		wg.Wait()
		close(reports)
	}()

	max := math.MinInt
	for result := range reports {
		if result > max {
			max = result
		}
	}

	return max

}
