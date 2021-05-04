package p0302smallestrectangleenclosingblackpixels

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minArea(t *testing.T) {
	for _, tc := range []struct {
		image []string
		x     int
		y     int
		want  int
	}{
		{[]string{"0010", "0110", "0100"}, 0, 2, 6},
		{[]string{"1"}, 0, 0, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.image), func(t *testing.T) {
			bs := make([][]byte, len(tc.image))
			for i := range tc.image {
				bs[i] = []byte(tc.image[i])
			}
			require.Equal(t, tc.want, minArea(bs, tc.x, tc.y))
		})
	}
}

func minArea(image [][]byte, x int, y int) int {
	m, n := len(image), len(image[0])

	// Find minY and maxY
	minY := 1<<31 - 1
	maxY := 0
	for row := range image {
		for _, n := range image[row] {
			if n == '1' {
				minY = min(minY, row)
				maxY = max(maxY, row)
				goto ContinueRowSearch
			}
		}
	ContinueRowSearch:
	}

	// For each column
	minX := 1<<31 - 1
	maxX := 0
	for col := 0; col < n; col++ {
		for row := 0; row < m; row++ {
			if image[row][col] == '1' {
				minX = min(minX, col)
				maxX = max(maxX, col)
				goto ContinueColSearch
			}
		}
	ContinueColSearch:
	}

	return (maxX - minX + 1) * (maxY - minY + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
