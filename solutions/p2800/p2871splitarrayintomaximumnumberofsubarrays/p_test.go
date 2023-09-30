package p2871splitarrayintomaximumnumberofsubarrays

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxSubarrays(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{22, 21, 29, 22}, 1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxSubarrays(tc.nums))
		})
	}
}

func maxSubarrays(nums []int) int {
	minPossible := nums[0]
	for _, x := range nums {
		minPossible &= x
	}

	if minPossible > 0 {
		return 1
	}

	// Greedily try to split whenever we get zero for a subsequence
	var count int
	res := nums[0]
	for i, x := range nums {
		res &= x
		if res != 0 {
			continue
		}
		// Split here.
		count++
		if i < len(nums)-1 {
			res = nums[i+1]
		}
	}

	return count
}
