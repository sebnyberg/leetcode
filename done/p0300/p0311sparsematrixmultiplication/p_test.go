package p0311sparsematrixmultiplication

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_multiply(t *testing.T) {
	for _, tc := range []struct {
		mat1 [][]int
		mat2 [][]int
		want [][]int
	}{
		{
			[][]int{
				{1, 0, 0},
				{-1, 0, 3},
			}, [][]int{
				{7, 0, 0},
				{0, 0, 0},
				{0, 0, 1},
			}, [][]int{
				{7, 0, 0},
				{-7, 0, 3},
			},
		},
		{
			[][]int{{0}},
			[][]int{{0}},
			[][]int{{0}},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.mat1), func(t *testing.T) {
			require.Equal(t, tc.want, multiply(tc.mat1, tc.mat2))
		})
	}
}

func multiply(mat1 [][]int, mat2 [][]int) [][]int {
	m := len(mat1)
	n := len(mat2[0])
	nn := len(mat2)
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for k := 0; k < nn; k++ {
				res[i][j] += mat1[i][k] * mat2[k][j]
			}
		}
	}
	return res
}
