package p0525contiguousarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMaxLength(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{0, 1, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findMaxLength(tc.nums))
		})
	}
}

func findMaxLength(nums []int) int {
	n := len(nums)
	// valIdx stores the first index of a sum
	valIdx := make(map[int]int, n/2)
	valIdx[0] = -1
	var res int
	var val int
	for i, n := range nums {
		if n == 0 {
			n = -1
		}
		val += n
		if idx, exists := valIdx[val]; exists {
			if d := i - idx; d > res {
				res = d
			}
		} else {
			valIdx[val] = i
		}
	}
	return res
}
