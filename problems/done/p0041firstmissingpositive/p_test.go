package p0041firstmissingpositive

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_firstMissingPositive(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{}, 1},
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 20}, 10},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, firstMissingPositive(tc.nums))
		})
	}
	incr := make([]int, 100)
	for i := range incr {
		incr[i] = i + 1
	}
	require.Equal(t, 101, firstMissingPositive(incr))
}

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	// Create a bitmap big enough to house the numbers
	bitmap := make([]uint64, len(nums)/64+1)
	for _, n := range nums {
		if n < 0 {
			continue
		}
		partition := n / 64
		if partition > len(bitmap) {
			continue
		}
		bitmap[partition] |= 1 << (n % 64)
	}

	// Walk through bits in the bitmap until we find an unset bit
	min := 1
	part := 0
	bitmap[0] >>= 1
	for part < len(bitmap) && bitmap[part]&1 == 1 {
		min++
		bitmap[part] >>= 1
		if min%64 == 0 {
			part++
		}
	}
	return min
}
