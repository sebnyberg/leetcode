package p2136earliestpossibledayoffullbloom

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_earliestFullBloom(t *testing.T) {
	for _, tc := range []struct {
		plantTime []int
		growTime  []int
		want      int
	}{
		{[]int{27, 5, 24, 17, 27, 4, 23, 16, 6, 26, 13, 17, 21, 3, 9, 10, 28, 26, 4, 10, 28, 2},
			[]int{26, 9, 14, 17, 6, 14, 23, 24, 11, 6, 27, 14, 13, 1, 15, 5, 12, 15, 23, 27, 28, 12}, 348},
		{[]int{1, 4, 3}, []int{2, 3, 1}, 9},
		{[]int{1, 2, 3, 2},
			[]int{2, 1, 2, 1}, 9},
		{[]int{1}, []int{1}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.plantTime), func(t *testing.T) {
			require.Equal(t, tc.want, earliestFullBloom(tc.plantTime, tc.growTime))
		})
	}
}

func earliestFullBloom(plantTime []int, growTime []int) int {
	var sums [10001]uint32
	var maxGrowTime int
	minGrowTime := math.MaxInt32
	for i := range plantTime {
		sums[growTime[i]] += uint32(plantTime[i])
		maxGrowTime = max(maxGrowTime, growTime[i])
		minGrowTime = min(minGrowTime, growTime[i])
	}
	var sum int
	var res int
	for i := maxGrowTime; i >= minGrowTime; i-- {
		sum += int(sums[i])
		res = max(res, sum+i)
	}

	return int(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
