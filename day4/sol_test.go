package day4

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
			args{lines: aoc.MustReadLinesFromFile("../input/day4.sample.txt")},
			4512,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day4.input.txt")},
			58412,
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
			args{lines: aoc.MustReadLinesFromFile("../input/day4.sample.txt")},
			1924,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day4.input.txt")},
			10030,
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
