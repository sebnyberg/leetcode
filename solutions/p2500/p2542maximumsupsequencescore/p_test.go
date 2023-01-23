package p2542maximumsupsequencescore

import (
	"container/heap"
	"sort"
)

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	// Sort nums1 by min value in nums2, then keep a running maximum sum of k
	// elements from nums2 while iterating over the values.
	//
	n := len(nums1)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] > nums2[idx[j]]
	})
	var sum int64
	var h minHeap
	var res int64
	for i, j := range idx {
		sum += int64(nums1[j])
		heap.Push(&h, nums1[j])
		if i >= k {
			sum -= int64(heap.Pop(&h).(int))
		}
		if i >= k-1 {
			res = max(res, int64(nums2[j])*sum)
		}
	}

	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

type minHeap []int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *minHeap) Push(x interface{}) {
	el := x.(int)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
