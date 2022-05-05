package p1845seatreservationmanager

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSeatManager(t *testing.T) {
	sm := Constructor(5)
	res := sm.Reserve()
	require.Equal(t, 1, res)
	res = sm.Reserve()
	require.Equal(t, 2, res)
	sm.Unreserve(2)
	res = sm.Reserve()
	require.Equal(t, 2, res)
	res = sm.Reserve()
	require.Equal(t, 3, res)
	res = sm.Reserve()
	require.Equal(t, 4, res)
	res = sm.Reserve()
	require.Equal(t, 5, res)
	sm.Unreserve(5)
}

type SeatManager struct {
	unreserved MinHeap
	pos        int
	max        int
}

func Constructor(n int) SeatManager {
	return SeatManager{
		unreserved: make(MinHeap, 0),
		pos:        1,
		max:        n,
	}
}

func (this *SeatManager) Reserve() int {
	if len(this.unreserved) > 0 {
		return heap.Pop(&this.unreserved).(int)
	}
	this.pos++
	return this.pos - 1
}

func (this *SeatManager) Unreserve(seatNumber int) {
	heap.Push(&this.unreserved, seatNumber)
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
