package p0895maxfreqstackv2

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFreqStack(t *testing.T) {
	t.Run("first test", func(t *testing.T) {
		stack := Constructor()
		for _, n := range []int{5, 7, 5, 7, 4, 5} {
			stack.Push(n)
		}
		for _, n := range []int{5, 7, 5, 4} {
			got := stack.Pop()
			require.Equal(t, n, got)
		}
	})

	t.Run("second", func(t *testing.T) {
		stack := Constructor()
		for _, n := range []int{4, 0, 9, 3, 4, 2} {
			stack.Push(n)
		}
		got := stack.Pop()
		require.Equal(t, 4, got)

		stack.Push(6)

		got = stack.Pop()
		require.Equal(t, 6, got)

		stack.Push(1)

		got = stack.Pop()
		require.Equal(t, 1, got)

		stack.Push(1)

		got = stack.Pop()
		require.Equal(t, 1, got)

		stack.Push(4)

		for _, n := range []int{4, 2, 3, 9, 0, 4} {
			got = stack.Pop()
			require.Equal(t, n, got)
		}
	})

	t.Run("third", func(t *testing.T) {
		stack := Constructor()
		for _, n := range []int{1, 0, 0, 1, 5, 4, 1, 5, 1, 6} {
			stack.Push(n)
		}
		for _, n := range []int{1, 1, 5, 1, 0, 6, 4, 5, 0, 1} {
			got := stack.Pop()
			require.Equal(t, n, got)
		}
	})
}

type FreqStack struct {
	h       ItemHeap      // heap of numbers, their frequency, and stack position
	npushes int           // used to create stack position for items
	items   map[int]*Item // map of items
}

type Item struct {
	val            int
	freq           int
	index          int
	stackPositions []int
}

type ItemHeap []*Item

func (s ItemHeap) Len() int { return len(s) }
func (s ItemHeap) Less(i, j int) bool {
	return s[i].freq > s[j].freq || (s[i].freq == s[j].freq && s[i].stackPositions[s[i].freq-1] > s[j].stackPositions[s[j].freq-1])
}
func (s ItemHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	s[i].index = i
	s[j].index = j
}
func (s *ItemHeap) Push(x interface{}) {
	n := len(*s)
	item := x.(*Item)
	item.index = n
	*s = append(*s, item)
}
func (s *ItemHeap) Pop() interface{} {
	old := *s
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*s = old[0 : n-1]
	return item
}

func Constructor() FreqStack {
	return FreqStack{
		h:     make(ItemHeap, 0),
		items: make(map[int]*Item),
	}
}

func (this *FreqStack) Push(x int) {
	if _, exists := this.items[x]; !exists {
		this.items[x] = &Item{
			val:            x,
			freq:           1,
			stackPositions: []int{this.npushes},
		}
		heap.Push(&this.h, this.items[x])
	} else {
		this.items[x].freq++
		this.items[x].stackPositions = append(this.items[x].stackPositions, this.npushes)
		heap.Fix(&this.h, this.items[x].index)
	}
	this.npushes++
}

func (this *FreqStack) Pop() int {
	topItem := this.h[0]

	if topItem.freq > 1 {
		// No need to pop before frequency is zero
		topItem.stackPositions = topItem.stackPositions[:len(topItem.stackPositions)-1]
		topItem.freq--
		heap.Fix(&this.h, topItem.index)
	} else {
		// Pop the item once its frequency has reached zero
		heap.Pop(&this.h)
		delete(this.items, topItem.val)
	}

	return topItem.val
}
