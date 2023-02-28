package p2577minimumtimetovisitacellinagrid

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumTime(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,1,3,2],[5,1,2,5],[4,3,8,6]]"),
			7,
		},
		// {
		// 	leetcode.ParseMatrix("[[0,2,4],[3,2,1],[1,0,4]]"),
		// 	-1,
		// },
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTime(tc.grid))
		})
	}
}

func minimumTime(grid [][]int) int {
	// Based on pen-and-paper, if standing next to a certain cell at time t,
	// then the earliest time to get to the next cell is to come there at the
	// lowest possible d such that t_cell >= t + 1 + 2*d
	//
	// The only time when there is no solution is when there is only one cell
	// and both neighbours > 1.
	if grid[0][1] > 1 && grid[1][0] > 1 {
		return -1
	}

	// Next, adjust the time when a square can be reached so that it matches
	// reality. Any cell which has a t that is not odd/even in accordance with
	// manhattan distance is adjusted.
	//
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
	curr := [][]int{{0, 0}}
	next := [][]int{}

	for dist := 1; len(curr) > 0; dist++ {
		next = next[:0]
		for _, x := range curr {
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				if !ok(ii, jj) || seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				if grid[ii][jj]&1 != dist&1 {
					grid[ii][jj]++
				}
				next = append(next, []int{ii, jj})
			}
		}
		curr, next = next, curr
	}

	// Finally, do another flood fill with union find to determine the moment
	// when top-left is combined with bottom-right. From that point on, keep
	// consuming cells until there are no more cells or the cell value is higher
	// than the current minimum time to reach the end.
	curr = curr[:0]
	curr = append(curr, []int{0, 0})
	h := make(minHeap, 0)

	for i := range seen {
		seen[i] = append(seen[i][:0], make([]bool, n)...)
	}
	seen[0][0] = true

	for t := 0; ; t++ {
		for len(h) > 0 && h[0].t <= t {
			x := heap.Pop(&h).(cell)
			curr = append(curr, []int{x.i, x.j})
		}
		next = next[:0]
		for _, x := range curr {
			// if x[0]+x[1]&1 != t&1 {
			// 	next = append(next, x)
			// 	continue
			// }
			if x[0] == m-1 && x[1] == n-1 {
				return t
			}
			for _, d := range dirs {
				ii := x[0] + d[0]
				jj := x[1] + d[1]
				if !ok(ii, jj) || seen[ii][jj] {
					continue
				}
				seen[ii][jj] = true
				if grid[ii][jj] > t+1 {
					heap.Push(&h, cell{ii, jj, grid[ii][jj]})
					continue
				}
				next = append(next, []int{ii, jj})
			}
		}
		curr, next = next, curr
	}
}

type cell struct {
	i, j int
	t    int
}

type minHeap []cell

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].heapIdx = i
	// h[j].heapIdx = j
}
func (h minHeap) Less(i, j int) bool {
	return h[i].t < h[j].t
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
