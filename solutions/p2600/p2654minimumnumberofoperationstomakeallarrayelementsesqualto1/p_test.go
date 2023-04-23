package p2654minimumnumberofoperationstomakeallarrayelementsesqualto1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{6, 10, 15}, 4},
		{[]int{1, 1}, 0},
		{[]int{2, 6, 3, 4}, 4},
		{[]int{2, 10, 6, 14}, -1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums))
		})
	}
}

func minOperations(nums []int) int {
	// I did not notice that the only valid operations where regarding adjacent
	// values... This is much simpler than I thought.
	// The goal is to find the shortest subarray such that the GCD of the
	// subarray is equal to 1.
	// For this we can simply GCD each possible subarray.
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	n := len(nums)
	var ones int
	for i := range nums {
		if nums[i] == 1 {
			ones++
		}
	}
	if ones > 0 {
		return n - ones
	}

	// Find the smallest subarray that has gcd == 1
	for k := 2; k <= len(nums); k++ {
		for l := 0; l < n-k+1; l++ {
			val := nums[l]
			for r := l + 1; r < l+k; r++ {
				val = gcd(val, nums[r])
			}
			if val == 1 {
				return n + k - 2
			}
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
