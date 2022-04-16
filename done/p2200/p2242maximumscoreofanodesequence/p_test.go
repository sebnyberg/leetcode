package p2242

import (
	"fmt"
	"leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumScore(t *testing.T) {
	for _, tc := range []struct {
		scores []int
		edges  [][]int
		want   int
	}{
		{
			[]int{5, 2, 9, 8, 4},
			leetcode.ParseMatrix("[[0,1],[1,2],[2,3],[0,2],[1,3],[2,4]]"),
			24,
		},
		{
			[]int{9, 20, 6, 4, 11, 12},
			leetcode.ParseMatrix("[[0,3],[5,3],[2,4],[1,3]]"),
			-1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.scores), func(t *testing.T) {
			require.Equal(t, tc.want, maximumScore(tc.scores, tc.edges))
		})
	}
}

func maximumScore(scores []int, edges [][]int) int {
	n := len(scores)
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	// Sort adj lists by score, keep only 4 largest
	for i := range adj {
		sort.Slice(adj[i], func(j, k int) bool {
			return scores[adj[i][j]] > scores[adj[i][k]]
		})
		adj[i] = adj[i][:min(len(adj[i]), 3)]
	}

	res := -1
	for _, e := range edges {
		b, c := e[0], e[1]
		for _, a := range adj[b] {
			for _, d := range adj[c] {
				if a == d || a == c || b == d {
					continue
				}
				res = max(res, scores[a]+scores[b]+scores[c]+scores[d])
			}
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
