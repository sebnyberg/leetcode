package p1016binarystringwithsubstringsrepresenting1ton

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_queryString(t *testing.T) {
	for i, tc := range []struct {
		s    string
		n    int
		want bool
	}{
		{"1", 1, true},
		{"10010111100001110010", 10, false},
		{"0110", 3, true},
		{"0110", 4, false},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, queryString(tc.s, tc.n))
		})
	}
}

func queryString(s string, n int) bool {
	// As a matter of fact, the string cannot represent numbers that are too
	// large. There are 10 bits in a 10^3 number, and so it would require at
	// least 1000 characters to cover all possible numbers smaller than 1000.

	// This means that we can simply try all numbers < n until we hit a brick
	// wall.
	var seen [1e4]bool
	markSeen := func(m int) {
		val := 0
		mod := (1 << m)
		for i := range s {
			val <<= 1
			val += int(s[i] - '0')
			if i < m-1 {
				continue
			}
			val %= mod
			seen[val] = true
		}
	}
	bs := fmt.Sprintf("%b", n)
	x := 1
	ub := 2
	for m := 1; m <= len(bs); m++ {
		markSeen(m)
		for x < min(n+1, ub) {
			if !seen[x] {
				return false
			}
			x++
		}
		if x > n {
			break
		}
		ub <<= 1
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
