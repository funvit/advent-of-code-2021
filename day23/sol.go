package day23

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
	"unicode"

	"aoc"
)

/*
	Note: output arg type of Part1 and Part2 can be freely changed.
*/

/*
	Game board example:

    #############
    #...........#
    ###A#B#C#D###
      #A#B#C#D#
      #########

*/

var podEnergyPerMove = map[string]int64{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}
var finishCols = map[string]int64{
	"A": 2,
	"B": 4,
	"C": 6,
	"D": 8,
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Pod struct {
	Kind        string
	Name        string
	EnergyTaken int64
	Pos         Point
	IsFinished  bool
	Moves       int
	Index       int
}

func (s *Pod) String() string {
	return s.Name
}

var _lastBoardId int64

type Board struct {
	HallwayPossiblePositions *PointSet
	BestFinishPositions      map[string]Point // pod kind => position
	Pods                     []*Pod

	id       int64
	origId   int64
	states   []stateAlt
	finished bool
}

func NewBoard() *Board {
	return &Board{
		HallwayPossiblePositions: NewPointSet(),
		BestFinishPositions:      map[string]Point{},
		Pods:                     nil,
		id:                       atomic.AddInt64(&_lastBoardId, 1),
	}
}

func (s *Board) State() string {
	var b strings.Builder

	var items []string
	for _, p := range s.Pods {
		p := p
		items = append(items, fmt.Sprintf(
			"%s-%d,%d",
			//p.Name,
			p.Kind,
			p.Pos.X,
			p.Pos.Y,
		))
	}
	sort.Strings(items)

	for k, v := range items {
		if k > 0 {
			b.WriteString(",")
		}
		b.WriteString(v)
	}

	return b.String()
}

func (s *Board) StateLong() string {
	var b strings.Builder

	var items []string
	for _, p := range s.Pods {
		p := p
		items = append(items, fmt.Sprintf(
			"%s-%d,%d(%d)",
			p.Name,
			//p.Kind,
			p.Pos.X,
			p.Pos.Y,
			p.Moves,
		))
	}
	sort.Strings(items)

	for k, v := range items {
		if k > 0 {
			b.WriteString(",")
		}
		b.WriteString(v)
	}

	return b.String()
}

type stateAlt struct {
	Cost   int64
	View   string
	State  string
	Id     int64
	OrigId int64
}

func (s *Board) StateAlt() stateAlt {
	var b strings.Builder

	for i := 0; i < 11; i++ {
		p := s.GetPodAtPos(Point{X: int64(i), Y: 0})
		if p == nil {
			b.WriteString(".")
			continue
		}
		b.WriteString(p.Kind)
	}
	b.WriteString(" ")

	podKinds := len(podEnergyPerMove)
	podKindDuplicates := len(s.Pods) / len(podEnergyPerMove)

	for j := 0; j < podKindDuplicates; j++ {
		for i := 0; i < podKinds; i++ {
			p := s.GetPodAtPos(Point{X: int64(i+1) * 2, Y: int64(j) + 1})
			if p == nil {
				b.WriteString(".")
				continue
			}
			b.WriteString(p.Kind)
		}
		b.WriteString(" ")
	}

	return stateAlt{
		Cost:   s.Cost(),
		View:   b.String(),
		State:  s.StateLong(),
		Id:     s.id,
		OrigId: s.origId,
	}
}

func (s *Board) maxY() int64 {
	var r int64
	for _, v := range s.BestFinishPositions {
		r = aoc.MaxInt64(r, v.Y)
	}
	return r
}

func (s *Board) Clone() *Board {
	b := NewBoard()

	b.origId = s.id

	for _, p := range s.Pods {
		v := *p
		b.Pods = append(b.Pods, &v)
	}
	s.HallwayPossiblePositions.Each(func(p Point) {
		b.HallwayPossiblePositions.Add(p)
	})
	for k, v := range s.BestFinishPositions {
		b.BestFinishPositions[k] = v
	}

	b.states = append(b.states, s.states...)

	return b
}

func (s *Board) getPodPath(p Pod, pos Point) []Point {

	var path []Point
	g := p.Pos // ghost

	// oh, those funky moves...

	if g.X != pos.X {
		for y := g.Y; y > 0; y-- {
			g.Y--
			path = append(path, g)
		}
	}

	for x := g.X; x < pos.X; x++ {
		g.X++
		path = append(path, g)
	}
	for x := g.X; x > pos.X; x-- {
		g.X--
		path = append(path, g)
	}

	if g.X == pos.X {
		for y := g.Y; y < pos.Y; y++ {
			g.Y++
			path = append(path, g)
		}
	}

	return path
}

func (s *Board) isPathBlocked(path []Point) (blocked bool, idx int) {
	for i, pos := range path {
		p := s.GetPodAtPos(pos)
		if p != nil {
			return true, i
		}
	}
	return false, -1
}

func (s *Board) MovePod(p *Pod, path []Point) {
	if len(path) == 0 {
		return
	}

	if s.Cost() == 0 {
		// store initial state
		s.states = append(s.states, s.StateAlt())
	}

	var eat int64
	for range path {
		eat += podEnergyPerMove[p.Kind]
	}

	p.Pos = path[len(path)-1]
	p.EnergyTaken += eat
	p.Moves++

	s.states = append(s.states, s.StateAlt())
}

func (s *Board) IsHallwayShortBlocked() bool {
	for _, p := range s.Pods {
		if s.HallwayPossiblePositions.Exists(p.Pos) && (p.Pos.X > 1 && p.Pos.X < 9) {
			return true
		}
	}
	return false
}

func (s *Board) GetMinCost() int64 {

	var bestSolId int64

	costs := map[string]int64{}
	ranks := map[string]int64{}
	boards := map[int64]*Board{}
	nodes := &Nodes{}

	// start state
	clone := s.Clone()
	nodes.Add(NodeWithCost{
		Board: clone,
		Cost:  math.MaxInt64,
		Id:    clone.State(),
	})
	startState := clone.State()

	var checked int
	for nodes.Len() > 0 {
		n := heap.Pop(nodes).(NodeWithCost)
		board := n.Board

		checked++
		if checked%50_000 == 0 {
			log.Println("checked:", checked)
		}

		//drawBoard(*board)

		func(s *Board) {
			var movedPods int
			for _, p := range s.Pods {
				p := p

				if p.Moves < 2 {
					// straight to finish
					fin, isTakenByWrongKind := s.GetFinishPos(*p)
					if !isTakenByWrongKind && fin != nil {
						path := s.getPodPath(*p, *fin)
						if len(path) == 0 {
							p.IsFinished = true
							movedPods++
							continue
						}
						isBlocked, _ := s.isPathBlocked(path)
						if !isBlocked {
							s.MovePod(p, path)
							p.IsFinished = true
							movedPods++
							continue
						}
					}
					if fin == nil {
						panic("cannot find finish for pod: " + p.Name)
					}
				}

				if p.Moves == 0 {
					// alt way: to hallway
					s.HallwayPossiblePositions.Each(func(pos Point) {
						path := s.getPodPath(*p, pos)
						if len(path) == 0 {
							return
						}
						isBlocked, _ := s.isPathBlocked(path)
						if isBlocked {
							return
						}

						b := s.Clone()
						bp := b.GetPodAtPos(p.Pos)
						if bp == nil {
							panic("cannot get pod in cloned board")
						}
						b.MovePod(bp, path)

						st := b.State()
						cost := b.Cost()

						boards[b.id] = b

						if r, ok := ranks[st]; !ok {
							ranks[st] = cost
							nodes.Add(NodeWithCost{
								Board: b,
								Cost:  cost,
								Id:    st,
							})
						} else if cost < r {
							ranks[st] = cost
							nodes.Update(NodeWithCost{
								Board: b,
								Cost:  cost,
								Id:    st,
							})
						}
					})
				}
			}

			if movedPods == 0 {
				// dead end
				return
			}

			st := s.State()
			cost := s.Cost()

			if s.IsAllPodsFinished() {
				if c, ok := costs[st]; (!ok || cost < c) && st != startState {
					costs[st] = cost
					bestSolId = s.id
				}
				return
			}

			if r, ok := ranks[st]; !ok {
				ranks[st] = cost
				nodes.Add(NodeWithCost{
					Board: s,
					Cost:  cost,
					Id:    st,
				})

				boards[s.id] = s
			} else if cost < r {
				ranks[st] = cost
				nodes.Update(NodeWithCost{
					Board: s,
					Cost:  cost,
					Id:    st,
				})

				boards[s.id] = s
			}
		}(board)
	}

	log.Println("total checked:", checked)

	var minCost int64 = math.MaxInt64

	b, ok := boards[bestSolId]
	if ok {
		minCost = b.Cost()
		nils := len(strconv.FormatInt(b.Cost(), 10))
		for _, v := range b.states {
			log.Printf(
				"%"+strconv.Itoa(nils)+"d %s | %s | Id:%d Orig:%d",
				v.Cost,
				v.View,
				v.State,
				v.Id,
				v.OrigId,
			)
		}
	}

	return minCost
}

func (s *Board) GetPodAtPos(pos Point) *Pod {
	for i := range s.Pods {
		p := s.Pods[i]
		if p.Pos == pos {
			return p
		}
	}
	return nil
}

// GetFinishPos returns finish pos for a pod kind and a bool,
// which is true if position is blocked by wrong kind pod.
//
// Can return nil as first out arg!
func (s *Board) GetFinishPos(pod Pod) (pos *Point, isTakenByWrongKind bool) {

	maxY := s.maxY()
	best := s.BestFinishPositions[pod.Kind]

	for i := maxY; i > 0; i-- {
		p := s.GetPodAtPos(best)
		if p == nil {
			v := best
			return &v, false
		}
		if p.Name == pod.Name {
			v := best
			return &v, false
		}
		if p.Kind == pod.Kind {
			best.Y--
			continue
		}
		// position is taken by wrong pod
		v := best
		return &v, true
	}

	v := best

	return &v, false
}

func (s *Board) Cost() int64 {
	var r int64
	for _, p := range s.Pods {
		r += p.EnergyTaken
	}
	return r
}

func (s *Board) IsAllPodsFinished() bool {

	if s.finished {
		return true
	}

	var finished int
	for _, p := range s.Pods {
		if !p.IsFinished {
			continue
		}
		finished++
	}

	if finished == len(s.Pods) {
		s.finished = true
		return true
	}

	return false
}

func (s *Board) allPossibleKindPositions(kind string) []Point {
	var positions []Point
	s.HallwayPossiblePositions.Each(func(p Point) {
		positions = append(positions, p)
	})

	for i := 0; i < len(s.Pods)/len(podEnergyPerMove); i++ {
		p := Point{X: s.BestFinishPositions[kind].X, Y: int64(i + 1)}
		positions = append(positions, p)
	}

	return positions
}

func Part1(lines []string) int64 {

	//
	// load data input
	//
	var podsStarts []*Pod

	b := NewBoard()

	indexes := map[string]int{}
	nextIndex := func(c string) int {
		if v, ok := indexes[c]; !ok {
			indexes[c] = 0
			return 0
		} else {
			v++
			indexes[c] = v
			return v
		}
	}

	// first: scan for pods
	for y := 2; y < len(lines)-1; y++ {
		for x := 0; x < len(lines[y]); x++ {
			c := lines[y][x]
			if !unicode.IsLetter(rune(c)) {
				continue
			}

			idx := nextIndex(string(c))
			podsStarts = append(podsStarts, &Pod{
				Kind: string(c),
				Name: fmt.Sprintf("%s%d", string(c), idx),
				Pos: Point{
					X: int64(x - 1),
					Y: int64(y - 1),
				},
				Index: idx,
			})

			col := finishCols[string(c)]
			b.BestFinishPositions[string(c)] = Point{X: col, Y: int64(len(lines)) - 3}
		}
	}
	b.Pods = podsStarts

	// second: scan for hallways possible positions (must exclude "doors")
	for i := 1; i < len(lines[1])-1; i++ {
		if lines[1][i] != '.' {
			continue
		}
		isRoomDoor := func(x int64) bool {
			for _, v := range b.BestFinishPositions {
				if v.X == x {
					return true
				}
			}
			return false
		}
		if isRoomDoor(int64(i - 1)) {
			continue
		}
		b.HallwayPossiblePositions.Add(Point{X: int64(i - 1), Y: int64(0)})
	}

	if b.HallwayPossiblePositions.Len() == 0 {
		panic("no hallway possible positions")
	}
	log.Println("hallway possible positions:", b.HallwayPossiblePositions.String())

	{
		log.Println("best pod positions:")
		var items []string
		for k, v := range b.BestFinishPositions {
			items = append(items, fmt.Sprintf("%s: %s", k, v))
		}
		sort.Strings(items)
		for _, v := range items {
			log.Println("-", v)
		}
	}

	log.Println("input pods:")
	for _, p := range podsStarts {
		log.Println("-", p.Name, "at", p.Pos)
	}

	//
	// process
	//
	cost := b.GetMinCost()

	return cost
}

func Part2(lines []string) int64 {
	return Part1(lines)
}

func drawBoard(b Board) {

	fmt.Printf("-[ board: %s ]-[ cost: %d ]-\n", b.State(), b.Cost())
	for i := 0; i < 11*4; i++ {
		fmt.Print("#")
	}
	fmt.Println()

	for y := 0; y < 3; y++ {
		for x := 0; x < 11; x++ {
			if y > 0 && (x < 2 || x > 8 || x == 3 || x == 5 || x == 7) {
				fmt.Print("### ")
				continue
			} else {
				p := b.GetPodAtPos(Point{X: int64(x), Y: int64(y)})
				if p == nil {
					fmt.Print("    ")
					continue
				}
				fmt.Printf("%s%d ", p.Name, p.Moves)
			}
		}
		fmt.Println()
	}

	for i := 0; i < 11*4; i++ {
		fmt.Print("#")
	}
	fmt.Println()
	fmt.Println()
}

func init() {
	// this registers solution to be able to run from ./cmd/main.go,
	// BUT you must add import like `_ "aoc/day23"` in "main.go" by hands!
	aoc.RegisterSolution(23, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
