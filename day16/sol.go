package day16

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"aoc"
)

type verSumScanner struct {
	verSumAcc int64
	idx       int
	bits      []byte
}

func VerSum(bits []byte) int64 {
	s := verSumScanner{bits: bits}
	s.verSum()
	return s.verSumAcc
}

func (s *verSumScanner) read(n int) []byte {
	s.idx += n
	return s.bits[s.idx-n : s.idx]
}

func (s *verSumScanner) verSum() {

	v := intFromBits(s.read(3))
	t := intFromBits(s.read(3))
	s.verSumAcc += v

	if t == 4 {
		for j := 0; j < len(s.bits); j += 5 {
			g := s.read(1)[0]
			_ = s.read(4)
			if g == '0' {
				break
			}
		}
	} else {
		switch s.read(1)[0] {
		case '0':
			const lenSize = 15
			subPacketsLen := int(intFromBits(s.read(lenSize)))
			end := s.idx + subPacketsLen
			for s.idx < end {
				s.verSum()
			}
		default:
			const lenSize = 11
			subPacketsCnt := int(intFromBits(s.read(lenSize)))
			for i := 0; i < subPacketsCnt; i++ {
				s.verSum()
			}
		}
	}
}

// Part1 returns sum of packet`s version numbers.
func Part1(lines []string) int64 {

	line := lines[0]

	var sb strings.Builder
	for i := range line {
		switch line[i] {
		case 'A', 'B', 'C', 'D', 'E', 'F':
			sb.WriteString(fmt.Sprintf("%b", line[i]-55))
		default:
			sb.WriteString(fmt.Sprintf("%04b", line[i]-'0'))
		}
	}
	bits := []byte(sb.String())

	log.Println("input len:", len(bits))

	return VerSum(bits)
}

type packet struct {
	Version  int64
	Type     int64
	Children []packet
	Value    int64 // used only with Type == 4
}

func (s *packet) Calculate() int64 {
	return s.calc(*s)
}

func (s *packet) calc(p packet) int64 {
	switch p.Type {
	case 0: // sum
		var r int64
		for _, sp := range p.Children {
			r += s.calc(sp)
		}
		return r

	case 1: // product
		var r int64 = 1
		for _, sp := range p.Children {
			r *= s.calc(sp)
		}
		return r

	case 2: // min
		var r int64
		for _, sp := range p.Children {
			v := s.calc(sp)
			if v < r || r == 0 {
				r = v
			}
		}
		return r

	case 3: // max
		var r int64
		for _, sp := range p.Children {
			v := s.calc(sp)
			if v > r || r == 0 {
				r = v
			}
		}
		return r

	case 4:
		return p.Value

	case 5: // gt
		v1 := s.calc(p.Children[0])
		v2 := s.calc(p.Children[1])
		if v1 > v2 {
			return 1
		}
		return 0

	case 6: // lt
		v1 := s.calc(p.Children[0])
		v2 := s.calc(p.Children[1])
		if v1 < v2 {
			return 1
		}
		return 0

	case 7: //eq
		v1 := s.calc(p.Children[0])
		v2 := s.calc(p.Children[1])
		if v1 == v2 {
			return 1
		}
		return 0

	default:
		panic("unknown packet type")
	}
}

// Part2 calculates expression result.
func Part2(lines []string) int64 {

	line := lines[0]

	var sb strings.Builder
	for i := range line {
		switch line[i] {
		case 'A', 'B', 'C', 'D', 'E', 'F':
			sb.WriteString(fmt.Sprintf("%b", line[i]-55))
		default:
			sb.WriteString(fmt.Sprintf("%04b", line[i]-'0'))
		}
	}
	bits := []byte(sb.String())

	log.Println("input len:", len(bits))

	p := Read(bits)

	//fmt.Printf("%+v\n", p)

	return p.Calculate()
}

type scanner struct {
	idx  int
	bits []byte
}

func Read(bits []byte) *packet {
	s := scanner{bits: bits}
	return s.scan()
}

func (s *scanner) read(n int) []byte {
	s.idx += n
	return s.bits[s.idx-n : s.idx]
}

func (s *scanner) scan() *packet {

	p := &packet{
		Version:  intFromBits(s.read(3)),
		Type:     intFromBits(s.read(3)),
		Children: nil,
		Value:    0,
	}

	if p.Type == 4 {
		var bitsAcc []byte
		for j := 0; j < len(s.bits); j += 5 {
			g := s.read(1)[0]
			bitsAcc = append(bitsAcc, s.read(4)...)
			if g == '0' {
				p.Value = intFromBits(bitsAcc)
				break
			}
		}
	} else {
		switch s.read(1)[0] {
		case '0':
			const lenSize = 15
			subPacketsLen := intFromBits(s.read(lenSize))
			end := s.idx + int(subPacketsLen)
			for s.idx < end {
				sp := s.scan()
				p.Children = append(p.Children, *sp)
			}
		default:
			const lenSize = 11
			subPacketsCnt := intFromBits(s.read(lenSize))
			for i := 0; i < int(subPacketsCnt); i++ {
				sp := s.scan()
				p.Children = append(p.Children, *sp)
			}
		}
	}

	return p
}

// bits example: 100101011010
func intFromBits(bits []byte) int64 {
	r, err := strconv.ParseInt(string(bits), 2, 64)
	if err != nil {
		panic(err)
	}
	return r
}

func init() {
	aoc.RegisterSolution(16, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
