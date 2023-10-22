package p2908minimumsumofmountaintripletsi

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumSum(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{8, 6, 1, 5, 3}, 9},
		{[]int{5, 4, 8, 7, 10, 2}, 13},
		{[]int{6, 5, 4, 3, 4, 5}, -1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumSum(tc.nums))
		})
	}
}

func minimumSum(nums []int) int {
	// For each index, keep track of the smallest value to it's right such that
	// it is part of a descending series.
	n := len(nums)
	right := make([]int, n)
	right[n-1] = math.MaxInt32
	for i := n - 2; i >= 0; i-- {
		right[i] = min(right[i+1], nums[i+1])
	}
	left := nums[0]
	res := math.MaxInt32
	for i := 1; i < n-1; i++ {
		if left < nums[i] && right[i] < nums[i] {
			res = min(res, left+nums[i]+right[i])
		}
		left = min(left, nums[i])
	}
	if res >= math.MaxInt32 {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
