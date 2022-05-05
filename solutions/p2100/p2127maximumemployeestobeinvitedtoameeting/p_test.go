package p2127maximumemployeestobeinvitedtoameeting

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumInvitations(t *testing.T) {
	for _, tc := range []struct {
		favorite []int
		want     int
	}{
		{[]int{8, 3, 4, 8, 0, 6, 7, 0, 5}, 5},
		{[]int{1, 0, 3, 2, 5, 6, 7, 4, 9, 8, 11, 10, 11, 12, 10}, 11},
		{[]int{1, 0, 0, 2, 1, 4, 7, 8, 9, 6, 7, 10, 8}, 6},
		{[]int{2, 2, 1, 2}, 3},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 0, 1, 4, 1}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.favorite), func(t *testing.T) {
			require.Equal(t, tc.want, maximumInvitations(tc.favorite))
		})
	}
}

func maximumInvitations(favorite []int) int {
	// Any node with an indegree of zero must be the starting point toward a
	// cycle

	// All nodes lead to a cycle or are part of a cycle
	// A cycle of length 2 is a 'chained cycle' and can be combined with other
	// chained cycles to form a large table.

	n := len(favorite)
	parents := make([][]int, len(favorite))
	for i, j := range favorite {
		parents[j] = append(parents[j], i)
	}
	visited := make([]bool, n)
	// First, find any node for which its favourite's favourite is the current
	// node.
	var chained int
	for i := range favorite {
		if visited[i] {
			continue
		}
		if favorite[favorite[i]] == i {
			// Mark both as visited and find maximum paths leading to these two nodes
			left, right := i, favorite[i]
			visited[left] = true
			visited[right] = true
			chained += findChain(visited, parents, 1, left)
			chained += findChain(visited, parents, 1, right)
		}
	}

	// For each unvisited node, find its cycle
	cycle := make([]bool, n)
	var maxCycleLen int
	for i, v := range visited {
		if v {
			continue
		}
		maxCycleLen = max(maxCycleLen, findCycle(visited, cycle, favorite, i))
	}

	// Return result
	return max(chained, maxCycleLen)
}

func findChain(visited []bool, parents [][]int, ntot, idx int) int {
	res := ntot
	for _, parent := range parents[idx] {
		if !visited[parent] {
			visited[parent] = true
			res = max(res, findChain(visited, parents, ntot+1, parent))
		}
	}
	return res
}

func findCycle(visited, cycle []bool, favorite []int, idx int) int {
	cur, next := idx, favorite[idx]
	stack := []int{cur}
	visited[cur] = true
	defer func() {
		for _, idx := range stack {
			cycle[idx] = true
		}
	}()
	for !visited[next] {
		cur, next = next, favorite[next]
		visited[cur] = true
		stack = append(stack, cur)
	}
	if cycle[next] {
		return 0
	}
	cycleLen := 1
	start := cur
	for next != start {
		cycleLen++
		cur, next = next, favorite[next]
	}
	return cycleLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
