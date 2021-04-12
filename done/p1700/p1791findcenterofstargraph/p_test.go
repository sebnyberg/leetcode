package p1791findcenterofstargraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findCenter(t *testing.T) {
	for _, tc := range []struct {
		edges [][]int
		want  int
	}{
		{[][]int{{1, 2}, {2, 3}, {4, 2}}, 2},
		{[][]int{{1, 2}, {5, 1}, {1, 3}, {1, 4}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, findCenter(tc.edges))
		})
	}
}

func findCenter(edges [][]int) int {
	n := len(edges) + 1
	edgeCount := make([]int, n+1)
	for _, edge := range edges {
		a, b := edge[0], edge[1]
		edgeCount[a]++
		edgeCount[b]++
		if edgeCount[a] == n-1 {
			return a
		}
		if edgeCount[b] == n-1 {
			return b
		}
	}
	return -1
}
