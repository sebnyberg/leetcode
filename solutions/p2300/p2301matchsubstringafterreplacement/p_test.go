package p2301matchsubstringafterreplacement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSubarrays(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int64
		want int64
	}{
		{[]int{2, 1, 4, 3, 5}, 10, 6},
		{[]int{1, 1, 1}, 5, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countSubarrays(tc.nums, tc.k))
		})
	}
}

func countSubarrays(nums []int, k int64) int64 {
	n := len(nums)
	var sum int
	var l int
	var res int64
	for r := 0; r < n; r++ {
		sum += nums[r]
		for l != r && sum*(r-l+1) >= int(k) {
			sum -= nums[l]
			l++
		}
		if sum*(r-l+1) < int(k) {
			res += int64(r - l + 1)
		}
	}
	return res
}
