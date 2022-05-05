package p1839longestsubstringovallvowelsinorder

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxFrequency(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 4}, 5, 3},
		{[]int{3, 9, 6}, 2, 1},
		{[]int{1, 4, 8, 13}, 5, 2},
		{
			[]int{9930, 9923, 9983, 9997, 9934, 9952, 9945, 9914, 9985, 9982, 9970, 9932, 9985, 9902, 9975, 9990, 9922, 9990, 9994, 9937, 9996, 9964, 9943, 9963, 9911, 9925, 9935, 9945, 9933, 9916, 9930, 9938, 10000, 9916, 9911, 9959, 9957, 9907, 9913, 9916, 9993, 9930, 9975, 9924, 9988, 9923, 9910, 9925, 9977, 9981, 9927, 9930, 9927, 9925, 9923, 9904, 9928, 9928, 9986, 9903, 9985, 9954, 9938, 9911, 9952, 9974, 9926, 9920, 9972, 9983, 9973, 9917, 9995, 9973, 9977, 9947, 9936, 9975, 9954, 9932, 9964, 9972, 9935, 9946, 9966},
			3056, 73,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxFrequency(tc.nums, tc.k))
		})
	}
}

func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)

	// Convert nums to be the distance between the previous element to the current
	n := len(nums)

	// The solution is the longest length in dist for which the sum is <= k
	l := 0
	r := 1
	sum := 0
	res := 0
	cur := nums[0]

	for r < n {
		switch {
		case sum <= k:
			res = max(res, r-l)
			// Add a new number
			if nums[r] != cur {
				sum += (r - l) * (nums[r] - cur)
				cur = nums[r]
			}
			r++
		case sum > k:
			sum -= (cur - nums[l])
			l++
		}
	}
	if sum <= k {
		res = max(res, r-l)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
