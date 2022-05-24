package p2267checkifthereisavalidpatenthesesstringpath

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hasValidPath(t *testing.T) {
	for _, tc := range []struct {
		grid [][]byte
		want bool
	}{
		{
			[][]byte{
				[]byte("((("),
				[]byte(")()"),
				[]byte("(()"),
				[]byte("(()"),
			},
			true,
		},
		{
			[][]byte{
				[]byte("))"),
				[]byte("(("),
			},
			false,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, hasValidPath(tc.grid))
		})
	}
}

func hasValidPath(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])
	mem := make(map[[3]int]bool, m*n)
	if grid[m-1][n-1] != ')' || grid[0][0] != '(' {
		return false
	}
	res := visit(mem, grid, m, n, m-1, n-1, 0)
	return res
}

// visit returns true if it is possible to end in the position with wantOpen
// number of opening parenthesis
func visit(mem map[[3]int]bool, grid [][]byte, m, n, i, j, wantOpen int) bool {
	k := [3]int{i, j, wantOpen}
	if v, exists := mem[k]; exists {
		return v
	}
	if i == 0 && j == 0 {
		return wantOpen == 1
	}
	if i < 0 || j < 0 || wantOpen < 0 {
		return false
	}
	// If this position has an opening parenthesis, then wantOpen decreases
	if grid[i][j] == '(' {
		wantOpen--
	} else {
		wantOpen++
	}
	res := visit(mem, grid, m, n, i-1, j, wantOpen) ||
		visit(mem, grid, m, n, i, j-1, wantOpen)
	mem[k] = res
	return mem[k]
}
