package p0485maxconsecutiveones

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMaxConsecutiveOnes(tc.nums))
		})
	}
}

func findMaxConsecutiveOnes(nums []int) int {
	var count int
	var maxCount int
	for i := range nums {
		if nums[i] == 0 {
			count = 0
			continue
		}
		count++
		if count > maxCount {
			maxCount = count
		}
	}
	return maxCount
}
