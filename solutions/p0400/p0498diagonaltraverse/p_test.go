package p0498diagonaltraverse

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findDiagonalOrder(t *testing.T) {
	for _, tc := range []struct {
		mat  [][]int
		want []int
	}{
		{leetcode.ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]"), []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		{leetcode.ParseMatrix("[[1,2],[3,4]]"), []int{1, 2, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat), func(t *testing.T) {
			require.Equal(t, tc.want, findDiagonalOrder(tc.mat))
		})
	}
}

func findDiagonalOrder(mat [][]int) []int {
	dirs := [][]int{{-1, 1}, {1, -1}}
	var i, j int
	m, n := len(mat), len(mat[0])
	var dir int
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	res := make([]int, 0, m*n)
	for count := 0; count != m*n; count++ {
		res = append(res, mat[i][j])
		ii, jj := i+dirs[dir][0], j+dirs[dir][1]
		if ok(ii, jj) {
			i, j = ii, jj
			continue
		}
		// A lot of cases here..
		if dir == 0 { // up
			if ii < 0 && jj < n {
				j++ // Move right
			} else { // Top right corner, or right side
				i++ // Move down
			}
		} else { // down
			if ii < m && jj < 0 {
				i++ // Move down instead
			} else { // Bottom left corner or bottom side
				j++ // Move right
			}
		}
		dir = (dir + 1) & 1
	}
	return res
}
