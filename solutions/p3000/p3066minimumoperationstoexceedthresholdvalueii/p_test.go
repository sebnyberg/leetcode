package p3066minimumoperationstoexceedthresholdvalueii

import "container/heap"

func minOperations(nums []int, k int) int {
	var h minHeap
	for _, x := range nums {
		h = append(h, x)
	}
	heap.Init(&h)
	var res int
	for len(h) >= 2 && h[0] < k {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		heap.Push(&h, min(x, y)*2+max(x, y))
		res++
	}
	return res
}

type minHeap []int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
