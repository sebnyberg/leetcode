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
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1534236469, 0},
		{1563847412, 0},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("%v", tc.x), func(t *testing.T) {
			require.Equal(t, tc.want, reverse(tc.x))
		})
	}
}

func reverse(x int) int {
	if x >= 1<<31 || x < -1<<31 {
		return 0
	}
	negative := x < 0
	if negative {
		x = -x
	}
	var res int
	for ; x > 0; x /= 10 {
		res = res*10 + x%10
	}
	if res >= 1<<31 || res < -1<<31 {
		return 0
	}
	if negative {
		return -res
	}
	return res
}
