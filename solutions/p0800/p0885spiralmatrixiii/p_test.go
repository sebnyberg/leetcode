package p0885spiralmatrixiii

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_spiralMatrixIII(t *testing.T) {
	for i, tc := range []struct {
		rows   int
		cols   int
		rStart int
		cStart int
		want   [][]int
	}{
		{5, 6, 1, 4, leetcode.ParseMatrix("[[0,0],[0,1],[0,2],[0,3]]")},
		{1, 4, 0, 0, leetcode.ParseMatrix("[[0,0],[0,1],[0,2],[0,3]]")},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, spiralMatrixIII(tc.rows, tc.cols, tc.rStart, tc.cStart))
		})
	}
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	i, j := rStart, cStart
	count := 1
	res := [][]int{{i, j}}
	const (
		hij = 0
		hii = 1
		loj = 2
		loi = 3
	)
	limits := []int{j + 1, i + 1, j - 1, i - 1}
	deltas := []int{1, 1, -1, -1}
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	k := 0
	for count < rows*cols {
		ii, jj := i+dirs[k][0], j+dirs[k][1]
		if ii < limits[loi] || ii > limits[hii] || jj < limits[loj] || jj > limits[hij] {
			limits[k] += deltas[k]
			k = (k + 1) % 4
			continue
		}
		if ii >= 0 && ii < rows && jj >= 0 && jj < cols {
			res = append(res, []int{ii, jj})
			count++
		}
		i = ii
		j = jj
	}
	return res
}
