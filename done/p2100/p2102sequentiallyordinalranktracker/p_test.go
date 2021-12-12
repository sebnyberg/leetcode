package p2102sequentiallyordinalranktracker

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSORTracker(t *testing.T) {
	st := Constructor()
	st.Add("bradford", 2)
	st.Add("branford", 3)
	require.Equal(t, "branford", st.Get())
	st.Add("alps", 2)
	require.Equal(t, "alps", st.Get())
	st.Add("orland", 2)
	require.Equal(t, "bradford", st.Get())
	st.Add("orlando", 3)
	require.Equal(t, "bradford", st.Get())
	st.Add("alpine", 2)
	require.Equal(t, "bradford", st.Get())
	require.Equal(t, "orland", st.Get())
}

func nicer(a, b Location) bool {
	if a.score == b.score {
		return a.name < b.name
	}
	return a.score > b.score
}

type SORTracker struct {
	min MinHeap
	max MaxHeap
	i   int
}

func Constructor() SORTracker {
	return SORTracker{
		min: MinHeap{},
		max: MaxHeap{},
	}
}

func (this *SORTracker) Add(name string, score int) {
	l := Location{score, name}
	if len(this.min) > 0 && nicer(l, this.min[0]) {
		heap.Push(&this.max, heap.Pop(&this.min).(Location))
		heap.Push(&this.min, l)
	} else {
		heap.Push(&this.max, l)
	}
}

func (this *SORTracker) Get() string {
	heap.Push(&this.min, heap.Pop(&this.max).(Location))
	return this.min[0].name
}

type Location struct {
	score int
	name  string
}

type MaxHeap []Location

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Less(i, j int) bool {
	return nicer(h[i], h[j])
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Location))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

type MinHeap []Location

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return !nicer(h[i], h[j])
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Location))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
