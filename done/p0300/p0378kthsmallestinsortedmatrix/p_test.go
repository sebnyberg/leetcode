package p0378kthsmallestinsortedmatrix

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthSmallest(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{
			[][]int{
				{1, 5, 9},
				{10, 11, 13},
				{12, 13, 15},
			},
			8, 13,
		},
		{
			[][]int{
				{1, 2},
				{3, 3},
			},
			2, 2,
		},
		{
			[][]int{
				{-5},
			},
			1, -5,
		},
		{
			[][]int{
				{1, 2},
				{1, 3},
			},
			2, 1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, kthSmallest(tc.matrix, tc.k))
		})
	}
}

func kthSmallest(matrix [][]int, k int) int {
	// This was not what I thought it was - a medium exercise.
	// Idea:
	// Keep a heap of numbers, then pop from the heap until finding the kth
	// element
	m, n := len(matrix), len(matrix[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	tovisit := make(posHeap, 0)
	tovisit = append(tovisit, &pos{0, 0, matrix[0][0]})
	visited[0][0] = true
	i := 0
	for len(tovisit) > 0 {
		x := heap.Pop(&tovisit).(*pos)
		i++
		if i == k {
			return x.val
		}
		if x.i < m-1 && !visited[x.i+1][x.j] {
			visited[x.i+1][x.j] = true
			heap.Push(&tovisit, &pos{x.i + 1, x.j, matrix[x.i+1][x.j]})
		}
		if x.j < n-1 && !visited[x.i][x.j+1] {
			visited[x.i][x.j+1] = true
			heap.Push(&tovisit, &pos{x.i, x.j + 1, matrix[x.i][x.j+1]})
		}
	}
	return -1
}

type pos struct {
	i, j, val int
}

type posHeap []*pos

func (h posHeap) Len() int { return len(h) }
func (h posHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h posHeap) Less(i, j int) bool {
	return h[i].val < h[j].val
}
func (h *posHeap) Push(x interface{}) {
	*h = append(*h, x.(*pos))
}
func (h *posHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
