package p0342poweroffour

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfFour(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want bool
	}{
		{16, true},
		{5, false},
		{1, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, isPowerOfFour(tc.n))
			require.Equal(t, tc.want, isPowerOfFourBits(tc.n))
		})
	}

}

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}
	for n > 1 {
		if n%4 != 0 {
			return false
		}
		n /= 4
	}
	return true
}

func isPowerOfFourBits(n int) bool {
	isPowerOfTwo := n&(n-1) == 0
	oneLessEvenlyDivisibleBy3 := (n-1)%3 == 0
	return n > 0 && isPowerOfTwo && oneLessEvenlyDivisibleBy3
}
