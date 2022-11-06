package p2462totalcosttohirekworkers

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_totalCost(t *testing.T) {
	for i, tc := range []struct {
		costs      []int
		k          int
		candidates int
		want       int64
	}{
		{[]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2, 423},
		{[]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4, 11},
		{[]int{1, 2, 4, 1}, 3, 3, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, totalCost(tc.costs, tc.k, tc.candidates))
		})
	}
}

func totalCost(costs []int, k int, candidates int) int64 {
	// Calculate total cost of hiring strictly from the left, then from the
	// right, and combine to form the best result

	n := len(costs)
	h := make(minHeap, 0, candidates)
	var l int
	r := n - 1
	for l < candidates {
		h = append(h, candidate{true, costs[l]})
		l++
	}
	for n-1-r < candidates && l <= r {
		h = append(h, candidate{false, costs[r]})
		r--
	}
	heap.Init(&h)
	var res int64
	for l <= r && k > 0 {
		x := heap.Pop(&h).(candidate)
		res += int64(x.val)
		if x.left {
			heap.Push(&h, candidate{true, costs[l]})
			l++
		} else {
			heap.Push(&h, candidate{false, costs[r]})
			r--
		}
		k--
	}
	for k > 0 {
		res += int64(heap.Pop(&h).(candidate).val)
		k--
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type candidate struct {
	left bool
	val  int
}

type minHeap []candidate

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	if h[i].val == h[j].val {
		return h[i].left
	}
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	el := x.(candidate)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
