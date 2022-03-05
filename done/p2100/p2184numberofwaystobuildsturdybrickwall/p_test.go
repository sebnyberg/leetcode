package p2184numberofwaystobuildsturdybrickwall

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_buildWall(t *testing.T) {
	for _, tc := range []struct {
		height int
		width  int
		bricks []int
		want   int
	}{
		{2, 3, []int{1, 2}, 2},
		{1, 1, []int{5}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.height), func(t *testing.T) {
			require.Equal(t, tc.want, buildWall(tc.height, tc.width, tc.bricks))
		})
	}
}

func buildWall(height int, width int, bricks []int) int {
	// skip any bricks wider than width
	var j int
	for i := 0; i < len(bricks); i++ {
		if bricks[i] <= width {
			bricks[j] = bricks[i]
			j++
		}
	}
	bricks = bricks[:j]
	if len(bricks) == 0 {
		return 0
	}
	// Find all possible configurations for a wall
	var masks []int
	getConfigs(bricks, 0, 0, width, &masks)

	// Then find all possible row configurations that meet the criteria.
	// Given a certain configuration, how many subconfigurations below that row
	// are valid? Well.. any mask which does not overlap with a selected mask
	// is valid. As such, we can count how many configurations are valid for
	// each possible mask.
	var sturdyWalls [101][1024]int
	res := findSturdyWalls(&sturdyWalls, height, 0, masks)
	return res
}

const mod = 1e9 + 7

func findSturdyWalls(sturdyWalls *[101][1024]int, height, prevRow int, masks []int) int {
	if height == 0 {
		return 1
	}
	if sturdyWalls[height][prevRow] != 0 {
		return sturdyWalls[height][prevRow]
	}
	// Find all ways that the previous row can be matched
	var res int
	for _, mask := range masks {
		if mask&prevRow > 0 { // overlapping edges -> not sturdy
			continue
		}
		res = (res + findSturdyWalls(sturdyWalls, height-1, mask, masks)) % mod
	}
	sturdyWalls[height][prevRow] = res
	return sturdyWalls[height][prevRow]
}

func getConfigs(bricks []int, mask, idx, width int, res *[]int) {
	if idx == width {
		*res = append(*res, mask^(1<<width))
		return
	}
	for _, b := range bricks {
		if idx+b <= width {
			getConfigs(bricks, mask|(1<<(idx+b)), idx+b, width, res)
		}
	}
}
