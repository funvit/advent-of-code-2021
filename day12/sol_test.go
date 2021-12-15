package day12

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
		want int
	}{
		{
			"sample",
			args{lines: aoc.MustReadLinesFromFile("../input/day12.sample.txt")},
			10,
		},
		{
			"sample 2",
			args{lines: aoc.MustReadLinesFromFile("../input/day12.sample2.txt")},
			19,
		},
		{
			"sample 3",
			args{lines: aoc.MustReadLinesFromFile("../input/day12.sample3.txt")},
			226,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day12.input.txt")},
			4792,
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
		want int
	}{
		{
			"sample",
			args{lines: aoc.MustReadLinesFromFile("../input/day12.sample.txt")},
			36,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day12.input.txt")},
			133360,
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
