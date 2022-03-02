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
		{[]int{1, 2, 3, 4, 5}, 0, 0},
		{[]int{3, 1, 4, 1, 5}, 2, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 4},
		{[]int{1, 3, 1, 5, 4}, 0, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findPairs(tc.nums, tc.k))
		})
	}
}

func findPairs(nums []int, k int) int {
	sort.Ints(nums)

	seen := make(map[int]struct{}, len(nums))
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[j] {
			if k == 0 {
				seen[nums[i]] = struct{}{}
			}
			continue
		}
		j++
		nums[j] = nums[i]
	}
	nums = nums[:j+1]
	if k == 0 {
		return len(seen)
	}

	var res int
	for _, num := range nums {
		if _, exists := seen[num-k]; exists {
			res++
		}
		seen[num] = struct{}{}
	}
	return res
}
