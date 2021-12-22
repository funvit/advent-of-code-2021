package day19

import (
	"testing"

	"aoc"
)

func TestPart1(t *testing.T) {
	type args struct {
		lines            []string
		leastSameBeacons int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"sample 1a",
			args{
				lines:            aoc.MustReadLinesFromFile("../input/day19.sample1a.txt"),
				leastSameBeacons: 6,
			},
			6,
		},
		{
			"sample 1b",
			args{
				lines:            aoc.MustReadLinesFromFile("../input/day19.sample1b.txt"),
				leastSameBeacons: 6,
			},
			7,
		},
		{
			"sample 2",
			args{
				lines:            aoc.MustReadLinesFromFile("../input/day19.sample2.txt"),
				leastSameBeacons: 12,
			},
			79,
		},
		{
			"sample 2a",
			args{
				lines:            aoc.MustReadLinesFromFile("../input/day19.sample2a.txt"),
				leastSameBeacons: 12,
			},
			38,
		},
		{
			"input",
			args{
				lines:            aoc.MustReadLinesFromFile("../input/day19.input.txt"),
				leastSameBeacons: 12,
			},
			335,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part1(tt.args.lines, tt.args.leastSameBeacons); got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		lines            []string
		leastSameBeacons int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"sample 2",
			args{
				lines:            aoc.MustReadLinesFromFile("../input/day19.sample2.txt"),
				leastSameBeacons: 12,
			},
			3621,
		},
		{
			"input",
			args{
				lines:            aoc.MustReadLinesFromFile("../input/day19.input.txt"),
				leastSameBeacons: 12,
			},
			10864,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Part2(tt.args.lines, tt.args.leastSameBeacons); got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
