package p0839similarstringgroups

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSimilarGroups(t *testing.T) {
	for _, tc := range []struct {
		strs []string
		want int
	}{
		{[]string{"tars", "rats", "arts", "start"}, 2},
		{[]string{"omv", "ovm"}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strs), func(t *testing.T) {
			require.Equal(t, tc.want, numSimilarGroups(tc.strs))
		})
	}
}

func numSimilarGroups(strs []string) int {
	// Brute-force w/ UF
	dsu := NewDSU(len(strs))
	for i := range strs {
		for j := range strs[i+1:] {
			if similar(strs[i], strs[i+j+1]) {
				dsu.union(i, i+j+1)
			}
		}
	}
	return dsu.ngroups
}

func similar(s, t string) bool {
	diff := 0
	for i := range s {
		if s[i] != t[i] {
			if diff == 2 {
				return false
			}
			diff++
		}
	}
	return true
}

type DSU struct {
	parent  []int
	size    []int
	ngroups int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent:  make([]int, n),
		size:    make([]int, n),
		ngroups: n,
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
	if a == b {
		return
	}
	d.ngroups--
	if d.size[a] < d.size[b] {
		a, b = b, a
	}
	d.parent[b] = a
	d.size[a] += d.size[b]
}
