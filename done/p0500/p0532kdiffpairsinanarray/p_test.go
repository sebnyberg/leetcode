package p0532kdiffpairsinanarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
		{[]int{1, 1}, 0, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findPairs(tc.nums, tc.k))
		})
	}
}

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	if k == 0 {
		var res int
		for i := 1; i < len(nums); i++ {
			if nums[i] == nums[i-1] && (i == 1 || nums[i-2] != nums[i]) {
				res++
			}
		}
		return res
	}

	// Remove duplicates from nums
	// if k == 0, count duplicates
	j := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			continue
		}
		nums[j] = nums[i]
		j++
	}
	nums = nums[:j]

	var res int
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, exists := seen[num-k]; exists {
			res++
		}
		seen[num] = struct{}{}
	}

	return res
}
