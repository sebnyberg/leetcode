package p0239slidingwindowmax

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSlidingWindow(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{
			[]int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7},
			7,
			[]int{9, 9, 10, 10, 10, 10, 10, 10, 10, 9, 9, 9, 8, 8},
		},
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, -2}, 2, []int{4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSlidingWindow(tc.nums, tc.k))
		})
	}
}

// Deque solution
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	deque := make([]int, 0, k)
	res := make([]int, 0, n-k+1)
	for i, n := range nums {
		// Remove items which have left the window
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// Remove tail items which are smaller than the current number
		var j int
		for j = len(deque) - 1; ; j-- {
			if j < 0 || nums[deque[j]] > n {
				break
			}
		}
		j++
		deque = deque[:j]

		// Add current item
		deque = append(deque, i)
		if i > k-2 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

// // Heap solution
// func maxSlidingWindow(nums []int, k int) []int {
// 	n := len(nums)
// 	items := make([]*Item, n)
// 	for i, n := range nums {
// 		items[i] = &Item{
// 			val: n,
// 			pos: min(k, i),
// 		}
// 	}
// 	var h MaxHeap = make([]*Item, k)
// 	copy(h, items[:k])
// 	heap.Init(&h)
// 	res := make([]int, 0, n-k+1)
// 	res = append(res, h[0].val)
// 	for i := k; i < n; i++ {
// 		heap.Push(&h, items[i])
// 		heap.Remove(&h, items[i-k].pos)
// 		res = append(res, h[0].val)
// 	}
// 	return res
// }

// type MaxHeap []*Item
// type Item struct {
// 	val int // value
// 	pos int // position within the heap, is updated as items are shuffled around
// }

// func (h *MaxHeap) Pop() interface{} {
// 	ret := (*h)[len(*h)-1]
// 	(*h) = (*h)[:len(*h)-1]
// 	return ret
// }
// func (h *MaxHeap) Push(a interface{}) {
// 	(*h) = append((*h), a.(*Item))
// }
// func (h MaxHeap) Len() int           { return len(h) }
// func (h MaxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
// func (h MaxHeap) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// 	h[i].pos, h[j].pos = i, j
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }
