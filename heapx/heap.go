package heap

// type CharCount struct {
// 	char    byte
// 	count   int
// 	maxIdx  int
// 	heapIdx int
// }

// // type MinHeap []*CharCount

// // func (h *MinHeap) Len() int { return len(*h) }
// // func (h *MinHeap) Less(i, j int) bool {
// // 	return (*h)[i].maxIdx < (*h)[j].maxIdx
// // }
// // func (h MinHeap) Swap(i, j int) {
// // 	h[i], h[j] = h[j], h[i]
// // 	h[i].heapIdx = i
// // 	h[j].heapIdx = j
// // }
// // func (h *MinHeap) Push(x interface{}) {
// // 	n := len(*h)
// // 	item := x.(*CharCount)
// // 	item.heapIdx = n
// // 	(*h) = append((*h), item)
// // }
// // func (h *MinHeap) Pop() interface{} {
// // 	n := len(*h)
// // 	res := (*h)[n-1]
// // 	(*h) = (*h)[:n-1]
// // 	return res
// // }

// func lengthOfLongestSubstringTwoDistinct(s string) int {
// 	var h MinHeap
// 	heap.Init(&h)
// 	// Plan: create a priority queue which puts the smallest max index
// 	// of a letter at the top. When the length of the PQ would become
// 	// greater than 2, calculate the total length from startIdx, pop
// 	// the smallest max index, update startIdx, and add the new letter
// 	// to the PQ.
// 	var startIdx int
// 	windowChars := make(map[byte])
// 	for i := range s {

// 	}

// 	return 0
// }
