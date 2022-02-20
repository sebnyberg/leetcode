package p2178

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumEvenSplit(t *testing.T) {
	for _, tc := range []struct {
		finalSum int64
		want     []int64
	}{
		{2, []int64{2}},
		{10, []int64{2, 8}},
		{14, []int64{2, 4, 8}},
		{12, []int64{2, 4, 6}},
		{7, []int64{}},
		// {28, []int64{6, 8, 2, 12}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.finalSum), func(t *testing.T) {
			require.Equal(t, tc.want, maximumEvenSplit(tc.finalSum))
		})
	}
}

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}
	// n sums add p to n(n+1)
	var lo, hi int64 = 0, math.MaxInt32
	for lo < hi {
		mid := (lo + hi) / 2
		res := mid * (mid + 1)
		if res > finalSum {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	left := finalSum - (lo-1)*(lo)
	var res []int64
	for i := int64(1); i < lo; i++ {
		res = append(res, i*2)
	}
	res[len(res)-1] += left
	return res
}
