package p1788maximizebeautyofthegarden

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumBeauty(t *testing.T) {
	for _, tc := range []struct {
		flowers []int
		want    int
	}{
		{[]int{-5, -1, 2, -1}, 0},
		{[]int{-1, -2, 0, -1}, -2},
		{[]int{1, 2, 3, 1, 2}, 8},
		{[]int{100, 1, 1, -3, 1}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.flowers), func(t *testing.T) {
			require.Equal(t, tc.want, maximumBeauty(tc.flowers))
		})
	}
}

func maximumBeauty(flowers []int) int {
	res := math.MinInt32
	prefixSum := 0
	sums := make(map[int]int)
	for _, flower := range flowers {
		if _, exists := sums[flower]; !exists {
			sums[flower] = prefixSum
			if flower > 0 {
				prefixSum += flower
			}
			continue
		}
		if flower >= 0 {
			prefixSum += flower
			res = max(res, prefixSum-sums[flower])
		} else {
			res = max(res, flower*2+prefixSum-sums[flower])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
