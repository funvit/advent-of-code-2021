package day19

import "aoc"

type Point struct {
	X, Y, Z int64
}

func NewPoint(x, y, z int64) *Point {
	return &Point{
		X: x,
		Y: y,
		Z: z,
	}
}

func (s *Point) EqualCoords(p Point) bool {
	return s.X == p.X && s.Y == p.Y && s.Z == p.Z
}

func (s *Point) Difference(p Point) Point {
	return Point{
		X: s.X - p.X,
		Y: s.Y - p.Y,
		Z: s.Z - p.Z,
	}
}

func (s *Point) Move(dist Point) Point {
	return Point{
		X: s.X + dist.X,
		Y: s.Y + dist.Y,
		Z: s.Z + dist.Z,
	}
}

func (s *Point) ManhattanDistance(p Point) int64 {
	return aoc.AbsInt64(s.X-p.X) + aoc.AbsInt64(s.Y-p.Y) + aoc.AbsInt64(s.Z-p.Z)
}
