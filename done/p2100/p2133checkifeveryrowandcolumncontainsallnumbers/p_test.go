package p2133checkifeveryrowandcolumncontainsallnumbers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkValid(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		want   bool
	}{
		{[][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}, true},
		{[][]int{{1, 1, 1}, {1, 2, 3}, {1, 2, 3}}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, checkValid(tc.matrix))
		})
	}
}

func checkValid(matrix [][]int) bool {
	n := len(matrix)
	// each row must contain only unique numbers
	for i := range matrix {
		seen := make([]bool, n+1)
		for _, v := range matrix[i] {
			if v > n || v <= 0 || seen[v] {
				return false
			}
			seen[v] = true
		}
	}
	// each col must contain only unique numbers
	for j := range matrix {
		seen := make([]bool, n+1)
		for i := range matrix {
			v := matrix[i][j]
			if v > n || v <= 0 || seen[v] {
				return false
			}
			seen[v] = true
		}
	}
	return true
}
