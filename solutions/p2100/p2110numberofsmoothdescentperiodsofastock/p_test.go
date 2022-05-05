package p2110numberofsmoothdescentperiodsofastock

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getDescentPeriods(t *testing.T) {
	for _, tc := range []struct {
		prices []int
		want   int64
	}{
		{[]int{3, 2, 1, 4}, 7},
		{[]int{8, 6, 7, 7}, 4},
		{[]int{1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.prices), func(t *testing.T) {
			require.Equal(t, tc.want, getDescentPeriods(tc.prices))
		})
	}
}

func getDescentPeriods(prices []int) int64 {
	var count int64
	var width int64 = 1
	prices = append(prices, math.MaxInt32)
	for i := 1; i < len(prices); i++ {
		if prices[i] != prices[i-1]-1 {
			for width > 0 {
				count += width
				width--
			}
			width = 1
		} else {
			width++
		}
	}
	return count
}
