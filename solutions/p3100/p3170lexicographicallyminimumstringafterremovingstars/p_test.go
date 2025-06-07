package p3170lexicographicallyminimumstringafterremovingstars

import "container/heap"

func clearStars(s string) string {
	h := minHeap{s: s}
	n := len(s)
	removed := make([]bool, n)
	for i, ch := range s {
		if ch != '*' {
			heap.Push(&h, i)
			continue
		}
		removed[i] = true
		if h.Len() > 0 {
			x := heap.Pop(&h).(int)
			removed[x] = true
		}
	}
	var res []byte
	for i := range s {
		if !removed[i] {
			res = append(res, s[i])
		}
	}
	return string(res)
}

type minHeap struct {
	s   string
	idx []int
}

func (h minHeap) Len() int { return len(h.idx) }
func (h minHeap) Swap(i, j int) {
	h.idx[i], h.idx[j] = h.idx[j], h.idx[i]
}
func (h minHeap) Less(i, j int) bool {
	a := h.s[h.idx[i]]
	b := h.s[h.idx[j]]
	if a == b {
		return h.idx[i] > h.idx[j]
	}
	return h.s[h.idx[i]] < h.s[h.idx[j]]
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
