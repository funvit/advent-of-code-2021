package day16

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
		want int64
	}{
		{
			"sample 0",
			args{lines: []string{"38006F45291200"}},
			9,
		},
		{
			"sample 1",
			args{lines: []string{"8A004A801A8002F478"}},
			16,
		},
		{
			"sample 2",
			args{lines: []string{"620080001611562C8802118E34"}},
			12,
		},
		{
			"sample 3",
			args{lines: []string{"C0015000016115A2E0802F182340"}},
			23,
		},
		{
			"sample 4",
			args{lines: []string{"A0016C880162017C3686B18A3D4780"}},
			31,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day16.input.txt")},
			871,
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
			args{lines: []string{"C200B40A82"}},
			3,
		},
		{
			"sample 2",
			args{lines: []string{"04005AC33890"}},
			54,
		},
		{
			"sample 3",
			args{lines: []string{"880086C3E88112"}},
			7,
		},
		{
			"sample 4",
			args{lines: []string{"CE00C43D881120"}},
			9,
		},
		{
			"sample 5",
			args{lines: []string{"D8005AC2A8F0"}},
			1,
		},
		{
			"sample 6",
			args{lines: []string{"F600BC2D8F"}},
			0,
		},
		{
			"sample 7",
			args{lines: []string{"9C005AC2F8F0"}},
			0,
		},
		{
			"sample 8",
			args{lines: []string{"9C0141080250320F1802104A08"}},
			1,
		},
		{
			"input",
			args{lines: aoc.MustReadLinesFromFile("../input/day16.input.txt")},
			68703010504,
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
