package p1062longestrepeatingsubstring

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestRepeatingSubstring(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"abcd", 0},
		{"abbaba", 2},
		{"aabcaabdaab", 3},
		{"aaaaa", 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestRepeatingSubstring(tc.s))
		})
	}
}

func longestRepeatingSubstring(s string) int {
	n := len(s)
	bs := []byte(s)
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if hasRepeatingSubstring(bs, mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}

func hasRepeatingSubstring(bs []byte, width int) bool {
	mod := 1_000_000_007
	pow := 1
	base := 29
	h := 0
	seen := make(map[int][]int, len(bs)-width)
	for i := range bs {
		h = (h*base + int(bs[i]-'a')) % mod
		if i < width {
			pow = pow * base % mod
		}
		if i >= width {
			h = (h - int(bs[i-width]-'a')*pow%mod + mod) % mod
		}
		if i >= width-1 {
			if _, exists := seen[h]; !exists {
				seen[h] = []int{i - width + 1}
				continue
			}
			for _, idx := range seen[h] {
				if bytes.Equal(bs[idx:idx+width], bs[i-width+1:i+1]) {
					return true
				}
			}
			seen[h] = append(seen[h], i-width+1)
		}
	}
	return false
}
