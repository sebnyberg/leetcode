package p0018foursum

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fourSum(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   [][]int
	}{
		// {[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{-2, -1, -1, 1, 1, 2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-1, -1, 1, 1}}},
		// {[]int{}, 0, [][]int{}},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.nums, tc.target), func(t *testing.T) {
			require.Equal(t, tc.want, fourSum(tc.nums, tc.target))
		})
	}
}

// Find unique quadruplets in the array for which the sum is equal to target.
func fourSum(nums []int, target int) [][]int {
	// 0 <= nums.length <= 200
	// -1e10 <= nums[i], target <= 1e10
	sort.Ints(nums)

	res := make([][]int, 0)

	var n1, n2, n3, n4 int

	// For each unique number
	for i := 0; i < len(nums)-3; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		n1 = nums[i]
		for j := i + 1; j < len(nums)-2; j++ {
			// Skip duplicates
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			n2 = nums[j]
			low, high := j+1, len(nums)-1
			for low < high {
				n3, n4 = nums[low], nums[high]
				sum := n1 + n2 + n3 + n4
				switch {
				case sum > target:
					high--
				case sum < target:
					low++
				default:
					// Skip duplicates
					for low++; low < high && nums[low] == nums[low-1]; low++ {
					}
					res = append(res, []int{n1, n2, n3, n4})
				}
			}
		}
	}

	return res
}
