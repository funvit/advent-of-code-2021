package day19

import (
	"log"
	"strings"

	"aoc"
)

type Scanner struct {
	Id     int64
	Points PointsSet
}

// Part1 returns total amount of beacons without duplicates.
func Part1(lines []string, leastSameBeacons int) int {

	const scannerStart = "--- scanner "

	//
	// read input
	//
	scannerById := map[int64]*Scanner{}
	var scannerIds []int64
	{
		var scanner *Scanner
		var sid int64
		for i, line := range lines {
			if strings.TrimSpace(line) == "" {
				scannerById[scanner.Id] = scanner
				continue
			}
			if strings.HasPrefix(line, scannerStart) {
				scanner = &Scanner{Id: sid, Points: *NewPointsSet()}
				scannerIds = append(scannerIds, sid)
				sid++
				continue
			}
			p := strings.SplitN(line, ",", 3)
			n := aoc.MustParseNumbers(p)
			pt := NewPoint(n[0], n[1], n[2])
			scanner.Points.Add(*pt)

			if i == len(lines)-1 {
				scannerById[scanner.Id] = scanner
			}
		}
	}
	log.Println("input read done, total scanners:", len(scannerIds))

	//
	// process
	//
	knownPoints := scannerById[0].Points.Clone()

	var unknownPoints []PointsSet
	for k, v := range scannerById {
		if k == 0 {
			continue
		}
		unknownPoints = append(unknownPoints, *v.Points.Clone())
	}

	for len(unknownPoints) > 0 {
		log.Println("known points len:", knownPoints.Len(), "unknown sets:", len(unknownPoints))

		overlap, data := isOverlaps(*knownPoints, unknownPoints[0], leastSameBeacons)

		if overlap {
			log.Println("...overlap with diff:", data.Diff, "duplicates:", data.Duplicates)

			rPoints := unknownPoints[0].Apply(data.PointTransform)

			// update known points
			knownPoints = knownPoints.Merge(*rPoints)
			log.Println("...known points after merge:", knownPoints.Len())

		} else {
			if len(unknownPoints) <= 2 {
				for _, v := range unknownPoints {
					knownPoints = knownPoints.Merge(v)
				}
				break
			}
			// fixme: detect full loop of unknown points? (possible forever loop)...
			if len(unknownPoints) > 2 {
				log.Println("...no overlap...")
				// put current unknown points set to end of sets.
				unknownPoints = append(unknownPoints[1:], unknownPoints[0])
				continue
			}
		}
		// next region
		unknownPoints = unknownPoints[1:]
	}

	return knownPoints.Len()
}

// Part2 returns the largest "Manhattan" distance between any two scanners.
func Part2(lines []string, leastSameBeacons int) int64 {

	const scannerStart = "--- scanner "

	//
	// read input
	//
	scannerById := map[int64]*Scanner{}
	var scannerIds []int64
	{
		var scanner *Scanner
		var sid int64
		for i, line := range lines {
			if strings.TrimSpace(line) == "" {
				scannerById[scanner.Id] = scanner
				continue
			}
			if strings.HasPrefix(line, scannerStart) {
				scanner = &Scanner{Id: sid, Points: *NewPointsSet()}
				scannerIds = append(scannerIds, sid)
				sid++
				continue
			}
			p := strings.SplitN(line, ",", 3)
			n := aoc.MustParseNumbers(p)
			pt := NewPoint(n[0], n[1], n[2])
			scanner.Points.Add(*pt)

			if i == len(lines)-1 {
				scannerById[scanner.Id] = scanner
			}
		}
	}
	log.Println("input read done, total scanners:", len(scannerIds))

	//
	// process
	//
	knownPoints := scannerById[0].Points.Clone()
	scannerDiffs := map[Point]struct{}{}

	var unknownPoints []PointsSet
	for k, v := range scannerById {
		if k == 0 {
			continue
		}
		unknownPoints = append(unknownPoints, *v.Points.Clone())
	}

	for len(unknownPoints) > 0 {
		overlap, data := isOverlaps(*knownPoints, unknownPoints[0], leastSameBeacons)

		if overlap {
			rPoints := unknownPoints[0].Apply(data.PointTransform)

			// update known points
			knownPoints = knownPoints.Merge(*rPoints)

			scannerDiffs[data.Diff] = struct{}{}

		} else {
			if len(unknownPoints) <= 2 {
				for _, v := range unknownPoints {
					knownPoints = knownPoints.Merge(v)
				}
				break
			}
			// fixme: detect full loop of unknown points? (possible forever loop)...
			if len(unknownPoints) > 2 {
				// put current unknown points set to end of sets.
				unknownPoints = append(unknownPoints[1:], unknownPoints[0])
				continue
			}
		}
		// next region
		unknownPoints = unknownPoints[1:]
	}

	var result int64
	for k1 := range scannerDiffs {
		for k2 := range scannerDiffs {
			if k2 == k1 {
				continue
			}
			d := k1.ManhattanDistance(k2)
			if d > result {
				result = d
			}
		}
	}

	return result
}

type overlapData struct {
	Diff           Point
	Duplicates     int
	PointTransform func(Point) Point
}

func isOverlaps(source PointsSet, other PointsSet, leastMostCommon int) (bool, *overlapData) {

	var i int

	for atIdx, at := range allAxisTransformations() {
		_ = atIdx
		for rotations := 0; rotations < 3; rotations++ {
			c := NewPointCounter(source.Len() + other.Len())

			source.Each(func(sp Point) {
				other.Each(func(op Point) {
					i++
					tp := rotateAxis(at(op), rotations)
					diff := sp.Difference(tp)
					c.Inc(diff)
				})
			})

			diff, mc := c.MostCommon()
			if mc >= leastMostCommon {
				log.Println("...isOverlaps iterations:", i)

				t := at
				r := rotations
				return true, &overlapData{
					Diff:       diff,
					Duplicates: mc,
					PointTransform: func(p Point) Point {
						// main trick is to store diff with applied transformation func.
						pp := rotateAxis(t(p), r)
						return pp.Move(diff)
					},
				}
			}
		}
	}

	return false, nil
}

func allAxisTransformations() []func(p Point) Point {
	/*
		Y
		^  .
		| / Z
		|/
		*----> X
	*/
	return []func(p Point) Point{
		// first: looking at plane X,Y by Z axis
		func(p Point) Point { return RotatePointZ(p, 0) },
		func(p Point) Point { return RotatePointZ(p, 1) },
		func(p Point) Point { return RotatePointZ(p, 2) },
		func(p Point) Point { return RotatePointZ(p, 3) },

		// second: looking by -Z, so X,Y become X,-Y
		func(p Point) Point { return RotatePointZ(Point{X: p.X, Y: -p.Y, Z: -p.Z}, 0) },
		func(p Point) Point { return RotatePointZ(Point{X: p.X, Y: -p.Y, Z: -p.Z}, 1) },
		func(p Point) Point { return RotatePointZ(Point{X: p.X, Y: -p.Y, Z: -p.Z}, 2) },
		func(p Point) Point { return RotatePointZ(Point{X: p.X, Y: -p.Y, Z: -p.Z}, 3) },
	}
}

func rotateAxis(p Point, n int) Point {
	for i := 0; i < n; i++ {
		p = Point{
			X: p.Y,
			Y: p.Z,
			Z: p.X,
		}
	}
	return p
}

func RotatePointZ(p Point, n int) Point {
	for i := 0; i < n; i++ {
		p = Point{
			X: p.Y,
			Y: -p.X,
			Z: p.Z,
		}
	}
	return p
}

func init() {
	aoc.RegisterSolution(19, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines, 12)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines, 12)
		},
	})
}
