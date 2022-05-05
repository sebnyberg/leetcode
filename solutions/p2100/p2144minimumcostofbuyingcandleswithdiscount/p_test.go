package p2144minimumcostofbuyingcandleswithdiscount

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumCost(t *testing.T) {
	for _, tc := range []struct {
		cost []int
		want int
	}{
		{[]int{1, 2, 3}, 5},
		{[]int{6, 5, 7, 9, 2, 2}, 23},
		{[]int{5, 5}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.cost), func(t *testing.T) {
			require.Equal(t, tc.want, minimumCost(tc.cost))
		})
	}
}

func minimumCost(cost []int) int {
	sort.Slice(cost, func(i, j int) bool {
		return cost[i] > cost[j]
	})
	var res int
	for i := range cost {
		if i%3 == 2 {
			continue
		}
		res += cost[i]
	}
	return res
}
