package day7

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
			args{lines: aoc.MustReadLinesFromFile("../input/day7.sample.txt")},
			37,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day7.input.txt")},
			349769,
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
			args{lines: aoc.MustReadLinesFromFile("../input/day7.sample.txt")},
			168,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day7.input.txt")},
			99540554,
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
