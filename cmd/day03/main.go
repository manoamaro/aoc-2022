package main

import (
	"fmt"
	"strings"

	"manoamaro.github.com/aoc-2022/internal"
)

func main() {
	input, err := internal.ReadInput(3)
	if err != nil {
		panic(err)
	}

	values := map[rune]int{}

	for r := 'a'; r <= 'z'; r++ {
		values[r] = int(r) - 96
	}

	for r := 'A'; r <= 'Z'; r++ {
		values[r] = (int(r) - 64) + 26
	}

	rucksacks := strings.Split(input, "\n")

	total_priority := 0

	for _, rucksack := range rucksacks {
		size := len(rucksack)
		c1 := rucksack[0 : size/2]
		c2 := rucksack[size/2:]

		r := map[rune]bool{}
		for _, c := range c1 {
			if strings.ContainsRune(c2, c) {
				r[c] = true
			}
		}

		p := 0

		for v, _ := range r {
			p += values[v]
		}

		total_priority += p
	}

	fmt.Printf("Part 1: %d\n", total_priority)

	total_priority_p2 := 0

	for i := 0; i < len(rucksacks)-3; i += 3 {
		group := rucksacks[i : i+3]

		r := map[rune]int{}

		for g, rucksack := range group {
			for _, c := range rucksack {
				r[c] |= (1 << g)
			}
		}

		for c, v := range r {
			if v == 7 {
				total_priority_p2 += values[c]
			}
		}
	}

	fmt.Printf("Part 2: %d\n", total_priority_p2)
}
