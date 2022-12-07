package p1039minimumscoretriangulationofpolygon

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minScoreTriangulation(t *testing.T) {
	for i, tc := range []struct {
		values []int
		want   int
	}{
		{[]int{1, 3, 1, 4, 1, 5}, 13},
		{[]int{3, 7, 4, 5}, 144},
		{[]int{35, 73, 90, 27, 71, 80, 21, 33, 33, 13, 48, 12, 68, 70, 80, 36, 66, 3, 70, 58}, 140295},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minScoreTriangulation(tc.values))
		})
	}
}

func minScoreTriangulation(values []int) int {
	// Fuck me... spend so many iterations trying to figure out how to most
	// efficiently "cut" the slice into parts when the easiest approach is to
	// select the first side and pair it with any other point.
	mem := make(map[[2]int]int)
	n := len(values)
	res := dfs(mem, values, 0, n-1)
	return res
}

func dfs(mem map[[2]int]int, values []int, i, j int) int {
	key := [2]int{i, j}
	if v, exists := mem[key]; exists {
		return v
	}
	if j-i <= 1 {
		return 0
	}
	// Match first side to any other point
	res := math.MaxInt32
	for k := i + 1; k < j; k++ {
		v := values[i] * values[j] * values[k]
		left := dfs(mem, values, i, k)
		right := dfs(mem, values, k, j)
		res = min(res, v+left+right)
	}

	mem[key] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
