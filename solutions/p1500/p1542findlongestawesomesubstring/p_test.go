package p1542findlongestawesomesubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestAwesome(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"185801630663498", 5},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, longestAwesome(tc.s))
		})
	}
}

func longestAwesome(s string) int {
	// Keep track of earliest occurrence of even/odd of each number, then try
	// all removals that would create either a) an even total count of all
	// characters, or b) an even total count for all characters except one.
	idx := make(map[int]int)
	idx[0] = -1 // all even at start
	var bm int
	res := 1
	for i, ch := range s {
		bm ^= (1 << int(ch-'0'))
		if j, exists := idx[bm]; exists {
			res = max(res, i-j)
		} else {
			idx[bm] = i
		}
		for k := 0; k <= 9; k++ {
			atleastoneodd := bm ^ (1 << k)
			if j, exists := idx[atleastoneodd]; exists {
				res = max(res, i-j)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
