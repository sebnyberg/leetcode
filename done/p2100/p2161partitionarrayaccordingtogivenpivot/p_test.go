package p2161partitionarrayaccordingtogivenpivot

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pivotArray(t *testing.T) {
	for _, tc := range []struct {
		nums  []int
		pivot int
		want  []int
	}{
		{[]int{9, 12, 5, 10, 14, 3, 10}, 10, []int{9, 5, 3, 10, 10, 12, 14}},
		{[]int{-3, 4, 3, 2}, 2, []int{-3, 2, 4, 3}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, pivotArray(tc.nums, tc.pivot))
		})
	}
}

func pivotArray(nums []int, pivot int) []int {
	var res []int
	for _, n := range nums {
		if n < pivot {
			res = append(res, n)
		}
	}
	for _, n := range nums {
		if n == pivot {
			res = append(res, n)
		}
	}
	for _, n := range nums {
		if n > pivot {
			res = append(res, n)
		}
	}
	return res
}
