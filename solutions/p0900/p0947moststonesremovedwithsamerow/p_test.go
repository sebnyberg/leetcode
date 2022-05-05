package p0947moststonesremovedwithsamerow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_removeStones(t *testing.T) {
	for _, tc := range []struct {
		stones [][]int
		want   int
	}{
		{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}, 5},
		{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}, 3},
		{[][]int{{0, 0}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.stones), func(t *testing.T) {
			require.Equal(t, tc.want, removeStones(tc.stones))
		})
	}
}

func removeStones(stones [][]int) int {
	n := len(stones)
	dsu := NewDSU(20000)
	for _, stone := range stones {
		dsu.union(stone[0], stone[1]+10000)
	}

	seen := make(map[int]struct{})
	groups := 0
	for _, stone := range stones {
		root := dsu.find(stone[0])
		if _, exists := seen[root]; !exists {
			groups++
		}
		seen[root] = struct{}{}
	}
	return n - groups
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
