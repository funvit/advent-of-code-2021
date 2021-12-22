package day19

type PointsSet struct {
	m map[Point]struct{}
}

func NewPointsSet() *PointsSet {
	return PointsSetWithSize(100)
}

func PointsSetWithSize(n int) *PointsSet {
	return &PointsSet{
		m: make(map[Point]struct{}, n),
	}
}

func (s *PointsSet) Add(p Point) {
	s.m[p] = struct{}{}
}

func (s *PointsSet) Each(fn func(p Point)) {
	for k := range s.m {
		fn(k)
	}
}

func (s *PointsSet) Move(dist Point) *PointsSet {
	r := PointsSetWithSize(s.Len())
	s.Each(func(p Point) {
		r.Add(p.Move(dist))
	})
	return r
}

func (s *PointsSet) Apply(fn func(Point) Point) *PointsSet {
	r := PointsSetWithSize(s.Len())
	s.Each(func(p Point) {
		r.Add(fn(p))
	})
	return r
}

func (s *PointsSet) Merge(set PointsSet) *PointsSet {
	r := PointsSetWithSize(s.Len() + set.Len())
	s.Each(func(p Point) {
		r.Add(p)
	})
	set.Each(func(p Point) {
		r.Add(p)
	})
	return r
}

func (s *PointsSet) Clone() *PointsSet {
	r := PointsSetWithSize(s.Len())
	s.Each(func(p Point) {
		r.Add(p)
	})
	return r
}

func (s *PointsSet) Len() int {
	return len(s.m)
}
