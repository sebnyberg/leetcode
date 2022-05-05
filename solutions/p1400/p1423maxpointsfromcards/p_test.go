package p1423maxpointsfromcards

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxScore(t *testing.T) {
	for _, tc := range []struct {
		cardPoints []int
		k          int
		want       int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
		{[]int{2, 2, 2}, 2, 4},
		{[]int{9, 7, 7, 9, 7, 7, 9}, 7, 55},
	} {
		t.Run(fmt.Sprintf("%+v", tc.cardPoints), func(t *testing.T) {
			require.Equal(t, tc.want, maxScore(tc.cardPoints, tc.k))
		})
	}
}

func maxScore(cardPoints []int, k int) int {
	var curSum int
	n := len(cardPoints)
	for i := 0; i < n-k; i++ {
		curSum += cardPoints[i]
	}
	minSum, totalSum := curSum, curSum
	for i := n - k; i < n; i++ {
		curSum -= cardPoints[i-(n-k)]
		curSum += cardPoints[i]
		totalSum += cardPoints[i]
		minSum = min(minSum, curSum)
	}
	return totalSum - minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
