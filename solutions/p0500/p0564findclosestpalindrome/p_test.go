package p0564findclosestpalindrome

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"12321", true},
		{"1", true},
		{"11", true},
		{"12", false},
		{"122", false},
		{"121", true},
		{"1223", false},
		{"1221", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, isPalindrome(tc.s))
		})
	}
}

func Test_nearestPalindromic(t *testing.T) {
	for _, tc := range []struct {
		n    string
		want string
	}{
		{"10", "9"},
		{"11", "9"},
		{"100", "99"},
		{"1", "0"},
		{"122", "121"},
		{"1222", "1221"},
		{"12", "11"},
		{"12345", "12321"},
		{"123345", "123321"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, nearestPalindromic(tc.n))
		})
	}
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func nearestPalindromic(n string) string {
	nlen := len(n)
	order := int(math.Pow10(nlen / 2))
	num, err := strconv.Atoi(n)
	check(err)

	noop := mirror(num)
	above := mirror(((num / order) * order) + order + 1)
	below := mirror((num/order)*order - 1)
	if noop > num {
		above = min(noop, above)
	} else if noop < num {
		below = max(noop, below)
	}
	if abs(above-num) < abs(below-num) {
		return strconv.Itoa(above)
	}
	return strconv.Itoa(below)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mirror(a int) int {
	s := []byte(strconv.Itoa(a))
	for l, r := 0, len(s)-1; l <= r; l, r = l+1, r-1 {
		s[r] = s[l]
	}
	n, err := strconv.Atoi(string(s))
	check(err)
	return n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func rev(s string) string {
	bs := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}

func isPalindrome(s string) bool {
	n := len(s)
	for l, r := (n-1)/2, n/2; l >= 0; l, r = l-1, r+1 {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}
