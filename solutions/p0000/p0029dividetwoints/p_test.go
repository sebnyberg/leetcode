package p0029dividetwoints

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divide(t *testing.T) {
	for _, tc := range []struct {
		dividend int
		divisor  int
		want     int
	}{
		{10, 3, 3},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
		{-1, 1, -1},
		{1, -1, -1},
		{-2147483648, -1, 2147483647},
	} {
		t.Run(fmt.Sprintf("%v/%v", tc.dividend, tc.divisor), func(t *testing.T) {
			require.Equal(t, tc.want, divide(tc.dividend, tc.divisor))
		})
	}
}

func divide(dividend int, divisor int) int {
	// Remove and store sign
	var negative bool
	if dividend < 0 {
		negative = true
		dividend = -dividend
	}
	if divisor < 0 {
		negative = !negative
		divisor = -divisor
	}

	quotient := 0
	for dividend >= divisor {
		// Find MSB of dividend
		msb := 1 << 30
		for ; !(dividend&msb == msb); msb >>= 1 {
		}

		// Shift divisor until aligned with dividend
		tmpquotient := 1
		tmpdivisor := divisor
		for ; !(tmpdivisor&msb == msb); tmpdivisor <<= 1 {
			tmpquotient <<= 1
		}

		// Ensure that divisor-dividend >= 0
		if tmpdivisor > dividend {
			tmpquotient >>= 1
			tmpdivisor >>= 1
		}

		dividend -= tmpdivisor
		quotient += tmpquotient
	}

	if negative {
		return -1 * quotient
	}
	if quotient > math.MaxInt32 {
		return math.MaxInt32
	}
	return quotient
}
