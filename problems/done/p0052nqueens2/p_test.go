package p0051nqueens

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solveNQueens(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{4, 2},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, totalNQueens(tc.n))
		})
	}
}

func totalNQueens(n int) int {
	// Place N queens on a NxN board, N <= 9
	// Brute-force, recursive
	// Create four integers to track the presence of a queen
	// on the same column, same row, diagonal and antidiagonal (/)
	// There will be max (9-1)*2+1=19 diagonals
	// which fits in an int with a bit-mask.
	var col, diags, antidiags int
	solutions := solve(n, 0, col, diags, antidiags)
	return solutions
}

func getDiag(n, i, j int) int {
	return n - 1 - j + i
}

func getAntiDiag(i, j int) int {
	return i + j
}

func solve(n, row, cols, diags, antidiags int) (solutions int) {
	// all queens have been placed
	if row == n {
		return 1
	}

	for col := 0; col < n; col++ {
		c := 1 << col
		diag := 1 << getDiag(n, row, col)
		antidiag := 1 << getAntiDiag(row, col)
		// check if position is available
		if (c&cols | diag&diags | antidiag&antidiags) > 0 {
			continue
		}
		// place the queen
		cols |= c
		diags |= diag
		antidiags |= antidiag
		solutions += solve(n, row+1, cols, diags, antidiags)
		// remove the queen
		cols &^= c
		diags &^= diag
		antidiags &^= antidiag
	}

	return solutions
}
