package p2911minimumchangestomakeksemipalindromes

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumChanges(t *testing.T) {
	for i, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"abcac", 2, 1},
		{"abcdef", 2, 2},
		{"aa55aa", 3, 0},
		{"abcc", 1, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumChanges(tc.s, tc.k))
		})
	}
}

func minimumChanges(s string, k int) int {
	n := len(s)

	// Pre-calculate the cost of making a semi-palindrome of s[i:j]
	palinCost := make([][]int, n+1)
	for i := range palinCost {
		palinCost[i] = make([]int, n+1)
		for j := range palinCost {
			if j > i {
				palinCost[i][j] = minPalinCost(s[i:j])
			} else {
				palinCost[i][j] = math.MaxInt32
			}
		}
	}

	// cost[i] = cost of making k substrings of the string s[i:]
	cost := make([]int, n+1)
	for i := range cost {
		// The cost of making 1 substring is given by the palindrome cost
		cost[i] = palinCost[i][n]
	}
	next := make([]int, n+1)

	// For an ever increasing amount of substrings
	for kk := 2; kk <= k; kk++ {
		for i := range next {
			next[i] = math.MaxInt32
		}

		// Cut the string s[i:] into two parts, s[i:j] and s[j:]
		//
		// The left side will have one substring, and s[j:] will contain the
		// other substrings.
		//
		// Try all such cuts and see which one is optimal.
		end := n - (kk - 1)
		for i := 0; i < end; i++ {
			for j := i + 1; j <= end; j++ {
				c := palinCost[i][j]
				next[i] = min(next[i], c+cost[j])
			}
		}

		cost, next = next, cost
	}

	return cost[0]
}

func minPalinCost(s string) int {
	n := len(s)
	res := math.MaxInt32
	for d := 1; d < n; d++ {
		if n%d != 0 {
			continue
		}
		var cost int
		for i := 0; i < d; i++ {
			for l, r := i, n+i-d; l < r; l, r = l+d, r-d {
				if s[l] != s[r] {
					cost++
				}
			}
		}
		res = min(res, cost)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
