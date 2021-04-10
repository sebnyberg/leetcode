package p1168optimizewaterdistributioninavillage

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCostToSupplyWater(t *testing.T) {
	for _, tc := range []struct {
		n     int
		wells []int
		pipes [][]int
		want  int
	}{
		{3, []int{1, 2, 2}, [][]int{{1, 2, 1}, {2, 3, 1}}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minCostToSupplyWater(tc.n, tc.wells, tc.pipes))
		})
	}
}

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
	for i, v := range wells {
		pipes = append(pipes, []int{0, i + 1, v})
	}
	sort.Slice(pipes, func(i, j int) bool {
		return pipes[i][2] < pipes[j][2]
	})
	cost := 0
	dsu := NewDSU(n + 1)
	for _, pipe := range pipes {
		a, b := dsu.find(pipe[0]), dsu.find(pipe[1])
		if a == b {
			continue
		}
		dsu.union(a, b)
		cost += pipe[2]
	}
	return cost
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
