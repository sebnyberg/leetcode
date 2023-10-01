package p2762continuoussubarrays

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_continuousSubarrays(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int64
	}{
		{[]int{5, 4, 2, 4}, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, continuousSubarrays(tc.nums))
		})
	}
}

func continuousSubarrays(nums []int) int64 {
	// Consider subarrays ending at a certain index.
	//
	// When would that subarray end sequences that used to be valid? When the
	// value changes the current range's max or min in such a way that max-min
	// is now > 2
	//
	// To make matters simple: lets add the current number to a min and max
	// heap. While the difference between the min and max heap is > 2, pop the
	// eldest element in the current range.
	min := minHeap{}
	max := maxHeap{}
	n := len(nums)
	items := make([]*item, n)
	for i := range nums {
		items[i] = &item{0, 0, nums[i]}
	}
	var l int
	var res int
	for i, x := range items {
		heap.Push(&min, x)
		heap.Push(&max, x)
		for max[0].val-min[0].val > 2 {
			heap.Remove(&min, items[l].minHeapIdx)
			heap.Remove(&max, items[l].maxHeapIdx)
			l++
		}
		res += i - l + 1
	}
	return int64(res)
}

type item struct {
	minHeapIdx int
	maxHeapIdx int
	val        int
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
	*h = (*h)[:n-1]
	return el
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
	*h = (*h)[:n-1]
	return el
}
