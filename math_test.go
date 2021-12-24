package aoc

import "testing"

func TestMaxInt64(t *testing.T) {
	type args struct {
		n []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"0, -50",
			args{n: []int64{0, -50}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxInt64(tt.args.n...); got != tt.want {
				t.Errorf("MaxInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInt64(t *testing.T) {
	type args struct {
		n []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"0, -50",
			args{n: []int64{0, -50}},
			-50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinInt64(tt.args.n...); got != tt.want {
				t.Errorf("MinInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
