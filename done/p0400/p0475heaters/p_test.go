package p0475heaters

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRadius(t *testing.T) {
	for _, tc := range []struct {
		houses  []int
		heaters []int
		want    int
	}{
		{[]int{1, 2, 3}, []int{2}, 1},
		{[]int{1, 2, 3, 4}, []int{1, 4}, 1},
		{[]int{1, 5}, []int{2}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.houses), func(t *testing.T) {
			require.Equal(t, tc.want, findRadius(tc.houses, tc.heaters))
		})
	}
}

func findRadius(houses []int, heaters []int) int {
	heaters = append(heaters, 0)
	copy(heaters[1:], heaters)
	heaters = append(heaters, math.MaxInt64)
	heaters[0] = math.MinInt32
	sort.Ints(houses)
	sort.Ints(heaters)

	prev, next := 0, 1
	var maxDist int
	for _, house := range houses {
		for house >= heaters[next] {
			prev, next = next, next+1
		}
		maxDist = max(maxDist, min(house-heaters[prev], heaters[next]-house))
	}
	return maxDist
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
