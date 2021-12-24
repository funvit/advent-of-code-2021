package day22

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
			"sample 1",
			args{lines: aoc.MustReadLinesFromFile("../input/day22.sample1.txt")},
			39,
		},
		{
			"sample 2",
			args{lines: aoc.MustReadLinesFromFile("../input/day22.sample2.txt")},
			590784,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day22.input.txt")},
			615700,
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
			"sample 1",
			args{lines: aoc.MustReadLinesFromFile("../input/day22.sample1.txt")},
			39,
		},
		{
			"sample 3",
			args{lines: aoc.MustReadLinesFromFile("../input/day22.sample3.txt")},
			2758514936282235,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day22.input.txt")},
			1236463892941356,
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
