package p0964leastoperatorstoexpressnumber

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_leastOpsExpressTarget(t *testing.T) {
	for i, tc := range []struct {
		x      int
		target int
		want   int
	}{
		{5, 501, 8},
		{100, 100000000, 3},
		{3, 19, 5},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, leastOpsExpressTarget(tc.x, tc.target))
		})
	}
}

func leastOpsExpressTarget(x int, target int) int {
	if target < x {
		return min(target*2-1, 2*(x-target))
	}
	if x == target {
		return 0
	}

	// create a power of x such that x^a >= target
	a := 1
	var p int
	for p = x; p < target; p *= x {
		a++
	}

	if p == target {
		return a - 1
	}

	// Now we can either choose the power below or above the target.
	// If we choose below, then we must find a way to construct the remainder
	// below
	res := math.MaxInt32
	if d := p - target; d < target {
		res = min(res, leastOpsExpressTarget(x, d)+a)
	}
	res = min(res, leastOpsExpressTarget(x, target-p/x)+a-1)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
