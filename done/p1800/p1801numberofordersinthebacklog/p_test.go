package p1801numberofordersinthebacklog

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getNumberOfBacklogOrders(t *testing.T) {
	for _, tc := range []struct {
		orders [][]int
		want   int
	}{
		// {[][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}, 6},
		// {[][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}, 999999984},
	} {
		t.Run(fmt.Sprintf("%+v", tc.orders), func(t *testing.T) {
			require.Equal(t, tc.want, getNumberOfBacklogOrders(tc.orders))
		})
	}
}

// TODO
func getNumberOfBacklogOrders(orders [][]int) int {
	return 0
}
