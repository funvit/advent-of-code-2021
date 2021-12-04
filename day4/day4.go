package day4

import (
	"strings"

	"aoc"
)

type BoardNumber struct {
	Num   int64
	Match bool
}

type NumberLine struct {
	Numbers [5]*BoardNumber
}
type Board struct {
	Index int
	Lines [5]*NumberLine
}

func (s *NumberLine) IsLineBingo() bool {
	var r int
	for _, v := range s.Numbers {
		if v.Match {
			r++
		}
	}
	return r == 5
}

func (s *Board) IsColumnBingo(col int) bool {
	var r int
	for _, l := range s.Lines {
		if l.Numbers[col].Match {
			r++
		}
	}
	return r == 5
}

func (s *Board) Score(lastNumber int64) int64 {
	var r int64
	for _, l := range s.Lines {
		for _, n := range l.Numbers {
			if !n.Match {
				r += n.Num
			}
		}
	}
	return r * lastNumber
}

func NewNumberLineFromSlice(v ...int64) *NumberLine {
	r := NumberLine{
		Numbers: [5]*BoardNumber{},
	}
	for i := 0; i < 5; i++ {
		r.Numbers[i] = &BoardNumber{
			Num:   v[i],
			Match: false,
		}
	}
	return &r
}

func WinnerScore(lines []string, findLast bool) int64 {
	var numbers []int64
	for _, s := range strings.Split(lines[0], ",") {
		numbers = append(numbers, aoc.MustParseInt64(s))
	}

	var boards []*Board
	var bi int
	for i := 2; i < len(lines); i += 6 {
		var b Board

		b.Index = bi

		b.Lines[0] = NewNumberLineFromSlice(parseLineToInt64s(lines[i])...)
		b.Lines[1] = NewNumberLineFromSlice(parseLineToInt64s(lines[i+1])...)
		b.Lines[2] = NewNumberLineFromSlice(parseLineToInt64s(lines[i+2])...)
		b.Lines[3] = NewNumberLineFromSlice(parseLineToInt64s(lines[i+3])...)
		b.Lines[4] = NewNumberLineFromSlice(parseLineToInt64s(lines[i+4])...)

		boards = append(boards, &b)
		bi++
	}

	var winBoard *Board
	var lastNumber int64
	skipBoards := map[int]bool{}
	for i := 0; i < len(numbers); i++ {
		n := numbers[i]

		//
		// find matches on boards
		//
		for bIdx, b := range boards {
			if skipBoards[bIdx] {
				continue
			}
			for _, l := range b.Lines {
				for _, ln := range l.Numbers {
					if ln.Num == n {
						ln.Match = true
					}
				}
			}
		}

		//
		// find win board
		//
		for bIdx, b := range boards {
			if skipBoards[bIdx] {
				continue
			}
			var win bool
			for i := 0; i < 5; i++ {
				if b.IsColumnBingo(i) {
					win = true
					break
				}
			}
			if !win {
				for _, l := range b.Lines {
					if l.IsLineBingo() {
						win = true
						break
					}
				}
			}
			if !win {
				continue
			}

			winBoard = b
			lastNumber = n

			if !findLast {
				goto exit
			}

			skipBoards[bIdx] = true

			if len(boards) == 0 {
				goto exit
			}

			continue
		}
	}

exit:
	if winBoard != nil {
		//log.Println("winning board:", winBoard)
		return winBoard.Score(lastNumber)
	}

	return -1
}

func Part1(lines []string) int64 {
	return WinnerScore(lines, false)
}
func Part2(lines []string) int64 {
	return WinnerScore(lines, true)
}

func parseLineToInt64s(s string) []int64 {
	s = strings.ReplaceAll(strings.TrimSpace(s), "  ", " ")
	p := strings.Split(s, " ")

	var r []int64

	for _, v := range p {
		r = append(r, aoc.MustParseInt64(v))
	}

	return r
}

func init() {
	aoc.RegisterSolution(4, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
