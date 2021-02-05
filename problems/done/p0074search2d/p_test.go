package p0074search2d

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchMatrix(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		target int
		want   bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.matrix, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, searchMatrix(tc.matrix, tc.target))
		})
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	var r int
	n := len(matrix)
	m := len(matrix[0])
	for {
		if r == n {
			return false
		}
		lo, hi := matrix[r][0], matrix[r][m-1]
		if target == lo || target == hi {
			return true
		}
		if target > lo && target < hi {
			break
		}
		r++
	}

	// Search for result within row
	row := matrix[r]
	lo := 0
	hi := m - 1
	for {
		mid := lo + ((hi - lo) / 2)
		n = row[mid]
		if n == target {
			return true
		}
		if lo == mid || hi == mid {
			return false
		}
		if n < target {
			lo = mid
			continue
		}
		if n > target {
			hi = mid
		}
	}
}
