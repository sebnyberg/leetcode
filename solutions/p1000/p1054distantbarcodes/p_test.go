package p1054distantbarcodes

import "container/heap"

func rearrangeBarcodes(barcodes []int) []int {
	count := make(map[int]int)
	for _, b := range barcodes {
		count[b]++
	}
	counts := make(maxHeap, 0, len(count))
	for x, c := range count {
		counts = append(counts, [2]int{x, c})
	}
	heap.Init(&counts)
	var res []int
	for len(counts) > 1 {
		x := heap.Pop(&counts).([2]int)
		y := heap.Pop(&counts).([2]int)
		if len(res) > 0 && res[len(res)-1] == x[0] {
			res = append(res, y[0])
			y[1]--
		} else {
			res = append(res, x[0])
			x[1]--
		}
		if x[1] != 0 {
			heap.Push(&counts, x)
		}
		if y[1] != 0 {
			heap.Push(&counts, y)
		}
	}
	res = append(res, heap.Pop(&counts).([2]int)[0])
	return res
}

type maxHeap [][2]int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h maxHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}
func (h *maxHeap) Push(x interface{}) {
	el := x.([2]int)
	*h = append(*h, el)
}
func (h *maxHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	*h = (*h)[:n-1]
	return el
}
