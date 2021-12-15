package day14

import (
	"log"
	"sort"
	"strings"

	"aoc"
)

func Part1(lines []string, apply int) int {

	tpl := []byte(lines[0])
	rules := map[[2]byte]byte{}

	for i := 2; i < len(lines); i++ {
		p := strings.SplitN(lines[i], " -> ", 2)
		rules[[2]byte{p[0][0], p[0][1]}] = p[1][0]
	}

	var formula []byte

	for j := 0; j < apply; j++ {
		var f []byte
		for i := 0; i < len(tpl)-1; i++ {
			rule := [2]byte{tpl[i], tpl[i+1]}
			f = append(f, tpl[i], rules[rule])
		}
		f = append(f, tpl[len(tpl)-1])
		tpl = f
		formula = f
	}

	counts := map[string]int{}
	for _, v := range formula {
		counts[string(v)]++
	}

	type cnt struct {
		S     string
		Count int
	}

	var countsArr []cnt
	for k, v := range counts {
		countsArr = append(countsArr, cnt{S: k, Count: v})
	}

	sort.Slice(countsArr, func(i, j int) bool {
		return countsArr[i].Count < countsArr[j].Count
	})

	return countsArr[len(countsArr)-1].Count - countsArr[0].Count
}

func Part2(lines []string, apply int) uint {

	tpl := []byte(lines[0])
	rules := map[[2]byte]byte{}

	for i := 2; i < len(lines); i++ {
		p := strings.SplitN(lines[i], " -> ", 2)
		rules[[2]byte{p[0][0], p[0][1]}] = p[1][0]
	}

	//
	// pairs
	//
	pairCounters := map[[2]byte]uint{}
	for i := 0; i < len(tpl)-1; i++ {
		k := [2]byte{tpl[i], tpl[i+1]}
		pairCounters[k] += 1
	}
	lastChar := tpl[len(tpl)-1]

	charCounters := map[byte]uint{}
	for i := 0; i < apply; i++ {
		charCounters = map[byte]uint{}
		newPairs := map[[2]byte]uint{}

		for k, v := range pairCounters {
			if b, ok := rules[k]; ok {
				newPairs[[2]byte{k[0], b}] += v
				newPairs[[2]byte{b, k[1]}] += v

				charCounters[k[0]] += v
				charCounters[b] += v
			}
		}

		pairCounters = newPairs
	}
	charCounters[lastChar] += 1

	log.Println("DEBUG")
	for k, v := range charCounters {
		log.Println(string(k), v)
	}

	var max uint
	for _, v := range charCounters {
		if v > max {
			max = v
		}
	}

	var min uint = max
	for _, v := range charCounters {
		if v < min {
			min = v
		}
	}

	return max - min
}

func init() {
	aoc.RegisterSolution(14, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines, 10)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines, 40)
		},
	})
}
