package main

import (
	"fmt"
	"strings"

	"manoamaro.github.com/aoc-2022/internal"
)

func main() {
	input, err := internal.ReadInput(4)
	if err != nil {
		panic(err)
	}

	pairs := strings.Split(input, "\n")

	overlaps := 0

	for _, p := range pairs {
		if len(p) == 0 {
			continue
		}

		s := strings.Split(p, ",")
		e1 := internal.MapToInt(strings.Split(s[0], "-"))
		e2 := internal.MapToInt(strings.Split(s[1], "-"))

		if (e1[0] >= e2[0] && e1[1] <= e2[1]) || (e2[0] >= e1[0] && e2[1] <= e1[1]) {
			overlaps += 1
		}
	}

	fmt.Printf("Part 1: %d\n", overlaps)

	overlaps = 0

	for _, p := range pairs {
		if len(p) == 0 {
			continue
		}

		s := strings.Split(p, ",")
		e1 := internal.MapToInt(strings.Split(s[0], "-"))
		e2 := internal.MapToInt(strings.Split(s[1], "-"))

		if e1[1] >= e2[0] && e1[0] <= e2[1] {
			overlaps += 1
		}
	}

	fmt.Printf("Part 2: %d\n", overlaps)
}
