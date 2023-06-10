package p1391checkifthereisavalidpathinagrid

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_hasValidPath(t *testing.T) {
	for i, tc := range []struct {
		grid [][]int
		want bool
	}{
		{
			leetcode.ParseMatrix("[[2,4,3],[6,5,2]]"),
			true,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, hasValidPath(tc.grid))
		})
	}
}

func hasValidPath(grid [][]int) bool {
	streets := []map[[2]int][2]int{
		1: {
			{0, 1}:  {0, 1},
			{0, -1}: {0, -1},
		},
		2: {
			{1, 0}:  {1, 0},
			{-1, 0}: {-1, 0},
		},
		3: {
			{0, 1}:  {1, 0},
			{-1, 0}: {0, -1},
		},
		4: {
			{-1, 0}: {0, 1},
			{0, -1}: {1, 0},
		},
		5: {
			{0, 1}: {-1, 0},
			{1, 0}: {0, -1},
		},
		6: {
			{1, 0}:  {0, 1},
			{0, -1}: {-1, 0},
		},
	}
	m := len(grid)
	n := len(grid[0])
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}

	check := func(dir [2]int) bool {
		seen := make([][]bool, m)
		for i := range seen {
			seen[i] = make([]bool, n)
		}
		i := 0
		j := 0
		seen[i][j] = true
		for {
			if i == m-1 && j == n-1 {
				return true
			}
			i += dir[0]
			j += dir[1]
			if !ok(i, j) || seen[i][j] {
				return false
			}
			nextDir, exists := streets[grid[i][j]][dir]
			if !exists {
				return false
			}
			dir = nextDir
			seen[i][j] = true
		}
	}

	for _, dir := range streets[grid[0][0]] {
		if check(dir) {
			return true
		}
	}
	return false
}
