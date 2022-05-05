package p1820maxnacceptedinvitations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumInvitations(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 1, 0},
			}, 9,
		},
		{[][]int{{1, 1, 1}, {1, 0, 1}, {0, 0, 1}}, 3},
		{[][]int{{1, 0, 1, 0}, {1, 0, 0, 0}, {0, 0, 1, 0}, {1, 1, 1, 0}}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, maximumInvitations(tc.grid))
		})
	}
}

func maximumInvitations(grid [][]int) int {
	nboys, ngirls := len(grid), len(grid[0])
	matches := make([]int, ngirls)
	for i := range matches {
		matches[i] = -1 // no match so far
	}

	count := 0
	// For each boy
	for boy := 0; boy < nboys; boy++ {
		seen := make([]bool, ngirls)
		// Match with each girl, forming an augmented path
		if dfs(matches, seen, grid, ngirls, boy) {
			count++
		}
	}
	return count
}

func dfs(matches []int, seen []bool, grid [][]int, ngirls int, boy int) bool {
	for girl := 0; girl < ngirls; girl++ {
		if grid[boy][girl] != 1 || seen[girl] {
			continue
		}
		seen[girl] = true
		// If this girl has not matched with any boy before,
		// or there is a path following this match which yields a new match,
		// update the matching to be between this girl and boy
		if matches[girl] == -1 || dfs(matches, seen, grid, ngirls, matches[girl]) {
			matches[girl] = boy
			return true
		}
	}
	return false
}
