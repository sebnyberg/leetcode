package p1044longestduplicatesubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestDupSubstring(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		{"dcsopfbhupztcyxctlyxocqwgcgydrxkbbeowdlkcehhslmidwphslbf", "hsl"},
		{"banana", "ana"},
		{"abcd", ""},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, longestDupSubstring(tc.s))
		})
	}
}

func longestDupSubstring(s string) string {
	n := len(s)
	l, r := 0, n
	bs := []byte(s)
	for l < r {
		mid := l + (r-l)/2
		if i := firstDuplicateOfLen(bs, mid); i != -1 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	idx := firstDuplicateOfLen(bs, l-1)
	return s[idx : idx+l-1]
}

func firstDuplicateOfLen(bs []byte, width int) int {
	if width == 0 {
		return 0
	}
	base := 29
	// mod := 1_000_000_007
	mod := 16777619341237
	pow := 1
	h := 0
	seen := make(map[int][]int, len(bs)-width)
	// seen := make(map[int]struct{}, len(bs)-width)
	for i, ch := range bs {
		h = (h*base + int(ch-'a'+1)) % mod
		if i < width {
			pow = pow * base % mod
		}
		if i >= width {
			// remove overflow element
			h = (h - int(bs[i-width]-'a'+1)*pow%mod + mod) % mod
		}
		if i >= width-1 {
			if _, exists := seen[h]; !exists {
				seen[h] = []int{i - width + 1}
				continue
			}
			// We have a hash match, either by collision or actual match, verify
			// using the indices already in the map
			for _, idx := range seen[h] {
				s := string(bs[i-width+1 : i+1])
				if s == string(bs[idx:idx+width]) {
					return idx
				}
			}
			seen[h] = append(seen[h], i-width+1)
		}
	}
	return -1
}
