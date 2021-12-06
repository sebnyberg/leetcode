package p2097validarrangementofpairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validArrangement(t *testing.T) {
	for _, tc := range []struct {
		pairs [][]int
		want  [][]int
	}{
		{[][]int{{5, 1}, {4, 5}, {11, 9}, {9, 4}}, [][]int{{11, 9}, {9, 4}, {4, 5}, {5, 1}}},
		{[][]int{{1, 3}, {3, 2}, {2, 1}}, [][]int{{1, 3}, {3, 2}, {2, 1}}},
		{[][]int{{1, 2}, {1, 3}, {2, 1}}, [][]int{{1, 2}, {2, 1}, {1, 3}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.pairs), func(t *testing.T) {
			require.Equal(t, tc.want, validArrangement(tc.pairs))
		})
	}
}

func validArrangement(pairs [][]int) [][]int {
	graph := make(map[int][]int)
	degree := make(map[int]int)
	for _, p := range pairs {
		graph[p[0]] = append(graph[p[0]], p[1])
		degree[p[0]]++
		degree[p[1]]--
	}

	start := pairs[0][0]
	for k, deg := range degree {
		if deg == 1 {
			start = k
		}
	}

	ans := make([]int, 0)
	var dfs func(x int)
	dfs = func(x int) {
		for len(graph[x]) > 0 {
			val := graph[x][len(graph[x])-1]
			graph[x] = graph[x][:len(graph[x])-1]
			dfs(val)
		}
		ans = append(ans, x)
	}

	dfs(start)
	for l, r := 0, len(ans)-1; l < r; l, r = l+1, r-1 {
		ans[l], ans[r] = ans[r], ans[l]
	}
	res := make([][]int, 0)
	for i := 0; i < len(ans)-1; i++ {
		res = append(res, []int{ans[i], ans[i+1]})
	}
	return res
}
