package day23

import (
	"testing"

	"aoc"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"sample",
			args{lines: aoc.MustReadLinesFromFile("../input/day23.sample.txt")},
			12521,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day23.input.txt")},
			15516,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.lines); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"sample",
			args{lines: aoc.MustReadLinesFromFile("../input/day23.sample2.txt")},
			44169,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day23.input2.txt")},
			45272,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.lines); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
