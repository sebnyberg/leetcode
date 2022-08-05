package p2354numberofexcellentpairs

import (
	"fmt"
	"math/bits"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countExcellentPairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{1, 2, 3, 1, 536870911}, 3, 12},
		{[]int{1, 2, 3, 1}, 3, 5},
		{[]int{1, 2, 3, 1}, 3, 5},
		{[]int{5, 1, 1}, 10, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countExcellentPairs(tc.nums, tc.k))
		})
	}
}

func countExcellentPairs(nums []int, k int) int64 {
	// De-duplicate
	sort.Ints(nums)
	var j int
	for i := range nums {
		if nums[i] == nums[j] {
			continue
		}
		j++
		nums[j] = nums[i]
	}
	nums = nums[:j+1]

	// The core insight is that AND + OR is the same as counting the bits
	// from both numbers.
	// This is true because IF the bit is set in one number, then OR will yield 1,
	// and if the bit is set in both numbers, then OR yields 1 and AND yields 1.
	var bitCount [33]int
	for _, x := range nums {
		bitCount[bits.OnesCount(uint(x))]++
	}
	var res int64
	for i := 0; i < 32; i++ {
		for j := max(0, k-i); j < 32; j++ {
			res += int64(bitCount[i] * bitCount[j])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
