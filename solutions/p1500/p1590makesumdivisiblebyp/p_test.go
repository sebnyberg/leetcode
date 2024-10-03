package p1590makesumdivisiblebyp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSubarray(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		p    int
		want int
	}{
		{[]int{1, 2, 3}, 7, -1},
		{[]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148, 7},
		{[]int{3, 1, 4, 2}, 6, 1},
		{[]int{6, 3, 5, 2}, 9, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minSubarray(tc.nums, tc.p))
		})
	}
}

func minSubarray(nums []int, p int) int {
	n := len(nums)

	var rest int
	for _, x := range nums {
		rest += x
		rest %= p
	}
	if rest == 0 {
		return 0
	}
	result := n

	lastIdx := map[int]int{}
	lastIdx[0] = -1
	var sum int
	for i := 0; i < n; i++ {
		sum = (sum + nums[i]) % p
		want := (sum - rest + p) % p
		if j, exists := lastIdx[want]; exists && i-j < result {
			result = i - j
		}
		lastIdx[sum] = i
	}
	if result == n {
		return -1
	}
	return result
}
