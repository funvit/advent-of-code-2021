package day14

import (
	"testing"

	"aoc"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines []string
		apply int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day14.sample.txt"),
				apply: 10,
			},
			1588,
		},
		{
			"input",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day14.input.txt"),
				apply: 10,
			},
			3259,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.lines, tt.args.apply); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		lines []string
		apply int
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		//{
		//	"sample x1",
		//	args{
		//		lines: aoc.MustReadLinesFromFile("../input/day14.sample.txt"),
		//		apply: 1,
		//	},
		//	1588,
		//},
		{
			"sample x10",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day14.sample.txt"),
				apply: 10,
			},
			1588,
		},
		{
			"sample x40",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day14.sample.txt"),
				apply: 40,
			},
			2188189693529,
		},
		{
			"input x40",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day14.input.txt"),
				apply: 40,
			},
			3459174981021,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.lines, tt.args.apply); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
