package p1144decreaseelementstomakearrayzigzag

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_movesToMakeZigzag(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{3, 10, 7, 9, 9, 3, 6, 9, 4}, 11},
		{[]int{1, 2, 3}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, movesToMakeZigzag(tc.nums))
		})
	}
}

func movesToMakeZigzag(nums []int) int {
	// We should try both alternatives.
	//
	// Consider that the current element should be smaller than the next and it
	// is currently not. E.g. 9, 6
	//
	// We know that 9 is smaller than any number that precedes it, so decreasing
	// it cannot change any prior actions, it only impacts its relationship with
	// the next number.
	//
	// We also know that we have to decrease it by at least 4 to continue
	// traversal.
	//
	// This gives us the solution. Try the following and collect the optimal
	// action:
	//
	// 1. Decrease odd indices until they are smaller than their neighbours
	// 2. Decrease even indices until they are smaller than their neighbours
	//
	if len(nums) == 1 {
		return 0
	}
	var res [2]int
	for i := range nums {
		// decrease current until it is smaller than left/right
		left := math.MaxInt32
		right := math.MaxInt32
		if i > 0 {
			left = nums[i-1]
		}
		if i < len(nums)-1 {
			right = nums[i+1]
		}
		d := nums[i] - min(left, right)
		if d >= 0 {
			res[i&1] += d + 1
		}
	}
	ret := math.MaxInt32
	for _, r := range res {
		ret = min(ret, r)
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
