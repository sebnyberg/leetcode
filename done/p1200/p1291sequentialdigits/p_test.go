package p1291sequentialdigits

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sequentialDigits(t *testing.T) {
	for _, tc := range []struct {
		low  int
		high int
		want []int
	}{
		{100, 300, []int{123, 234}},
		{1000, 13000, []int{1234, 2345, 3456, 4567, 5678, 6789, 12345}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.low), func(t *testing.T) {
			require.Equal(t, tc.want, sequentialDigits(tc.low, tc.high))
		})
	}
}

func sequentialDigits(low int, high int) []int {
	nums := []int{12, 23, 34, 45, 56, 67, 78, 89}
	var res []int
	for i := 0; len(nums) > 0 && nums[i] <= high; {
		if nums[i] >= low {
			res = append(res, nums[i])
		}
		i++
		if i == len(nums) {
			i = 0
			nums = nums[:len(nums)-1]
			for j := range nums {
				nums[j] = nums[j]*10 + nums[j]%10 + 1
			}
		}
	}
	return res
}
