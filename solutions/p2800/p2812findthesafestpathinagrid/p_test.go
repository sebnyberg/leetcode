package p2812findthesafestpathinagrid

import (
	"container/heap"
	"math"
)

func maximumSafenessFactor(grid [][]int) int {
	topology := createTopologyGrid(grid)

	n := len(grid)
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < n && j < n
	}

	h := maxHeap{
		topology: topology,
	}
	heap.Push(&h, []int{0, 0})
	minDist := topology[0][0]
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, n)
	}
	seen[0][0] = true
	for {
		x := heap.Pop(&h).([]int)
		minDist = min(minDist, topology[x[0]][x[1]])
		if x[0] == n-1 && x[1] == n-1 {
			return minDist
		}
		for _, d := range dirs {
			ii := x[0] + d[0]
			jj := x[1] + d[1]
			if !ok(ii, jj) || seen[ii][jj] {
				continue
			}
			seen[ii][jj] = true
			heap.Push(&h, []int{ii, jj})
		}
	}
}

var dirs = [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func createTopologyGrid(grid [][]int) [][]int {
	n := len(grid)
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res {
			if grid[i][j] != 1 {
				res[i][j] = math.MaxInt32
			}
		}
	}

	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < n && j < n
	}

	// Simply do bfs from each thief position.
	var curr, next [][]int
	bfs := func(i, j int) {
		curr = append(curr[:0], []int{i, j})
		for k := 1; len(curr) > 0; k++ {
			next = next[:0]
			for _, x := range curr {
				for _, d := range dirs {
					xx := x[0] + d[0]
					yy := x[1] + d[1]
					if !ok(xx, yy) || res[xx][yy] <= k {
						continue
					}
					res[xx][yy] = k
					next = append(next, []int{xx, yy})
				}
			}
			curr, next = next, curr
		}
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				bfs(i, j)
			}
		}
	}

	return res
}

type maxHeap struct {
	topology  [][]int
	positions [][]int
}

func (h maxHeap) Len() int { return len(h.positions) }
func (h maxHeap) Swap(i, j int) {
	h.positions[i], h.positions[j] = h.positions[j], h.positions[i]
}
func (h maxHeap) Less(i, j int) bool {
	x := h.topology[h.positions[i][0]][h.positions[i][1]]
	y := h.topology[h.positions[j][0]][h.positions[j][1]]
	return x > y
}
func (h *maxHeap) Push(x interface{}) {
	h.positions = append(h.positions, x.([]int))
}
func (h *maxHeap) Pop() interface{} {
	n := len(h.positions)
	it := h.positions[n-1]
	h.positions = h.positions[:n-1]
	return it
}
