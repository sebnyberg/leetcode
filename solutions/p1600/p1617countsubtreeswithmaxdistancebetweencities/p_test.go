package p1617countsubtreeswithmaxdistancebetweencities

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_countSubgraphsForEachDiameter(t *testing.T) {
	for i, tc := range []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{
			4,
			leetcode.ParseMatrix("[[1,2],[2,3],[2,4]]"),
			[]int{3, 4, 0},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, countSubgraphsForEachDiameter(tc.n, tc.edges))
		})
	}
}

func countSubgraphsForEachDiameter(n int, edges [][]int) []int {
	adj := make([][]int, n)
	for i := range edges {
		a := edges[i][0] - 1
		b := edges[i][1] - 1
		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}
	curr := []int{}
	next := []int{}
	maxLen := func(i, bm int) int {
		var k int
		curr = append(curr[:0], i)
		bm ^= 1 << i
		for len(curr) > 0 {
			next = next[:0]
			for _, x := range curr {
				for _, y := range adj[x] {
					if bm&(1<<y) == 0 {
						continue
					}
					bm &^= 1 << y
					next = append(next, y)
				}
			}
			curr, next = next, curr
			k++
		}
		if bm != 0 {
			return -1
		}
		return k
	}

	res := make([]int, n-1)
outer:
	for x := 1; x < (1 << n); x++ {
		// Lazy version - BFS from every possible start node..
		var maxDist int
		for i := 0; i < n; i++ {
			if x&(1<<i) > 0 {
				d := maxLen(i, x)
				maxDist = max(maxDist, d)
				if d == -1 {
					continue outer
				}
			}
		}
		if maxDist > 1 {
			res[maxDist-2]++
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
