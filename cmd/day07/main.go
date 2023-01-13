package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"manoamaro.github.com/aoc-2022/internal"
)

type Obj struct {
	name     string
	parent   *Obj
	children []*Obj
	size     int
}

var cdReg = regexp.MustCompile(`\$ cd (.+)`)
var lsReg = regexp.MustCompile(`\$ ls`)
var dirReg = regexp.MustCompile(`dir (.+)`)
var fileReg = regexp.MustCompile(`(\d+) (.+)`)

func main() {
	input, err := internal.ReadInput(7)
	if err != nil {
		panic(err)
	}

	root := &Obj{}

	currentDir := root

	lines := strings.Split(input, "\n")

	readingLS := false

	for _, l := range lines {
		if cdReg.MatchString(l) {
			readingLS = false
			name := cdReg.FindStringSubmatch(l)[1]
			if name == ".." {
				currentDir = currentDir.parent
			} else {
				newDir := currentDir.GetChild(name)
				if newDir == nil {
					newDir = &Obj{
						name:   name,
						parent: currentDir,
					}
					currentDir.children = append(currentDir.children, newDir)
				}
				currentDir = newDir
			}
		} else if lsReg.MatchString(l) {
			readingLS = true
		} else if readingLS {
			if dirReg.MatchString(l) {
				name := dirReg.FindStringSubmatch(l)[1]
				if !currentDir.Exists(name) {
					newDir := Obj{
						name:   name,
						parent: currentDir,
					}
					currentDir.children = append(currentDir.children, &newDir)
				}
			} else if fileReg.MatchString(l) {
				p := fileReg.FindStringSubmatch(l)
				size, _ := strconv.Atoi(p[1])
				name := p[2]
				newFile := &Obj{
					name:   name,
					parent: currentDir,
					size:   size,
				}
				currentDir.children = append(currentDir.children, newFile)
			}
		}
	}
	CalculateSizes(root)

	part1 := part1(root)
	fmt.Printf("Part 1: %v\n", part1)

	part2 := part2(root)
	fmt.Printf("Part 2: %v\n", part2)
}

func part1(o *Obj) int {
	total := 0
	if len(o.children) > 0 && o.size <= 100000 {
		total += o.size
	}

	for _, v := range o.children {
		total += part1(v)
	}
	return total
}

func part2(root *Obj) int {
	total := root.size
	freeSpace := 70000000 - total
	spaceToBeFree := 30000000 - freeSpace

	dirs := part2_1(root, spaceToBeFree)
	sort.SliceStable(dirs, func(i, j int) bool {
		return dirs[i].size < dirs[j].size
	})

	return dirs[0].size
}

func part2_1(o *Obj, s int) []*Obj {
	r := make([]*Obj, 0)
	if len(o.children) > 0 && o.size >= s {
		r = append(r, o)
	}

	for _, v := range o.children {
		r = append(r, part2_1(v, s)...)
	}

	return r
}

func (o Obj) Println(l int) {
	acc := ""
	for i := 0; i < l; i++ {
		acc += "-"
	}
	acc += fmt.Sprintf("%s (%d)", o.name, o.size)
	println(acc)
	for _, v := range o.children {
		v.Println(l + 1)
	}
}

func (o *Obj) Exists(name string) bool {
	return o.GetChild(name) != nil
}

func (o *Obj) GetChild(name string) *Obj {
	for _, v := range o.children {
		if v.name == name {
			return v
		}
	}
	return nil
}

func CalculateSizes(o *Obj) {
	total := 0
	for _, v := range o.children {
		CalculateSizes(v)
		total += v.size
	}
	o.size += total
}
