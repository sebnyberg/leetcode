package p2334subarraywithelementsgreaterthanvaryingthreshold

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validSubarraySize(t *testing.T) {
	for _, tc := range []struct {
		nums      []int
		threshold int
		want      int
	}{
		{[]int{1, 3, 4, 3, 1}, 6, 3},
		{[]int{6, 5, 6, 5, 8}, 7, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, validSubarraySize(tc.nums, tc.threshold))
		})
	}
}

func validSubarraySize(nums []int, threshold int) int {
	// The key insight in this exercise is that a candidate subarray will start
	// and end with small numbers. The reason being that if you have a subarray
	// which can be expanded in either direction without increasing the maximum
	// value of the range, then you might as well add those values as well. It
	// will improve the probability of finding an answer.
	//
	// Using a monotonically increasing stack, we can quickly find subarrays of
	// large values between small values.
	//
	stack := []int{}
	nums = append(nums, 0)
	m := 0
	for i, v := range nums {
		for m > 0 && v < nums[stack[m-1]] {
			x := nums[stack[m-1]]
			stack = stack[:m-1]
			m--
			j := -1
			if m > 0 {
				j = stack[m-1]
			}
			if x > threshold/(i-j-1) {
				return i - j - 1
			}
		}
		stack = append(stack, i)
		m++
	}
	return -1
}
