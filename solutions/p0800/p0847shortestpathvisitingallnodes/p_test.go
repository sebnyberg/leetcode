package p0847shortestpathvisitingallnodes

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
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

type state struct {
	bm int
	i  int
}

func shortestPathLength(graph [][]int) int {
	// Let's perform a large BFS of the graph and hope memory doesn't choke.
	//
	// We start in all positions simultaneously, marking their bitmasked state as
	// seen.
	//
	// Whenever we are given the option to revisit a prior state we refuse to do
	// so.
	//
	// Whenever the current bitmask becomes full, we exit with the number of steps
	// of BFS.

	n := len(graph)
	if n == 1 {
		return 0
	}
	seen := make(map[state]struct{})
	curr := []state{}
	next := []state{}
	hasSeen := func(s state) bool {
		if _, exists := seen[s]; exists {
			return true
		}
		return false
	}
	for i := 0; i < n; i++ {
		s := state{1 << i, i}
		seen[s] = struct{}{}
		curr = append(curr, s)
	}
	done := (1 << n) - 1
	for steps := 1; ; steps++ {
		next = next[:0]
		for _, x := range curr {
			for _, nei := range graph[x.i] {
				nextState := state{x.bm | (1 << nei), nei}
				if hasSeen(nextState) {
					continue
				}
				if nextState.bm == done {
					return steps
				}
				seen[nextState] = struct{}{}
				next = append(next, nextState)
			}
		}
		curr, next = next, curr
	}
}
