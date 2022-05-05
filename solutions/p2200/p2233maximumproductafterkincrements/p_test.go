package p2233maximumproductafterkincrements

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumProduct(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1}, 1, 2},
		{[]int{4, 0, 7, 8, 8, 0, 2, 4, 5}, 1, 0},
		{[]int{0, 4}, 5, 20},
		{[]int{9, 7, 8}, 9, 1331},
		{[]int{6, 3, 3, 2}, 2, 216},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumProduct(tc.nums, tc.k))
		})
	}
}

const mod = 1e9 + 7

func maximumProduct(nums []int, k int) int {
	// My hypothesis is that an optimal solution can be found by incrementing the
	// smallest number in nums.
	// So the problem is rather to fill numbers in nums such that a max delta of
	// k has been used.

	// We may use binary search to find the minimum value in nums after
	// incrementing all of k. Then we fill and calculate the result.
	var maxNum int
	for _, x := range nums {
		maxNum = max(maxNum, x)
	}

	l, r := 0, math.MaxInt32
	for l < r {
		mid := l + (r-l)/2
		// Check if mid is the smallest number
		kk := k
		for _, x := range nums {
			if x < mid {
				kk -= mid - x
			}
			if kk < 0 {
				break
			}
		}
		if kk < 0 {
			r = mid
		} else {
			l = mid + 1
		}
	}
	minVal := l - 1
	rest := k
	for i, x := range nums {
		if x < minVal {
			nums[i] = minVal
			rest -= minVal - x
		}
	}
	if minVal == 0 {
		return 0
	}
	// The minimum number in nums is l
	var res int
	for _, x := range nums {
		if x == minVal && rest > 0 {
			if rest > 0 {
				x++
				rest--
			}
		}
		res = (max(res, 1) * x) % mod
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
