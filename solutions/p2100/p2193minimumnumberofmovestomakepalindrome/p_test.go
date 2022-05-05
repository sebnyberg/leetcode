package p2193

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minMovesToMakePalindrome(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"eqvvhtcsaaqtqesvvqch", 17},
		{"aabb", 2},
		{"letelt", 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minMovesToMakePalindrome(tc.s))
		})
	}
}

func minMovesToMakePalindrome(s string) int {
	n := len(s)
	bs := []byte(s)
	var swaps int
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		if bs[l] == bs[r] {
			continue
		}
		// ll is l's closest counterpart on the right side
		var ll int
		for ll = r; ll > l && bs[ll] != bs[l]; ll-- {
		}
		// rr is r's closest counterpart on the left side
		var rr int
		for rr = l; rr < r && bs[rr] != bs[r]; rr++ {
		}
		if r-ll < rr-l {
			// Keep bs[l], move bs[ll] into position
			swaps += r - ll
			copy(bs[ll:], bs[ll+1:])
		} else {
			// Keep bs[r], move bs[rr] into position
			swaps += rr - l
			copy(bs[1:], bs[:rr])
		}
	}

	return swaps
}
