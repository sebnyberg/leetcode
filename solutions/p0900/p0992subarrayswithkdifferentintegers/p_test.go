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
	// Convert this into two sub-problems: finding the number of subarrays with
	// at most K distinct characters, then take count(K) - count(K-1) to get the
	// answer.
	a := countSubarraysWithAtMostKDistinct(nums, k-1)
	b := countSubarraysWithAtMostKDistinct(nums, k)
	return b - a
}

func countSubarraysWithAtMostKDistinct(nums []int, k int) int {
	if k == 0 {
		return 0
	}
	l := 0
	numCount := make(map[int]int)
	ndistinct := 0
	res := 0
	for r, num := range nums {
		if numCount[num] == 0 {
			ndistinct++
		}
		numCount[num]++
		for l < r && ndistinct > k {
			numCount[nums[l]]--
			if numCount[nums[l]] == 0 {
				ndistinct--
			}
			l++
		}
		res += r - l + 1
	}
	return res
}
