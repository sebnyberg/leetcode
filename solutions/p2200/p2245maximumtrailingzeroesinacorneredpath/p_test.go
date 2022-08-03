package p2245maximumtrailingzeroesinacorneredpath

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxTrailingZeros(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[23,17,15,3,20],[8,1,20,27,11],[9,4,6,2,21],[40,9,1,10,6],[22,7,4,5,3]]"),
			3,
		},
		{
			leetcode.ParseMatrix("[[4,3,2],[7,6,1],[8,8,8]]"),
			0,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, maxTrailingZeros(tc.grid))
		})
	}
}

func maxTrailingZeros(grid [][]int) int {
	// Can a sum of any numbers result in a leading zero if it does not contain a
	// leading zero itself?
	// Yeah.. but only 2*5 = 10

	// What about other numbers? Do they hurt?
	// No.

	// So. The goal is to find a horizontal + vertical path that includes as many
	// 2s, 5s, and numbers with leading zeroes as possible.

	// The problem is then the 2s and 5s.
	// Essentially, the matrix could be normalized to contain only 10s, 2s, and
	// 5s. The goal is to hunt for as many 10s + min(2s,5s) as possible.

	countFactors := func(x int) (int, int) {
		var a, b int
		for ; x%2 == 0; x /= 2 {
			a++
		}
		for ; x%5 == 0; x /= 5 {
			b++
		}
		return a, b
	}

	m, n := len(grid), len(grid[0])
	preRows := make([][][2]int, m)
	for r := range preRows {
		preRows[r] = make([][2]int, n+1)
		for c, v := range grid[r] {
			preRows[r][c+1] = preRows[r][c]
			a, b := countFactors(v)
			preRows[r][c+1][0] += a
			preRows[r][c+1][1] += b
		}
	}
	preCols := make([][][2]int, n)
	for c := range preCols {
		preCols[c] = make([][2]int, m+1)
		for r := 0; r < m; r++ {
			v := grid[r][c]
			preCols[c][r+1] = preCols[c][r]
			a, b := countFactors(v)
			preCols[c][r+1][0] += a
			preCols[c][r+1][1] += b
		}
	}
	const (
		dirLeft  = 0
		dirRight = 1
		dirUp    = 2
		dirDown  = 3
	)
	calc := func(r, c int, dir1, dir2 int) int {
		a, b := countFactors(grid[r][c])
		counts := [2]int{a, b}
		for _, d := range []int{dir1, dir2} {
			var delta [2]int
			switch d {
			case dirLeft:
				delta = preRows[r][c]
			case dirUp:
				delta = preCols[c][r]
			case dirRight:
				delta = preRows[r][n]
				for i := 0; i < 2; i++ {
					delta[i] -= preRows[r][c+1][i]
				}
			case dirDown:
				delta = preCols[c][m]
				for i := 0; i < 2; i++ {
					delta[i] -= preCols[c][r+1][i]
				}
			}
			for i := 0; i < 2; i++ {
				counts[i] += delta[i]
			}
		}
		return min(counts[0], counts[1])
	}

	// Visit each cell, checking the maximum possible result.
	var res int
	for r := range grid {
		for c := range grid[r] {
			// There are six kinds of combinations
			res = max(res, calc(r, c, dirLeft, dirUp))
			res = max(res, calc(r, c, dirLeft, dirRight))
			res = max(res, calc(r, c, dirLeft, dirDown))
			res = max(res, calc(r, c, dirUp, dirRight))
			res = max(res, calc(r, c, dirUp, dirDown))
			res = max(res, calc(r, c, dirRight, dirDown))
		}
	}
	return res
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
