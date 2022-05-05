package p1724checkingexistenceofedgelengthlimitedpaths2

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDistanceLimitedPathExists(t *testing.T) {
	d := Constructor(6, [][]int{{0, 2, 4}, {0, 3, 2}, {1, 2, 3}, {2, 3, 1}, {4, 5, 5}})
	res := d.Query(2, 3, 2)
	require.Equal(t, true, res)
	res = d.Query(1, 3, 3)
	require.Equal(t, false, res)
	res = d.Query(2, 0, 3)
	require.Equal(t, true, res)
	res = d.Query(0, 5, 6)
	require.Equal(t, false, res)
}

type DistanceLimitedPathsExist struct {
	dsus []DSUDistance
}

type DSUDistance struct {
	d        *DSU
	distance int
}

func Constructor(n int, edgeList [][]int) DistanceLimitedPathsExist {
	// Sort edgeList by distance
	sort.Slice(edgeList, func(i, j int) bool {
		return edgeList[i][2] < edgeList[j][2]
	})

	res := DistanceLimitedPathsExist{}

	dsu := NewDSU(n)
	var j int
	m := len(edgeList)
	for i := 0; i < m; {
		for j = i + 1; j < m && edgeList[i][2] == edgeList[j][2]; j++ {
		}
		for _, edge := range edgeList[i:j] {
			dsu.union(edge[0], edge[1])
		}
		res.dsus = append(res.dsus, DSUDistance{
			d:        dsu,
			distance: edgeList[i][2],
		})
		dsu = copyDSU(dsu)
		i = j
	}

	return res
}

func (this *DistanceLimitedPathsExist) Query(p int, q int, limit int) bool {
	if this.dsus[0].distance >= limit {
		return false
	}
	i := 0
	for ; i < len(this.dsus) && this.dsus[i].distance < limit; i++ {
	}
	i--
	return this.dsus[i].d.find(p) == this.dsus[i].d.find(q)
}

/**
 * Your DistanceLimitedPathsExist object will be instantiated and called as such:
 * obj := Constructor(n, edgeList);
 * param_1 := obj.Query(p,q,limit);
 */

type DSU struct {
	parent []int
	size   []int
}

func copyDSU(d *DSU) *DSU {
	cpy := &DSU{
		parent: make([]int, len(d.parent)),
		size:   make([]int, len(d.size)),
	}
	copy(cpy.parent, d.parent)
	copy(cpy.size, d.size)
	return cpy
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
