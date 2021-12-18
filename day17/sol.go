package day17

import (
	"fmt"
	"log"
	"strings"

	"aoc"
)

func Part1(lines []string) int64 {

	line := strings.TrimPrefix(lines[0], "target area: ")
	p := strings.SplitN(line, ", ", 2)
	xp := strings.SplitN(strings.TrimPrefix(p[0], "x="), "..", 2)
	yp := strings.SplitN(strings.TrimPrefix(p[1], "y="), "..", 2)

	x1, x2 := aoc.MustParseInt64(xp[0]), aoc.MustParseInt64(xp[1])
	y1, y2 := aoc.MustParseInt64(yp[1]), aoc.MustParseInt64(yp[0])

	log.Printf("input: %d,%d - %d,%d", x1, y1, x2, y2)

	// Russian:
	// 1. Ускорение по оси X не должно превышать x2, иначе снаряд перелетит зону.
	// 2. Ускорение по оси Y не должно превышать abs(y2), иначе снаряд
	//    перелетит зону (ибо "шаг" имеет некоторую кратность).
	// 3. Ускорение по оси Y рассматриваем только больше 0, иначе высота будет нулевая.
	// 4. Нужна система ограничений. Т.е. останавливать поиск попадания, если снаряд
	//    явно идет мимо.

	var maxH int64

	for vx := int64(1); vx <= x2; vx++ {
		for vy := int64(1); vy <= absInt64(y2); vy++ {
			v := vel{X: vx, Y: vy}
			h, ok := isHit(
				pos{X: x1, Y: y1},
				pos{X: x2, Y: y2},
				pos{},
				v,
			)
			if ok {
				if maxH == 0 || maxH < h {
					maxH = h
				}
			}
		}
	}

	return maxH
}

func Part2(lines []string) int {

	line := strings.TrimPrefix(lines[0], "target area: ")
	p := strings.SplitN(line, ", ", 2)
	xp := strings.SplitN(strings.TrimPrefix(p[0], "x="), "..", 2)
	yp := strings.SplitN(strings.TrimPrefix(p[1], "y="), "..", 2)

	x1, x2 := aoc.MustParseInt64(xp[0]), aoc.MustParseInt64(xp[1])
	y1, y2 := aoc.MustParseInt64(yp[1]), aoc.MustParseInt64(yp[0])

	log.Printf("input: %d,%d - %d,%d", x1, y1, x2, y2)

	var hits int

	for vx := int64(1); vx <= x2; vx++ {
		for vy := int64(y2); vy <= absInt64(y2); vy++ {
			v := vel{X: vx, Y: vy}
			_, ok := isHit(
				pos{X: x1, Y: y1},
				pos{X: x2, Y: y2},
				pos{},
				v,
			)
			if ok {
				hits++
			}
		}
	}

	return hits
}

func isHit(rectMin, rectMax, start pos, v vel) (height int64, ok bool) {

	x1, y1 := rectMin.X, rectMin.Y
	x2, y2 := rectMax.X, rectMax.Y

	c := start
	startVel := v

	var h int64
	log.Println("checking vel:", v)

	for {
		if c.X < x1 && c.Y < y2 {
			// too near
			log.Println("- pos too near", c)
			break
		}
		if (c.X > x2 && c.Y <= y1) || (c.X >= x1 && c.Y < y2) {
			// too far
			log.Println("- pos too far", c)
			break
		}

		if v.Y > 0 {
			h += v.Y
		}

		for inRect(pos{X: x1, Y: y1}, pos{X: x2, Y: y2}, c) {
			log.Println(
				"! in rect, start vel:", startVel,
				"max height:", h,
			)
			return h, true
		}
		c, v = next(c, v)
	}

	return 0, false
}

type vel struct {
	X, Y int64
}

func (s vel) String() string {
	return fmt.Sprintf("(%d,%d)", s.X, s.Y)
}

type pos struct {
	X, Y int64
}

func (s pos) String() string {
	return fmt.Sprintf("(%d,%d)", s.X, s.Y)
}

func next(p pos, v vel) (pos, vel) {

	x := p.X + v.X
	y := p.Y + v.Y

	v_ := vel{X: v.X - 1, Y: v.Y - 1}
	if v_.X < 0 {
		v_.X = 0
	}

	return pos{X: x, Y: y}, v_
}

func inRect(start, end, p pos) bool {
	return minInt64(start.X, end.X) <= p.X && p.X <= maxInt64(start.X, end.X) &&
		minInt64(start.Y, end.Y) <= p.Y && p.Y <= maxInt64(start.Y, end.Y)
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func absInt64(v int64) int64 {
	if v < 0 {
		return -v
	}
	return v
}

func init() {
	aoc.RegisterSolution(17, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
