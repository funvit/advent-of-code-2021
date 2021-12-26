package aoc

import "strconv"

func MustParseInt64(s string) int64 {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func MustParseNumbers(strs []string) []int64 {
	var r []int64

	for _, v := range strs {
		n, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, n)
	}

	return r
}
