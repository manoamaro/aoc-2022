package main

import (
	"fmt"

	"manoamaro.github.com/aoc-2022/internal"
)

func main() {
	input, err := internal.ReadInput(6)
	if err != nil {
		panic(err)
	}

	marker := 0

	for i := 4; i < len(input); i++ {
		s := input[i-4 : i]
		m := map[rune]internal.Void{}
		for _, v := range s {
			m[v] = internal.Void{}
		}
		if len(m) == 4 {
			marker = i
			break
		}
	}

	fmt.Printf("Part 1: %v\n", marker)

	for i := 14; i < len(input); i++ {
		s := input[i-14 : i]
		m := map[rune]internal.Void{}
		for _, v := range s {
			m[v] = internal.Void{}
		}
		if len(m) == 14 {
			marker = i
			break
		}
	}

	fmt.Printf("Part 2: %v\n", marker)

}
