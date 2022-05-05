package p0202smalleststringwithswaps

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestStringWithSwaps(t *testing.T) {
	for _, tc := range []struct {
		s     string
		pairs [][]int
		want  string
	}{
		{"dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}, "abcd"},
		{"dcab", [][]int{{0, 3}, {1, 2}}, "bacd"},
		{"cba", [][]int{{0, 1}, {1, 2}}, "abc"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, smallestStringWithSwaps(tc.s, tc.pairs))
		})
	}
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	dsu := NewDSU(len(s))
	for _, pair := range pairs {
		dsu.union(pair[0], pair[1])
	}

	groups := make(map[int][]int)
	for i := range s {
		r := dsu.find(i)
		groups[r] = append(groups[r], i)
	}

	res := []byte(s)
	for _, indices := range groups {
		chs := make([]byte, len(indices))
		for i, idx := range indices {
			chs[i] = s[idx]
		}
		sort.Ints(indices)
		sort.Slice(chs, func(i, j int) bool { return chs[i] < chs[j] })
		for i, idx := range indices {
			res[idx] = chs[i]
		}
	}

	// 5. Return
	return string(res)
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
