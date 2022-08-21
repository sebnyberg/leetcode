package p0840magicsquaresingrid

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_numMagicSquaresInside(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{leetcode.ParseMatrix("[[4,3,8,4],[9,5,1,9],[2,7,6,2]]"), 1},
		{leetcode.ParseMatrix("[[10,3,5],[1,6,11],[7,9,2]]"), 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, numMagicSquaresInside(tc.grid))
		})
	}
}

func numMagicSquaresInside(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	if m < 3 || n < 3 {
		return 0
	}
	isMagic := func(i, j int) int {
		if i+3 > m || j+3 > n {
			return 0
		}
		var seen int
		var rowSum [3]int
		var colSum [3]int
		var diagSum [2]int
		diagSum[0] = grid[i][j] + grid[i+1][j+1] + grid[i+2][j+2]
		diagSum[1] = grid[i][j+2] + grid[i+1][j+1] + grid[i+2][j]
		for ii := i; ii < i+3; ii++ {
			for jj := j; jj < j+3; jj++ {
				if grid[ii][jj] >= 10 || grid[ii][jj] <= 0 {
					return 0
				}
				if seen&(1<<grid[ii][jj]) > 0 {
					return 0
				}
				seen |= (1 << grid[ii][jj])
				rowSum[ii-i] += grid[ii][jj]
				colSum[jj-j] += grid[ii][jj]
			}
		}
		if rowSum[1] != rowSum[0] || rowSum[2] != rowSum[1] {
			return 0
		}
		if colSum[1] != colSum[0] || colSum[2] != colSum[1] {
			return 0
		}
		if diagSum[0] != diagSum[1] {
			return 0
		}
		if diagSum[0] != rowSum[0] || diagSum[0] != colSum[0] {
			return 0
		}
		return 1
	}
	var res int
	for i := 0; i+3 <= m; i++ {
		for j := 0; j+3 <= n; j++ {
			res += isMagic(i, j)
		}
	}
	return res
}
