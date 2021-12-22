package day06

import (
	"log"
	"strings"

	"aoc"
)

const (
	part1Days = 80
	part2Days = 256
)

func Part1(lines []string) int {
	initial := aoc.MustParseNumbers(func() []string {
		return strings.Split(lines[0], ",")
	}())

	source := initial

	var nextFish, addon []int64
	for i := 0; i < part1Days; i++ {
		for _, x := range source {
			a, b := nextTimer(x)
			if a == -1 {
				log.Fatalln("unexpected value:", a, "iteration:", i)
			}
			nextFish = append(nextFish, a)
			if b != -1 {
				addon = append(addon, b)
			}
		}
		nextFish = append(nextFish, addon...)

		source = nextFish

		nextFish = make([]int64, 0)
		addon = addon[:0]
	}

	return len(source)
}

func Part2(lines []string) int {
	initial := aoc.MustParseNumbers(func() []string {
		return strings.Split(lines[0], ",")
	}())

	m := make(map[int64]int, 8)
	for _, v := range initial {
		m[v]++
	}

	for i := 0; i < part2Days; i++ {
		m_ := make(map[int64]int, 8)
		for k := range m {
			if k == 0 {
				m_[8] += m[0]
				m_[6] += m[0]
				continue
			}
			m_[k-1] += m[k]
		}
		m = m_
	}

	var sum int
	for _, v := range m {
		sum += v
	}

	return sum
}

func nextTimer(v int64) (newVal, newFish int64) {
	switch v {
	case 3:
		return 2, -1
	case 2:
		return 1, -1
	case 1:
		return 0, -1
	case 0:
		return 6, 8
	case 5:
		return 4, -1
	case 4:
		return 3, -1
	case 8:
		return 7, -1
	case 7:
		return 6, -1
	case 6:
		return 5, -1
	default:
		return -1, -1
	}
}

func init() {
	aoc.RegisterSolution(6, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
