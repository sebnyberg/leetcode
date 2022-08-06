package p0775globalandlocalinversions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isIdealPermutation(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want bool
	}{
		{[]int{1, 0, 2}, true},
		{[]int{1, 2, 0}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, isIdealPermutation(tc.nums))
		})
	}
}

func isIdealPermutation(nums []int) bool {
	// Local inversions are easy - just count
	// Global inversions are trickier - you must know how many elements are
	// than the current one for a given position.
	// We can use a BIT to count number of occurrences below a certain value but
	// it feels overkill for a medium problem
	// If there is a local inversion, then there is also a global inversion.
	// However, not all global inversions are local inversions.
	// For example, if nums[i] > nums[j] and j = i + 2, then there are more global
	// inversions.
	// This gives us the solution: if there is a number smaller than a number at
	// least two steps prior, then there are more global than local inversions.
	n := len(nums)
	currMin := nums[n-1]
	for i := n - 3; i >= 0; i-- {
		if nums[i] > currMin {
			return false
		}
		currMin = min(currMin, nums[i+1])
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
