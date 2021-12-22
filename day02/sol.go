package day02

import (
	"fmt"
	"strings"

	"aoc"
)

const (
	ForwardCmd = "forward"
	DownCmd    = "down"
	UpCmd      = "up"
)

type Position struct {
	Horizontal int64
	Depth      int64
	Aim        int64
}

func (s *Position) Result() int64 {
	return s.Horizontal * s.Depth
}

func Part1(lines []string) (*Position, error) {

	var r Position

	for _, v := range lines {
		p := strings.SplitN(v, " ", 2)

		val := aoc.MustParseInt64(p[1])

		switch p[0] {
		case ForwardCmd:
			r.Horizontal += val
		case DownCmd:
			r.Depth += val
		case UpCmd:
			r.Depth -= val
		default:
			return nil, fmt.Errorf("unknown cmd: %s", p[0])
		}
	}

	return &r, nil
}

func Part2(lines []string) (*Position, error) {

	var r Position

	for _, v := range lines {
		p := strings.SplitN(v, " ", 2)

		val := aoc.MustParseInt64(p[1])

		switch p[0] {
		case ForwardCmd:
			r.Horizontal += val
			r.Depth += r.Aim * val
		case DownCmd:
			r.Aim += val
		case UpCmd:
			r.Aim -= val
		default:
			return nil, fmt.Errorf("unknown cmd: %s", p[0])
		}
	}

	return &r, nil
}

func init() {
	aoc.RegisterSolution(2, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			v, err := Part1(lines)
			if err != nil {
				panic(err)
			}
			return v.Result()
		},
		Part2: func(lines []string) interface{} {
			v, err := Part2(lines)
			if err != nil {
				panic(err)
			}
			return v.Result()
		},
	})
}
