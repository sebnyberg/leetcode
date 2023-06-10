package p2732findagoodsubsetofthematrix

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_goodSubsetofBinaryMatrix(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want []int
	}{
		{leetcode.ParseMatrix("[[0,1,1,0],[0,0,0,1],[1,1,1,1]]"), []int{0, 1}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, goodSubsetofBinaryMatrix(tc.grid))
		})
	}
}

func goodSubsetofBinaryMatrix(grid [][]int) []int {
	// Shit problem tbh. Apparently when n < 6, there exists a solution if a
	// selection of two rows have mutually exclusive bits, so 700 people just
	// tried finding a solution with 1 or 2 numbers and got accepted.
	//
	// I on the other hand tried to figure out a way that works for n >= 6 etc,
	// i.e. not relying on the random assumption that two rows will somehow
	// always work.
	var rows []int
	for i := range grid {
		var x int
		for j := range grid[i] {
			x <<= 1
			x += grid[i][j]
		}
		if x == 0 {
			return []int{i}
		}
		rows = append(rows, x)
	}
	for i := range rows {
		for j := i + 1; j < len(rows); j++ {
			if rows[i]&rows[j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}
