package day22

import (
	"log"
	"strings"

	"aoc"
)

/*
	Note: output arg type of Part1 and Part2 can be freely changed.
*/

// Part1 .
//
// Uses simple solution - iterate over all 3D points...
func Part1(lines []string) int {

	minMax := [2]int64{-50, 50}
	minMaxBox := BoxFromCoords([2]Point3D{
		{X: minMax[0], Y: minMax[0], Z: minMax[0]},
		{X: minMax[1], Y: minMax[1], Z: minMax[1]},
	})
	filterSelections := func(items []SignedBox) []SignedBox {
		var r []SignedBox
		for _, s := range items {
			if !s.Box.IsOverlaps(*minMaxBox) {
				log.Println("skipping:", s)
				continue
			}

			s.Box.EachCoordinatePair(func(p [2]*int64) {
				if n := aoc.MaxInt64(*p[0], minMax[0]); n != *p[0] {
					log.Println("limiting", *p[0], "to", n)
					*p[0] = n
				}
				if n := aoc.MinInt64(*p[1], minMax[1]); n != *p[1] {
					log.Println("limiting", *p[1], "to", n)
					*p[1] = n
				}
			})

			r = append(r, s)
		}
		return r
	}
	_ = filterSelections

	sels := parseSteps(lines)
	sels = filterSelections(sels)
	log.Println("using steps:", len(sels))

	//
	// process
	//
	grid := map[Point3D]struct{}{}

	for _, s := range sels {
		for x := s.Box.Coords[0].X; x <= s.Box.Coords[1].X; x++ {
			for y := s.Box.Coords[0].Y; y <= s.Box.Coords[1].Y; y++ {
				for z := s.Box.Coords[0].Z; z <= s.Box.Coords[1].Z; z++ {
					p := Point3D{X: x, Y: y, Z: z}
					if s.Add {
						grid[p] = struct{}{}
					} else {
						delete(grid, p)
					}
				}
			}
		}
	}

	return len(grid)
}

// Part2 .
//
// Simple solution (iterate) can lead to out of memory!
//
// Instead, collect boxes and compensation boxes...
//
/*
	2D example:


        *---------*
        | A       |
        |      *--+----*
        |      | C|    |
        *------+--*    |
               |       |
               |     B |
               *-------*

	Result = A + B - C

*/
func Part2(lines []string) int64 {

	steps := parseSteps(lines)
	log.Println("using steps:", len(steps))

	//
	// process
	//
	var boxes []SignedBox

	for _, s := range steps {
		var compensations []SignedBox

		if len(boxes) == 0 {
			boxes = append(boxes, s)
			continue
		}

		for _, b := range boxes {
			if !b.Box.IsOverlaps(s.Box) {
				continue
			}
			if b.Add && s.Add {
				r := SignedBox{
					Box: *b.Box.GetIntersect(s.Box),
					Add: false,
				}
				compensations = append(compensations, r)
				continue
			}
			if b.Add && !s.Add {
				r := SignedBox{
					Box: *b.Box.GetIntersect(s.Box),
					Add: false,
				}
				compensations = append(compensations, r)
				continue
			}
			if !b.Add && !s.Add {
				r := SignedBox{
					Box: *b.Box.GetIntersect(s.Box),
					Add: true,
				}
				compensations = append(compensations, r)
				continue
			}
			if !b.Add && s.Add {
				r := SignedBox{
					Box: *b.Box.GetIntersect(s.Box),
					Add: true,
				}
				compensations = append(compensations, r)
				continue
			}
		}
		if s.Add {
			boxes = append(boxes, s)
		}
		boxes = append(boxes, compensations...)
	}

	var r int64
	for _, b := range boxes {
		switch b.Add {
		case true:
			r += b.Box.CoordsSum()
		default:
			r -= b.Box.CoordsSum()
		}
	}

	return r
}

func parseSteps(lines []string) []SignedBox {
	var sels []SignedBox
	for _, line := range lines {
		var s SignedBox

		p := strings.SplitN(line, " ", 2)
		if p[0] == "on" {
			s.Add = true
		}

		c := strings.SplitN(p[1], ",", 3)

		var i int
		s.Box.EachCoordinatePair(func(p [2]*int64) {
			cc := c[i][2:]
			v := aoc.MustParseNumbers(strings.SplitN(cc, "..", 2))
			*p[0] = v[0]
			*p[1] = v[1]
			i++
		})
		sels = append(sels, s)
	}
	return sels
}

type Point3D struct {
	X, Y, Z int64
}

type SignedBox struct {
	Box Box
	Add bool // true = add, false = subtract
}

type Box struct {
	Coords [2]Point3D
}

func BoxFromCoords(c [2]Point3D) *Box {
	return &Box{Coords: c}
}

func (s *Box) EachCoordinatePair(fn func(p [2]*int64)) {
	for _, p := range [][2]*int64{
		{&s.Coords[0].X, &s.Coords[1].X},
		{&s.Coords[0].Y, &s.Coords[1].Y},
		{&s.Coords[0].Z, &s.Coords[1].Z},
	} {
		fn(p)
	}
}

func (s *Box) IsOverlaps(b Box) bool {

	x := isRangesOverlap([2]int64{s.Coords[0].X, s.Coords[1].X}, [2]int64{b.Coords[0].X, b.Coords[1].X})
	y := isRangesOverlap([2]int64{s.Coords[0].Y, s.Coords[1].Y}, [2]int64{b.Coords[0].Y, b.Coords[1].Y})
	z := isRangesOverlap([2]int64{s.Coords[0].Z, s.Coords[1].Z}, [2]int64{b.Coords[0].Z, b.Coords[1].Z})

	return x && y && z
}

func (s *Box) CoordsSum() int64 {
	var r int64
	s.EachCoordinatePair(func(p [2]*int64) {
		n := aoc.AbsInt64(*p[1]-*p[0]) + 1
		if r == 0 {
			r = n
			return
		}
		r *= n
	})
	return r
}

func (s *Box) GetIntersect(b Box) *Box {
	x, _ := getRangesIntersection(
		[2]int64{s.Coords[0].X, s.Coords[1].X},
		[2]int64{b.Coords[0].X, b.Coords[1].X},
	)
	y, _ := getRangesIntersection(
		[2]int64{s.Coords[0].Y, s.Coords[1].Y},
		[2]int64{b.Coords[0].Y, b.Coords[1].Y},
	)
	z, _ := getRangesIntersection(
		[2]int64{s.Coords[0].Z, s.Coords[1].Z},
		[2]int64{b.Coords[0].Z, b.Coords[1].Z},
	)
	return &Box{
		Coords: [2]Point3D{
			{
				X: x[0],
				Y: y[0],
				Z: z[0],
			},
			{
				X: x[1],
				Y: y[1],
				Z: z[1],
			},
		},
	}
}

func (s *Box) GetSum(b Box) *Box {
	x, _ := getRangesSum(
		[2]int64{s.Coords[0].X, s.Coords[1].X},
		[2]int64{b.Coords[0].X, b.Coords[1].X},
	)
	y, _ := getRangesSum(
		[2]int64{s.Coords[0].Y, s.Coords[1].Y},
		[2]int64{b.Coords[0].Y, b.Coords[1].Y},
	)
	z, _ := getRangesSum(
		[2]int64{s.Coords[0].Z, s.Coords[1].Z},
		[2]int64{b.Coords[0].Z, b.Coords[1].Z},
	)
	return &Box{
		Coords: [2]Point3D{
			{
				X: x[0],
				Y: y[0],
				Z: z[0],
			},
			{
				X: x[1],
				Y: y[1],
				Z: z[1],
			},
		},
	}
}

func isRangesOverlap(a, b [2]int64) bool {
	return a[0] <= b[1] && a[1] >= b[0]
}

func getRangesIntersection(a, b [2]int64) ([2]int64, bool) {
	if !isRangesOverlap(a, b) {
		return [2]int64{}, false
	}

	return [2]int64{
			aoc.MaxInt64(a[0], b[0]),
			aoc.MinInt64(a[1], b[1]),
		},
		true
}

func getRangesSum(a, b [2]int64) ([2]int64, bool) {
	return [2]int64{
			aoc.MinInt64(a[0], b[0]),
			aoc.MaxInt64(a[1], b[1]),
		},
		true
}

func init() {
	// this registers solution to be able to run from ./cmd/main.go,
	// BUT you must add import like `_ "aoc/day22"` in "main.go" by hands!
	aoc.RegisterSolution(22, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
