package p1722minimizehammingdistanceafterswap

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumHammingDistance(t *testing.T) {
	for _, tc := range []struct {
		source       []int
		target       []int
		allowedSwaps [][]int
		want         int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}, [][]int{}, 2},
		{[]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.source), func(t *testing.T) {
			require.Equal(t, tc.want, minimumHammingDistance(tc.source, tc.target, tc.allowedSwaps))
		})
	}
}

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	n := len(source)
	dsu := NewDSU(n)
	for _, swap := range allowedSwaps {
		dsu.union(swap[0], swap[1])
	}

	swapGroups := make(map[int][]int, len(source))
	groupNums := make(map[int]map[int]int, len(source))
	for i := range source {
		root := dsu.find(i)
		swapGroups[root] = append(swapGroups[root], i)
		if _, exists := groupNums[root]; !exists {
			groupNums[root] = make(map[int]int)
		}
		groupNums[root][source[i]]++
	}

	hammingDist := 0
	for root, indices := range swapGroups {
		for _, idx := range indices {
			if groupNums[root][target[idx]] > 0 {
				groupNums[root][target[idx]]--
			} else {
				hammingDist++
			}
		}
	}

	return hammingDist
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

func (d *DSU) sameSet(a, b int) bool {
	return d.find(a) == d.find(b)
}
