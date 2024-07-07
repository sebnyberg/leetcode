package p1518waterbottles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numWaterBottles(t *testing.T) {
	for _, tc := range []struct {
		numBottles  int
		numExchange int
		want        int
	}{
		{9, 3, 13},
		{15, 4, 19},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numBottles), func(t *testing.T) {
			require.Equal(t, tc.want, numWaterBottles(tc.numBottles, tc.numExchange))
		})
	}
}

func numWaterBottles(numBottles int, numExchange int) int {
	var res int
	var empty int
	for numBottles > 0 {
		res += numBottles
		numBottles += empty
		empty = numBottles % numExchange
		numBottles /= numExchange
	}
	return res
}
