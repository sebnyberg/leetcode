package p0547numberofprovinces

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findCircleNum(t *testing.T) {
	for _, tc := range []struct {
		isConnected [][]int
		want        int
	}{
		{[][]int{{1, 0, 0, 1}, {0, 1, 1, 0}, {0, 1, 1, 1}, {1, 0, 1, 1}}, 1},
		// {[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.isConnected), func(t *testing.T) {
			require.Equal(t, tc.want, findCircleNum(tc.isConnected))
		})
	}
}

func findCircleNum(isConnected [][]int) int {
	dsu := NewDSU(len(isConnected))
	for i := range isConnected {
		for j, n := range isConnected[i] {
			if j == i || n == 0 {
				continue
			}
			dsu.Union(i, j)
		}
	}

	n := len(dsu.Groups())

	return n
}

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
	}
	for i := range dsu.parent {
		dsu.parent[i] = i
	}
	return dsu
}

// Point x -> y
func (d *DSU) Union(x, y int) {
	d.parent[d.Find(x)] = d.Find(y)
}

// Find x in the disjoint set union
func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		return d.Find(d.parent[x])
	}
	return d.parent[x]
}

func (d *DSU) Groups() []int {
	parentIDs := make(map[int]struct{})
	for id := range d.parent {
		parentIDs[d.Find(id)] = struct{}{}
	}
	parentIDsList := make([]int, 0, len(parentIDs))
	for parentID := range parentIDs {
		parentIDsList = append(parentIDsList, parentID)
	}
	sort.Ints(parentIDsList)
	return parentIDsList
}
