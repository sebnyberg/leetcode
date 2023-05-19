package p0562longestline

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestLine(t *testing.T) {
	for _, tc := range []struct {
		M    [][]int
		want int
	}{
		{[][]int{
			{0, 0, 1, 1},
			{0, 0, 1, 0},
			{0, 1, 0, 1},
		}, 3},
		{[][]int{
			{0, 1, 1, 0},
			{0, 1, 1, 0},
			{0, 0, 0, 1},
		}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.M), func(t *testing.T) {
			require.Equal(t, tc.want, longestLine(tc.M))
		})
	}
}

func longestLine(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	m, n := len(M), len(M[0])
	var maxRunLen int

	// rows
	for i := range M {
		var curLen int
		for j := range M[i] {
			switch {
			case M[i][j] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[i][j] == 0:
				curLen = 0
			}
		}
	}

	// cols
	for j := range M[0] {
		var curLen int
		for i := range M {
			switch {
			case M[i][j] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[i][j] == 0:
				curLen = 0
			}
		}
	}

	validIndex := func(i, j int) bool {
		return i >= 0 && i < m && j >= 0 && j < n
	}

	// Diagonal
	for j := -m + 1; j < n; j++ {
		var curLen int
		for k := 0; k < m; k++ {
			if !validIndex(k, j+k) {
				continue
			}
			switch {
			case M[k][j+k] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[k][j+k] == 0:
				curLen = 0
			}
		}
	}

	// Antidiagonal
	for j := n - 1 + m - 1; j >= 0; j-- {
		var curLen int
		for k := 0; k < m; k++ {
			if !validIndex(k, j-k) {
				continue
			}
			switch {
			case M[k][j-k] == 1:
				curLen++
				maxRunLen = max(maxRunLen, curLen)
			case M[k][j-k] == 0:
				curLen = 0
			}
		}
	}

	return maxRunLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
