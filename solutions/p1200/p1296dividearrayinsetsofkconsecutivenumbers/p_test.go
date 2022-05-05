package p1296dividearrayinsetsofkconsecutivenumbers

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPossibleDivide(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 3, 4, 4, 5, 6}, 4, true},
		{[]int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}, 3, true},
		{[]int{1, 2, 3, 4}, 3, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, isPossibleDivide(tc.nums, tc.k))
		})
	}
}

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	// First, perform RLE of nums
	sort.Ints(nums)
	count := make([]int32, len(nums))
	j := -1
	for i := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			count[j]++
			continue
		}
		j++
		nums[j] = nums[i]
		count[j]++
	}
	nums = nums[:j+1]
	count = count[:j+1]

	// Then, for each index in nums for which count is non-zero, reduce count for
	// k elements by the count. If there is no capacity, return false
	for i := 0; i < len(nums); i++ {
		if count[i] == 0 {
			continue
		}
		rem := count[i]
		count[i] -= rem
		for j := i + 1; j < i+k; j++ {
			if j == len(nums) || rem > count[j] || nums[j] != nums[j-1]+1 {
				return false
			}
			count[j] -= rem
		}
	}
	return true
}
