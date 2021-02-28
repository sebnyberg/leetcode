package p0895maxfreqstack

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFreqStack(t *testing.T) {
	stack := Constructor()
	// for _, n := range []int{5, 7, 5, 7, 4, 5} {
	// 	stack.Push(n)
	// }
	// for _, n := range []int{5, 7, 5, 4} {
	// 	got := stack.Pop()
	// 	require.Equal(t, n, got)
	// }

	for _, n := range []int{4, 0, 9, 3, 4, 2} {
		stack.Push(n)
	}
	got := stack.Pop()
	require.Equal(t, 4, got)

	stack.Push(6)

	got = stack.Pop()
	require.Equal(t, 6, got)
}

type FreqStack struct {
	numFreq map[int]int // frequency of a given number
	items   ItemHeap    // max-heap based on frequency and sequential index
	npushes int         // used to create stack position for items
}

type Item struct {
	val      int
	freq     int // number of identical values in the stack when this was added
	seqIndex int // global sequential index of all items added to the heap
}

type ItemHeap []Item

func (s ItemHeap) Len() int { return len(s) }
func (s ItemHeap) Less(i, j int) bool {
	return s[i].freq > s[j].freq ||
		(s[i].freq == s[j].freq && s[i].seqIndex > s[j].seqIndex)
}
func (s ItemHeap) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s *ItemHeap) Push(x interface{}) {
	*s = append(*s, x.(Item))
}
func (s *ItemHeap) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func Constructor() FreqStack {
	return FreqStack{
		numFreq: make(map[int]int),
		items:   make(ItemHeap, 0),
	}
}

func (this *FreqStack) Push(x int) {
	this.numFreq[x]++
	heap.Push(&this.items, Item{
		val:      x,
		freq:     this.numFreq[x],
		seqIndex: this.npushes,
	})
	this.npushes++
}

func (this *FreqStack) Pop() int {
	res := heap.Pop(&this.items).(Item)
	this.numFreq[res.val]--
	return res.val
}
