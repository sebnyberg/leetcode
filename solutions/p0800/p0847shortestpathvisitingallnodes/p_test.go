package p0847shortestpathvisitingallnodes

import (
	"fmt"
	"leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestPathLength(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		want  int
	}{
		{leetcode.ParseMatrix("[[1],[0,2,4],[1,3],[2],[1,5],[4]]"), 6},
		{leetcode.ParseMatrix("[[1,2,3],[0],[0],[0]]"), 4},
		{leetcode.ParseMatrix("[[1],[0,2,4],[1,3,4],[2],[1,2]]"), 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.graph), func(t *testing.T) {
			require.Equal(t, tc.want, shortestPathLength(tc.graph))
		})
	}
}

func shortestPathLength(graph [][]int) int {
	n := len(graph)
	end := (1 << n) - 1
	mem := make([][]int, n)
	for i := range mem {
		mem[i] = make([]int, 1<<n)
	}
	seen := make([][]bool, n)
	for i := range seen {
		seen[i] = make([]bool, 1<<n)
	}
	curr := make([][2]int, n)
	for i := 0; i < n; i++ {
		curr[i] = [2]int{i, 1 << i}
	}
	next := make([][2]int, 0, n)
	for steps := 1; len(curr) > 0; steps++ {
		next = next[:0]

		for _, x := range curr {
			node, mask := x[0], x[1]
			for _, nei := range graph[node] {
				nextMask := mask | (1 << nei)
				if nextMask == end {
					return steps
				}
				if seen[nei][nextMask] {
					continue
				}
				seen[nei][nextMask] = true
				next = append(next, [2]int{nei, nextMask})
			}
		}

		curr, next = next, curr
	}

	return 0
}
