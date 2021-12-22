package day08

import (
	"strings"

	"aoc"
)

/*
 aaaa
b    c
b    c
 dddd
e    f
e    f
 gggg

Easy to guess: 1, 4, 7, 8
Hard: 2,3,5,6,9

By segments amount:
- 2: 1
- 3: 7
- 4: 4
- 5: 2, 3, 5
- 6: 0, 6, 9
- 7: 8
*/

func Part1(lines []string) int {

	var r int

	for _, l := range lines {
		p := strings.SplitN(l, "|", 2)
		//variants := strings.Fields(p[0])
		guess := strings.Fields(p[1])

		if v := findByLen(guess, 2); len(v) != 0 {
			r += len(v)
		}
		if v := findByLen(guess, 3); len(v) != 0 {
			r += len(v)
		}
		if v := findByLen(guess, 4); len(v) != 0 {
			r += len(v)
		}
		if v := findByLen(guess, 7); len(v) != 0 {
			r += len(v)
		}
	}

	return r
}

func Part2(lines []string) int {

	var r int

	for _, l := range lines {
		r += guessValue(l)
	}

	return r
}

func guessValue(input string) int {

	p := strings.SplitN(input, "|", 2)
	variants := strings.Fields(p[0])

	digits := make(map[int]string, 10)
	if v := findByLen(variants, 2); len(v) != 0 {
		digits[1] = v[0]
	}
	if v := findByLen(variants, 3); len(v) != 0 {
		digits[7] = v[0]
	}
	if v := findByLen(variants, 4); len(v) != 0 {
		digits[4] = v[0]
	}
	if v := findByLen(variants, 7); len(v) != 0 {
		digits[8] = v[0]
	}

	eg := removeAny(digits[8], digits[7])
	eg = removeAny(eg, digits[4])

	bd := removeAny(digits[4], digits[7])

	// find 5
	if vars := findByLen(variants, 5); len(vars) > 0 {
		for _, v := range vars {
			if containsAll(v, bd) {
				digits[5] = v
			}
		}
	}

	e := removeAny(eg, digits[5])
	// find 6
	if vars := findByLen(variants, 6); len(vars) > 0 {
		for _, v := range vars {
			if containsAll(v, bd) && containsAll(v, e) {
				digits[6] = v
				break
			}
		}
	}
	// find 2
	if vars := findByLen(variants, 5); len(vars) > 0 {
		for _, v := range vars {
			if containsAll(v, e) && !containsExactly(v, digits[5]) {
				digits[2] = v
			}
		}
	}
	// find 3
	if vars := findByLen(variants, 5); len(vars) > 0 {
		for _, v := range vars {
			if !containsExactly(v, digits[2]) && !containsExactly(v, digits[5]) {
				digits[3] = v
			}
		}
	}

	// find 9
	if vars := findByLen(variants, 6); len(vars) > 0 {
		for _, v := range vars {
			if notContainsAll(v, e) {
				digits[9] = v
			}
		}
	}

	// find 0
	if vars := findByLen(variants, 6); len(vars) > 0 {
		for _, v := range vars {
			if !containsExactly(v, digits[6]) && !containsExactly(v, digits[9]) {
				digits[0] = v
			}
		}
	}

	//
	// guess
	//
	var r int
	guess := strings.Fields(p[1])
	for _, v := range guess {
		for val, seg := range digits {
			if containsExactly(v, seg) {
				r *= 10
				r += val
				break
			}
		}
	}

	return r
}

func findByLen(segmentGroups []string, l int) []string {
	var r []string
	for _, v := range segmentGroups {
		if len(v) == l {
			r = append(r, v)
		}
	}

	return r
}

func removeAny(source, chars string) string {
	var b strings.Builder

	for _, s := range source {
		if !strings.ContainsAny(chars, string(s)) {
			b.WriteString(string(s))
		}
	}

	return b.String()
}

func notContainsAll(source, chars string) bool {
	for _, s := range source {
		if strings.ContainsAny(chars, string(s)) {
			return false
		}
	}
	return true
}

func containsAll(source, chars string) bool {
	var i int
	for _, s := range chars {
		if strings.Contains(source, string(s)) {
			i++
		}
	}
	return i == len(chars)
}

func containsExactly(source, chars string) bool {
	var i int
	for _, s := range chars {
		if strings.Contains(source, string(s)) {
			i++
		}
	}
	return i == len(chars) && len(chars) == len(source)
}

func init() {
	aoc.RegisterSolution(8, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
