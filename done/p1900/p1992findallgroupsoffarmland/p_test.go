package p1992findallgroupsoffarmland

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findFarmland(t *testing.T) {
	for _, tc := range []struct {
		land [][]int
		want [][]int
	}{
		{[][]int{{1, 0, 0}, {0, 1, 1}, {0, 1, 1}}, [][]int{{0, 0, 0, 0}, {1, 1, 2, 2}}},
		{[][]int{{1, 1}, {1, 1}}, [][]int{{0, 0, 1, 1}}},
		{[][]int{{0}}, [][]int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.land), func(t *testing.T) {
			require.Equal(t, tc.want, findFarmland(tc.land))
		})
	}
}

func findFarmland(land [][]int) [][]int {
	// 0 -> forested
	// 1 -> farmland
	// Find farmland groups, print out the top left / bottom right corner
	// for each group.
	m, n := len(land), len(land[0])
	flatIdx := func(i, j int) int {
		return i*n + j
	}
	dsu := NewDSU(m*n + 1)
	for i := range land {
		for j := range land[i] {
			if land[i][j] != 1 {
				continue
			}
			for _, near := range [][2]int{
				{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
			} {
				ii, jj := near[0], near[1]
				if ii < 0 || jj < 0 || ii >= m || jj >= n || land[ii][jj] != 1 {
					continue
				}
				// join land in DSU
				dsu.union(flatIdx(i, j), flatIdx(ii, jj))
			}
		}
	}
	// For each unique root idx in the DSU, find top left / bottom right corner
	landCorners := make(map[int][4]int)
	for i := range land {
		for j := range land[i] {
			if land[i][j] != 1 {
				continue
			}
			root := dsu.find(flatIdx(i, j))
			if _, exists := landCorners[root]; !exists {
				landCorners[root] = [4]int{i, i, j, j}
				continue
			}
			landCorners[root] = [4]int{
				min(landCorners[root][0], i),
				max(landCorners[root][1], i),
				min(landCorners[root][2], j),
				max(landCorners[root][3], j),
			}
		}
	}
	res := make([][]int, 0, len(landCorners))
	for _, corners := range landCorners {
		res = append(res, []int{
			corners[0], corners[2], corners[1], corners[3],
		})
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type DSU struct {
	parent []int
	size   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.size[i] = 1
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if d.parent[a] == a {
		return a
	}
	root := d.find(d.parent[a])
	d.parent[a] = root
	return root
}

func (d *DSU) union(a, b int) {
	a = d.find(a)
	b = d.find(b)
	if a != b {
		if d.size[a] < d.size[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.size[a] += d.size[b]
	}
}
