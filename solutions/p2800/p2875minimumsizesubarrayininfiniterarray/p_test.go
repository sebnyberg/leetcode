package p2875minimumsizesubarrayininfiniterarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSizeSubarray(t *testing.T) {
	for i, tc := range []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3}, 5, 2},
		{[]int{1, 1, 1, 2, 3}, 4, 2},
		{[]int{2, 4, 6, 8}, 3, -1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minSizeSubarray(tc.nums, tc.target))
		})
	}
}

func minSizeSubarray(nums []int, target int) int {
	var sum int
	for _, x := range nums {
		sum += x
	}
	n := len(nums)

	// First we discount any runs as long as the original array
	k := target / sum
	target -= k * sum
	var res int
	res += k * n

	// Then we find the smallest subarray in nums that contains the target.
	//
	// We can do this using a stack and double the original array length. For
	// ease, I will just duplicate nums
	nums = append(nums, make([]int, n)...)
	copy(nums[n:], nums[:n])
	subres := math.MaxInt32
	var l int
	target -= nums[l]
	if target == 0 {
		return res + 1
	}
	for r := 1; r < len(nums); r++ {
		target -= nums[r]
		for target < 0 {
			target += nums[l]
			l++
		}
		if target == 0 {
			subres = min(subres, r-l+1)
		}
	}
	if subres == math.MaxInt32 {
		return -1
	}
	return subres + res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
