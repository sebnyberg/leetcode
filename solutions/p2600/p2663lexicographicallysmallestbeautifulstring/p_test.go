package p2663lexicographicallysmallestbeautifulstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_smallestBeautifulString(t *testing.T) {
	for i, tc := range []struct {
		s    string
		k    int
		want string
	}{
		{"abcz", 26, "abda"},
		{"dc", 4, ""},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, smallestBeautifulString(tc.s, tc.k))
		})
	}
}

func smallestBeautifulString(s string, k int) string {
	b := []byte(s)
	var i int
	for i = len(s) - 1; i >= 0; i-- {
		var prevPrev byte = 0
		var prev byte = 1
		if i >= 2 {
			prevPrev = s[i-2]
		}
		if i >= 1 {
			prev = s[i-1]
		}
		// Set this position to the smallest value > s[i] such that it does not
		// become the same as prev or prevPrev
		ch := s[i] + 1
		for ch < byte('a'+k) && (ch == prev || ch == prevPrev) {
			ch++
		}
		if ch >= byte('a'+k) {
			b[i] = 0 // needs to be filled out later.
		} else {
			// success!
			b[i] = ch
			break
		}
	}
	if b[0] == 0 {
		return ""
	}
	for j := i + 1; j < len(s); j++ {
		var prevPrev byte = 0
		var prev byte = 1
		if j >= 2 {
			prevPrev = b[j-2]
		}
		if j >= 1 {
			prev = b[j-1]
		}
		var x int
		for x = 0; x < k; x++ {
			ch := byte(x + 'a')
			if ch != prev && ch != prevPrev {
				b[j] = ch
				break
			}
		}
		if b[j] == 0 {
			return ""
		}
	}
	return string(b)
}
