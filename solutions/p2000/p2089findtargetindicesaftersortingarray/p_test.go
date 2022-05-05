package p2089findtargetindicesaftersortingarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_targetIndices(t *testing.T) {
	for _, tc := range []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{1, 2, 5, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 5, 2, 3}, 3, []int{3}},
		{[]int{1, 2, 5, 2, 3}, 5, []int{4}},
		{[]int{1, 2, 5, 2, 3}, 4, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, targetIndices(tc.nums, tc.target))
		})
	}
}

func targetIndices(nums []int, target int) []int {
	var idx, count int
	for _, n := range nums {
		if n < target {
			idx++
		}
		if n == target {
			count++
		}
	}
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = i + idx
	}
	return res
}
