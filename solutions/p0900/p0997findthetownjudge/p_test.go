package p0997findthetownjudge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findJudge(t *testing.T) {
	for _, tc := range []struct {
		n     int
		trust [][]int
		want  int
	}{
		{2, [][]int{{1, 2}}, 2},
		{3, [][]int{{1, 3}, {2, 3}}, 3},
		{3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findJudge(tc.n, tc.trust))
		})
	}
}

func findJudge(n int, trust [][]int) int {
	deg := make([]int, n)
	for _, t := range trust {
		deg[t[1]-1]++
		deg[t[0]-1]--
	}
	for idx, count := range deg {
		if count == n-1 {
			return idx + 1
		}
	}
	return -1
}
