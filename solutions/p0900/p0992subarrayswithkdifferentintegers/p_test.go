package p0992subarrayswithkdifferentintegers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subarraysWithKDistinct(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{2, 1, 1, 1, 2}, 1, 8},
		{[]int{1, 2}, 1, 2},
		{[]int{1, 2, 1, 2, 3}, 2, 7},
		{[]int{1, 2, 1, 3, 4}, 3, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, subarraysWithKDistinct(tc.nums, tc.k))
		})
	}
}

func subarraysWithKDistinct(nums []int, k int) int {
	// Greedy solution
	// Keep track of the number of distinct numbers from nums[l:i]
	// Whenever possible, move l and increment the number of variants
	// The variants keeps track of the redundant prefix of the window
	n := len(nums)
	count := make([]int, n+1)
	var ndistinct int
	var l int
	variants := 1
	var res int
	for _, x := range nums {
		if count[x] == 0 {
			ndistinct++
		}
		count[x]++
		if ndistinct < k {
			continue
		}
		for ndistinct > k {
			variants = 1
			count[nums[l]]--
			if count[nums[l]] == 0 {
				ndistinct--
			}
			l++
		}
		// count how many numbers could be removed from the left while keeping k
		// distinct numbers around
		for count[nums[l]] > 1 {
			variants++
			count[nums[l]]--
			l++
		}
		res += variants
	}
	return res
}
