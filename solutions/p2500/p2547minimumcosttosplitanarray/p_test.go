package p2547minimumcosttosplitanarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minCost(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 1, 2, 1, 3, 3}, 2, 8},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minCost(tc.nums, tc.k))
		})
	}
}

func minCost(nums []int, k int) int {
	n := len(nums)
	mem := make([]int, n)
	res := dp(mem, nums, 0, n, k)
	return res
}

func dp(mem, nums []int, i, n, k int) int {
	if i == n {
		return 0
	}
	if mem[i] != 0 {
		return mem[i]
	}
	count := make([]int, n+1)
	count[nums[i]] = 1
	res := k + dp(mem, nums, i+1, n, k)
	var subn int
	for j := i + 1; j < n; j++ {
		if count[nums[j]] == 1 {
			subn += 2
		}
		if count[nums[j]] > 1 {
			subn++
		}
		count[nums[j]]++
		res = min(res, subn+k+dp(mem, nums, j+1, n, k))
	}

	mem[i] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
