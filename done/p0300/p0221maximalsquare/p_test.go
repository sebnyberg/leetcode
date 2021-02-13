package p0221maximalsquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximalSquare(t *testing.T) {
	for _, tc := range []struct {
		matrix [][]byte
		want   int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		},
		{[][]byte{}, 0},
		{[][]byte{{'0'}}, 0},
		{[][]byte{{'1'}}, 1},
		{[][]byte{{'0', '0'}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.matrix), func(t *testing.T) {
			require.Equal(t, tc.want, maximalSquare((tc.matrix)))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m := len(matrix[0])
	var maxArea int
	hist := make([]int, m)
	for _, row := range matrix {
		for j, ch := range row {
			if ch == '0' {
				hist[j] = 0
				continue
			}
			hist[j]++
			maxArea = max(maxArea, 1)
			minHeight := hist[j]
			for k := j + 1; k < m && (k-j+1) <= minHeight && row[k] != '0'; k++ {
				minHeight = min(minHeight, hist[k]+1)
				if minHeight == (k - j + 1) {
					maxArea = max(maxArea, minHeight*(k-j+1))
				}
			}
		}
	}
	return maxArea
}
