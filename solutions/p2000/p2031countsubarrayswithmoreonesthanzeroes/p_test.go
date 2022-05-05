package p2031countsubarrayswithmoreonesthanzeroes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subarraysWithMoreZerosThanOnes(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 1, 0, 1}, 9},
		{[]int{0}, 0},
		{[]int{1}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, subarraysWithMoreZerosThanOnes(tc.nums))
		})
	}
}

const (
	mod      = 1e9 + 7
	maxItems = 1e5
)

func subarraysWithMoreZerosThanOnes(nums []int) int {
	var sumCount [maxItems*2 + 1]uint32
	var offset int = maxItems
	sumCount[offset] = 1
	var validCount, totalCount int
	var sum int
	for i := range nums {
		if nums[i] == 0 {
			validCount = (validCount - int(sumCount[sum-1+offset])) % mod
			sum--
		} else {
			validCount = (validCount + int(sumCount[sum+offset])) % mod
			sum++
		}
		sumCount[int(sum)+offset]++
		totalCount = (totalCount + validCount) % mod
	}
	return int(totalCount)
}
