package p1074nsubmatricesthatsumtotarget

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSubmatrixSumTarget(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]int
		target int
		want   int
	}{
		{[][]int{{0, 1, 0}, {1, 1, 1}, {0, 1, 0}}, 0, 4},
		{[][]int{{1, -1}, {-1, 1}}, 0, 5},
		{[][]int{{904}}, 0, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, numSubmatrixSumTarget(tc.matrix, tc.target))
		})
	}
}

func numSubmatrixSumTarget(matrix [][]int, target int) int {
	// Four variables:
	// xStart, xLen, yStart, yLen
	// Try brute-force first
	m, n := len(matrix), len(matrix[0])
	var count int
	for yStart := 0; yStart < m; yStart++ {
		for yLen := 1; yLen <= m-yStart; yLen++ {
			for xStart := 0; xStart < n; xStart++ {
				for xLen := 1; xLen <= n-xStart; xLen++ {
					res := sum(matrix, xStart, xLen, yStart, yLen)
					if res == target {
						count++
					}
				}
			}
		}
	}
	return count
}

func sum(matrix [][]int, xStart, xLen, yStart, yLen int) int {
	var res int
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			res += matrix[yStart+y][xStart+x]
		}
	}
	return res
}
