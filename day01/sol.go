package day01

import "aoc"

type dto1 struct {
	value     int64
	increased bool
}

func CountIncreases(values []int64) int {
	m := make([]dto1, len(values))

	for i := range values {
		if i == 0 {
			m = append(m, dto1{
				value:     values[i],
				increased: false,
			})
			continue
		}
		m = append(m, dto1{
			value:     values[i],
			increased: values[i] > values[i-1],
		})
	}

	var r int
	for _, v := range m {
		if v.increased {
			r++
		}
	}

	return r
}

func GetThreeSums(values []int64) []int64 {

	var r []int64

	for i := range values {
		if len(values)-i < 3 {
			break
		}
		r = append(r,
			values[i]+values[i+1]+values[i+2],
		)
	}

	return r
}

func Part1(lines []string) int {
	return CountIncreases(aoc.MustParseNumbers(lines))
}

func Part2(lines []string) int {
	values := GetThreeSums(aoc.MustParseNumbers(lines))
	return CountIncreases(values)
}

func init() {
	aoc.RegisterSolution(1, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
