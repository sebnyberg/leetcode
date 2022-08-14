package p2375countsmallestnumberfromdistring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestNumber(t *testing.T) {
	for _, tc := range []struct {
		pattern string
		want    string
	}{
		{"IIIDIDDD", "123549876"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.pattern), func(t *testing.T) {
			require.Equal(t, tc.want, smallestNumber(tc.pattern))
		})
	}
}

func smallestNumber(pattern string) string {
	// Let's try everything
	var res string
	n := len(pattern)
	buf := make([]byte, n+1)
	dfs(pattern, buf, 0, 0, &res)
	return res
}

func dfs(pattern string, buf []byte, bm, i int, res *string) {
	if i == len(pattern)+1 {
		if s := string(buf); *res == "" || s < *res {
			*res = s
		}
		return
	}
	if i == 0 {
		// Just try all things
		for x, ch := 1, byte('1'); x <= 9; x, ch = x+1, ch+1 {
			buf[i] = ch
			dfs(pattern, buf, bm|(1<<x), i+1, res)
		}
		return
	}
	// Have to follow pattern
	if pattern[i-1] == 'D' {
		// Smaller than previous value
		for x, ch := 1, byte('1'); x <= 8; x, ch = x+1, ch+1 {
			if ch < buf[i-1] && bm&(1<<x) == 0 {
				buf[i] = byte(ch)
				dfs(pattern, buf, bm|(1<<x), i+1, res)
			}
		}
	} else {
		// Larger than previous value
		for x, ch := 2, byte('2'); x <= 9; x, ch = x+1, ch+1 {
			if ch > buf[i-1] && bm&(1<<x) == 0 {
				buf[i] = byte(ch)
				dfs(pattern, buf, bm|(1<<x), i+1, res)
			}
		}
	}
}
