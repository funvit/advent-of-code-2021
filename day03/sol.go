package day03

import (
	"strconv"

	"aoc"
)

func MostCommon(lines []string) string {

	var r []byte

	for i := 0; i < len(lines[0]); i++ {
		var setBits int
		for _, l := range lines {
			if l[i] == '1' {
				setBits++
			}
		}
		if setBits > len(lines)-setBits {
			r = append(r, '1')
			continue
		}
		r = append(r, '0')
	}

	return string(r[:])
}

func LeastCommon(lines []string) string {

	var r []byte

	for i := 0; i < len(lines[0]); i++ {
		var setBits int
		for _, l := range lines {
			if l[i] == '1' {
				setBits++
			}
		}
		if setBits > len(lines)-setBits {
			r = append(r, '0')
			continue
		}
		r = append(r, '1')
	}

	return string(r[:])
}

func BinaryToUint64(str string) uint64 {
	v, err := strconv.ParseUint(str, 2, 64)
	if err != nil {
		panic(err)
	}

	return v
}

func OxyRating(lines []string) uint64 {

	data := make([]string, 0, len(lines))
	data = append(data, lines...)

	for pos := 0; pos < len(data[0]); pos++ {

		var setBits int
		for _, l := range data {
			if l[pos] == '1' {
				setBits++
			}
		}

		var mostCommon uint8
		if setBits >= len(data)-setBits {
			mostCommon = '1'
		} else {
			mostCommon = '0'
		}

		// filter data
		filteredData := make([]string, 0, len(data))
		for _, v := range data {
			if v[pos] == mostCommon {
				filteredData = append(filteredData, v)
			}
		}

		data = filteredData

		if len(data) == 1 {
			break
		}
	}

	return BinaryToUint64(data[0])
}

func CO2Rating(lines []string) uint64 {

	data := make([]string, 0, len(lines))
	data = append(data, lines...)

	for pos := 0; pos < len(data[0]); pos++ {

		var setBits int
		for _, l := range data {
			if l[pos] == '1' {
				setBits++
			}
		}

		var leastCommon uint8
		if setBits >= len(data)-setBits {
			leastCommon = '0'
		} else {
			leastCommon = '1'
		}

		// filter data
		filteredData := make([]string, 0, len(data))
		for _, v := range data {
			if v[pos] == leastCommon {
				filteredData = append(filteredData, v)
			}
		}

		data = filteredData

		if len(data) == 1 {
			break
		}
	}

	return BinaryToUint64(data[0])
}

func Part1(lines []string) uint64 {

	m := MostCommon(lines)
	l := LeastCommon(lines)

	return BinaryToUint64(m) * BinaryToUint64(l)
}

func Part2(lines []string) uint64 {

	oxy := OxyRating(lines)
	co2 := CO2Rating(lines)

	return oxy * co2
}

func init() {
	aoc.RegisterSolution(3, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
