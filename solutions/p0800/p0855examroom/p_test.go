package p0855examroom

import (
	"container/heap"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExamRoom(t *testing.T) {
	// e := Constructor(10)
	// require.Equal(t, 0, e.Seat())
	// require.Equal(t, 9, e.Seat())
	// require.Equal(t, 4, e.Seat())
	// require.Equal(t, 2, e.Seat())
	// e.Leave(4)
	// require.Equal(t, 5, e.Seat())

	// e := Constructor(10)
	// require.Equal(t, 0, e.Seat())
	// require.Equal(t, 9, e.Seat())
	// require.Equal(t, 4, e.Seat())
	// e.Leave(0)
	// e.Leave(4)

	e := Constructor(4)
	require.Equal(t, 0, e.Seat())
	require.Equal(t, 3, e.Seat())
	require.Equal(t, 1, e.Seat())
	require.Equal(t, 2, e.Seat())
	e.Leave(1)
	e.Leave(3)
	require.Equal(t, 1, e.Seat())
}

// The idea behind this solution is to keep a linked-list containing both
// candidates and seated people.
// Whenever someone sits down at a candidate position, it creates new candidates
// on both sides of the seat (if possible)
// When someone leaves a seat, then that seat is removed along with any
// neighbouring candidates, and a new candidate is formed between the two
// occupied neighbouring seats on both sides.
type ExamRoom struct {
	n          int
	m          int
	seats      map[int]*seat
	candidates candidateHeap
}

func Constructor(n int) ExamRoom {
	var r ExamRoom
	r.n = n
	r.m = 0
	r.seats = make(map[int]*seat)
	s1 := &seat{
		pos:      math.MinInt32,
		occupied: true,
	}
	s2 := &seat{
		pos:      2*n + 1,
		occupied: true,
	}
	s1.right = s2
	s2.left = s1
	r.seats[s1.pos] = s1
	r.seats[s2.pos] = s2
	r.addCandidate(s1.pos, s2.pos)
	return r
}

type seat struct {
	left, right *seat
	pos         int
	heapIdx     int
	occupied    bool
}

func (s seat) dist() int {
	return min(s.pos-s.left.pos, s.right.pos-s.pos)
}

func (this *ExamRoom) addCandidate(i, j int) {
	mid := i + (j-i)/2
	mid = max(0, mid)
	mid = min(this.n-1, mid)
	if _, exists := this.seats[mid]; exists {
		return
	}
	this.seats[mid] = &seat{
		left:     this.seats[i],
		right:    this.seats[j],
		pos:      mid,
		occupied: false,
	}
	this.seats[i].right = this.seats[mid]
	this.seats[j].left = this.seats[mid]
	heap.Push(&this.candidates, this.seats[mid])
}

func (this *ExamRoom) Seat() int {
	if this.m == this.n {
		return -1
	}

	// Pick the optimal candidate
	cand := heap.Pop(&this.candidates).(*seat)
	cand.occupied = true

	// Create new candidates on both sides (if possible)
	this.addCandidate(cand.left.pos, cand.pos)
	this.addCandidate(cand.pos, cand.right.pos)

	this.m++
	return cand.pos
}

func (this *ExamRoom) Leave(p int) {
	// Whena a person leaves, then the previous and next candidate should be
	// removed (if any such candidates exist).
	s := this.seats[p]
	prev := s.left
	if !prev.occupied {
		prevPrev := prev.left
		prevPrev.right = s
		heap.Remove(&this.candidates, prev.heapIdx)
		delete(this.seats, prev.pos)
		prev = prevPrev
	}
	next := s.right
	if !next.occupied {
		nextNext := next.right
		heap.Remove(&this.candidates, next.heapIdx)
		delete(this.seats, next.pos)
		next = nextNext
	}
	prev.right = next
	next.left = prev
	delete(this.seats, s.pos)
	this.addCandidate(prev.pos, next.pos)
	this.m--
}

type candidateHeap []*seat

func (h candidateHeap) Len() int { return len(h) }
func (h candidateHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h candidateHeap) Less(i, j int) bool {
	d1, d2 := h[i].dist(), h[j].dist()
	if d1 == d2 {
		return h[i].pos < h[j].pos
	}
	return d1 > d2
}
func (h *candidateHeap) Push(x interface{}) {
	a := x.(*seat)
	a.heapIdx = len(*h)
	*h = append(*h, a)
}
func (h *candidateHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
