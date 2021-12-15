package day13

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
			args{lines: aoc.MustReadLinesFromFile("../input/day13.sample.txt")},
			17,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day13.input.txt")},
			775,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.lines, 1); got != tt.want {
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
		want string
	}{
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day13.input.txt")},
			`
XXX  XXXX X  X XXX  X  X XXX  X  X XXX   
X  X X    X  X X  X X  X X  X X X  X  X  
X  X XXX  X  X X  X X  X X  X XX   X  X  
XXX  X    X  X XXX  X  X XXX  X X  XXX   
X X  X    X  X X    X  X X    X X  X X   
X  X XXXX  XX  X     XX  X    X  X X  X  
                                         
`,
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
