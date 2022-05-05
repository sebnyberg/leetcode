package p1833maximumicecreambars

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxIceCream(t *testing.T) {
	for _, tc := range []struct {
		costs []int
		coins int
		want  int
	}{
		{[]int{1, 3, 2, 4, 1}, 7, 4},
		{[]int{10, 6, 8, 7, 7, 8}, 5, 0},
		{[]int{1, 6, 3, 1, 2, 5}, 20, 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.costs), func(t *testing.T) {
			require.Equal(t, tc.want, maxIceCream(tc.costs, tc.coins))
		})
	}
}

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	var buys int
	for i := 0; i < len(costs) && costs[i] <= coins; i++ {
		coins -= costs[i]
		buys++
	}
	return buys
}
