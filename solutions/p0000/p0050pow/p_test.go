package p0050pow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_myPow(t *testing.T) {
	for _, tc := range []struct {
		x    float64
		n    int
		want float64
	}{
		{2, -2147483648, 0},
		{2, 10, 1024},
		{2, -2, 0.2500},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.x, tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, myPow(tc.x, tc.n))
		})
	}
}

func myPow(x float64, n int) (res float64) {
	switch {
	case n < 0:
		return 1 / x * myPow(1/x, -(n+1))
	case n == 0:
		return 1
	case n == 1:
		return x
	case n == 2:
		return x * x
	case n%2 == 0:
		return myPow(myPow(x, n/2), 2)
	default:
		return x * myPow(myPow(x, n/2), 2)
	}
}
