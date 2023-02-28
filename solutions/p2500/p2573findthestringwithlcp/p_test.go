package p2573findthestringwithlcp

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_findTheString(t *testing.T) {
	for i, tc := range []struct {
		lcp  [][]int
		want string
	}{
		{
			leetcode.ParseMatrix("[[4,0,2,0],[0,3,0,1],[2,0,2,0],[0,1,0,1]]"),
			"abab",
		},
		{
			leetcode.ParseMatrix("[[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,1]]"),
			"aaaa",
		},
		{
			leetcode.ParseMatrix("[[4,3,2,1],[3,3,2,1],[2,2,2,1],[1,1,1,3]]"),
			"",
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, findTheString(tc.lcp))
		})
	}
}

func findTheString(lcp [][]int) string {
	// If lcp[i][j] >= 0, then a[i] == a[j]
	//
	// To form the lexicographically smallest string, a[0] must be 'a'. Then any
	// time lcp[0][j] >= 0, a[j] must also be 'a'. If a[0][1] == 0, then a[1] !=
	// a[0], and a[1] must be 'b'.
	//
	// If there are no more characters left in the alphabet, or if there is some
	// inconsistency in the assignment, e.g. a[0] == a[1] and a[0] == a[2] but
	// a[1] != a[2], then there is no solution.
	//
	//
	n := len(lcp)
	a := make([]byte, n)
	var c byte = 'a'
	for i := range lcp {
		if a[i] == 0 {
			if c > 'z' {
				return ""
			}
			a[i] = c
			c++
		}
		for j := range lcp[i] {
			if lcp[i][j] == 0 {
				continue
			}
			if a[j] != 0 && a[j] != a[i] {
				return ""
			}
			a[j] = a[i]
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if a[i] == a[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			}
			if dp[i][j] != lcp[i][j] {
				return ""
			}
		}
	}

	return string(a)
}
