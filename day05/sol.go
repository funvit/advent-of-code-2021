package day05

import (
	"strings"

	"aoc"
)

type Point struct {
	X, Y int
}

func ParseLine(s string) (Point, Point) {
	p := strings.SplitN(s, " -> ", 2)

	p1 := strings.SplitN(p[0], ",", 2)
	p2 := strings.SplitN(p[1], ",", 2)

	return Point{
			X: int(aoc.MustParseInt64(p1[0])),
			Y: int(aoc.MustParseInt64(p1[1])),
		}, Point{
			X: int(aoc.MustParseInt64(p2[0])),
			Y: int(aoc.MustParseInt64(p2[1])),
		}

}

// Part1 .
//
// Determine the number of points where at least two lines overlap.
//
// Trick: do not use geometry, just fill map with points and values!
func Part1(lines []string) int {

	field := map[Point]int{}

	for _, l := range lines {
		p1, p2 := ParseLine(l)

		// skip, if not vertical or horizontal
		if p1.X != p2.X && p1.Y != p2.Y {
			continue
		}

		// set values in field
		if p1.X == p2.X {
			// horizontal
			start, end := p1.Y, p2.Y
			if p1.Y > p2.Y {
				start, end = p2.Y, p1.Y
			}
			for i := start; i <= end; i++ {
				field[Point{X: p1.X, Y: i}]++
			}
		}
		if p1.Y == p2.Y {
			// vertical
			start, end := p1.X, p2.X
			if p1.X > p2.X {
				start, end = p2.X, p1.X
			}
			for i := start; i <= end; i++ {
				field[Point{X: i, Y: p1.Y}]++
			}
		}
	}

	//// debug
	//for y := 0; y < 10; y++ {
	//	for x := 0; x < 10; x++ {
	//		v := field[Point{X: x, Y: y}]
	//		if v == 0 {
	//			fmt.Print(".")
	//			continue
	//		}
	//		fmt.Print(v)
	//	}
	//	fmt.Print("\n")
	//}

	// count values in fields with value >=2
	var r int
	for _, v := range field {
		if v >= 2 {
			r++
		}
	}

	return r
}

// Part2 .
//
// Same as Part1, but with diagonal lines.
//
// Note: if diagonal lines exists, they expected to be at 45 degree.
func Part2(lines []string) int {

	field := map[Point]int{}

	for _, l := range lines {
		p1, p2 := ParseLine(l)

		// set values in field
		var x, y int

		var xDone, yDone bool

		for {
			field[Point{X: p1.X + x, Y: p1.Y + y}]++

			// note: respect direction!
			if p1.X+x == p2.X {
				xDone = true
			} else {
				if p2.X > p1.X {
					x++
				} else {
					x--
				}
			}
			if p1.Y+y == p2.Y {
				yDone = true
			} else {
				if p2.Y > p1.Y {
					y++
				} else {
					y--
				}
			}

			if xDone && yDone {
				break
			}
		}
	}

	//// debug
	//for y := 0; y < 10; y++ {
	//	for x := 0; x < 10; x++ {
	//		v := field[Point{X: x, Y: y}]
	//		if v == 0 {
	//			fmt.Print(".")
	//			continue
	//		}
	//		fmt.Print(v)
	//	}
	//	fmt.Print("\n")
	//}

	// count values in fields with value >=2
	var r int
	for _, v := range field {
		if v >= 2 {
			r++
		}
	}

	return r
}

func init() {
	aoc.RegisterSolution(5, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
