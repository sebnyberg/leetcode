package p1878getbiggestthreerhombussums

import (
	"container/heap"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getBiggestThree(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want []int
	}{
		{
			[][]int{
				{8, 20, 12, 14, 12, 18},
				{8, 17, 6, 4, 4, 10},
				{20, 13, 7, 18, 19, 16},
				{3, 17, 20, 5, 18, 13},
				{2, 7, 17, 12, 6, 7},
				{19, 2, 10, 11, 20, 2},
				{1, 1, 17, 1, 14, 4},
				{9, 11, 2, 14, 11, 16},
				{14, 7, 2, 5, 11, 20}},
			[]int{111, 103, 101},
		},
		{
			[][]int{
				{20, 17, 9, 13, 5, 2, 9, 1, 5},
				{14, 9, 9, 9, 16, 18, 3, 4, 12},
				{18, 15, 10, 20, 19, 20, 15, 12, 11},
				{19, 16, 19, 18, 8, 13, 15, 14, 11},
				{4, 19, 5, 2, 19, 17, 7, 2, 2},
			},
			[]int{107, 103, 102},
		},
		{
			[][]int{{3, 4, 5, 1, 3}, {3, 3, 4, 2, 3}, {20, 30, 200, 40, 10}, {1, 5, 5, 4, 1}, {4, 3, 2, 2, 5}},
			[]int{228, 216, 211},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{20, 9, 8},
		},
		{
			[][]int{{7, 7, 7}},
			[]int{7},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			printGrid(tc.grid)
			require.Equal(t, tc.want, getBiggestThree(tc.grid))
		})
	}
}

func printGrid(grid [][]int) {
	for i := range grid {
		for _, n := range grid[i] {
			fmt.Fprintf(os.Stdout, "%v\t", n)
		}
		fmt.Fprintf(os.Stdout, "\n")
	}
}

func getBiggestThree(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	maxSumHeap := make(IntHeap, 0)
	maxSize := (min(m, n) - 1) / 2
	for size := 0; size <= maxSize; size++ {
		for rowOffset := 0; rowOffset+size < m; rowOffset++ {
			for colOffset := 0; colOffset+size < n; colOffset++ {
				// Check sum
				sum := countSum(grid, rowOffset, colOffset, size)
				if sum > 0 {
					heap.Push(&maxSumHeap, sum)
				}
			}
		}
	}

	i := 0
	res := make([]int, 0, 3)
	for len(maxSumHeap) > 0 && i < 3 {
		x := heap.Pop(&maxSumHeap).(int)
		if i == 0 {
			res = append(res, x)
		} else {
			if res[i-1] == x {
				continue
			}
			res = append(res, x)
		}
		i++
	}

	return res
}

func ok(i, j, m, n int) bool {
	return i >= 0 && j >= 0 && i < m && j < n
}

func countSum(grid [][]int, row, col, size int) int {
	if size == 0 {
		return grid[row][col]
	}
	m, n := len(grid), len(grid[0])
	var sum int
	for _, incr := range [][2]int{{-1, 1}, {1, 1}, {1, -1}, {-1, -1}} {
		for k := 0; k < size; k++ {
			row += incr[0]
			col += incr[1]
			if !ok(row, col, m, n) {
				return -1
			}
			sum += grid[row][col]
		}
	}
	return sum
}

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
