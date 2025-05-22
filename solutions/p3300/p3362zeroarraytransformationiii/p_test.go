package p3362zeroarraytransformationiii

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxRemoval(t *testing.T) {
	for _, tc := range []struct {
		nums    []int
		queries [][]int
		want    int
	}{
		{[]int{0, 0, 3}, [][]int{{0, 2}, {1, 1}, {0, 0}, {0, 0}}, -1},
		{[]int{2, 0, 2}, [][]int{{0, 2}, {0, 2}, {1, 1}}, 1},
		{[]int{1, 1, 1, 1}, [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxRemoval(tc.nums, tc.queries))
		})
	}
}

func maxRemoval(nums []int, queries [][]int) int {
	n := len(nums)

	// sort by start date
	sort.Slice(queries, func(i, j int) bool {
		return queries[i][0] < queries[j][0]
	})

	var h maxHeap
	sub := make([]int, n+1)
	var delta int
	for i := range nums {
		delta -= sub[i]

		// Push new queries to the heap
		for len(queries) > 0 && queries[0][0] == i {
			// pop the first query
			query := queries[0]
			queries = queries[1:]

			// push the query to the heap
			heap.Push(&h, [2]int{query[0], query[1]})
		}

		for nums[i] > delta {
			if len(h) == 0 {
				return -1
			}
			q := heap.Pop(&h).([2]int)
			if q[1] < i {
				continue
			}
			sub[q[1]+1]++
			delta++
		}
	}
	return len(h)
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
