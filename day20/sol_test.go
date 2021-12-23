package day20

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
			args{
				lines: aoc.MustReadLinesFromFile("../input/day20.sample.txt"),
			},
			35,
		},
		{
			"input",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day20.input.txt"),
			},
			5573,
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
			args{
				lines: aoc.MustReadLinesFromFile("../input/day20.sample.txt"),
			},
			3351,
		},
		{
			"input",
			args{
				lines: aoc.MustReadLinesFromFile("../input/day20.input.txt"),
			},
			20097,
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
