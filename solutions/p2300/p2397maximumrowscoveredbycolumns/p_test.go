package p2397maximumrowscoveredbycolumns

import (
	"fmt"
	"math/bits"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_maxmumRows(t *testing.T) {
	for _, tc := range []struct {
		mat  [][]int
		cols int
		want int
	}{
		{leetcode.ParseMatrix("[[0,0,0],[1,0,1],[0,1,1],[0,0,1]]"), 2, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, maximumRows(tc.mat, tc.cols))
		})
	}
}

func maximumRows(mat [][]int, cols int) int {
	// There aren't that many ways to select columns, so we can just try them all.
	m := len(mat)
	n := len(mat[0])
	rows := make([]int, m)
	for i := range mat {
		for _, v := range mat[i] {
			rows[i] <<= 1
			rows[i] += v
		}
	}
	var maxRes int
	for x := 1; x < (1 << n); x++ {
		if bits.OnesCount(uint(x)) != cols {
			continue
		}
		// Count rows
		var res int
		for _, row := range rows {
			if row&x == row {
				res++
			}
		}
		maxRes = max(maxRes, res)
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
