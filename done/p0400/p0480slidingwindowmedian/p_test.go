package p0480slidingwindowmedian

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_medianSlidingWindow(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []float64
	}{
		{[]int{7, 0, 3, 9, 9, 9, 1, 7, 2, 3}, 6, []float64{8.00000, 6.00000, 8.00000, 8.00000, 5.00000}},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1.00000, -1.00000, -1.00000, 3.00000, 5.00000, 6.00000}},
		{[]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3, []float64{2.00000, 3.00000, 3.00000, 3.00000, 2.00000, 3.00000, 2.00000}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, medianSlidingWindow(tc.nums, tc.k))
		})
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	below := make(maxHeap, 0)
	above := make(minHeap, 0)
	items := make([]*item, len(nums))
	for i, num := range nums {
		items[i] = &item{val: num, heapIdx: 0}
	}

	res := make([]float64, 0, len(nums)-k)
	for i, it := range items {
		heap.Push(&above, it)
		if i >= k {
			// Remove item
			x := items[i-k]
			if x.below {
				heap.Remove(&below, x.heapIdx)
			} else {
				heap.Remove(&above, x.heapIdx)
			}
		}
		// Fix errors in heaps
		for len(below) > 0 && len(above) > 0 && below[0].val > above[0].val {
			left := heap.Pop(&below).(*item)
			left.below = false
			right := heap.Pop(&above).(*item)
			right.below = true
			heap.Push(&below, right)
			heap.Push(&above, left)
		}
		if len(above)-len(below) > k%2 {
			// Move from above to below
			x := heap.Pop(&above).(*item)
			x.below = true
			heap.Push(&below, x)
		}

		if i >= k-1 {
			if k%2 == 0 {
				res = append(res, (float64(below[0].val)+float64(above[0].val))/2)
			} else {
				res = append(res, float64(above[0].val))
			}
		}
	}
	return res
}

type item struct {
	val     int
	heapIdx int
	below   bool
}

type maxHeap []*item

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Swap(i, j int) {
	h[i].heapIdx = j
	h[j].heapIdx = i
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}
func (h *maxHeap) Push(x interface{}) {
	it := x.(*item)
	it.heapIdx = len(*h)
	*h = append(*h, it)
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

type minHeap []*item

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i].heapIdx = j
	h[j].heapIdx = i
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	it := x.(*item)
	it.heapIdx = len(*h)
	*h = append(*h, it)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
