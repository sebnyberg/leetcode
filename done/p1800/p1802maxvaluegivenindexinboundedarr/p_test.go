package p1802maxvaluegivenindexinboundedarr

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getSum(t *testing.T) {
	for _, tc := range []struct {
		left        int
		right       int
		targetValue int
		want        int
	}{
		{2, 0, 1, 1},
		{2, 0, 2, 3},
		{2, 0, 3, 6},
		{2, 0, 4, 9},
		{2, 0, 5, 12},
	} {
		t.Run(fmt.Sprintf("%+v", tc.left), func(t *testing.T) {
			require.Equal(t, tc.want, getSum(tc.left, tc.right, tc.targetValue))
		})
	}
}

func Test_maxValue(t *testing.T) {
	for _, tc := range []struct {
		n      int
		index  int
		maxSum int
		want   int
	}{
		{4, 0, 4, 1},
		{1, 0, 24, 24},
		{3, 2, 18, 7},
		{4, 2, 6, 2},
		{6, 1, 10, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, maxValue(tc.n, tc.index, tc.maxSum))
		})
	}
}

func maxValue(n int, index int, maxSum int) int {
	leftWidth := index
	rightWidth := n - index - 1

	// Find ceil
	res := sort.Search(2*maxSum, func(i int) bool {
		return getSum(leftWidth, rightWidth, i) > maxSum
	})
	a := getSum(leftWidth, rightWidth, res)
	b := getSum(leftWidth, rightWidth, res+1)
	c := getSum(leftWidth, rightWidth, res-1)
	_ = a
	_ = b
	_ = c

	return res - 1
}

func getSum(left, right, targetValue int) int {
	leftWidth := min(targetValue-1, left)
	rightWidth := min(targetValue-1, right)
	leftSum := (leftWidth * (leftWidth + 1)) / 2
	rightSum := (rightWidth * (rightWidth + 1)) / 2
	leftDiff := targetValue - 1 - left
	if leftDiff > 0 {
		leftSum += left * leftDiff
	} else if leftDiff < 0 {
		leftSum += 1 * -leftDiff
	}
	rightDiff := targetValue - 1 - right
	if rightDiff > 0 {
		rightSum += right * rightDiff
	} else if rightDiff < 0 {
		leftSum += 1 * -rightDiff
	}
	return leftSum + rightSum + targetValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
