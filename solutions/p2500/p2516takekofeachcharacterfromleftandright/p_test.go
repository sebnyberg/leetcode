package p2516takekofeachcharacterfromleftandright

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_takeCharacters(t *testing.T) {
	for i, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"aabaaacaabc", 2, 8},
		{"a", 1, -1},
		{"a", 0, -1},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, takeCharacters(tc.s, tc.k))
		})
	}
}

func takeCharacters(s string, k int) int {
	if k == 0 {
		return 0
	}
	var count [3]int
	for i := range s {
		count[s[i]-'a']++
	}
	if count[0] < k || count[1] < k || count[2] < k {
		return -1
	}

	// Read from count until the count satisfies the request
	count = [3]int{}
	n := len(s)
	r := n - 1
	var got int
	for {
		count[s[r]-'a']++
		if count[s[r]-'a'] == k {
			got++
			if got == 3 {
				break
			}
		}
		r--
	}
	res := n - r

	// Remove from right until unsatisfied, then add from left until satisfied
	// again.
	for l := range s {
		// Remove from right while ok
		count[s[l]-'a']++
		for r < n && count[s[r]-'a'] > k {
			count[s[r]-'a']--
			r++
		}
		res = min(res, l+1+(n-r))
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
