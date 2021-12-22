package day19

type PointCounter struct {
	m map[Point]int
}

func NewPointCounter(size int) *PointCounter {
	return &PointCounter{
		m: make(map[Point]int, size),
	}
}

func (s *PointCounter) Inc(p Point) {
	s.m[p] = s.m[p] + 1
}

func (s *PointCounter) MostCommon() (point Point, n int) {

	for k, v := range s.m {
		if n < v {
			point = k
			n = v
		}
	}

	return point, n
}
