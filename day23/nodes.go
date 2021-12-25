package day23

import (
	"container/heap"
)

type NodeWithCost struct {
	Board *Board
	Cost  int64
	Id    string
}

type Nodes struct {
	storage []NodeWithCost
}

func (s Nodes) Len() int {
	return len(s.storage)
}

func (s Nodes) Less(i, j int) bool {
	return s.storage[i].Cost < s.storage[j].Cost
}

func (s Nodes) Swap(i, j int) {
	s.storage[i], s.storage[j] = s.storage[j], s.storage[i]
}

// Push pushes a NodeWithCost.
func (s *Nodes) Push(x interface{}) {
	s.storage = append(s.storage, x.(NodeWithCost))
}

// Pop pops last item and returns as NodeWithCost.
func (s *Nodes) Pop() interface{} {

	if len(s.storage) == 0 {
		return nil
	}
	old := s.storage
	l := len(old)
	x := old[l-1]
	s.storage = old[0 : l-1]

	return x
}

func (s *Nodes) Add(n NodeWithCost) {
	heap.Push(s, n)
	s.Update(n)
}

func (s *Nodes) AddMax(n NodeWithCost) {
	s.storage = append(s.storage, n)
}

func (s *Nodes) Update(n NodeWithCost) {

	for i, v := range s.storage {
		if v.Id == n.Id {
			s.storage[i] = n
			heap.Fix(s, i)
			return
		}
	}
}
