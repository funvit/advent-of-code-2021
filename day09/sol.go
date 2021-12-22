package day09

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"aoc"
)

type valueAtPos struct {
	Val int64
	Pos pos
}

func (s valueAtPos) String() string {
	return fmt.Sprintf("%d at %d,%d", s.Val, s.Pos.X, s.Pos.Y)
}

type pos struct {
	X, Y int
}

func (s pos) String() string {
	return fmt.Sprintf("%d,%d", s.X, s.Y)
}

func Part1(lines []string) int64 {
	var heightsYX [][]int64

	for _, l := range lines {
		lh := aoc.MustParseNumbers(strings.Split(l, ""))
		heightsYX = append(heightsYX, lh)
	}

	var riskLevels []int64
	for y := 0; y < len(heightsYX); y++ {
		for x := 0; x < len(heightsYX[0]); x++ {
			if isLowestHeight(heightsYX, pos{X: x, Y: y}) {
				riskLevels = append(riskLevels, 1+heightsYX[y][x])
				fmt.Print(heightsYX[y][x])
				continue
			}
			fmt.Print(".")
		}
		fmt.Print("\n")
	}

	var sum int64
	for _, v := range riskLevels {
		sum += v
	}

	return sum
}

func Part2(lines []string) int {
	var heightsYX [][]int64

	for _, l := range lines {
		lh := aoc.MustParseNumbers(strings.Split(l, ""))
		heightsYX = append(heightsYX, lh)
	}
	maxX := len(heightsYX[0]) - 1
	maxY := len(heightsYX) - 1

	var centers []valueAtPos
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if isLowestHeight(heightsYX, pos{X: x, Y: y}) {
				centers = append(centers, valueAtPos{
					Val: heightsYX[y][x],
					Pos: pos{X: x, Y: y},
				})
				continue
			}
		}
	}

	log.Println("found centers:", len(centers))

	var basins [][]valueAtPos
	exclude := map[pos]struct{}{}
	for _, c := range centers {
		exclude[c.Pos] = struct{}{}
	}

	for i, c := range centers {
		log.Println("processing center at idx:", i, "pos:", c.Pos.String())
		l := leak(heightsYX, c.Pos, exclude)
		if len(l) == 0 {
			continue
		}
		log.Println("found leak positions:", len(l))
		l[c.Pos] = c

		var values []valueAtPos
		for _, v := range l {
			values = append(values, v)
			exclude[v.Pos] = struct{}{}
		}
		basins = append(basins, values)
	}

	sort.Slice(basins, func(i, j int) bool {
		return len(basins[i]) < len(basins[j])
	})

	var r int
	for i := 0; i < 3; i++ {
		l := len(basins[len(basins)-3+i])
		if r == 0 {
			r = l
			continue
		}
		r *= l
	}

	return r
}

// leak .
//
// Note: exclude changed in-place to allow data accumulation.
func leak(heightsYX [][]int64, start pos, exclude map[pos]struct{}) map[pos]valueAtPos {

	r := map[pos]valueAtPos{}

	nearest := nearestHeights(heightsYX, start.X, start.Y)
	nearest = filterHeights(nearest, 9)
	var nearest_ []valueAtPos
	for _, v := range nearest {
		if v.Val <= heightsYX[start.Y][start.X] {
			continue
		}
		nearest_ = append(nearest_, v)
	}
	nearest = nearest_

	for _, n := range nearest {
		r[n.Pos] = n
	}

	for _, n := range nearest {
		if _, ok := exclude[n.Pos]; ok {
			continue
		}
		exclude[n.Pos] = struct{}{}

		rr := leak(heightsYX, n.Pos, exclude)
		for k, v := range rr {
			r[k] = v
		}
	}

	return r
}

func nearestHeights(heightsYX [][]int64, x, y int) []valueAtPos {

	r := make([]valueAtPos, 0, 4)
	add := func(v int64, x, y int) {
		r = append(r, valueAtPos{
			Val: v,
			Pos: pos{
				X: x,
				Y: y,
			},
		})
	}

	if x > 0 {
		add(heightsYX[y][x-1], x-1, y)
	}
	if y > 0 {
		add(heightsYX[y-1][x], x, y-1)

	}
	if x < len(heightsYX[y])-1 {
		add(heightsYX[y][x+1], x+1, y)
	}
	if y < len(heightsYX)-1 {
		add(heightsYX[y+1][x], x, y+1)
	}

	return r
}

// isLowestHeight .
func isLowestHeight(heightsYX [][]int64, at pos) bool {
	valAtPos := heightsYX[at.Y][at.X]

	nearest := nearestHeights(heightsYX, at.X, at.Y)
	sort.Slice(nearest, func(i, j int) bool {
		return nearest[i].Val < nearest[j].Val
	})

	var i int
	r := true
	for _, v := range nearest {
		if v.Val == -1 {
			continue
		}

		if valAtPos >= v.Val {
			i++
			r = false
			break
		}
		i++
	}

	return r && i > 0
}

func filterHeights(heights []valueAtPos, val int64) []valueAtPos {

	var r []valueAtPos
	for _, v := range heights {
		if v.Val == val {
			continue
		}
		r = append(r, v)
	}
	return r
}

func init() {
	aoc.RegisterSolution(9, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
