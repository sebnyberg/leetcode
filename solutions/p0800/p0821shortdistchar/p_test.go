package p0821shortdistchar

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_shortestToChar(t *testing.T) {
	for _, tc := range []struct {
		s    string
		c    byte
		want []int
	}{
		{"loveleetcode", 'e', []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"aaab", 'b', []int{3, 2, 1, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, shortestToChar(tc.s, tc.c))
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func shortestToChar(s string, c byte) []int {
	res := make([]int, utf8.RuneCountInString(s))
	prev, next := math.MinInt32, strings.IndexByte(s, c)
	for i := range s {
		res[i] = min(i-prev, next-i)
		if i == next {
			prev = next
			next = strings.IndexByte(s[prev+1:], c)
			if next == -1 {
				next = math.MaxInt32
			} else {
				next += prev + 1
			}
		}
	}
	return res
}
