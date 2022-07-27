package p0007reverseinteger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverse(t *testing.T) {
	tcs := []struct {
		x    int
		want int
	}{
		{1534236469, 0},
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1563847412, 0},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.x), func(t *testing.T) {
			require.Equal(t, tc.want, reverse(tc.x))
		})
	}
}

func reverse(x int) int {
	// This solution is an actual safe implementation on a system that does not
	// have any numbers beyond 32 bits. Most people on Leetcode just use a 64-bit
	// number and check near the end..
	val := int32(x)
	if int(val) != x {
		return 0
	}

	// There are more negative values than positive, so we work with negative
	// numbers and fix the sign at the end. That way, we can manage up to 2^31
	// values. If the sign is negative, and we get 2^31, then we've overflowed a
	// positive 32-bit integer.
	sign := 1
	if x >= 0 {
		sign = -1
		val = -val
	}

	var res int32
	var intMin int32 = -(1 << 31)
	for val < 0 {
		if res*10/10 != res {
			return 0
		}
		res *= 10
		if res < intMin-(val%10) {
			return 0
		}
		res += val % 10
		val /= 10
	}
	if sign == -1 && res == intMin {
		return 0
	}
	return sign * int(res)
}
