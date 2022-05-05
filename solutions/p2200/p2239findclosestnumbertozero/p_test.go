package p2239

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findClosestNumber(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{-4, -2, 1, 4, 8}, 1},
		{[]int{2, -1, 1}, 1},
		{[]int{-10000}, -10000},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findClosestNumber(tc.nums))
		})
	}
}

func findClosestNumber(nums []int) int {
	res := math.MaxInt64
	minDist := math.MaxInt64
	for _, x := range nums {
		if abs(x) == minDist {
			if x > res {
				res = x
			}
		}
		if abs(x) < minDist {
			minDist = abs(x)
			res = x
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
