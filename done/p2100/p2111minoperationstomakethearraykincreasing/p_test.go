package p2111minoperationstomakethearraykincreasing

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{5, 4, 3, 2, 1}, 4},
		{[]int{1, 2, 3, 15, 7, 9}, 1},
		{[]int{3, 3, 3, 3, 2, 3, 15, 3, 3}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.arr))
		})
	}
}

func Test_kIncreasing(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{12, 6, 12, 6, 14, 2, 13, 17, 3, 8, 11, 7, 4, 11, 18, 8, 8, 3}, 1, 12},
		// {[]int{5, 4, 3, 2, 1}, 1, 4},
		// {[]int{4, 1, 5, 2, 6, 2}, 2, 0},
		// {[]int{4, 1, 5, 2, 6, 2}, 3, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, kIncreasing(tc.arr, tc.k))
		})
	}
}

func kIncreasing(arr []int, k int) int {
	// k dictates which slices of elements must be considered
	// These slices are independent from each-other, and such the solution to the
	// min number of operations for the general case of k=1 can be applied to all
	// other k as well.

	// We can always pick the most common element and adjust all others to match
	// that element for a cost of len(slice)-count[mostFrequentNumber]

	// We can find the longest non-decreasing sequence and adjust any numbers
	// not matching the sequence...
	// Let's try and get some more test-cases back to see what's best
	sub := make([]int, 0)
	var res int
	for offset := 0; offset < k; offset++ {
		sub = sub[:0]
		for i := offset; i < len(arr); i += k {
			sub = append(sub, arr[i])
		}
		res += minOperations(sub)
	}
	return res
}

func minOperations(arr []int) int {
	return len(arr) - lengthOfLIS(arr)
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, 0, len(nums))
	for _, num := range nums {
		insertPos := sort.Search(len(dp), func(i int) bool {
			return dp[i] > num
		})
		if insertPos == len(dp) {
			dp = append(dp, num)
		} else {
			dp[insertPos] = num
		}
	}
	return len(dp)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
