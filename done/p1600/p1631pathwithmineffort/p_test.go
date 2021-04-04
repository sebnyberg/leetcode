package p1613pathwithmineffort

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumEffortPath(t *testing.T) {
	for _, tc := range []struct {
		heights [][]int
		want    int
	}{
		{[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}, 2},
		{[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}, 1},
		{[][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.heights), func(t *testing.T) {
			require.Equal(t, tc.want, minimumEffortPath(tc.heights))
		})
	}
}

type point struct {
	i, j int
}

type pointDistance struct {
	p        point
	distance int
}

type pointHeap []pointDistance

func (h pointHeap) Len() int           { return len(h) }
func (h pointHeap) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h pointHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *pointHeap) Push(x interface{}) {
	*h = append(*h, x.(pointDistance))
}

func (h *pointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights[0]), len(heights)
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	h := pointHeap{pointDistance{point{0, 0}, 0}}
	heap.Init(&h)
	for {
		cur := heap.Pop(&h).(pointDistance)
		if cur.p.i == m-1 && cur.p.j == n-1 {
			return cur.distance
		}
		visited[cur.p.i][cur.p.j] = true
		for _, near := range []point{
			{cur.p.i - 1, cur.p.j},
			{cur.p.i + 1, cur.p.j},
			{cur.p.i, cur.p.j - 1},
			{cur.p.i, cur.p.j + 1},
		} {
			if near.i < 0 || near.j < 0 || near.i >= m || near.j >= n || visited[near.i][near.j] {
				continue
			}
			heap.Push(&h, pointDistance{
				p: near,
				distance: max(
					cur.distance,
					abs(heights[near.i][near.j]-heights[cur.p.i][cur.p.j]),
				),
			})
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
