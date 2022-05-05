package p0765couplesholdinghands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSwapsCouples(t *testing.T) {
	for _, tc := range []struct {
		row  []int
		want int
	}{
		{[]int{9, 12, 2, 10, 11, 0, 13, 6, 4, 5, 3, 8, 1, 7}, 5},
		{[]int{0, 2, 1, 3}, 1},
		{[]int{3, 2, 0, 1}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.row), func(t *testing.T) {
			require.Equal(t, tc.want, minSwapsCouples(tc.row))
		})
	}
}

func minSwapsCouples(row []int) int {
	n := len(row)

	pos := make([]int, n)
	for i, v := range row {
		pos[v] = i / 2
	}

	seen := make([]bool, n/2)
	dsu := NewDSU(n / 2)
	var visit func(i int, target int)
	visit = func(coupleIdx int, target int) {
		if seen[coupleIdx] {
			return
		}
		a, b := row[coupleIdx*2], row[coupleIdx*2+1]
		// we want to visit the non-target's ^1
		if a == target {
			a = b
		}
		nextTarget := a ^ 1
		seen[coupleIdx] = true
		dsu.union(coupleIdx, pos[nextTarget])
		visit(pos[nextTarget], nextTarget)
	}
	for i := 0; i < n; i += 2 {
		a, b := row[i], row[i+1]
		if a == b^1 || seen[i/2] {
			continue
		}
		visit(i/2, a)
	}
	nswaps := 0
	for _, s := range dsu.size {
		if s > 1 {
			nswaps += s - 1
		}
	}

	return nswaps
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
