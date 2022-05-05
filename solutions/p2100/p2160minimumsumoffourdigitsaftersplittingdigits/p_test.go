package p2160minimumsumoffourdigitsaftersplittingdigits

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumSum(t *testing.T) {
	for _, tc := range []struct {
		num  int
		want int
	}{
		{2932, 52},
		{4009, 13},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, minimumSum(tc.num))
		})
	}
}

func minimumSum(num int) int {
	// Since num is small we can try all possible combinations
	nums := make([]int, 0)
	for num > 0 {
		nums = append(nums, num%10)
		num /= 10
	}
	return dfs(nums, 0, 0, 0, 0, len(nums))
}

func dfs(nums []int, first, second, bm, count, n int) int {
	if count == n {
		return first + second
	}
	res := math.MaxInt32
	for i := 0; i < n; i++ {
		if bm&(1<<i) > 0 {
			continue
		}
		res = min(res, dfs(nums, first*10+nums[i], second, bm|(1<<i), count+1, n))
		res = min(res, dfs(nums, first, second*10+nums[i], bm|(1<<i), count+1, n))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
