package p0315countofsmallernumbersafterself

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSmaller(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1}, []int{0}},
		{[]int{-1, -1}, []int{0, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countSmaller(tc.nums))
		})
	}
}

func countSmaller(nums []int) []int {
	// Idea: shift each number in nums by the smallest possible number
	// Then create a BIT where the index is the number, and the value at each
	// index is the number of occurrences at that index
	offset := 10000
	size := 20000 + 2
	tree := make([]uint32, size)

	query := func(idx uint32) uint32 {
		var res uint32
		for idx > 0 {
			res += tree[idx]
			idx -= idx & -idx // remove MSB
		}
		return res
	}
	update := func(idx int, val uint32) {
		idx++ // BIT index is one-indexed
		for idx < size {
			tree[idx] += val
			idx += idx & -idx
		}
	}

	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[i] = int(query(uint32(nums[i] + offset)))
		update(nums[i]+offset, 1)
	}
	return res
}
