package p0373findkpairswithsmallestsums

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kSmallestPairs(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{[]int{-15, 20, 45, 117, 223}, []int{-15, 20, 45, 117, 223, 546, 663, 714, 749, 801}, 2,
			[][]int{{-15, -15}, {20, -15}},
		},
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}},
		{[]int{1, 2}, []int{3}, 3, [][]int{{1, 3}, {2, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			res := kSmallestPairs(tc.nums1, tc.nums2, tc.k)
			cur := math.MinInt32
			for _, pair := range res {
				sum := pair[0] + pair[1]
				require.GreaterOrEqual(t, sum, cur)
				cur = sum
			}
		})
	}
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// Create min-heap containing pair indices and their sum.
	// Keep the first index of each pair fixed and increase the second index as
	// pairs are popped from the heap. Once the second index goes out of bounds,
	// completely remove the pair from the heap.
	pairs := make(pairHeap, len(nums1))
	for i := range nums1 {
		pairs[i] = &pair{i, 0, nums1[i] + nums2[0]}
	}
	heap.Init(&pairs) // O(n)

	res := make([][]int, 0, k)
	for ; len(pairs) > 0 && k > 0; k-- {
		p := pairs[0] // *pair => safe to perform immediate updates
		res = append(res, []int{nums1[p.i], nums2[p.j]})
		if p.j+1 == len(nums2) { // no more numbers to pair with in nums2
			heap.Pop(&pairs) // O(n*logn)
		} else {
			p.j++
			p.sum = nums1[p.i] + nums2[p.j]
			heap.Fix(&pairs, 0) // O(n*logn)
		}
	}
	return res
}

type pair struct {
	i   int // i is from nums1 and may not be changed
	j   int // j is from nums2 and is changed over time
	sum int
}

type pairHeap []*pair

func (h pairHeap) Len() int            { return len(h) }
func (h pairHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h pairHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(*pair)) }
func (h *pairHeap) Pop() interface{} {
	n := len(*h)
	p := (*h)[n-1]
	*h = (*h)[:n-1]
	return p
}
