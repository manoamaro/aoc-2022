package main

import (
	"fmt"
	"regexp"
	"strings"

	"manoamaro.github.com/aoc-2022/internal"
)

func main() {
	input, err := internal.ReadInput(5)
	if err != nil {
		panic(err)
	}

	parts := strings.Split(input, "\n\n")

	stacks := parseStack(parts[0])

	commands_raw := strings.Split(parts[1], "\n")

	commands := [][]int{}

	regex_cmd := regexp.MustCompile(`move (\d{1,2}) from (\d{1,2}) to (\d{1,2})`)
	for _, c := range commands_raw {
		if len(c) == 0 {
			continue
		}
		cv := internal.MapToInt(regex_cmd.FindStringSubmatch(c)[1:])
		commands = append(commands, cv)
	}

	for _, c := range commands {
		qtd := c[0]
		f := c[1] - 1
		t := c[2] - 1

		fs := stacks[f][len(stacks[f])-qtd : len(stacks[f])]

		fs = internal.Reverse(fs)

		stacks[f] = stacks[f][0 : len(stacks[f])-qtd]
		stacks[t] = append(stacks[t], fs...)
	}

	fmt.Printf("Part 1: %v\n", getTop(stacks))

	stacks = parseStack(parts[0])

	for _, c := range commands {
		qtd := c[0]
		f := c[1] - 1
		t := c[2] - 1

		fs := stacks[f][len(stacks[f])-qtd : len(stacks[f])]

		stacks[f] = stacks[f][0 : len(stacks[f])-qtd]
		stacks[t] = append(stacks[t], fs...)
	}

	fmt.Printf("Part 2: %v\n", getTop(stacks))

}

var regex = regexp.MustCompile(`(\[[A-Z]\])|(\s{4})`)

func parseStack(input string) [][]string {
	stacks := [][]string{}

	stacks_raw := strings.Split(input, "\n")

	for i := len(stacks_raw) - 2; i >= 0; i-- {
		v := regex.FindAllString(stacks_raw[i], -1)
		for j, v2 := range v {
			if strings.TrimSpace(v2) != "" {
				if len(stacks) <= j {
					stacks = append(stacks, []string{})
				}
				stacks[j] = append(stacks[j], v2)
			}
		}
	}
	return stacks
}

func getTop(stacks [][]string) string {
	top := ""
	for _, v := range stacks {
		top += v[len(v)-1]
	}
	top = strings.ReplaceAll(top, "[", "")
	top = strings.ReplaceAll(top, "]", "")
	return top
}
