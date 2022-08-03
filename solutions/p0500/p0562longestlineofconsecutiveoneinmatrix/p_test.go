package p0562longestlineofconsecutiveoneinmatrix

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestLine(t *testing.T) {
	for _, tc := range []struct {
		mat  [][]int
		want int
	}{
		{leetcode.ParseMatrix("[[0,0,0],[0,1,0],[0,0,0]]"), 1},
		{leetcode.ParseMatrix("[[0,1,1,0],[0,1,1,0],[0,0,0,1]]"), 3},
		{leetcode.ParseMatrix("[[1,1,1,1],[0,1,1,0],[0,0,0,1]]"), 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, longestLine(tc.mat))
		})
	}
}

func longestLine(mat [][]int) int {
	diag := []int{1, 1}
	antiDiag := []int{1, -1}
	row := []int{0, 1}
	col := []int{1, 0}

	m, n := len(mat), len(mat[0])

	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	scan := func(i, j int, dir []int) int {
		var maxCount int
		var count int
		for ; ok(i, j); i, j = i+dir[0], j+dir[1] {
			if mat[i][j] != 1 {
				count = 0
				continue
			}
			count++
			maxCount = max(maxCount, count)
		}
		return maxCount
	}

	var res int
	for j := 0; j < n; j++ {
		res = max(res, scan(0, j, col))      // Scan column
		res = max(res, scan(0, j, diag))     // Scan diagonal
		res = max(res, scan(0, j, antiDiag)) // Scan anti-diag
	}

	for i := 0; i < m; i++ {
		res = max(res, scan(i, 0, row))        // Scan row
		res = max(res, scan(i, 0, diag))       // Scan diagonal
		res = max(res, scan(i, n-1, antiDiag)) // Scan anti-diag
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
