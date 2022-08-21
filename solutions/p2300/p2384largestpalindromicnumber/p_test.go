package p2384largestpalindromicnumber

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestPalindromic(t *testing.T) {
	for _, tc := range []struct {
		num  string
		want string
	}{
		{"444947137", "7449447"},
		{"00009", "9"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.num), func(t *testing.T) {
			require.Equal(t, tc.want, largestPalindromic(tc.num))
		})
	}
}

func largestPalindromic(num string) string {
	var digitCount [10]int
	for _, digit := range num {
		digitCount[digit-'0']++
	}

	// Let's start by choosing the middle
	// It should be the largest number which has an odd count
	middle := ""
	for d, count := range digitCount {
		if count&1 == 1 {
			middle = fmt.Sprintf("%d", d)
		}
	}

	right := make([]byte, 0)
	for d, count := range digitCount {
		if count >= 2 {
			right = append(right, bytes.Repeat([]byte{byte(d + '0')}, count/2)...)
		}
	}
	if len(right) > 0 && right[len(right)-1] == '0' {
		if middle != "" {
			return middle
		}
		return "0"
	}
	left := make([]byte, len(right))
	copy(left, right)
	for l, r := 0, len(left)-1; l < r; l, r = l+1, r-1 {
		left[l], left[r] = left[r], left[l]
	}
	ss := string(left) + middle + string(right)
	return ss
}
