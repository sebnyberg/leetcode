package p0576outofboundarypaths

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPaths(t *testing.T) {
	for _, tc := range []struct {
		m, n, maxMove, startRow, startColumn int
		want                                 int
	}{
		{2, 2, 2, 0, 0, 6},
		{1, 3, 3, 0, 1, 12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			require.Equal(t, tc.want, findPaths(tc.m, tc.n, tc.maxMove, tc.startRow, tc.startColumn))
		})
	}
}

const mod = 1_000_000_007

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	if maxMove == 0 {
		return 0
	}
	// dp[k][i][j] = the number of ways in which the ball ends up in spot
	// (i,j) given k rounds
	dp := make([][][]int, maxMove)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}

	dp[0][startRow][startColumn] = 1

	for k := 1; k < maxMove; k++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				for _, nei := range [][2]int{
					{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1},
				} {
					a, b := nei[0], nei[1]
					if a < 0 || b < 0 || a >= m || b >= n {
						continue
					}
					dp[k][i][j] += dp[k-1][a][b]
					dp[k][i][j] %= mod
				}
			}
		}
	}

	var res int
	// for each vertical row, add the number of ways in which the ball could end
	// up in that position
	for i := 0; i < m; i++ {
		for k := 0; k < maxMove; k++ {
			res += dp[k][i][0]
			res += dp[k][i][n-1]
			res %= mod
		}
	}
	// for each horizontal row
	for i := 0; i < n; i++ {
		for k := 0; k < maxMove; k++ {
			res += dp[k][0][i]
			res += dp[k][m-1][i]
			res %= mod
		}
	}
	return res % mod
}
