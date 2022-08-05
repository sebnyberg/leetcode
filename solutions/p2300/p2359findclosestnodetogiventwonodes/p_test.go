package p3

import (
	"fmt"
	"math"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_closestMeetingNode(t *testing.T) {
	for _, tc := range []struct {
		edges []int
		node1 int
		node2 int
		want  int
	}{
		{[]int{5, 3, 1, 0, 2, 3, 4}, 3, 2, 3},
		{[]int{5, 3, 1, 0, 2, 4, 5}, 3, 2, 3},
		{[]int{4, 4, 8, -1, 9, 8, 4, 4, 1, 1}, 5, 6, 1},
		{[]int{2, 2, 3, -1}, 0, 1, 2},
		{[]int{1, 2, -1}, 0, 2, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, closestMeetingNode(tc.edges, tc.node1, tc.node2))
		})
	}
}

func init() {
	debug.SetGCPercent(-1)
}

// https://www.github.com/sebnyberg/leetcode

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	type state struct {
		seen []bool
		node int
		res  int
	}
	newState := func(node, n int) *state {
		return &state{
			seen: make([]bool, n),
			node: node,
			res:  math.MaxInt32,
		}
	}
	n := len(edges)
	var s [2]*state
	s[0] = newState(node1, n)
	s[1] = newState(node2, n)

	// Alternate between the two nodes, moving a node one step forward each time
	for i := 0; s[0].node != -1 || s[1].node != -1; i = i ^ 1 {
		j := i ^ 1
		curr := s[i]
		other := s[j]
		// Important! Only break due to finding a result on every other iteration.
		// Otherwise, we would miss out finding the smallest out of two results.
		if i == 0 && (s[0].res != math.MaxInt32 || s[1].res != math.MaxInt32) {
			break
		}
		if curr.node == -1 {
			continue
		}
		if other.seen[curr.node] { // solution
			curr.res = curr.node
			curr.node = -1
			continue
		}
		if curr.seen[curr.node] { // cycle
			curr.node = -1
			continue
		}
		curr.seen[curr.node] = true
		curr.node = edges[curr.node]
	}
	if res := min(s[0].res, s[1].res); res != math.MaxInt32 {
		return res
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
