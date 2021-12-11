package day10

import (
	"testing"

	"aoc"
)

func TestPart(t *testing.T) {
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
			args{lines: aoc.MustReadLinesFromFile("../input/day10.sample.txt")},
			26397,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day10.input.txt")},
			399153,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.lines); got != tt.want {
				t.Errorf("Part() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateLine(t *testing.T) {
	tests := []struct {
		name                   string
		arg                    string
		wantInvalidClosingChar string
		wantValid              bool
	}{
		// TODO: Add test cases.
		{
			"1",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"}",
			false,
		},
		{
			"2",
			"[[<[([]))<([[{}[[()]]]",
			")",
			false,
		},
		{
			"3",
			"[{[{({}]{}}([{[{{{}}([]",
			"]",
			false,
		},
		{
			"4",
			"[<(<(<(<{}))><([]([](",
			")",
			false,
		},
		{
			"5",
			"<{([([[(<>()){}]>(<<{{",
			">",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotInvalidClosing, gotValid := validateLine(tt.arg)
			if gotInvalidClosing.InvalidClosingChar != tt.wantInvalidClosingChar {
				t.Errorf("validateLine() gotInvalidClosing = %v, want %v", gotInvalidClosing, tt.wantInvalidClosingChar)
			}
			if gotValid != tt.wantValid {
				t.Errorf("validateLine() gotValid = %v, want %v", gotValid, tt.wantValid)
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
			args{lines: aoc.MustReadLinesFromFile("../input/day10.sample.txt")},
			288957,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day10.input.txt")},
			2995077699,
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
