package day6

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
			args{lines: aoc.MustReadLinesFromFile("../input/day6.sample.txt")},
			5934,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day6.input.txt")},
			360761,
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
			args{lines: aoc.MustReadLinesFromFile("../input/day6.sample.txt")},
			26984457539,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day6.input.txt")},
			1632779838045,
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
