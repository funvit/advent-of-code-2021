package aoc

import "math"

type Int64Counter struct {
	m map[int64]int
}

func NewInt64Counter() *Int64Counter {
	return &Int64Counter{
		m: map[int64]int{},
	}
}

func (s *Int64Counter) Inc(v int64) {
	s.m[v] = s.m[v] + 1
}

func (s *Int64Counter) MostCommon() (val int64, n int) {

	var mcVal int64
	var mc int

	for k, v := range s.m {
		if mc < v {
			mcVal = k
			mc = v
		}
	}

	return mcVal, mc
}

func (s *Int64Counter) GetByN(n int) (int64, bool) {
	for k, v := range s.m {
		if n == v {
			return k, true
		}
	}
	return 0, false
}

func MinInt(n ...int) int {
	var r int = math.MaxInt64

	for _, v := range n {
		if r > v {
			r = v
		}
	}

	return r
}

func MinInt64(n ...int64) int64 {
	var r int64 = math.MaxInt64

	for _, v := range n {
		if r > v {
			r = v
		}
	}

	return r
}

func MaxInt(n ...int) int {
	var r int = math.MinInt64

	for _, v := range n {
		if r < v {
			r = v
		}
	}

	return r
}

func MaxInt64(n ...int64) int64 {
	var r int64 = math.MinInt64

	for _, v := range n {
		if r < v {
			r = v
		}
	}

	return r
}

func AbsInt64(n int64) int64 {
	if n < 0 {
		return -n
	}
	return n
}
