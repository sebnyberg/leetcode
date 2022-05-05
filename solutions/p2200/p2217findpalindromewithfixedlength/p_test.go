package p2217findpalindromewithfixedlength

import (
	"fmt"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kthPalindrome(t *testing.T) {
	for _, tc := range []struct {
		queries   []int
		intLength int
		want      []int64
	}{
		{[]int{1, 2, 3, 4, 5, 90}, 3, []int64{101, 111, 121, 131, 141, 999}},
		{[]int{2, 4, 6}, 4, []int64{1111, 1331, 1551}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.queries), func(t *testing.T) {
			require.Equal(t, tc.want, kthPalindrome(tc.queries, tc.intLength))
		})
	}
}

func kthPalindrome(queries []int, intLength int) []int64 {
	// Another fucking palindrome problem.... jeez

	// As usual with palindromes, we don't actually have to consider both sides,
	// only one side. In this scenario, the left side is preferable.
	// For odd intLengths, we have to disregard the right-most number when forming
	// the result.

	// So with e.g. intlength 4, the smallest possible first integer is 10, then
	// 11, 12, ... 99.
	// For 5, it's 100, 101, 102, ... 999
	// For 6, it's the same as for 5
	digits := ((intLength + 1) / 2) - 1
	start := int(math.Pow10(digits))
	end := int(math.Pow10(digits+1)) - 1
	maxIdx := end - start + 1
	res := make([]int64, len(queries))
	for i, q := range queries {
		if q > maxIdx {
			res[i] = -1
			continue
		}
		s := fmt.Sprint(start + q - 1)
		ss := s + rev(s[:len(s)-intLength&1])
		x, _ := strconv.Atoi(ss)
		res[i] = int64(x)
	}

	return res
}

func rev(s string) string {
	b := []byte(s)
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
	return string(b)
}
