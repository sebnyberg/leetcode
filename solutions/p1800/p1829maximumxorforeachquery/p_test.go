package p1829maximumxorforeachquery

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getMaximumXor(t *testing.T) {
	for _, tc := range []struct {
		nums       []int
		maximumBit int
		want       []int
	}{
		{[]int{0, 1, 1, 3}, 2, []int{0, 3, 2, 3}},
		{[]int{2, 3, 4, 7}, 3, []int{5, 2, 6, 5}},
		{[]int{0, 1, 2, 2, 5, 7}, 3, []int{4, 3, 6, 4, 6, 7}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, getMaximumXor(tc.nums, tc.maximumBit))
		})
	}
}

func getMaximumXor(nums []int, maximumBit int) []int {
	var val int
	res := make([]int, len(nums))
	mask := 0
	for i := 1; i <= maximumBit; i++ {
		mask <<= 1
		mask += 1
	}
	for i, n := range nums {
		val ^= n
		maxVal := ^val & mask
		res[len(nums)-1-i] = maxVal
	}
	return res
}
