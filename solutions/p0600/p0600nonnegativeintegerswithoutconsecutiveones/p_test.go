package p0600nonnegativeintegerswithoutconsecutiveones

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findIntegers(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{3, 3},
		{5, 5},
		{1, 2},
		{2, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findIntegers(tc.n))
		})
	}
}

func findIntegers(n int) int {
	// 0, and 1 are valid
	// Adding a zero to the start cannot change the number of valid combinations,
	// so the number of valid combinations for 0xxx is simply f[xxx]
	// Adding a 10 to the start for the result 10xx is simply f[xx]

	// f[1] is 2
	// f[2] is 3
	// f[3] is 0xx + 10x = f[2] + f[1] = 5
	// f[4] is 0xxx + 10xx = f[3] + f[2] = 5 + 3

	// This gives us the fibonacci series from 2 and forward.
	var f [32]int
	f[0] = 1
	f[1] = 2
	for i := 2; i < 32; i++ {
		f[i] = f[i-2] + f[i-1]
	}

	// How do we then count the number of valid integers <= n?

	// For each position, we have to check whether the 0xxxx combination or 10xxx
	// combination is guaranteed to fit the number

	// What if we have 1110101?

	// When there is a 1, we can consider all possible matches in the range
	// prior to that position. For example, with
	//   1110101
	// We consider the range
	// 1110101
	// 0000000-0111111
	// Then we can continue to the second number, and keep going until we are
	// done. In the end, we have tried all numbers smaller than 1110101 and the
	// only remaining number is the number itself.

	var res int
	s := fmt.Sprintf("%02b", n)
	for i, b := range s {
		if b == '0' {
			continue
		}
		res += f[len(s)-i-1]
		if i > 0 && s[i-1] == '1' && s[i] == '1' {
			return res
		}
	}
	return res + 1
}
