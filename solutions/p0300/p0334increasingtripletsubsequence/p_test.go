package p0334increasingtripletsubsequence

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_increasingTriplet(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 1, 1, 1}, false},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{2, 1, 5, 0, 4, 6}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, increasingTriplet(tc.nums))
		})
	}
}

func increasingTriplet(nums []int) bool {
	// The only thing that matters is nums[j]..
	// What I mean by that is, for any given index in nums, if there exists a
	// prior index such that its value is smaller than nums[j] and nums[j] is
	// less than the current second value, then we update the second value.
	minVal := math.MaxInt32
	midVal := math.MaxInt32
	for _, x := range nums {
		if x < minVal {
			minVal = x
			continue
		}
		if midVal < x {
			return true
		}
		if x > minVal && x < midVal {
			midVal = x
		}
	}
	return false
}
