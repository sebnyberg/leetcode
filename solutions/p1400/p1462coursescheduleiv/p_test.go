package p1462coursescheduleiv

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_checkIfPrerequisite(t *testing.T) {
	for _, tc := range []struct {
		numCourses    int
		prerequisites [][]int
		queries       [][]int
		want          []bool
	}{
		{
			4,
			leetcode.ParseMatrix("[[2,3],[2,1],[0,3],[0,1]]"),
			leetcode.ParseMatrix("[[0,1],[0,3],[2,3],[3,0],[2,0],[0,2]]"),
			[]bool{true, true, true, false, false, false},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numCourses), func(t *testing.T) {
			require.Equal(t, tc.want, checkIfPrerequisite(tc.numCourses, tc.prerequisites, tc.queries))
		})
	}
}

func checkIfPrerequisite(numCourses int, prerequisites [][]int, queries [][]int) []bool {
	adj := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	prereq := make([][]bool, numCourses)
	for i := range prereq {
		prereq[i] = make([]bool, numCourses)
	}

	for _, p := range prerequisites {
		a := p[0]
		b := p[1]
		adj[a] = append(adj[a], b)
		indeg[b]++
	}
	curr := []int{}
	next := []int{}
	for i, d := range indeg {
		prereq[i][i] = true // a course is a prereq of itself
		if d == 0 {
			curr = append(curr, i)
		}
	}
	for len(curr) > 0 {
		next = next[:0]
		for _, x := range curr {
			for _, y := range adj[x] {
				for i := range prereq[x] {
					prereq[y][i] = prereq[y][i] || prereq[x][i]
				}
				indeg[y]--
				if indeg[y] == 0 {
					next = append(next, y)
				}
			}
		}
		curr, next = next, curr
	}
	res := make([]bool, len(queries))
	for i, q := range queries {
		a := q[0]
		b := q[1]
		if prereq[b][a] {
			res[i] = true
		}
	}
	return res
}
