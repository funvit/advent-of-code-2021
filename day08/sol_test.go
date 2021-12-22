package day08

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
			args{lines: aoc.MustReadLinesFromFile("../input/day8.sample.txt")},
			26,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day8.input.txt")},
			387,
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
		// TODO: Add test cases.
		{
			"sample2",
			args{lines: aoc.MustReadLinesFromFile("../input/day8.sample2.txt")},
			5353,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day8.input.txt")},
			986034,
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

func Test_guessValue(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"sample2",
			args{input: aoc.MustReadLinesFromFile("../input/day8.sample2.txt")[0]},
			5353,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := guessValue(tt.args.input); got != tt.want {
				t.Errorf("guessValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
