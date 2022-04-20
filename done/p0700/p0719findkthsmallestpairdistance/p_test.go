package p0719findkthsmallestpairdistance

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestDistancePair(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 3, 1}, 1, 0},
		{[]int{1, 1, 1}, 2, 0},
		{[]int{1, 6, 1}, 3, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, smallestDistancePair(tc.nums, tc.k))
		})
	}
}

func smallestDistancePair(nums []int, k int) int {
	// nums[i] is too large for a frequency count
	// Ordering does not matter (i.e. we can sort)
	// With sort + a stack, we could efficiently count how many distances exist <=
	// some value.
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]*3) // sentinel
	countSmallerThan := func(d int) int {
		var l int
		var res int
		for r := 1; r < len(nums); r++ {
			// Move left pointer until range is OK
			for r-l+1 >= 2 && nums[r]-nums[l] > d {
				l++
			}

			// Count valid distances ending in r
			if r-l+1 >= 2 {
				res += r - l
			}
		}
		return res
	}

	lo, hi := 0, nums[len(nums)-1]-nums[0]+1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if countSmallerThan(mid) < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}
