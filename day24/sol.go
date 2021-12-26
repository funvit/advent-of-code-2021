package day24

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync/atomic"

	"aoc"
)

/*
	Note: output arg type of Part1 and Part2 can be freely changed.
*/

// number have 14 digits
const digits = 14

/*
	Idea: collect "div z", `add x -?\d+`, `add x \d+` before first eql
*/

var _lastAluId int64

const (
	inp = uint8(iota + 1)
	add
	mul
	div
	mod
	eql
)

type ALU struct {
	X, Y, W, Z      int64 // registers
	InstructionSets []*InstructionSet
	id              int64
}

func (a *ALU) Clone() *ALU {
	c := &ALU{
		id: atomic.AddInt64(&_lastAluId, 1),
	}

	for _, v := range a.InstructionSets {
		s := NewInstructionSet(c)
		s.Set = append(s.Set, v.Set...)
		s.inputInto = s.src(v.inputIntoRegName)
		c.InstructionSets = append(c.InstructionSets, s)
	}

	return c
}

type InstructionSet struct {
	x, y, w, z       *int64
	Set              []Instruction
	inputInto        *int64
	inputIntoRegName string // need for alu cloning
}
type Instruction struct {
	Kind  uint8
	Apply func(set *InstructionSet)
}

func NewInstructionSet(a *ALU) *InstructionSet {
	return &InstructionSet{
		x: &a.X,
		y: &a.Y,
		w: &a.W,
		z: &a.Z,
	}
}

func (is *InstructionSet) Input(v int64) {
	*is.inputInto = v
}

func (is *InstructionSet) src(regName string) *int64 {
	var v *int64
	switch regName {
	case "x":
		v = is.x
	case "y":
		v = is.y
	case "w":
		v = is.w
	case "z":
		v = is.z
	default:
		panic("cannot find source " + regName)
	}
	return v
}

func inpInst(s string, is *InstructionSet) Instruction {
	is.inputInto = is.src(string(s[4]))
	is.inputIntoRegName = string(s[4])

	return Instruction{
		Kind: inp,
		Apply: func(is *InstructionSet) {
			// dummy
		},
	}
}

func addInst(s string) Instruction {
	return Instruction{
		Kind: add,
		Apply: func(is *InstructionSet) {
			var src *int64

			b, err := strconv.ParseInt(s[6:], 10, 64)
			if err != nil {
				name := s[6:]
				src = is.src(name)
			} else {
				src = &b
			}

			name := string(s[4])
			switch name {
			case "x":
				*is.x += *src
			case "y":
				*is.y += *src
			case "w":
				*is.w += *src
			case "z":
				*is.z += *src
			default:
				panic("cannot store into " + name)
			}
		},
	}
}

func mulInst(s string) Instruction {
	return Instruction{
		Kind: mul,
		Apply: func(is *InstructionSet) {
			var src *int64

			b, err := strconv.ParseInt(s[6:], 10, 64)
			if err != nil {
				name := s[6:]
				src = is.src(name)
			} else {
				src = &b
			}

			name := string(s[4])
			switch name {
			case "x":
				*is.x *= *src
			case "y":
				*is.y *= *src
			case "w":
				*is.w *= *src
			case "z":
				*is.z *= *src
			default:
				panic("cannot store into " + name)
			}
		},
	}
}

func divInst(s string) Instruction {
	return Instruction{
		Kind: div,
		Apply: func(is *InstructionSet) {
			var src *int64

			b, err := strconv.ParseInt(s[6:], 10, 64)
			if err != nil {
				name := s[6:]
				src = is.src(name)
			} else {
				src = &b
			}

			name := string(s[4])
			switch name {
			case "x":
				*is.x /= *src
			case "y":
				*is.y /= *src
			case "w":
				*is.w /= *src
			case "z":
				*is.z /= *src
			default:
				panic("cannot store into " + name)
			}
		},
	}
}

func modInst(s string) Instruction {
	return Instruction{
		Kind: mod,
		Apply: func(is *InstructionSet) {
			var src *int64

			b, err := strconv.ParseInt(s[6:], 10, 64)
			if err != nil {
				name := s[6:]
				src = is.src(name)
			} else {
				src = &b
			}

			name := string(s[4])
			switch name {
			case "x":
				*is.x %= *src
			case "y":
				*is.y %= *src
			case "w":
				*is.w %= *src
			case "z":
				*is.z %= *src
			default:
				panic("cannot store into " + name)
			}
		},
	}
}

func eqlInst(s string) Instruction {
	return Instruction{
		Kind: eql,
		Apply: func(is *InstructionSet) {
			var src *int64

			b, err := strconv.ParseInt(s[6:], 10, 64)
			if err != nil {
				name := s[6:]
				src = is.src(name)
			} else {
				src = &b
			}

			name := string(s[4])
			switch name {
			case "x":
				if *is.x == *src {
					*is.x = 1
				} else {
					*is.x = 0
				}
			case "y":
				if *is.y == *src {
					*is.y = 1
				} else {
					*is.y = 0
				}
			case "w":
				if *is.w == *src {
					*is.w = 1
				} else {
					*is.w = 0
				}
			case "z":
				if *is.z == *src {
					*is.z = 1
				} else {
					*is.z = 0
				}
			default:
				panic("cannot store into " + name)
			}
		},
	}
}

func (is *InstructionSet) Process() {
	for _, v := range is.Set {
		v.Apply(is)
	}
}

func Part1(lines []string) int64 {

	xas, yas, zms := parseInputConsts(lines)
	ws := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := getNumber(xas, yas, zms, ws)

	return r
}

func Part2(lines []string) int64 {

	xas, yas, zms := parseInputConsts(lines)
	ws := []int64{9, 8, 7, 6, 5, 4, 3, 2, 1}
	r := getNumber(xas, yas, zms, ws)

	return r
}

func getNumber(xas []int64, yas []int64, zms []int64, ws []int64) int64 {
	zs := []int64{0} // ending
	results := map[int64][]int64{}

	log.Println("note: processing in reverse")

	for i := len(xas) - 1; i >= 0; i-- {
		log.Println("processing index:", i, "known z values:", len(zs))
		xa := xas[i]
		ya := yas[i]
		zm := zms[i]

		var newZs []int64

		for _, w := range ws {
			for _, z := range zs {

				zs_ := backwardZ(xa, ya, zm, z, w)
				for _, v := range zs_ {
					newZs = append(newZs, v)
					results[v] = append([]int64{w}, results[z]...)
				}
			}
		}
		zs = newZs
	}

	var r int64 = -1
	if n, ok := results[0]; ok {
		r = 0
		for i := 0; i < len(n); i++ {
			r = r*10 + n[i]
		}
	}
	return r
}

var ( // magic constants
	xar = regexp.MustCompile(`add x (-?\d+)`)
	yar = regexp.MustCompile(`add y w\nadd y (\d+)\nmul y x`)
	zmr = regexp.MustCompile(`div z (\d+)`)
)

func parseInputConsts(lines []string) ([]int64, []int64, []int64) {
	var xas, yas, zms []int64

	text := strings.Join(lines, "\n")
	parts := strings.Split(text, "inp w")

	for _, p := range parts {
		if p == "" {
			continue
		}
		xa := aoc.MustParseInt64(xar.FindStringSubmatch(p)[1])
		ya := aoc.MustParseInt64(yar.FindStringSubmatch(p)[1])
		zm := aoc.MustParseInt64(zmr.FindStringSubmatch(p)[1])

		xas = append(xas, xa)
		yas = append(yas, ya)
		zms = append(zms, zm)
	}
	return xas, yas, zms
}

func calculateZ(xa, ya, zm, z, w int64) int64 {
	z_ := z / zm
	if w == z%26+xa {
		return z_
	}
	return 26 + z_ + w + ya
}

func backwardZ(xa, ya, zm, z, w int64) []int64 {

	var zs []int64
	x := z - w - ya
	if x%26 == 0 {
		zs = append(zs, x/26*zm)
	}
	if 0 <= w-xa && w-xa < 26 {
		zs = append(zs, w-xa+z*zm)
	}
	return zs
}

func parseALU(lines []string) *ALU {

	var alu ALU
	var startFrom int

	for i := 0; i < digits; i++ {
		var g *InstructionSet
		for lineIdx, line := range lines[startFrom:] {
			insName := line[0:3]
			switch insName {
			case "inp":
				if g != nil {
					startFrom = lineIdx * (i + 1)
					goto out
				}
				g = NewInstructionSet(&alu)
				g.Set = append(g.Set, inpInst(line, g))
			case "add":
				g.Set = append(g.Set, addInst(line))
			case "mul":
				g.Set = append(g.Set, mulInst(line))
			case "div":
				g.Set = append(g.Set, divInst(line))
			case "mod":
				g.Set = append(g.Set, modInst(line))
			case "eql":
				g.Set = append(g.Set, eqlInst(line))
			default:
				panic("unknown instruction at line " + strconv.Itoa(lineIdx*i))
			}
		}
	out:
		alu.InstructionSets = append(alu.InstructionSets, g)
		g = nil
	}

	return &alu
}

func checkModelNumber(num int64, alu *ALU) bool {

	a := strconv.FormatInt(num, 10)
	if strings.Contains(a, "0") {
		return false
	}

	var number []int64 = aoc.MustParseNumbers(strings.SplitN(
		strconv.FormatInt(num, 10),
		"",
		14,
	))

	for i := 0; i < len(number); i++ {
		n := number[i]
		g := alu.InstructionSets[i]

		g.Input(n)
		g.Process()
	}

	return alu.Z == 0
}

func init() {
	// this registers solution to be able to run from ./cmd/main.go,
	// BUT you must add import like `_ "aoc/day24"` in "main.go" by hands!
	aoc.RegisterSolution(24, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
