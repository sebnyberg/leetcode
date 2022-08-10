package p0802findeventualsafestates

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_eventualSafeNodes(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		want  []int
	}{
		{
			leetcode.ParseMatrix("[[1,2],[2,3],[5],[0],[5],[],[]]"),
			[]int{2, 4, 5, 6},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.graph), func(t *testing.T) {
			require.Equal(t, tc.want, eventualSafeNodes(tc.graph))
		})
	}
}

func eventualSafeNodes(graph [][]int) []int {
	// It appears we want to discard any nodes which has a path that leads to a
	// cycle, and all nodes which have a path to a node which leads to a cycle.
	// We can do this with a stack/map-based approach.
	// For each node, check if it's been seen before. If it hasn't, start DFS.
	n := len(graph)
	adj := graph
	result := make([]byte, n)
	const (
		unseen     = 0
		inProgress = 1
		safe       = 2
		unsafe     = 3
	)
	var safeNodes []int
	var check func(i int) bool
	check = func(i int) bool {
		if len(adj) == 0 {
			result[i] = safe
			safeNodes = append(safeNodes, i)
		}
		if result[i] >= 2 {
			return result[i] == safe
		}
		if result[i] == inProgress {
			result[i] = unsafe
			return false
		}
		result[i] = inProgress
		for _, nei := range adj[i] {
			if !check(nei) {
				result[i] = unsafe
				return false
			}
		}
		safeNodes = append(safeNodes, i)
		result[i] = safe
		return true
	}

	for i, v := range result {
		if v != 0 {
			continue
		}
		check(i)
	}
	sort.Ints(safeNodes)
	return safeNodes
}
