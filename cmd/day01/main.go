package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"manoamaro.github.com/aoc-2022/internal"
)

func main() {

	input, err := internal.ReadInput(1)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(input, "\n")

	elves := make([]int, 0)

	tempElf := 0
	for _, line := range lines {
		if len(line) == 0 {
			elves = append(elves, tempElf)
			tempElf = 0
		} else {
			if v, err := strconv.Atoi(line); err != nil {
			} else {
				tempElf += v
			}
		}
	}
	sort.Ints(elves)
	fatElf := elves[len(elves)-1]

	fmt.Printf("Part 1: %d\n", fatElf)

	fatElves := elves[len(elves)-3:]
	sum := internal.SumInts(fatElves)

	fmt.Printf("Part 2: %d\n", sum)
}
