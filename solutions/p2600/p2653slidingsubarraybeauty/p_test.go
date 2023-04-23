package p2653slidingsubarraybeauty

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getSubarrayBeauty(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		x    int
		want []int
	}{
		{[]int{1, -1, -3, -2, 3}, 3, 2, []int{-1, -2, -2}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, getSubarrayBeauty(tc.nums, tc.k, tc.x))
		})
	}
}

func getSubarrayBeauty(nums []int, k int, x int) []int {
	// keep a max heap of size x and a min heap
	//
	// keep a max heap of negative numbers
	min := minHeap{}
	max := maxHeap{}
	n := len(nums)
	items := make([]*item, n)
	for i := range nums {
		items[i] = &item{
			val:        nums[i],
			minHeapIdx: -1,
			maxHeapIdx: -1,
		}
	}

	res := make([]int, 0, n-k+1)
	for i := range nums {
		if i >= k {
			// prune old items
			if items[i-k].maxHeapIdx != -1 {
				heap.Remove(&max, items[i-k].maxHeapIdx)
			}
			if items[i-k].minHeapIdx != -1 {
				heap.Remove(&min, items[i-k].minHeapIdx)
			}
		}
		if nums[i] < 0 {
			heap.Push(&min, items[i])
		}
		for len(max) < x && len(min) > 0 {
			// Move from min to max
			heap.Push(&max, heap.Pop(&min))
		}
		if len(max) > 0 && len(min) > 0 {
			for min[0].val < max[0].val {
				a := heap.Pop(&min)
				heap.Push(&min, heap.Pop(&max))
				heap.Push(&max, a)
			}
		}
		if i < k-1 {
			continue
		}
		if len(max) > x {
			panic("hahah")
		}
		if len(max) >= x {
			res = append(res, max[0].val)
		} else {
			res = append(res, 0)
		}
	}
	return res
}

type item struct {
	minHeapIdx int
	maxHeapIdx int
	val        int
}

type maxHeap []*item

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].maxHeapIdx = i
	h[j].maxHeapIdx = j
}
func (h maxHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}
func (h *maxHeap) Push(x interface{}) {
	el := x.(*item)
	el.maxHeapIdx = len(*h)
	*h = append(*h, el)
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	el.maxHeapIdx = -1
	*h = (*h)[:n-1]
	return el
}

type minHeap []*item

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].minHeapIdx = i
	h[j].minHeapIdx = j
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	el := x.(*item)
	el.minHeapIdx = len(*h)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	el.minHeapIdx = -1
	*h = (*h)[:n-1]
	return el
}
