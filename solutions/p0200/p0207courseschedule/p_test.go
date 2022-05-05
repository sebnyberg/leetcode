package p0207courseschedule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canFinish(t *testing.T) {
	for _, tc := range []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numCourses), func(t *testing.T) {
			require.Equal(t, tc.want, canFinish(tc.numCourses, tc.prerequisites))
		})
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	degree := make([]int, numCourses)
	for _, p := range prerequisites {
		adj[p[1]] = append(adj[p[1]], p[0])
		degree[p[0]]++
	}

	independent := make([]int, 0)
	for i, d := range degree {
		if d == 0 {
			independent = append(independent, i)
		}
	}

	for len(independent) > 0 {
		courseIndex := independent[len(independent)-1]
		independent = independent[:len(independent)-1]
		numCourses--
		for _, near := range adj[courseIndex] {
			degree[near]--
			if degree[near] == 0 {
				independent = append(independent, near)
			}
		}
	}
	return numCourses == 0
}
