package day21

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
		// TODO: Add test cases.
		{
			"sample",
			args{lines: aoc.MustReadLinesFromFile("../input/day21.sample.txt")},
			739785,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day21.input.txt")},
			1004670,
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
			args{lines: aoc.MustReadLinesFromFile("../input/day21.sample.txt")},
			444356092776315,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day21.input.txt")},
			492043106122795,
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
