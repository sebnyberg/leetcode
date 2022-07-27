package p2318numberofdistinctrollsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_distinctSequences(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{4, 184},
		{2, 22},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, distinctSequences(tc.n))
		})
	}
}

func distinctSequences(n int) int {
	if n == 1 {
		return 6
	}
	if n == 2 {
		return 22
	}
	valid := [7][7]bool{}
	for a := 1; a < 6; a++ {
		for b := a + 1; b <= 6; b++ {
			c := a
			d := b
			for d != 0 {
				c, d = d, c%d
			}
			if c == 1 {
				valid[a][b] = true
				valid[b][a] = true
			}
		}
	}

	const mod = 1e9 + 7
	var dp [2][7][7]int
	for a := 1; a <= 6; a++ {
		for b := 1; b <= 6; b++ {
			if valid[a][b] {
				dp[0][a][b] = 1
			}
		}
	}

	for k := 0; k < n-2; k++ {
		for a := 1; a <= 6; a++ {
			for b := 1; b <= 6; b++ {
				if a == b || !valid[a][b] {
					continue
				}
				for c := 1; c <= 6; c++ {
					if c == a || c == b || !valid[b][c] {
						continue
					}
					dp[1][b][c] = (dp[1][b][c] + dp[0][a][b]) % mod
				}
			}
		}
		dp[0] = dp[1]
		dp[1] = [7][7]int{}
	}
	var res int
	for a := 1; a <= 6; a++ {
		for b := 1; b <= 6; b++ {
			res += (dp[0][a][b]) % mod
		}
	}

	return res % mod
}
