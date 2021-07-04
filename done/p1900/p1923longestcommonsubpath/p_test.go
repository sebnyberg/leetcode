package p1923longestcommonsubpath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommonSubpath(t *testing.T) {
	for _, tc := range []struct {
		n     int
		paths [][]int
		want  int
	}{
		{5, [][]int{{0, 1, 2, 3, 4}, {2, 3, 4}, {4, 0, 1, 2, 3}}, 2},
		{3, [][]int{{0}, {1}, {2}}, 0},
		{5, [][]int{{0, 1, 2, 3, 4}, {4, 3, 2, 1, 0}}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, longestCommonSubpath(tc.n, tc.paths))
		})
	}
}

func longestCommonSubpath(n int, paths [][]int) int {
	// Idea:
	//
	// Set first path as the match path.
	// Then for each path i > 0
	// 1. Match using rolling hash. If no match,
	// 2. Update match to longest shared sequence between match and path
	return 0
}
