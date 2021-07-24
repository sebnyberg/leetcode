package p0611validtrianglenumber

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_triangleNumber(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 0, 0}, 0},
		{[]int{2, 2, 3, 4}, 3},
		{[]int{4, 2, 3, 4}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, triangleNumber(tc.nums))
		})
	}
}

func triangleNumber(nums []int) int {
	// A valid triangle is a triplet such that the sum of two of its numbers is
	// greater than the third number.
	n := len(nums)
	sort.Ints(nums)
	var res int
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			// Find first index that is greater than nums[i]+nums[j]
			firstGreater := sort.SearchInts(nums, nums[i]+nums[j])
			if firstGreater <= j {
				continue
			}
			res += firstGreater - j - 1
		}
	}
	return res
}
