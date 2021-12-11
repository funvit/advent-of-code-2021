package day11

import (
	"testing"

	"aoc"
)

// Note: both parts are too easy!

func TestPart1(t *testing.T) {
	type args struct {
		lines []string
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample, 10 steps",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day11.sample.txt"),
				steps: 10,
			},
			204,
		},
		{
			"sample 100 steps",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day11.sample.txt"),
				steps: 100,
			},
			1656,
		},
		{
			"input",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day11.input.txt"),
				steps: 100,
			},
			1546,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.lines, tt.args.steps); got != tt.want {
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
		// TODO: Add test cases.
		{
			"sample",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day11.sample.txt"),
			},
			195,
		},
		{
			"input",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day11.input.txt"),
			},
			471,
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
