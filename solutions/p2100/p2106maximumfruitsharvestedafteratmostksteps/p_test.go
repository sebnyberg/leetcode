package p2106maximumfruitsharvestedafteratmostksteps

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxTotalFruits(t *testing.T) {
	for _, tc := range []struct {
		fruits      [][]int
		startPos, k int
		want        int
	}{
		{[][]int{{200000, 10000}}, 200000, 0, 10000},
		{[][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, 5, 4, 14},
		{[][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4, 9},
		{[][]int{{0, 3}, {6, 4}, {8, 5}}, 3, 2, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.fruits), func(t *testing.T) {
			require.Equal(t, tc.want, maxTotalFruits(tc.fruits, tc.startPos, tc.k))
		})
	}
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	var l, sum, maxSum int
	for l < len(fruits) && fruits[l][0] < startPos-k {
		l++
	}

	for r := l; r < len(fruits) && fruits[r][0] <= startPos+k; r++ {
		sum += fruits[r][1]
		for min(startPos-2*fruits[l][0]+fruits[r][0], 2*fruits[r][0]-startPos-fruits[l][0]) > k {
			sum -= fruits[l][1]
			l++
		}
		maxSum = max(maxSum, sum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
