package p2593findscoreofanarrayaftermarkingallelements

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findScore(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{29, 20, 36, 39, 50, 42, 46, 34, 47}, 135},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, findScore(tc.nums))
		})
	}
}

func findScore(nums []int) int64 {
	// Use a PQ to keep track of unmarked, low-value (and low-index) elements.
	// Poll until the prioritized element is already marked. Then all elements
	// must be marked.
	n := len(nums)
	h := make(minHeap, n)
	items := make([]*item, n)
	for i, x := range nums {
		y := &item{
			idx:     i,
			val:     x,
			marked:  false,
			heapIdx: i,
		}
		h[i] = y
		items[i] = y
	}
	heap.Init(&h)
	var score int
	for !h[0].marked {
		h[0].marked = true
		a := h[0]
		score += h[0].val
		heap.Fix(&h, 0)
		l := max(0, a.idx-1)
		items[l].marked = true
		heap.Fix(&h, items[l].heapIdx)
		r := min(n-1, a.idx+1)
		items[r].marked = true
		heap.Fix(&h, items[r].heapIdx)
	}
	return int64(score)
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

type item struct {
	idx     int
	val     int
	heapIdx int
	marked  bool
}

type minHeap []*item

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIdx = i
	h[j].heapIdx = j
}
func (h minHeap) Less(i, j int) bool {
	if h[i].marked != h[j].marked {
		return !h[i].marked
	}
	if h[i].val == h[j].val {
		return h[i].idx < h[j].idx
	}
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	el := x.(*item)
	el.heapIdx = len(*h)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	// el = nil
	*h = (*h)[:n-1]
	return el
}
