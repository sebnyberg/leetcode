package p0645setmismatch

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findErrorNums(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{8, 7, 3, 5, 3, 6, 1, 4}, []int{3, 2}},
		{[]int{2, 3, 2}, []int{2, 1}},
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{3, 2, 2}, []int{2, 1}},
		{[]int{1, 1}, []int{1, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, findErrorNums(tc.nums))
		})
	}
}

func findErrorNums(nums []int) (res []int) {
	var a int
	for i, n := range nums {
		a ^= n     // XOR once per number in nums
		a ^= i + 1 // XOR once per expected number
	}

	// find lowest bit, it belongs either to the missing or double value
	lowbit := a & -a
	b := a
	for i, n := range nums {
		// filter by low bit
		// this will skip either the missing or double value
		if n&lowbit > 0 {
			b ^= n
		}
		if (i+1)&lowbit > 0 {
			b ^= i + 1
		}
	}

	// XOR a with b to remove b from a
	a ^= b

	// At this stage, a and b are the double and missing values, but we don't
	// know which one is which! Do a final loop to figure out the return order.
	for _, n := range nums {
		if n == a {
			return []int{a, b}
		}
		if n == b {
			return []int{b, a}
		}
	}
	return nil
}
