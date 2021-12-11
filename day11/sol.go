package day11

import (
	"fmt"
	"os"
	"strings"

	"aoc"
)

var debug = os.Getenv("DEBUG") != ""

type octo struct {
	Pos           pos
	Energy        int64
	FlashedAtStep int
	TotalFlashes  int
}
type pos struct {
	X, Y int
}

func Part1(lines []string, steps int) int {
	octos := map[pos]octo{}

	for row, l := range lines {
		for col, v := range strings.Split(l, "") {
			at := pos{X: col, Y: row}
			octos[at] = octo{
				Energy: aoc.MustParseInt64(v),
				Pos:    at,
			}
		}
	}

	for i := 1; i <= steps; i++ {
		if debug {
			fmt.Println("step:", i)
		}
		for k := range octos {
			octos = incEnergy(octos, k, i)
		}
		for k := range octos {
			octos = flash(octos, k, i)
		}
		// debug
		if debug {
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					fmt.Print(octos[pos{X: x, Y: y}].Energy)
				}
				fmt.Println("")
			}
		}
	}

	var r int
	for _, v := range octos {
		r += v.TotalFlashes
	}

	return r
}

func Part2(lines []string) int {
	octos := map[pos]octo{}

	for row, l := range lines {
		for col, v := range strings.Split(l, "") {
			at := pos{X: col, Y: row}
			octos[at] = octo{
				Energy: aoc.MustParseInt64(v),
				Pos:    at,
			}
		}
	}

	for i := 1; i < 1_000_000; i++ {
		if debug {
			fmt.Println("step:", i)
		}
		for k := range octos {
			octos = incEnergy(octos, k, i)
		}
		for k := range octos {
			octos = flash(octos, k, i)
		}
		// debug
		if debug {
			for y := 0; y < 10; y++ {
				for x := 0; x < 10; x++ {
					fmt.Print(octos[pos{X: x, Y: y}].Energy)
				}
				fmt.Println("")
			}
		}

		var flashedCnt int
		for _, v := range octos {
			if v.Energy != 0 {
				continue
			}
			flashedCnt++
		}
		if flashedCnt == len(octos) {
			return i
		}
	}

	return -1
}

func incEnergy(octos map[pos]octo, at pos, step int) map[pos]octo {
	v, ok := octos[at]
	if !ok || v.FlashedAtStep == step {
		return octos
	}

	if v.Energy > 9 {
		return octos
	}

	v.Energy++
	octos[at] = v

	return octos
}

func flash(octos map[pos]octo, at pos, step int) map[pos]octo {
	v, ok := octos[at]
	if !ok || v.FlashedAtStep == step {
		return octos
	}

	if v.Energy < 10 {
		return octos
	}

	v.Energy = 0
	v.TotalFlashes++
	v.FlashedAtStep = step

	octos[at] = v

	for _, o := range nearest(octos, at) {
		octos = incEnergy(octos, o.Pos, step)
	}
	for _, o := range nearest(octos, at) {
		octos = flash(octos, o.Pos, step)
	}

	return octos
}

func nearest(octos map[pos]octo, at pos) []octo {
	var r []octo

	var maxPos pos
	for k := range octos {
		if maxPos.X < k.X {
			maxPos.X = k.X
		}
		if maxPos.Y < k.Y {
			maxPos.Y = k.Y
		}
	}

	start := pos{
		X: at.X - 1,
		Y: at.Y - 1,
	}
	if start.X < 0 {
		start.X = 0
	}
	if start.Y < 0 {
		start.Y = 0
	}

	end := pos{
		X: at.X + 1,
		Y: at.Y + 1,
	}
	if end.X > maxPos.X {
		end.X = maxPos.X
	}
	if end.Y > maxPos.Y {
		end.Y = maxPos.Y
	}

	for x := start.X; x <= end.X; x++ {
		for y := start.Y; y <= end.Y; y++ {
			p := pos{
				X: x,
				Y: y,
			}
			if p == at {
				continue
			}
			r = append(r, octos[p])
		}
	}

	return r
}

func init() {
	aoc.RegisterSolution(11, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines, 100)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
