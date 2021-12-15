package day13

import (
	"strings"

	"aoc"
)

const (
	horizDirection = iota
	vertDirection
)

type pos struct {
	X, Y int64
}
type foldInstruction struct {
	Direction int
	Index     int64
}

func Part1(lines []string, applyFolds int) int {

	var maxX, maxY int64 = 0, 0

	dots := map[pos]bool{}
	var instructions []*foldInstruction

	var m bool
	for _, l := range lines {
		if l == "" || m {
			if l == "" {
				m = true
				continue
			}

			p := strings.SplitN(l[strings.LastIndex(l, " ")+1:], "=", 2)

			d := foldInstruction{
				Index: aoc.MustParseInt64(p[1]),
			}
			switch p[0] {
			case "x":
				d.Direction = vertDirection

				if maxX < d.Index*2 {
					maxX = d.Index * 2
				}

			case "y":
				d.Direction = horizDirection

				if maxY < d.Index*2 {
					maxY = d.Index * 2
				}

			default:
				panic("unknown direction: " + p[0])
			}

			instructions = append(instructions, &d)

			continue
		}

		n := aoc.MustParseNumbers(strings.SplitN(l, ",", 2))
		dots[pos{X: n[0], Y: n[1]}] = true
	}

	//
	// fold
	//
	for i := 0; i < applyFolds; i++ {
		fi := instructions[i]

		switch fi.Direction {
		case horizDirection:
			for x := int64(0); x <= maxX; x++ {
				for y := int64(0); y <= fi.Index; y++ {
					dest := pos{X: x, Y: y}
					source := pos{X: x, Y: 2*fi.Index - y}

					if !dots[dest] && dots[source] {
						dots[dest] = true
					}
					dots[source] = false
				}
			}

			maxY = fi.Index

		case vertDirection:
			for x := int64(0); x <= fi.Index; x++ {
				for y := int64(0); y <= maxY; y++ {
					dest := pos{X: x, Y: y}
					source := pos{X: 2*fi.Index - x, Y: y}

					if !dots[dest] && dots[source] {
						dots[dest] = true
					}
					dots[source] = false
				}
			}

			maxX = fi.Index
		}

		// debug
		//for y := int64(0); y <= maxY; y++ {
		//	for x := int64(0); x <= maxX; x++ {
		//		v := dots[pos{X: x, Y: y}]
		//		if v {
		//			fmt.Print("X")
		//			continue
		//		}
		//		fmt.Print(".")
		//	}
		//	fmt.Println("")
		//}
		//fmt.Println("")
	}

	var r int
	for _, v := range dots {
		if v {
			r++
		}
	}

	return r
}

func Part2(lines []string) string {

	var maxX, maxY int64 = 0, 0

	dots := map[pos]bool{}
	var instructions []*foldInstruction

	var m bool
	for _, l := range lines {
		if l == "" || m {
			if l == "" {
				m = true
				continue
			}

			p := strings.SplitN(l[strings.LastIndex(l, " ")+1:], "=", 2)

			d := foldInstruction{
				Index: aoc.MustParseInt64(p[1]),
			}
			switch p[0] {
			case "x":
				d.Direction = vertDirection

				if maxX < d.Index*2 {
					maxX = d.Index * 2
				}

			case "y":
				d.Direction = horizDirection

				if maxY < d.Index*2 {
					maxY = d.Index * 2
				}

			default:
				panic("unknown direction: " + p[0])
			}

			instructions = append(instructions, &d)

			continue
		}

		n := aoc.MustParseNumbers(strings.SplitN(l, ",", 2))
		dots[pos{X: n[0], Y: n[1]}] = true
	}

	//
	// fold
	//
	for _, fi := range instructions {

		switch fi.Direction {
		case horizDirection:
			for x := int64(0); x <= maxX; x++ {
				for y := int64(0); y <= fi.Index; y++ {
					dest := pos{X: x, Y: y}
					source := pos{X: x, Y: 2*fi.Index - y}

					if !dots[dest] && dots[source] {
						dots[dest] = true
					}
					dots[source] = false
				}
			}

			maxY = fi.Index

		case vertDirection:
			for x := int64(0); x <= fi.Index; x++ {
				for y := int64(0); y <= maxY; y++ {
					dest := pos{X: x, Y: y}
					source := pos{X: 2*fi.Index - x, Y: y}

					if !dots[dest] && dots[source] {
						dots[dest] = true
					}
					dots[source] = false
				}
			}

			maxX = fi.Index
		}

	}

	//debug
	var b strings.Builder

	b.WriteString("\n")
	for y := int64(0); y <= maxY; y++ {
		for x := int64(0); x <= maxX; x++ {
			v := dots[pos{X: x, Y: y}]
			if v {
				b.WriteString("X")
				continue
			}
			b.WriteString(" ")
		}
		b.WriteString("\n")
	}
	//b.WriteString("\n")

	return b.String()
}

func init() {
	aoc.RegisterSolution(13, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines, 1)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
