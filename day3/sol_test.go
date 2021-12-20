package day3

import (
	"testing"

	"aoc"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want uint64
	}{
		{
			"sample",
			aoc.MustReadLinesFromFile("../input/day3.sample.txt"),
			198,
		},
		{
			"input",
			aoc.MustReadLinesFromFile("../input/day3.input.txt"),
			2583164,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part1(tt.arg)
			if got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want uint64
	}{
		{
			"sample",
			aoc.MustReadLinesFromFile("../input/day3.sample.txt"),
			230,
		},
		{
			"input",
			aoc.MustReadLinesFromFile("../input/day3.input.txt"),
			2784375,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Part2(tt.arg)
			if got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBinaryToUint64(t *testing.T) {
	tests := []struct {
		arg  string
		want uint64
	}{
		{
			"10110",
			22,
		},
		{
			"01001",
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			if got := BinaryToUint64(tt.arg); got != tt.want {
				t.Errorf("BinaryToUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOxyRating(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want uint64
	}{
		{
			"sample",
			aoc.MustReadLinesFromFile("../input/day3.sample.txt"),
			23,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OxyRating(tt.arg); got != tt.want {
				t.Errorf("OxyRating() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCO2Rating(t *testing.T) {
	tests := []struct {
		name string
		arg  []string
		want uint64
	}{
		{
			"sample",
			aoc.MustReadLinesFromFile("../input/day3.sample.txt"),
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CO2Rating(tt.arg); got != tt.want {
				t.Errorf("CO2Rating() = %v, want %v", got, tt.want)
			}
		})
	}
}
