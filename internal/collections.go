package internal

import "strconv"

func SumInts(v []int) int {
	r := 0
	for _, i := range v {
		r += i
	}
	return r
}

func MapToInt(in []string) []int {
	r := make([]int, len(in))
	for i, v := range in {
		r[i], _ = strconv.Atoi(v)
	}
	return r
}
