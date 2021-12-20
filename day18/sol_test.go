package day18

import (
	"reflect"
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
			"sample 1",
			args{lines: []string{
				"[[[[4,3],4],4],[7,[[8,4],9]]]",
				"[1,1]",
			}},
			1384,
		},
		{
			"sample 2",
			args{lines: aoc.MustReadLinesFromFile("../input/day18.sample2.txt")},
			4140,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day18.input.txt")},
			4347,
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
			"sample 2",
			args{lines: aoc.MustReadLinesFromFile("../input/day18.sample2.txt")},
			3993,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day18.input.txt")},
			4721,
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

func Test_reduce(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"sample 1",
			args{s: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"},
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
		{
			"sample 1a",
			args{s: "[1,[[[[3,8],1],2],3]]"},
			"[4,[[[0,9],2],3]]",
		},
		{
			"sample 2",
			args{s: "[7,[6,[5,[4,[3,2]]]]]"},
			"[7,[6,[5,[7,0]]]]",
		},
		{
			"sample 2a",
			args{s: "[[6,[5,[4,[3,2]]]],1]"},
			"[[6,[5,[7,0]]],3]",
		},
		{
			"sample 3",
			args{s: "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]"},
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			"sample 3a",
			args{s: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			"my ex 1",
			args{s: "[[1,9],[[[[2,2],9],9],9]]"},
			"[[1,[5,6]],[[[5,0],[7,8]],9]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduce(tt.args.s); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findPairs(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []pairAt
	}{
		{
			"1",
			args{s: "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]"},
			[]pairAt{
				{
					Pair:       pair{4, 3},
					IndexStart: 5,
					IndexEnd:   8,
				},
				{
					Pair:       pair{8, 4},
					IndexStart: 21,
					IndexEnd:   24,
				},
				{
					Pair:       pair{1, 1},
					IndexStart: 32,
					IndexEnd:   35,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPairs(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumLines(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"ex 1",
			args{lines: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
			}},
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			"ex 2",
			args{lines: []string{
				"[1,1]",
				"[2,2]",
				"[3,3]",
				"[4,4]",
				"[5,5]",
				"[6,6]",
			}},
			"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			"sample 1 sum 1",
			args{lines: []string{
				"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]",
				"[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]",
			}},
			"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
		},
		{
			"sample 1 sum 2",
			args{lines: []string{
				"[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]",
				"[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]",
			}},
			"[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]",
		},
		{
			"sample 1 sum 6",
			args{lines: []string{
				"[[[[6,6],[6,6]],[[6,0],[6,7]]],[[[7,7],[8,9]],[8,[8,1]]]]",
				"[2,9]",
			}},
			"[[[[6,6],[7,7]],[[0,7],[7,7]]],[[[5,5],[5,6]],9]]",
		},
		{
			"sample file 1",
			args{lines: aoc.MustReadLinesFromFile("../input/day18.sample1.txt")},
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
		{
			"sample file 2",
			args{lines: aoc.MustReadLinesFromFile("../input/day18.sample2.txt")},
			"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumLines(tt.args.lines); got != tt.want {
				t.Errorf("sumLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
