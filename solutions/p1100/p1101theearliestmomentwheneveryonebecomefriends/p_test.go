package p1101theearliestmomentwheneveryonebecomefriends

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_earliestAcq(t *testing.T) {
	for _, tc := range []struct {
		logs [][]int
		N    int
		want int
	}{
		{[][]int{{20190101, 0, 1}, {20190104, 3, 4}, {20190107, 2, 3}, {20190211, 1, 5}, {10290224, 2, 4}, {20190301, 0, 3}, {20190312, 1, 2}, {20190322, 4, 5}}, 6, 20190301},
	} {
		t.Run(fmt.Sprintf("%+v", tc.logs), func(t *testing.T) {
			require.Equal(t, tc.want, earliestAcq(tc.logs, tc.N))
		})
	}
}

func earliestAcq(logs [][]int, N int) int {
	dsu := NewDSU(N + 1)
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})
	need := N - 1
	for _, log := range logs {
		if dsu.find(log[1]) == dsu.find(log[2]) {
			continue
		}
		dsu.union(log[1], log[2])
		need--
		if need == 0 {
			return log[0]
		}
	}
	return -1
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
