package day25

import (
	"fmt"
	"strings"

	"aoc"
)

/*
	Note: output arg type of Part1 and Part2 can be freely changed.
*/

const (
	cucumberE  = ">"
	cucumberS  = "v"
	emptySpace = "."
)

func Part1(lines []string) int {
	//todo: write code here

	maxWidth := len(lines[0])
	maxHeight := len(lines)

	var cucumbers []string
	for _, line := range lines {
		cucumbers = append(cucumbers, strings.Split(line, "")...)
	}

	g := grid{
		MaxWidth:  maxWidth,
		MaxHeight: maxHeight,
	}

	//log.Println("input data:")
	//g.DrawFn(cucumbers, func(s string) {
	//	log.Println(s)
	//})

	var moved int = -1
	var steps int
	for moved > 0 || moved == -1 {
		moved = 0
		newState := append([]string{}, cucumbers...)
		for y := 0; y < maxHeight; y++ {
			for x := 0; x < maxWidth; x++ {
				curIdx := g.GetIdx(x, y)
				cur := cucumbers[curIdx]

				nextIdx := g.GetNextEIdx(x, y)
				next := cucumbers[nextIdx]

				if next == emptySpace && cur == cucumberE {
					newState[nextIdx] = cur
					newState[curIdx] = emptySpace
					moved++
				}
			}
		}
		cucumbers = append([]string{}, newState...)

		//log.Println("step", steps, "first part data:")
		//g.DrawFn(cucumbers, func(s string) {
		//	log.Println(s)
		//})

		for x := 0; x < maxWidth; x++ {
			for y := 0; y < maxHeight; y++ {
				curIdx := g.GetIdx(x, y)
				cur := cucumbers[curIdx]

				nextIdx := g.GetNextSIdx(x, y)
				next := cucumbers[nextIdx]

				if next == emptySpace && cur == cucumberS {
					newState[nextIdx] = cur
					newState[curIdx] = emptySpace
					moved++
				}
			}
		}
		cucumbers = append([]string{}, newState...)

		steps++

		//log.Println("step", steps, "data:")
		//g.DrawFn(cucumbers, func(s string) {
		//	log.Println(s)
		//})
	}

	return steps
}

func Part2(lines []string) int {
	//Note: no actual part two :)

	fmt.Println()
	fmt.Println("Congratulations! You've finished every puzzle in Advent of Code 2021!")
	fmt.Println()

	fmt.Print(`
    _\/_
     /\
     /\
    /  \
    /~~\o
   /o   \
  /~~*~~~\
 o/    o \
 /~~~~~~~~\~'
/__*_______\
     ||
   \====/
    \__/ Art by Shanaka Dias

`)

	return -1
}

type grid struct {
	MaxWidth  int
	MaxHeight int
}

func (s *grid) GetIdx(x, y int) int {
	return x + y*s.MaxWidth
}

func (s *grid) GetNextEIdx(x, y int) int {
	return (x+1)%s.MaxWidth + y*s.MaxWidth
}

func (s *grid) GetNextSIdx(x, y int) int {
	return x + (y+1)%s.MaxHeight*s.MaxWidth
}

func (s *grid) DrawFn(cucumbers []string, fn func(s string)) {

	for y := 0; y < s.MaxHeight; y++ {
		var b strings.Builder
		for x := 0; x < s.MaxWidth; x++ {
			c := cucumbers[s.GetIdx(x, y)]
			if c == emptySpace {
				b.WriteString(emptySpace)
				continue
			}
			b.WriteString(c)
		}
		fn(fmt.Sprint(b.String()))
	}
}

func init() {
	// this registers solution to be able to run from ./cmd/main.go,
	// BUT you must add import like `_ "aoc/day25"` in "main.go" by hands!
	aoc.RegisterSolution(25, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
