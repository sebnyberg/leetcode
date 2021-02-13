package p0059spiralmatrix2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateMatrix(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, generateMatrix(tc.n))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	cur := 1
	var i, j int
	for {
		// right
		if res[i][j] != 0 {
			return res
		}
		for res[i][min(j, n-1)] == 0 {
			res[i][j] = cur
			cur++
			j++
		}
		j--
		i++
		if res[min(i, n-1)][j] != 0 {
			return res
		}
		for res[min(i, n-1)][j] == 0 {
			res[i][j] = cur
			cur++
			i++
		}
		i--
		j--
		if res[i][j] != 0 {
			return res
		}
		for res[i][max(0, j)] == 0 {
			res[i][j] = cur
			cur++
			j--
		}
		j++
		i--
		if res[i][j] != 0 {
			return res
		}
		for res[i][j] == 0 {
			res[i][j] = cur
			cur++
			i--
		}
		i++
		j++
	}
}
