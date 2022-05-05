package p0689maximumsumof3nonoverlappingsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSumOfThreeSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{4, 3, 2, 1}, 1, []int{0, 1, 2}},
		{[]int{1, 2, 1, 2, 6, 7, 5, 1}, 2, []int{0, 3, 5}},
		{[]int{1, 2, 1, 2, 1, 2, 1, 2, 1}, 2, []int{0, 2, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxSumOfThreeSubarrays(tc.nums, tc.k))
		})
	}
}

func maxSumOfThreeSubarrays(nums []int, k int) []int {
	// It feels like a DP exercise.
	n := len(nums)
	dp := make([][3]int, n)
	start := make([][3][3]int, n)
	var sum int
	// res := []int{0, 0, 0}
	for i, n := range nums {
		sum += n
		if i >= k {
			sum -= nums[i-k]
		}
		if i >= k-1 {
			if i == 0 || sum > dp[i-1][0] {
				dp[i][0] = sum
				start[i][0] = [3]int{i - k + 1, 0, 0}
			} else {
				dp[i][0] = dp[i-1][0]
				start[i][0] = start[i-1][0]
			}
		}
		if i >= 2*k-1 {
			if sum+dp[i-k][0] > dp[i-1][1] {
				dp[i][1] = sum + dp[i-k][0]
				start[i][1] = [3]int{start[i-k][0][0], i - k + 1, 0}
			} else {
				dp[i][1] = dp[i-1][1]
				start[i][1] = start[i-1][1]
			}
		}
		if i >= 3*k-1 {
			if sum+dp[i-k][1] > dp[i-1][2] {
				dp[i][2] = sum + dp[i-k][1]
				start[i][2] = start[i-k][1]
				start[i][2][2] = i - k + 1
			} else {
				dp[i][2] = dp[i-1][2]
				start[i][2] = start[i-1][2]
			}
		}
	}
	res := start[n-1][2][:]
	return res
}
