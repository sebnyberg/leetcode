package p2022convert1darrayinto2darray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_construct2DArray(t *testing.T) {
	for _, tc := range []struct {
		original []int
		m, n     int
		want     [][]int
	}{
		{[]int{1, 1, 1, 1}, 4, 1, [][]int{{1}, {1}, {1}, {1}}},
		{[]int{1, 2, 3, 4}, 2, 2, [][]int{{1, 2}, {3, 4}}},
		{[]int{1, 2, 3}, 1, 3, [][]int{{1, 2, 3}}},
		{[]int{1, 2}, 1, 1, [][]int{}},
		{[]int{3}, 1, 2, [][]int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.original), func(t *testing.T) {
			require.Equal(t, tc.want, construct2DArray(tc.original, tc.m, tc.n))
		})
	}
}

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return [][]int{}
	}
	res := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		row := make([]int, 0, n)
		for j := 0; j < n; j++ {
			row = append(row, original[i*n+j])
		}
		res = append(res, row)
	}
	return res
}
