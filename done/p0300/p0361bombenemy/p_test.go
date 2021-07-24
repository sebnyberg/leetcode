package p0361bombenemy

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxKilledEnemies(t *testing.T) {
	for _, tc := range []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				{'0', 'E', '0', '0'},
				{'E', '0', 'W', 'E'},
				{'0', 'E', '0', '0'},
			}, 3,
		},
		{
			[][]byte{
				{'W', 'W', 'W'},
				{'0', '0', '0'},
				{'E', 'E', 'E'},
			}, 1,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, maxKilledEnemies(tc.grid))
		})
	}
}

func maxKilledEnemies(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	// This is a classic "scan right, scan left" exercise, just in 2D
	// The idea is to scan "enemiesAbove", "enemiesBelow", etc uint16o matrices
	// then find the max combination of all four matrices
	above := make([][]uint16, m)
	below := make([][]uint16, m)
	left := make([][]uint16, m)
	right := make([][]uint16, m)
	for i := range above {
		above[i] = make([]uint16, n)
		below[i] = make([]uint16, n)
		left[i] = make([]uint16, n)
		right[i] = make([]uint16, n)
	}

	// Scan from top to bottom, scanning enemies above
	for row := 1; row < m; row++ {
		for col := 0; col < n; col++ {
			above[row][col] = above[row-1][col]
			switch grid[row-1][col] {
			case 'E':
				above[row][col]++
			case 'W':
				above[row][col] = 0
			}
		}
	}
	// Scan from bottom to top
	for row := m - 2; row >= 0; row-- {
		for col := 0; col < n; col++ {
			below[row][col] = below[row+1][col]
			switch grid[row+1][col] {
			case 'E':
				below[row][col]++
			case 'W':
				below[row][col] = 0
			}
		}
	}

	// Left to right
	for col := 1; col < n; col++ {
		for row := 0; row < m; row++ {
			left[row][col] = left[row][col-1]
			switch grid[row][col-1] {
			case 'E':
				left[row][col]++
			case 'W':
				left[row][col] = 0
			}
		}
	}

	// Right to left
	for col := n - 2; col >= 0; col-- {
		for row := 0; row < m; row++ {
			right[row][col] = right[row][col+1]
			switch grid[row][col+1] {
			case 'E':
				right[row][col]++
			case 'W':
				right[row][col] = 0
			}
		}
	}

	// Find max sum
	var maxKilled uint16
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '0' {
				maxKilled = max(maxKilled,
					above[i][j]+below[i][j]+left[i][j]+right[i][j],
				)
			}
		}
	}
	return int(maxKilled)
}

func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}
