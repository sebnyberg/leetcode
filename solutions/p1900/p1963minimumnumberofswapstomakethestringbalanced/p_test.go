package p1963minimumnumberofswapstomakethestringbalanced

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minSwaps(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"][][", 1},
		{"]]][[[", 2},
		{"[]", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, minSwaps(tc.s))
		})
	}
}

func minSwaps(s string) int {
	n := len(s)
	ss := []byte(s)
	revb := []byte(s)
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		revb[l], revb[r] = revb[r], revb[l]
	}
	var llparen int
	var rrparen int
	var swaps int
	l, r := 0, 0
	for {
		// Move l until finding an unmatched bracket
		for ; l < n-r; l++ {
			if ss[l] == ']' {
				if llparen > 0 {
					llparen--
				} else {
					break
				}
			} else {
				llparen++
			}
		}
		// Move r until finding an unmatched bracket
		for ; r < n-l; r++ {
			if revb[r] == '[' {
				if rrparen > 0 {
					rrparen--
				} else {
					break
				}
			} else {
				rrparen++
			}
		}
		if l+r >= n {
			break
		}

		// Swap
		ss[l] = '['
		revb[r] = ']'
		swaps++
	}
	return swaps
}
