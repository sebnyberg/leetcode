package p0632smallestrangecoveringelementsfromklists

import (
	"container/heap"
	"fmt"
	"github.com/sebnyberg/leetcode"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestRange(t *testing.T) {
	for _, tc := range []struct {
		nums [][]int
		want []int
	}{
		{leetcode.ParseMatrix("[[10,10],[11,11]]"), []int{10, 11}},
		{leetcode.ParseMatrix("[[4,10,15,24,26],[0,9,12,20],[5,18,22,30]]"), []int{20, 24}},
		{leetcode.ParseMatrix("[[1,2,3],[1,2,3],[1,2,3]]"), []int{1, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, smallestRange(tc.nums))
		})
	}
}

func smallestRange(nums [][]int) []int {
	// h will store the minimum time from the input nums
	h := make(MinHeap, 0, len(nums))
	for i := range nums {
		for j := range nums[i] {
			h = append(h, item{i: i, t: nums[i][j]})
		}
	}
	heap.Init(&h)
	// r stores the minimum times taken from the input nums (current range)
	r := make(MinHeap, 0)
	listCount := make([]int, len(nums))
	var count int
	var maxVal int
	res := []int{0, 0}
	minDiff := math.MaxInt32

	for len(h) > 0 {
		// Pick items from h until the range covers all lists
		for len(h) > 0 {
			x := heap.Pop(&h).(item)
			listCount[x.i]++
			heap.Push(&r, x)
			maxVal = max(maxVal, x.t)
			if listCount[x.i] == 1 {
				count++
			}
			if count == len(nums) {
				break
			}
		}
		if count < len(nums) {
			break
		}
		// Remove from range until no longer covering all lists
		for len(r) > 0 && listCount[r[0].i] > 1 {
			listCount[r[0].i]--
			heap.Pop(&r)
		}
		if d := maxVal - r[0].t; d < minDiff {
			minDiff = d
			res[0] = r[0].t
			res[1] = maxVal
		}
		listCount[r[0].i]--
		heap.Pop(&r)
		count--
	}
	return res
}

type item struct {
	t int
	i int
}

type MinHeap []item

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].t < h[j].t
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
