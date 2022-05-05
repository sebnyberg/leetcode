package p1508rangesumofsortedsubarraysums

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_rangeSum(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		n     int
		left  int
		right int
		want  int
	}{
		{[]int{1, 2, 3, 4}, 4, 1, 5, 13},
		{[]int{1, 2, 3, 4}, 4, 3, 4, 6},
		{[]int{1, 2, 3, 4}, 4, 1, 10, 50},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, rangeSum(tc.nums, tc.n, tc.left, tc.right))
		})
	}
}

const mod = 1000000007

func rangeSum(nums []int, n int, left int, right int) int {
	// brute-force
	sums := make([]int, 0, n)
	for i := range nums {
		sums = append(sums, nums[i])
		cur := nums[i]
		for j := i + 1; j < n; j++ {
			cur += nums[j]
			sums = append(sums, cur)
		}
	}
	sort.Ints(sums)
	var res int
	for i := left - 1; i <= right-1; i++ {
		res += sums[i]
		res %= mod
	}
	return res % mod
}
