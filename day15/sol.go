package day15

import (
	"container/heap"
	"log"
	"math"
	"strings"

	"aoc"
)

type pos struct {
	X, Y int
}

func Part1(lines []string) int {
	b := board{}
	var start, end pos

	start = pos{}
	end = pos{X: 0, Y: len(lines) - 1}

	_ = start

	for y, l := range lines {
		for x, n := range aoc.MustParseNumbers(strings.Split(l, "")) {
			b[pos{X: x, Y: y}] = int(n)
			if end.X < x {
				end.X = x
			}
		}
	}

	log.Printf("Board %d,%d - %d,%d", start.X, start.Y, end.X, end.Y)

	rank := b.getMinPathRank(start, end)

	return rank
}

func Part2(lines []string) int {
	const mul = 5

	b := board{}
	var start, end pos

	start = pos{}
	end = pos{
		X: len(lines[0])*mul - 1,
		Y: len(lines)*mul - 1,
	}

	_ = start

	for dy := 0; dy < mul; dy++ {
		for dx := 0; dx < mul; dx++ {

			for y, l := range lines {
				for x, n := range aoc.MustParseNumbers(strings.Split(l, "")) {
					x_ := x + (dx * len(lines[0]))
					y_ := y + (dy * len(lines))

					n_ := int(n) + dx + dy
					if n_ != 9 {
						n_ %= 9
					}

					b[pos{X: x_, Y: y_}] = n_
				}
			}
		}
	}

	log.Printf("Board %d,%d - %d,%d", start.X, start.Y, end.X, end.Y)

	rank := b.getMinPathRank(start, end)

	return rank
}

type board map[pos]int

func (s board) getMinPathRank(start, end pos) int {
	d := map[pos]int{}

	q := &priorityQueue{
		items: make([]pos, 0, len(s)),
		idx:   make(map[pos]int, len(s)),
		pri:   make(map[pos]int, len(s)),
	}
	for k := range s {
		if k != start {
			d[k] = math.MaxInt64
		}
		q.AddRankPos(k, d[k])
	}

	for len(q.items) > 0 {
		p := heap.Pop(q).(pos)

		for _, v := range next(p, start, end) {
			r := d[p] + int(s[v])
			if r < d[v] {
				d[v] = r
				q.Update(v, r)
			}
		}
	}

	return d[end]
}

func next(p pos, min, max pos) []pos {
	var r []pos
	for _, d := range []pos{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
		p_ := pos{X: p.X + d.X, Y: p.Y + d.Y}
		if p_.X >= min.X && p_.Y >= min.Y && p_.X <= max.X && p_.Y <= max.Y {
			r = append(r, p_)
		}
	}
	return r
}

type priorityQueue struct {
	items []pos
	pri   map[pos]int
	idx   map[pos]int
}

func (q *priorityQueue) AddRankPos(p pos, priority int) {
	heap.Push(q, p)
	q.Update(p, priority)
}

func (q *priorityQueue) Update(p pos, priority int) {
	q.pri[p] = priority
	heap.Fix(q, q.idx[p])
}

func (q *priorityQueue) Len() int {
	return len(q.items)
}

func (q *priorityQueue) Less(i, j int) bool {
	return q.pri[q.items[i]] < q.pri[q.items[j]]
}

func (q *priorityQueue) Swap(i, j int) {
	q.items[i], q.items[j] = q.items[j], q.items[i]
	q.idx[q.items[i]] = i
	q.idx[q.items[j]] = j
}

func (q *priorityQueue) Push(x interface{}) {
	item := x.(pos)
	q.idx[item] = len(q.items)
	q.items = append(q.items, item)
}

func (q *priorityQueue) Pop() interface{} {
	curr := q.items
	l := len(curr)
	item := curr[l-1]
	q.idx[item] = -1
	q.items = curr[:l-1]

	return item
}

func init() {
	aoc.RegisterSolution(15, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
