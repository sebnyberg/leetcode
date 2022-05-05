package p2077pathsinmazethatleadtosameroom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfPaths(t *testing.T) {
	for _, tc := range []struct {
		n         int
		corridors [][]int
		want      int
	}{
		{3, [][]int{{1, 3}, {2, 3}, {2, 1}}, 1},
		{5, [][]int{{1, 2}, {5, 2}, {4, 1}, {2, 4}, {3, 1}, {3, 4}}, 2},
		{4, [][]int{{1, 2}, {3, 4}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfPaths(tc.n, tc.corridors))
		})
	}
}

func numberOfPaths(n int, corridors [][]int) int {
	adj := make([][]int, n+1)
	nei := make([][]bool, n+1)
	for i := range nei {
		nei[i] = make([]bool, n+1)
	}
	for _, corr := range corridors {
		a, b := corr[0], corr[1]
		nei[a][b] = true
		nei[b][a] = true
		if a < b {
			adj[a] = append(adj[a], b)
		} else {
			adj[b] = append(adj[b], a)
		}
	}
	var res int
	for i := 1; i <= n; i++ {
		for j := 0; j < len(adj[i]); j++ {
			a := adj[i][j]
			for k := j + 1; k < len(adj[i]); k++ {
				b := adj[i][k]
				if nei[a][b] || nei[b][a] {
					res++
				}
			}
		}
	}
	return res
}
