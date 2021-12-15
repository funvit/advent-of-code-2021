package day12

import (
	"strings"

	"aoc"
)

const (
	startName = "start"
	endName   = "end"
)

type cave struct {
	name  string
	isBig bool
}

func (s cave) Name() string {
	return s.name
}

func (s cave) IsBig() bool {
	return s.isBig
}

func (s *cave) String() string {
	return s.name
}

func NewCave(name string) *cave {
	return &cave{
		name:  name,
		isBig: name[0] == strings.Title(string(name[0]))[0],
	}
}

type connection struct {
	From, To *cave
}

func Part1(lines []string) int {

	var start, end *cave
	var connections []connection
	caves := make(map[string]*cave, len(lines)*2)

	for _, l := range lines {
		p := strings.SplitN(l, "-", 2)

		a := NewCave(p[0])
		caves[a.Name()] = a

		b := NewCave(p[1])
		caves[b.Name()] = b

		connections = append(connections, connection{From: a, To: b})

		if a.Name() == startName {
			start = a
		}
		if b.Name() == endName {
			end = b
		}
	}

	if start == nil {
		panic("start cave not found")
	}
	if end == nil {
		panic("end cave not found")
	}

	w := new(walker)
	paths := w.FindPaths(*start, connections, isOnceCaveVisitPossible)

	// debug
	//log.Println("DEBUG")
	//for _, path := range paths {
	//	log.Println(pathToString(path))
	//}

	return len(paths)
}

func Part2(lines []string) int {

	var start, end *cave
	var connections []connection
	caves := make(map[string]*cave, len(lines)*2)

	for _, l := range lines {
		p := strings.SplitN(l, "-", 2)

		a := NewCave(p[0])
		caves[a.Name()] = a

		b := NewCave(p[1])
		caves[b.Name()] = b

		connections = append(connections, connection{From: a, To: b})

		if a.Name() == startName {
			start = a
		}
		if b.Name() == endName {
			end = b
		}
	}

	if start == nil {
		panic("start cave not found")
	}
	if end == nil {
		panic("end cave not found")
	}

	w := new(walker)
	paths := w.FindPaths(*start, connections, isTwiceCaveVisitPossible)

	// debug
	//log.Println("DEBUG")
	//for _, path := range paths {
	//	log.Println(pathToString(path))
	//}

	return len(paths)
}

type walker struct {
	relations   []connection
	currentPath []cave
	knownPaths  [][]cave
	foundPaths  [][]cave
	checker     checkerFunc
}
type checkerFunc func(cave_ cave, path []cave) bool

func (s *walker) FindPaths(start cave, relations []connection, checker checkerFunc) [][]cave {
	s.relations = relations
	s.checker = checker

	s.walk(start)

	return s.foundPaths
}

func (s *walker) walk(start cave) {

	s.currentPath = append(s.currentPath, start)

	for _, rp := range getRelatedCaves(s.relations, start) {
		if rp.Name() == startName {
			continue
		}
		if rp.Name() == endName {
			if !s.checker(rp, s.currentPath) {
				continue
			}
			path := append([]cave{}, s.currentPath...)
			path = append(s.currentPath, rp)
			s.foundPaths = append(s.foundPaths, path)

			continue
		}
		if !s.checker(rp, s.currentPath) {
			continue
		}

		s.walk(rp)
	}

	if len(s.currentPath) > 0 {
		s.currentPath = s.currentPath[0 : len(s.currentPath)-1]
	}
}

func isOnceCaveVisitPossible(cave_ cave, currentPath []cave) bool {

	if cave_.Name() == startName {
		return false
	}

	for _, v := range currentPath {
		if v.Name() == cave_.Name() && !cave_.IsBig() {
			return false
		}
	}

	return true
}

func isTwiceCaveVisitPossible(cave_ cave, currentPath []cave) bool {

	if cave_.Name() == startName {
		return false
	}

	counters := map[string]int{} // cave name => visit counters
	for _, v := range currentPath {
		if v.Name() == startName {
			continue
		}
		if v.Name() == endName {
			continue
		}
		if !v.IsBig() {
			counters[v.Name()]++
		}
	}
	for _, v := range currentPath {
		if !cave_.IsBig() && v.Name() == cave_.Name() {
			counters[v.Name()]++
			break
		}
	}

	var cavesVisitedTwice int
	for _, v := range counters {
		if v > 1 {
			cavesVisitedTwice += v
		}
	}

	return cavesVisitedTwice <= 2
}

func pathToString(path []cave) string {
	var b strings.Builder
	for i, p := range path {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(p.Name())
	}
	return b.String()
}

func getRelatedCaves(all []connection, source cave) []cave {
	var r []cave
	for _, v := range all {
		if v.From.Name() == source.Name() {
			r = append(r, *v.To)
			continue
		}
		if v.To.Name() == source.Name() {
			r = append(r, *v.From)
			continue
		}
	}

	return r
}

func init() {
	aoc.RegisterSolution(12, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
