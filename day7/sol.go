package day7

import (
	"log"
	"sort"
	"strings"

	"aoc"
)

func Part1(lines []string) int64 {

	items := aoc.MustParseNumbers(strings.Split(lines[0], ","))

	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})

	median := items[len(items)/2]
	if len(items)%2 == 0 {
		log.Println("median can be wrong")
	}

	bestPos := median

	var fuelTotal int64
	for _, v := range items {
		if v > bestPos {
			fuelTotal += v - bestPos
			continue
		}
		fuelTotal += bestPos - v
	}

	return fuelTotal
}

func Part2(lines []string) int64 {

	items := aoc.MustParseNumbers(strings.Split(lines[0], ","))
	sort.Slice(items, func(i, j int) bool {
		return items[i] < items[j]
	})

	// brute force.
	// calculate total fuel for all crabs at all positions and return min value.
	max := items[len(items)-1]
	fuels := make([]int64, 0, max)
	for i := int64(0); i < max; i++ {
		var f int64
		for _, v := range items {
			var steps int64
			if v > i {
				steps = v - i
			} else {
				steps = i - v
			}

			for i := int64(1); i <= steps; i++ {
				f += i
			}
		}
		fuels = append(fuels, f)
	}

	sort.Slice(fuels, func(i, j int) bool {
		return fuels[i] < fuels[j]
	})

	// min is better
	return fuels[0]
}

func init() {
	aoc.RegisterSolution(7, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
