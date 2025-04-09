package p3375minimumoperationstomakearrayvaluesequal

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{5, 2, 5, 4, 5}, 2, 2},
		{[]int{2, 1, 2}, 2, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums, tc.k))
		})
	}
}

func minOperations(nums []int, k int) int {
	m := make(map[int]struct{})
	for _, x := range nums {
		m[x] = struct{}{}
	}
	vals := make([]int, 0, len(m))
	for k := range m {
		vals = append(vals, k)
	}
	sort.Ints(vals)
	if vals[0] < k {
		return -1
	}
	if vals[0] == k {
		vals = vals[1:]
	}
	var ops int
	for i := len(vals) - 1; i >= 0; i-- {
		ops++
	}
	return ops
}
