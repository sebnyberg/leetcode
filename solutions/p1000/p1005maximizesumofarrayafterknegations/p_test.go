package p1005maximizesumofarrayafterknegations

import "container/heap"

func largestSumAfterKNegations(nums []int, k int) int {
	h := make(minHeap, 0)
	for i := range nums {
		h = append(h, nums[i])
	}
	heap.Init(&h)
	var sum int
	for _, x := range nums {
		sum += x
	}
	for k > 0 {
		sum -= h[0]
		sum += -h[0]
		h[0] = -h[0]
		heap.Fix(&h, 0)
		k--
	}
	return sum
}

type minHeap []int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].heapIdx = i
	// h[j].heapIdx = j
}
func (h minHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *minHeap) Push(x interface{}) {
	el := x.(int)
	// el.heapIdx = len(*h)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	// el = nil
	*h = (*h)[:n-1]
	return el
}
