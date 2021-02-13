package p0069sqrtx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mySqrt(t *testing.T) {
	for _, tc := range []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.x), func(t *testing.T) {
			require.Equal(t, tc.want, mySqrt(tc.x))
		})
	}
}

func mySqrt(x int) int {
	c := x
	// Using Newton-Raphson
	// We are looking for the number x for which y=x^2=c
	// Start with guessing x_n = c
	// The tangent line at x_n is y = 2*x_n*x - x_n^2
	// The next guess is the place where the tangent line meets y=c, i.e.
	// c = 2 * x_n * x - x_n^2
	// <=>
	// c + x_n^2 = 2 * x_n * x
	// <=>
	// (1/2) (c + (x_n^2 / x_n))
	// <=>
	// (x_n + (c/x_n))/2
	xn := c
	for xn*xn > c {
		xn = (xn + (c / xn)) / 2
	}
	return xn
}
