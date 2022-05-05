package p1217mincosttomovechipstosameposition

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCostToMoveChips(t *testing.T) {
	for _, tc := range []struct {
		position []int
		want     int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{2, 2, 2, 3, 3}, 2},
		{[]int{1, 1000000000}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.position), func(t *testing.T) {
			require.Equal(t, tc.want, minCostToMoveChips(tc.position))
		})
	}
}

func minCostToMoveChips(position []int) int {
	// All chips can be gathered so that they end up as two piles next to each
	// other for free
	var chipCount [2]int
	for _, pos := range position {
		chipCount[pos%2]++
	}
	// Then the smallest pile should be moved to the larger pile
	if chipCount[1] < chipCount[0] {
		return chipCount[1]
	}
	return chipCount[0]
}
