package day10

import (
	"bytes"
	"sort"
	"strings"

	"aoc"
)

func Part1(lines []string) int {

	illegal := map[string]int{}

	for _, l := range lines {
		r, ok := validateLine(l)
		if !ok && r.InvalidClosingChar != "" {
			illegal[r.InvalidClosingChar]++
		}
	}

	var r int
	for k, v := range illegal {
		switch k {
		case ")":
			r += v * 3
		case "]":
			r += v * 57
		case "}":
			r += v * 1197
		case ">":
			r += v * 25137
		}
	}

	return r
}

func Part2(lines []string) int {

	type invalid struct {
		Line             string
		ValidationResult validationResult
	}

	var invalidLines []invalid

	for _, l := range lines {
		r, ok := validateLine(l)
		if !ok && r.InvalidClosingChar != "" {
			continue
		}
		invalidLines = append(invalidLines, invalid{
			Line:             l,
			ValidationResult: r,
		})
	}

	var scores []int
	for _, v := range invalidLines {
		var r int
		for i := len(v.ValidationResult.NonClosed) - 1; i >= 0; i-- {
			c := v.ValidationResult.NonClosed[i]
			r *= 5
			switch c {
			case getPair(')'):
				r += 1
			case getPair(']'):
				r += 2
			case getPair('}'):
				r += 3
			case getPair('>'):
				r += 4
			}
		}
		scores = append(scores, r)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

type validationResult struct {
	InvalidClosingChar string
	NonClosed          []byte
}

func (s *validationResult) String() string {
	var r string

	if len(s.NonClosed) > 0 {
		r += string(bytes.Join([][]byte{s.NonClosed}, []byte{','}))
	}
	if s.InvalidClosingChar != "" {
		r += " non-closed: " + s.InvalidClosingChar
	}
	return r
}

func validateLine(line string) (res validationResult, valid bool) {

	var openings []byte

	for i := 0; i < len(line); i++ {
		b := line[i]
		c := string(b)

		if strings.ContainsAny(c, "({[<") {
			openings = append(openings, b)
			continue
		}
		if strings.ContainsAny(c, ")}]>") {
			expected := getPair(openings[len(openings)-1])
			if b != expected {
				return validationResult{
					InvalidClosingChar: c,
					NonClosed:          openings,
				}, false
			}
			openings = openings[:len(openings)-1]
			continue
		}
	}

	return validationResult{
			NonClosed: openings,
		},
		len(openings) == 0
}

func getPair(start uint8) uint8 {
	switch start {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	case '>':
		return '<'
	default:
		panic("unexpected char")
	}
}

func init() {
	aoc.RegisterSolution(10, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
