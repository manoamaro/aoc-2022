package main

import (
	"fmt"
	"strings"

	"manoamaro.github.com/aoc-2022/internal"
)

const (
	RockMove     = 0
	PaperMove    = -1
	ScissorsMove = 1
	Rock         = 1
	Paper        = 2
	Scissors     = 3
	Win          = 6
	Draw         = 3
	Lost         = 0
)

var Moves map[string]int = map[string]int{
	"A": RockMove,
	"B": PaperMove,
	"C": ScissorsMove,
	"X": RockMove,
	"Y": PaperMove,
	"Z": ScissorsMove,
}

var MovesValues map[string]int = map[string]int{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func main() {
	input, err := internal.ReadInput(2)
	if err != nil {
		panic(err)
	}

	rounds := strings.Split(input, "\n")

	totalResult := 0

	for _, round := range rounds {
		if len(round) > 0 {
			moves := strings.Split(round, " ")
			o := Moves[moves[0]]
			p := Moves[moves[1]]
			result := Draw
			if p != o {
				if (p == RockMove && o == ScissorsMove) ||
					(p == ScissorsMove && o == PaperMove) ||
					(p == PaperMove && o == RockMove) {
					result = Win
				} else {
					result = Lost
				}
			}

			totalResult += result + MovesValues[moves[1]]

		}
	}

	fmt.Printf("Part 1: %d\n", totalResult)

	totalResult = 0

	for _, round := range rounds {
		if len(round) > 0 {
			moves := strings.Split(round, " ")
			o := moves[0]
			p := moves[1]

			if p == "Y" {
				// Draw
				totalResult += Draw + MovesValues[o]
			} else {
				if p == "Z" {
					// Win
					totalResult += Win
					if Moves[o] == RockMove {
						totalResult += Paper
					} else if Moves[o] == PaperMove {
						totalResult += Scissors
					} else {
						totalResult += Rock
					}
					println(totalResult)
				} else {
					// Lose
					totalResult += Lost
					if Moves[o] == RockMove {
						totalResult += Scissors
					} else if Moves[o] == PaperMove {
						totalResult += Rock
					} else {
						totalResult += Paper
					}
					println(totalResult)
				}
			}
		}
	}

	fmt.Printf("Part 2: %d\n", totalResult)

}
