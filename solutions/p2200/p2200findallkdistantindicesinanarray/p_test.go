package p2200findallkdistantindicesinanarray

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findKDistantIndices(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		key  int
		k    int
		want []int
	}{
		{[]int{3, 4, 9, 1, 3, 9, 5}, 9, 1, []int{1, 2, 3, 4, 5, 6}},
		{[]int{2, 2, 2, 2, 2}, 2, 2, []int{0, 1, 2, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findKDistantIndices(tc.nums, tc.key, tc.k))
		})
	}
}

func findKDistantIndices(nums []int, key int, k int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)
	for i := range left {
		right[i] = math.MaxInt32
		left[i] = math.MaxInt32
	}
	for i := range nums {
		if nums[i] == key {
			left[i] = 0
		} else if i > 0 {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] == key {
			right[i] = 0
		} else if i < n-1 {
			right[i] = right[i+1] + 1
		}
	}
	var res []int
	for i := range nums {
		if left[i] <= k || right[i] <= k {
			res = append(res, i)
		}
	}
	return res
}
