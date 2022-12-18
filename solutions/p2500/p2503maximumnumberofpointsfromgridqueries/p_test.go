package p2503maximumnumberofpointsfromgridqueries

import (
	"container/heap"
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxPoints(t *testing.T) {
	for i, tc := range []struct {
		grid    [][]int
		queries []int
		want    []int
	}{
		{
			leetcode.ParseMatrix("[[1,2,3],[2,5,7],[3,5,1]]"),
			[]int{5, 6, 2},
			[]int{5, 8, 1},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxPoints(tc.grid, tc.queries))
		})
	}
}

func maxPoints(grid [][]int, queries []int) []int {
	// We can calculate how many cells can be visited for each value using a
	// min-heap of available cell values.
	//
	// Then we can binary search over the array to answer the queries.
	//
	h := minHeap{cell{0, 0, grid[0][0]}}
	dirs := [][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	m := len(grid)
	n := len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	seen := make([][]bool, m)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[0][0] = true

	var i int
	counts := []int{0}
	vals := []int{0}
	for len(h) > 0 {
		x := heap.Pop(&h).(cell)
		if x.val > vals[i] {
			vals = append(vals, x.val)
			counts = append(counts, counts[i]+1)
			i++
		} else {
			counts[i]++
		}
		for _, d := range dirs {
			ii := x.i + d[0]
			jj := x.j + d[1]
			if !ok(ii, jj) || seen[ii][jj] {
				continue
			}
			seen[ii][jj] = true
			heap.Push(&h, cell{ii, jj, grid[ii][jj]})
		}
	}

	// Now there's a sorted list of values and counts such that given a value
	// v > values[i], there are count[i] values reachable in the grid.
	// So we can binary search to find the first value that is smaller than the
	// query, and add that to the result.
	k := len(queries)
	res := make([]int, k)
	for i := range queries {
		j := sort.Search(len(vals), func(jj int) bool {
			return vals[jj] >= queries[i]
		})
		j -= 1
		// if j == len(vals) {
		// 	res[i] = m * n
		// } else {
		res[i] = counts[j]
		// }
	}
	return res
}

type cell struct {
	i, j, val int
}

type minHeap []cell

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].heapIdx = i
	// h[j].heapIdx = j
}
func (h minHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *minHeap) Push(x interface{}) {
	el := x.(cell)
	// el.heapIdx = len(*h)
	*h = append(*h, el)
}
func (h *minHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	// el = nil
	*h = (*h)[:n-1]
	return el
}
