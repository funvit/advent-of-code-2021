package day21

import (
	"log"
	"strings"

	"aoc"
)

/*
	Note: output arg type of Part1 and Part2 can be freely changed.
*/

const (
	maxBoardPos        = 10
	diceRollsPerPlayer = 3
)

func Part1(lines []string) int {

	const maxScore = 1000
	const maxDiceSides = 100

	line1 := lines[0]
	line2 := lines[1]

	p1 := Player{
		Id:  1,
		Pos: int(aoc.MustParseInt64(line1[strings.LastIndex(line1, " ")+1:])),
	}
	p2 := Player{
		Id:  2,
		Pos: int(aoc.MustParseInt64(line2[strings.LastIndex(line2, " ")+1:])),
	}

	log.Println("player 1 pos:", p1.Pos)
	log.Println("player 2 pos:", p2.Pos)

	g := Game{
		Players:  []*Player{&p1, &p2},
		Dice:     NewDice(maxDiceSides),
		MaxScore: maxScore,
	}

	return g.PlayDeterministic()
}

func Part2(lines []string) int {

	const maxScore = 21
	const maxDiceSides = 3

	line1 := lines[0]
	line2 := lines[1]

	p1 := Player{
		Id:  1,
		Pos: int(aoc.MustParseInt64(line1[strings.LastIndex(line1, " ")+1:])),
	}
	p2 := Player{
		Id:  2,
		Pos: int(aoc.MustParseInt64(line2[strings.LastIndex(line2, " ")+1:])),
	}

	log.Println("player 1 pos:", p1.Pos)
	log.Println("player 2 pos:", p2.Pos)

	g := Game{
		Players:  []*Player{&p1, &p2},
		Dice:     NewDice(maxDiceSides),
		MaxScore: maxScore,
	}

	w := g.PlayQuantum(0)

	return aoc.MaxInt(w[0], w[1])
}

type Game struct {
	Players  []*Player
	Dice     *Dice
	MaxScore int
	winner   *Player
}

func (s *Game) Clone() *Game {
	g := &Game{
		Dice:     s.Dice,
		MaxScore: s.MaxScore,
	}

	for _, p := range s.Players {
		g.Players = append(g.Players, &Player{
			Id:    p.Id,
			Pos:   p.Pos,
			Score: p.Score,
		})
	}

	return g
}

func (s *Game) isPlayerWon(p *Player) bool {
	if s.MaxScore == 0 {
		panic("game max score is not set")
	}
	return p.Score >= s.MaxScore
}

func (s *Game) PlayDeterministic() int {

	for {
		for pIdx, p := range s.Players {
			var rollSum int
			for i := 0; i < diceRollsPerPlayer; i++ {
				rollSum += s.Dice.Next()
			}
			p.MoveToPos(GetNextBoardPos(p.Pos, rollSum))

			if s.isPlayerWon(p) {
				s.winner = p
				return s.Players[(pIdx+1)%2].Score * s.Dice.Rolls()
			}
		}
	}
}

// PlayQuantum returns total won universes for two players.
func (s *Game) PlayQuantum(playerIdx int) [2]int {

	m := [2]int{}

	for r, u := range allPossibleUniverses {
		g := s.Clone()
		p := g.Players[playerIdx]
		p.MoveToPos(GetNextBoardPos(p.Pos, r))

		if s.isPlayerWon(p) {
			m[playerIdx] += u
			continue
		}

		nextPlayerIdx := 1 - playerIdx // 0,1,0...

		w := g.playQuantumCached(nextPlayerIdx)
		//w := g.PlayQuantum(nextPlayerIdx)
		m[0] += u * w[0]
		m[1] += u * w[1]
	}

	return m
}

type QuantumCache map[quantumCacheKey][2]int // val is a won universes per player idx
type quantumCacheKey struct {
	currentPlayerIdx             int
	playerOnePos, playerOneScore int
	playerTwoPos, playerTwoScore int
}

var quantumCache = QuantumCache{}

func (s *Game) playQuantumCached(playerIdx int) [2]int {

	k := quantumCacheKey{
		currentPlayerIdx: playerIdx,
		playerOnePos:     s.Players[0].Pos,
		playerOneScore:   s.Players[0].Score,
		playerTwoPos:     s.Players[1].Pos,
		playerTwoScore:   s.Players[1].Score,
	}
	w, ok := quantumCache[k]
	if ok {
		return w
	}

	w = s.PlayQuantum(playerIdx)
	quantumCache[k] = w

	return w
}

type Player struct {
	Id    int
	Pos   int
	Score int
}

func (s *Player) MoveToPos(pos int) {
	s.Pos = pos
	s.Score += s.Pos
}

type Dice struct {
	n     int
	max   int
	rolls int
}

func NewDice(maxSides int) *Dice {
	return &Dice{
		n:   0,
		max: maxSides,
	}
}

func (s *Dice) Rolls() int {
	return s.rolls
}

func (s *Dice) Val() int {
	return s.n
}

func (s *Dice) Next() int {
	s.n++
	if s.n > s.max {
		s.n = 1
	}
	s.rolls++
	return s.n
}

func GetNextBoardPos(current int, step int) int {
	n := (current + step) % maxBoardPos
	if n == 0 {
		n = maxBoardPos
	}
	return n
}

// roll=>universes
var allPossibleUniverses map[int]int

func init() {
	allPossibleUniverses = make(map[int]int, 27)
	for a := 0; a < 3; a++ {
		for b := 0; b < 3; b++ {
			for c := 0; c < 3; c++ {
				allPossibleUniverses[a+b+c+3]++
			}
		}
	}
}

func init() {
	// this registers solution to be able to run from ./cmd/main.go,
	// BUT you must add import like `_ "aoc/day21"` in "main.go" by hands!
	aoc.RegisterSolution(21, aoc.DaySolution{
		Part1: func(lines []string) interface{} {
			return Part1(lines)
		},
		Part2: func(lines []string) interface{} {
			return Part2(lines)
		},
	})
}
