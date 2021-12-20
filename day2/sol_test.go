package day2

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
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"sample",
			args{lines: aoc.MustReadLinesFromFile("../input/day2.sample.txt")},
			150,
			false,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day2.input.txt")},
			1604850,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part1(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Result(), tt.want) {
				t.Errorf("Part1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			"one",
			args{lines: aoc.MustReadLinesFromFile("../input/day2.sample.txt")},
			900,
			false,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day2.input.txt")},
			1685186100,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Part2(tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("Part2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Result(), tt.want) {
				t.Errorf("Part2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
