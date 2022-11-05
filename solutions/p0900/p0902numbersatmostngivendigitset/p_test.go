package p0902numbersatmostngivendigitset

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_atMostNGivenDigitSet(t *testing.T) {
	for i, tc := range []struct {
		digits []string
		n      int
		want   int
	}{
		{[]string{"2", "7", "8", "9"}, 1378912, 5460},
		{[]string{"1", "3", "5", "6", "7", "8"}, 62774961, 1222386},
		{[]string{"1", "7"}, 231, 10},
		{[]string{"1", "4", "9"}, 1000000000, 29523},
		{[]string{"1", "4", "9"}, 1000, 39},
		{[]string{"1", "4", "9"}, 10, 3},
		{[]string{"1", "3", "5", "7"}, 100, 20},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, atMostNGivenDigitSet(tc.digits, tc.n))
		})
	}
}

func atMostNGivenDigitSet(digits []string, n int) int {
	ds := strings.Join(digits, "")
	s := fmt.Sprint(n)
	var res int

	// Add all valid numbers with length < len(n)
	// For example, with n = 783 and digits = {1, 7} that would include
	// 1, 7, 11, 17, 71, 77
	combs := 1
	for j := len(s) - 1; j >= 1; j-- {
		combs *= len(ds)
		res += combs
	}

	// For each position i in s, calculate the number of valid numbers for which
	// the first digit is less than s[i]
	//
	// If the first digit is not a valid number, then it will have to be
	// reduced, in which case we are done.
	//
	// For example, with n = 783 and digits = {1, 7}
	//
	// add 1xx
	// continue with 7 fixed in place (7xx)
	//
	// add 71x
	// add 77x
	// 8 is not a valid number, so we stop
	//
	for i := 0; i < len(s); i++ {
		var j int
		for j < len(ds) && ds[j] < s[i] {
			res += combs
			j++
		}
		if j == len(ds) || s[i] != ds[j] {
			return res
		}
		combs /= len(ds)
	}
	// If we reached this far, then the entire number is valid
	// Count it, too
	return res + 1
}
