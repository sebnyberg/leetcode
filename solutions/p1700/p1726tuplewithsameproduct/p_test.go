package p1726tuplewithsameproduct

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_tupleSameProduct(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 4, 6}, 8},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, tupleSameProduct(tc.nums))
		})
	}
}

func tupleSameProduct(nums []int) int {
	// An immediate insight is that there are ~10^12 unique quadruplets of
	// numbers, which is not feasible to explore.
	// However, there are only 10^6 different pairs, and checking equality between
	// pairs is easy - it's just the geometric sum of unique products * 8 (perms).
	prodCount := make(map[int]int)
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			prodCount[nums[i]*nums[j]]++
		}
	}
	var res int
	for _, c := range prodCount {
		res += 8 * c * (c - 1) / 2
	}
	return res
}
