package p2905findindiceswithindexandvaluedifferenceii

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findIndices(t *testing.T) {
	for i, tc := range []struct {
		nums            []int
		indexDifference int
		valueDifference int
		want            []int
	}{
		{[]int{5, 1, 4, 1}, 2, 4, []int{0, 1}},
		{[]int{2, 1}, 0, 0, []int{0, 0}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, findIndices(tc.nums, tc.indexDifference, tc.valueDifference))
		})
	}
}

func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	minh := minHeap{
		nums: nums,
	}
	maxh := maxHeap{
		nums: nums,
	}

	for i := indexDifference; i < len(nums); i++ {
		heap.Push(&minh, i-indexDifference)
		heap.Push(&maxh, i-indexDifference)
		max := maxh.idx[0]
		min := minh.idx[0]
		if abs(nums[max]-nums[i]) >= valueDifference {
			res := []int{i, max}
			sort.Ints(res)
			return res
		}
		if abs(nums[min]-nums[i]) >= valueDifference {
			res := []int{i, min}
			sort.Ints(res)
			return res
		}
	}
	return []int{-1, -1}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type minHeap struct {
	nums []int
	idx  []int
}

func (h minHeap) Len() int { return len(h.idx) }
func (h minHeap) Swap(i, j int) {
	h.idx[i], h.idx[j] = h.idx[j], h.idx[i]
}
func (h minHeap) Less(i, j int) bool {
	return h.nums[h.idx[i]] < h.nums[h.idx[j]]
}
func (h *minHeap) Push(x interface{}) {
	el := x.(int)
	h.idx = append(h.idx, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(h.idx)
	el := (h.idx)[n-1]
	h.idx = h.idx[:n-1]
	return el
}

type maxHeap struct {
	nums []int
	idx  []int
}

func (h maxHeap) Len() int { return len(h.idx) }
func (h maxHeap) Swap(i, j int) {
	h.idx[i], h.idx[j] = h.idx[j], h.idx[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h.nums[h.idx[i]] > h.nums[h.idx[j]]
}
func (h *maxHeap) Push(x interface{}) {
	el := x.(int)
	h.idx = append(h.idx, el)
}
func (h *maxHeap) Pop() interface{} {
	n := len(h.idx)
	el := (h.idx)[n-1]
	h.idx = h.idx[:n-1]
	return el
}
