package day18

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"aoc"
)

var pairRegexp = regexp.MustCompile(`\[(\d+,\d+)\]`)
var regNumAtRightRE = regexp.MustCompile(`^\]+,\[*(\d+)`)
var regNumAtLeftRE = regexp.MustCompile(`(\d+)\]*,\[+$`)
var gte10RE = regexp.MustCompile(`(\d\d+)`)

func Part1(lines []string) int64 {

	var s string = lines[0]

	for i := 1; i < len(lines); i++ {
		s = "[" + s + "," + lines[i] + "]"
		s = reduce(s)
	}

	log.Println("reduce result:", s)

	return magnitude(s)
}

func Part2(lines []string) int64 {

	var maxMagnitude int64

	for i, a := range lines {
		for j, b := range lines {
			if j == i {
				continue
			}
			s := "[" + a + "," + b + "]"
			s = reduce(s)
			m := magnitude(s)

			if maxMagnitude < m {
				maxMagnitude = m
			}
		}
	}

	for i, a := range lines {
		for j, b := range lines {
			if j == i {
				continue
			}
			s := "[" + b + "," + a + "]"
			s = reduce(s)
			m := magnitude(s)

			if maxMagnitude < m {
				maxMagnitude = m
			}
		}
	}

	return maxMagnitude
}

func magnitude(s string) int64 {

	var r int64

	for {
		pairs := findPairs(s)
		if len(pairs) == 0 {
			return r
		}

		p := pairs[0]
		r = p.Pair.A*3 + p.Pair.B*2

		s = s[:p.IndexStart-1] + strconv.FormatInt(r, 10) + s[p.IndexEnd+1:]

		continue
	}
}

func sumLines(lines []string) string {

	var s string = lines[0]
	for i := 1; i < len(lines); i++ {
		s = "[" + s + "," + lines[i] + "]"
		s = reduce(s)
	}

	return s
}

type pair struct {
	A, B int64
}

type pairAt struct {
	Pair                 pair
	IndexStart, IndexEnd int
}

func findPairs(s string) []pairAt {
	locs := pairRegexp.FindAllStringSubmatchIndex(s, -1)

	if len(locs) == 0 {
		return nil
	}

	var r []pairAt

	for _, v := range locs {
		n := aoc.MustParseNumbers(strings.SplitN(s[v[2]:v[3]], ",", 2))
		r = append(r, pairAt{
			Pair: pair{
				A: n[0],
				B: n[1],
			},
			IndexStart: v[2],
			IndexEnd:   v[3],
		})
	}

	return r
}

func isInsideFourPairs(p pairAt, source string) (ok bool, depth int) {
	l := strings.Count(source[:p.IndexStart], "[") - strings.Count(source[:p.IndexStart], "]")
	return l > 4, l
}

func reduce(s string) string {

	log.Println("reducing:  ", s)

	var repeat bool = true
	for repeat {
		pairs := findPairs(s)
		repeat = false

		for _, p := range pairs {
			if ok, _ := isInsideFourPairs(p, s); !ok {
				continue
			}

			//log.Println("pair", p, "is inside of four pairs")

			nextNum, nextStart, nextEnd := findRegularNumberAtRight(s[p.IndexEnd:])
			_ = nextEnd
			//log.Println("regular number at right", nextNum)

			prevNum, prevStart, prevEnd := findRegularNumberAtLeft(s[:p.IndexStart])
			_ = prevEnd
			_ = prevStart
			//log.Println("regular number at left", prevNum)

			var prev, next newRegular

			if prevNum != -1 {
				prev.Set = true
				prev.Value = p.Pair.A + prevNum
				prev.Suffix = s[p.IndexStart-prevEnd : p.IndexStart-1]
			} else {
				prevStart = 1
			}
			if nextNum != -1 {
				next.Set = true
				next.Value = p.Pair.B + nextNum
				next.Prefix = s[p.IndexEnd+1 : p.IndexEnd+nextStart]
			} else {
				nextEnd = 1
			}

			s = s[:p.IndexStart-prevStart] + prev.String() + "0" + next.String() + s[p.IndexEnd+nextEnd:]
			repeat = true

			break
		}
		if !repeat {
			s, repeat = splitAnyOnce(s)
		}
	}

	return s
}

type newRegular struct {
	Value          int64
	Set            bool
	Prefix, Suffix string
}

func (s *newRegular) String() string {

	if !s.Set {
		return ""
	}

	var b strings.Builder

	b.WriteString(s.Prefix)
	b.WriteString(strconv.FormatInt(s.Value, 10))
	b.WriteString(s.Suffix)

	return b.String()
}

func findRegularNumberAtRight(s string) (number int64, start, end int) {

	v := regNumAtRightRE.FindStringSubmatchIndex(s)
	if len(v) >= 4 && v[2] != -1 && v[3] != -1 {
		return aoc.MustParseInt64(s[v[2]:v[3]]), v[2], v[3]
	}

	return -1, 0, 0
}

func findRegularNumberAtLeft(s string) (number int64, start, end int) {

	v := regNumAtLeftRE.FindStringSubmatchIndex(s)
	if len(v) >= 4 && v[2] != -1 && v[1] != -1 {
		return aoc.MustParseInt64(s[v[2]:v[3]]), v[1] - v[2], v[1] - v[3]
	}

	return -1, 0, 0
}

func splitAnyOnce(s string) (string, bool) {

	loc := gte10RE.FindStringIndex(s)
	if len(loc) == 0 {
		return s, false
	}

	n := aoc.MustParseInt64(s[loc[0]:loc[1]])
	h := n / 2
	sub := fmt.Sprintf("[%d,%d]", h, n-h)

	return s[:loc[0]] + sub + s[loc[1]:], true
}

func init() {
	aoc.RegisterSolution(18, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
