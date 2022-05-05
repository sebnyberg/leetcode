package p0215kthlargestinarr

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKthLargest(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{2, 1}, 2, 1},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.nums, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, findKthLargest(tc.nums, tc.k))
		})
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums[:k])
	heap.Init(&h)

	for i := k; i < len(nums); i++ {
		n := nums[i]
		if n > h[0] {
			heap.Pop(&h)
			heap.Push(&h, n)
		}
	}

	return h[0]
}
