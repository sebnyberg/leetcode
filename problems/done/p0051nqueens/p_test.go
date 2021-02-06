package p0051nqueens

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_solveNQueens(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want [][]string
	}{
		// {4, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q..."}}},
		{1, [][]string{{"Q"}}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, solveNQueens(tc.n))
		})
	}
}

func getDiag(n, i, j int) int {
	return n - 1 - j + i
}

func getAntiDiag(i, j int) int {
	return i + j
}

type pos struct {
	i int
	j int
}

func solveNQueens(n int) [][]string {
	// Place N queens on a NxN board, N <= 9
	// Brute-force, recursive
	// Create four integers to track the presence of a queen
	// on the same column, same row, diagonal and antidiagonal (/)
	// There will be max (9-1)*2+1=19 diagonals
	// which fits in an int with a bit-mask.
	var col, diags, antiDiags int
	solutions := make([][]pos, 0)
	cur := make([]pos, 0)
	solve(&solutions, &cur, n, 0, col, diags, antiDiags)
	rowStr := strings.Repeat(".", n)
	solStrs := make([][]string, len(solutions))
	for i, sol := range solutions {
		solStrs[i] = make([]string, n)
		for _, pos := range sol {
			solStrs[i][pos.i] = rowStr
			solStrs[i][pos.i] = solStrs[i][pos.i][:pos.j] + string('Q') + solStrs[i][pos.i][pos.j+1:]
		}
	}
	return solStrs
}
func solve(solutions *[][]pos, cur *[]pos, n, row, cols, diags, antiDiags int) {
	if row == n {
		// If done, add it to the list of solutions and exit
		*solutions = append(*solutions, make([]pos, n))
		copy((*solutions)[len(*solutions)-1], *cur)
		return
	}

	for col := 0; col < n; col++ {
		c := 1 << col
		diag := 1 << getDiag(n, row, col)
		antiDiag := 1 << getAntiDiag(row, col)
		// check if position is available
		if (c&cols | diag&diags | antiDiag&antiDiags) > 0 {
			continue
		}
		// pick the position
		*cur = append(*cur, pos{row, col})
		cols |= c
		diags |= diag
		antiDiags |= antiDiag
		solve(solutions, cur, n, row+1, cols, diags, antiDiags)
		cols &^= c
		diags &^= diag
		antiDiags &^= antiDiag
		*cur = (*cur)[:len(*cur)-1]
	}
}
