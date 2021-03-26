package p0210courseschedule2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findOrder(t *testing.T) {
	for _, tc := range []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{3, [][]int{{0, 1}, {0, 2}, {1, 2}}, []int{2, 1, 0}},
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 2, 1, 3}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.numCourses, tc.prerequisites), func(t *testing.T) {
			require.Subset(t, tc.want, findOrder(tc.numCourses, tc.prerequisites))
		})
	}
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, req := range prerequisites {
		adj[req[1]] = append(adj[req[1]], req[0])
		indegree[req[0]]++
	}

	todos := make([]int, 0)
	seen := make([]bool, numCourses)
	res := make([]int, 0)
	for i, deg := range indegree {
		if deg == 0 {
			todos = append(todos, i)
			seen[i] = true
			res = append(res, i)
			numCourses--
		}
	}

	for len(todos) > 0 {
		newTodos := make([]int, 0)
		for _, todoIdx := range todos {
			for _, adjIdx := range adj[todoIdx] {
				indegree[adjIdx]--
				if seen[adjIdx] || indegree[adjIdx] != 0 {
					continue
				}
				seen[adjIdx] = true
				newTodos = append(newTodos, adjIdx)
				res = append(res, adjIdx)
				numCourses--
			}
		}
		todos = newTodos
	}

	if numCourses > 0 {
		return nil
	}

	return res
}
