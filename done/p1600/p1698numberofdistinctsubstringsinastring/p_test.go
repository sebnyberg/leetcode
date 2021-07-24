package p1698numberofdistinctsubstringsinastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countDistinct(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"abcdefg", 28},
		{"aabbaba", 21},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, countDistinct(tc.s))
		})
	}
}

func countDistinct(s string) int {
	n := len(s)
	mod := 1_000_000_007
	base := 29
	seen := make(map[int][][2]int)
	for l := 0; l < n; l++ {
		h := 0
		for i := l; i < n; i++ {
			h = (h*base + int(s[i]-'a'+1)) % mod
			seen[h] = append(seen[h], [2]int{l, i})
		}
	}
	var count int
	for _, positions := range seen {
		if len(positions) <= 1 {
			count++
			continue
		}
		vals := make(map[string]struct{})
		for _, pos := range positions {
			start, end := pos[0], pos[1]
			vals[s[start:end+1]] = struct{}{}
		}
		count += len(vals)
	}
	return count
}
