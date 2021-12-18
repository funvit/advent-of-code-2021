package day17

import "testing"

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
			args{lines: []string{"target area: x=20..30, y=-10..-5"}},
			45,
		},
		{
			"input",
			args{lines: []string{"target area: x=94..151, y=-156..-103"}},
			12090,
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
			args{lines: []string{"target area: x=20..30, y=-10..-5"}},
			12*9 + 4,
		},
		{
			"input",
			args{lines: []string{"target area: x=94..151, y=-156..-103"}},
			5059,
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
