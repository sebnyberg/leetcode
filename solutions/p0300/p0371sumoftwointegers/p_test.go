package p0371sumoftwointegers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getSum(t *testing.T) {
	for _, tc := range []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.a), func(t *testing.T) {
			require.Equal(t, tc.want, getSum(tc.a, tc.b))
		})
	}
}

func getSum(a int, b int) int {
	x, y := abs(a), abs(b)
	if x < y {
		x, y = y, x
		a, b = b, a
	}
	sign := 1
	if a < 0 {
		sign = -1
	}
	if a*b >= 0 {
		for y > 0 {
			x, y = x^y, (x&y)<<1
		}
	} else {
		for y > 0 {
			x, y = x^y, (^x&y)<<1
		}
	}
	return x * sign
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
