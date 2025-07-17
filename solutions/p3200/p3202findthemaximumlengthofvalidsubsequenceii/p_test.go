package p3202findthemaximumlengthofvalidsubsequenceii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumLength(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 5},
		{[]int{1, 4, 2, 3, 1, 4}, 3, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maximumLength(tc.nums, tc.k))
		})
	}
}

func maximumLength(nums []int, k int) int {
	// The goal is to select numbers from nums such that the combination of
	// numbers mod k is equal.
	//
	// First thing to note is that k <= 10^3, and for any numbers x and y, (x+y)%k
	// is equal to (x%k + y%k) % k. This means that we can take all numbers in
	// nums mod k before finding the solution.

	for i := range nums {
		nums[i] %= k
	}

	// There are two types of sequences: either all numbers are equal, or all
	// numbers alternate between two numbers.
	// So, at each number, the longest possible subsequence is the longest out of:
	// (1) count of prior occurrences of that same number, or
	// (2) count of prior subsequences which contain the same number and last
	//     ended with a different number.

	// Let's store sequence lengths in a map where the key is either a single
	// number, or a pair of numbers.
	seqLen := make(map[[2]int]int)
	seen := make(map[int]struct{})

	for _, x := range nums {
		seen[x] = struct{}{}

		// For each seen number
		for y := range seen {
			if y == x {
				// And add the current number to a "dummy" sequence containing only this
				// number.
				seqLen[[2]int{x, x}]++
				continue
			}

			// Start a new sequence or add onto a prior one
			seqLen[[2]int{y, x}] = max(2, seqLen[[2]int{x, y}]+1)
		}
	}

	var res int
	for _, c := range seqLen {
		res = max(res, c)
	}
	return res
}
