package p2360longestcycleinagraph

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCycle(t *testing.T) {
	for _, tc := range []struct {
		edges []int
		want  int
	}{
		{[]int{-1, 4, -1, 2, 0, 4}, -1},
		{[]int{3, 3, 4, 2, 3}, 3},
		{[]int{2, -1, 3, 1}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.edges), func(t *testing.T) {
			require.Equal(t, tc.want, longestCycle(tc.edges))
		})
	}
}

// https://www.github.com/sebnyberg/leetcode

func longestCycle(edges []int) int {
	n := len(edges)
	times := make([]int, n)
	visited := func(i int) bool {
		return times[i] > 0
	}
	t := 1
	res := -1
	for i := range edges {
		t0 := t
		if visited(i) {
			continue
		}
		times[i] = t
		t++
		// Iterate until end or reaching a previously visited node
		first, second := i, edges[i]
		for second != -1 && !visited(second) {
			times[second] = t
			first = second
			second = edges[second]
			t++
		}
		if second != -1 && times[second] >= t0 {
			res = max(res, times[first]-times[second]+1)
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
