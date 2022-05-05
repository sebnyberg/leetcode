package p1627graphconnectivitywiththreshold

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_areConnected(t *testing.T) {
	for _, tc := range []struct {
		n         int
		threshold int
		queries   [][]int
		want      []bool
	}{
		{6, 2, [][]int{{1, 4}, {2, 5}, {3, 6}}, []bool{false, false, true}},
		{6, 0, [][]int{{4, 5}, {3, 4}, {3, 2}, {2, 6}, {1, 3}}, []bool{true, true, true, true, true}},
		{5, 1, [][]int{{4, 5}, {4, 5}, {3, 2}, {2, 3}, {3, 4}}, []bool{false, false, false, false, false}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, areConnected(tc.n, tc.threshold, tc.queries))
		})
	}
}

func areConnected(n int, threshold int, queries [][]int) []bool {
	dsu := NewDSU(n + 1)
	for k := threshold + 1; 2*k <= n; k++ {
		for j := 2 * k; j <= n; j += k {
			dsu.union(k, j)
		}
	}

	res := make([]bool, len(queries))
	for i, query := range queries {
		res[i] = dsu.find(query[0]) == dsu.find(query[1])
	}
	return res
}

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
		rank:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
		dsu.rank[i] = 1
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
		if d.rank[a] < d.rank[b] {
			a, b = b, a
		}
		d.parent[b] = a
		d.rank[a]++
	}
}
