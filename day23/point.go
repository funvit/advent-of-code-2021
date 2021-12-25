package day23

import (
	"fmt"
	"strings"
)

type Point struct {
	X, Y int64
}

func (s Point) String() string {
	return fmt.Sprintf("(%d,%d)", s.X, s.Y)
}

type PointSet struct {
	m     map[Point]struct{}
	order []*Point
}

func (s *PointSet) String() string {
	var b strings.Builder
	var i int
	s.Each(func(p Point) {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(p.String())
		i++
	})
	return b.String()
}

func NewPointSet() *PointSet {
	return &PointSet{m: map[Point]struct{}{}}
}

func (s *PointSet) Add(p Point) {
	s.m[p] = struct{}{}
	s.order = append(s.order, &p)
}

func (s *PointSet) Each(fn func(Point)) {
	for _, v := range s.order {
		fn(*v)
	}
}

func (s *PointSet) First(matcher func(Point) bool, found func(Point)) {
	for _, v := range s.order {
		if matcher(*v) {
			found(*v)
			break
		}
	}
}

func (s *PointSet) Delete(p Point) {
	delete(s.m, p)

	var idx int = -1
	for i := range s.order {
		if *s.order[i] == p {
			idx = i
			break
		}
	}
	if idx != -1 {
		s.order = append(s.order[:idx], s.order[idx+1:]...)
	}
}

func (s *PointSet) Exists(p Point) bool {
	_, ok := s.m[p]
	return ok
}

func (s *PointSet) Len() int {
	return len(s.order)
}
