package p2981findlongestspecialsubstringthatoccursthricei

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumLength(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"aaaa", 2},
		{"abcdef", -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, maximumLength(tc.s))
		})
	}
}

func maximumLength(s string) int {
	// We may use binary search to find the maximum length substring that occurs
	// thrice
	check := func(n int) bool {
		var count [26]int
		var l, r int
		for r < len(s) {
			for r < len(s) && s[r] == s[l] {
				r++
			}
			ch := int(s[l] - 'a')
			m := r - l
			count[ch] += max(0, m-n)
			if count[ch] >= 3 {
				return true
			}
			l = r
		}
		return false
	}
	lo := 0
	hi := math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if hi == 0 {
		return -1
	}
	return hi
}
