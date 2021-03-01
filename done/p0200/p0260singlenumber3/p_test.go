package p0260singlenumber3

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_singleNumber(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, singleNumber(tc.nums))
		})
	}
}

func singleNumber(nums []int) []int {
	// Xoring a number with itself becomes zero
	// There are two of all numbers except the two target numbers,
	// c below will become a ^ b
	var c int
	for _, n := range nums {
		c ^= n
	}

	// Since a != b, at least one bit in c is 1, it can be found with the
	// right-most bit trick
	rightmostBit := c & -c
	var a int
	for _, n := range nums {
		if n&rightmostBit > 0 {
			a ^= n
		}
	}
	b := c ^ a
	return []int{a, b}
}
