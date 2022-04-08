package p0703kthlargestinastream

import "container/heap"

type KthLargest struct {
	k int
	h MaxHeap
}

func Constructor(k int, nums []int) KthLargest {
	h := make(MaxHeap, 0, k)
	for _, n := range nums {
		h = append(h, n)
	}
	heap.Init(&h)
	for len(h) > k {
		heap.Pop(&h)
	}
	return KthLargest{h: h, k: k}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(&this.h, val)
	for len(this.h) > this.k {
		heap.Pop(&this.h)
	}
	return this.h[0]
}

type MaxHeap []int

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
