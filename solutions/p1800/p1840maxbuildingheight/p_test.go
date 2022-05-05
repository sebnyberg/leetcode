package p1840maxbuildingheight

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxBuilding(t *testing.T) {
	for _, tc := range []struct {
		n            int
		restrictions [][]int
		want         int
	}{
		{10, [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}, 5},
		{10, [][]int{{6, 0}, {5, 2}, {7, 0}, {9, 1}, {2, 4}, {3, 4}, {4, 0}, {8, 2}, {10, 0}}, 1},
		{10, [][]int{{8, 5}, {9, 0}, {6, 2}, {4, 0}, {3, 2}, {10, 0}, {5, 3}, {7, 3}, {2, 4}}, 2},
		{5, [][]int{{2, 1}, {4, 1}}, 2},
		{6, [][]int{}, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maxBuilding(tc.n, tc.restrictions))
		})
	}
}

const (
	pos    = 0
	height = 1
)

func maxBuilding(n int, restrictions [][]int) int {
	// Too annoying to type
	builds := restrictions

	// Add implicit restriction for the first building and sort by position
	builds = append(builds, []int{1, 0})
	sort.Slice(builds, func(i, j int) bool {
		return builds[i][pos] < builds[j][pos]
	})

	// Add auxiliary last building if not part of the list already
	if builds[len(builds)-1][pos] != n {
		builds = append(builds, []int{n, n + 1})
	}

	// From left to right, calculate the height each building would have
	// if left unhindered by the right side
	for i := range builds {
		if i == 0 {
			continue
		}
		builds[i][height] = min(
			builds[i][height],
			builds[i-1][height]+builds[i][pos]-builds[i-1][pos],
		)
	}

	// Then from right to left
	nbuilds := len(builds)
	for i := nbuilds - 2; i >= 0; i-- {
		builds[i][height] = min(
			builds[i][height],
			builds[i+1][height]+builds[i+1][pos]-builds[i][pos],
		)
	}

	// Heights of all buildings are correct according to restrictions
	// Iterate over pairs of buildings, checking if there is a long
	// gap that could be utilized to build higher than the restricted buildings
	maxHeight := 0
	for i := range builds {
		if i == 0 {
			continue
		}
		maxHeight = max(maxHeight, builds[i][height])

		heightDiff := abs(builds[i][height] - builds[i-1][height])
		buildingsBetween := builds[i][pos] - builds[i-1][pos] - 1
		extra := buildingsBetween - heightDiff + 1
		if extra == 0 {
			continue
		}
		// there is a long gap between buildings, and it can be utilized to
		// build higher buildings
		maxHeight = max(maxHeight, max(builds[i][height], builds[i-1][height])+extra/2)
	}

	return maxHeight
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
