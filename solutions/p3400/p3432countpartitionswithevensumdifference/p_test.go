package p3432countpartitionswithevensumdifference

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countPartitions(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{10, 10, 3, 7, 6}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countPartitions(tc.nums))
		})
	}
}

func countPartitions(nums []int) int {
	var sum int
	for i := range nums {
		sum += nums[i]
	}
	var left int
	var res int
	for i := range len(nums) - 1 {
		left += nums[i]
		sum -= nums[i]
		if (left-sum)%2 == 0 {
			res++
		}
	}
	return res
}
