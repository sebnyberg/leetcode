package p2386findtheksumofanarray

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kSum(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{2, 4, -2}, 5, 2},
		{[]int{1, -2, 3, 4, -10, 12}, 16, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, kSum(tc.nums, tc.k))
		})
	}
}

func kSum(nums []int, k int) int64 {
	// The largest value is maxSum
	//
	// The second largest value is:
	// 1. maxSum - the smallest positive value, OR
	// 2. maxsum + the largest negative value.
	//
	// Both of these are equivalent to doing maxSum - smallest absolute value
	//
	// The next largest value is going to be the result of removing either the
	// second smallest absolute value, which is either the second, or a
	// combination of the first and second element.
	//
	// The approach is therefore to find the k'th smallest subsequence sum, and
	// remove it from the maximum sum.
	//
	// To find the k'th smallest subsequence sum, we need to find a way to iterate
	// over alternatives in an efficient way.
	//
	// Given that nums is sorted, we know that the first element is smallest.
	// Then it's either the second, or the first + second.
	// This gives us an interesting recursion. Given a current sum and
	// next-to-include index i, let's say 18 and i=7, then we can either include
	// the current index and add the next, or skip the current index and include
	// the next. This way we enumerate all possible sums.
	//
	var maxSum int
	for i, x := range nums {
		if x > 0 {
			maxSum += x
		} else {
			nums[i] = -x
		}
	}
	sort.Ints(nums)
	h := minHeap{item{nums[0], 0}}
	res := maxSum
	for i := 1; i < k; i++ {
		x := heap.Pop(&h).(item)
		res = maxSum - x.sum
		if x.idx+1 < len(nums) {
			heap.Push(&h, item{x.sum - nums[x.idx] + nums[x.idx+1], x.idx + 1})
			heap.Push(&h, item{x.sum + nums[x.idx+1], x.idx + 1})
		}
	}
	return int64(res)
}

type item struct {
	sum int
	idx int
}

type minHeap []item

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h minHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(item))
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
