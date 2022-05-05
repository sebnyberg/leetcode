package p2050parallelcourses3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumTime(t *testing.T) {
	for _, tc := range []struct {
		n         int
		relations [][]int
		time      []int
		want      int
	}{
		{3, [][]int{{1, 3}, {2, 3}}, []int{3, 2, 5}, 8},
		{5, [][]int{{1, 5}, {2, 5}, {3, 5}, {3, 4}, {4, 5}}, []int{1, 2, 3, 4, 5}, 12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTime(tc.n, tc.relations, tc.time))
		})
	}
}

func minimumTime(n int, relations [][]int, time []int) int {
	// Just do topo sort of courses and count time
	adj := make([][]int, n)
	indeg := make([]int, n)
	for _, rel := range relations {
		a, b := rel[0]-1, rel[1]-1
		adj[a] = append(adj[a], b)
		indeg[b]++
	}
	tovisit := []int{}
	for i, deg := range indeg {
		if deg == 0 {
			tovisit = append(tovisit, i)
		}
	}
	next := []int{}
	maxTimes := make([]int, n)
	for len(tovisit) > 0 {
		next = next[:0]
		for _, u := range tovisit {
			for _, v := range adj[u] {
				maxTimes[v] = max(maxTimes[v], maxTimes[u]+time[u])
				indeg[v]--
				if indeg[v] == 0 {
					next = append(next, v)
				}
			}
		}
		tovisit, next = next, tovisit
	}
	var maxTime int
	for i := range maxTimes {
		maxTime = max(maxTime, maxTimes[i]+time[i])
	}
	return maxTime
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
