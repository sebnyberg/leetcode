package p2073timeneededtobuytickets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_timeRequiredToBuy(t *testing.T) {
	for _, tc := range []struct {
		tickets []int
		k       int
		want    int
	}{
		{[]int{2, 3, 2}, 2, 6},
		{[]int{5, 1, 1, 1}, 0, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tickets), func(t *testing.T) {
			require.Equal(t, tc.want, timeRequiredToBuy(tc.tickets, tc.k))
		})
	}
}

func timeRequiredToBuy(tickets []int, k int) int {
	var time int
	for {
		for i := 0; i < len(tickets); i++ {
			if i == k && tickets[k] == 1 {
				return time + 1
			}
			if tickets[i] > 0 {
				tickets[i]--
				time++
			}
		}
	}
}
