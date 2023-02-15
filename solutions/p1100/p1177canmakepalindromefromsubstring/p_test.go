package p1177canmakepalindromefromsubstring

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_canMakePaliQueries(t *testing.T) {
	for i, tc := range []struct {
		s       string
		queries [][]int
		want    []bool
	}{
		{
			"abcda",
			leetcode.ParseMatrix("[[3,3,0],[1,2,0],[0,3,1],[0,3,2],[0,4,1]]"),
			[]bool{true, false, false, true, true},
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, canMakePaliQueries(tc.s, tc.queries))
		})
	}
}

func canMakePaliQueries(s string, queries [][]int) []bool {
	// Since the letters can be rearranged freely, we are only concerned with
	// odd counts of letters. If there are m odd letters, then k must be >= m-1
	//
	// To efficiently count letters, we can use a count-prefix sum
	n := len(s)
	pre := make([][26]int, n+1)
	var count [26]int
	for i := range s {
		count[s[i]-'a']++
		pre[i+1] = count
	}

	nq := len(queries)
	res := make([]bool, nq)
	for i, q := range queries {
		l, r := q[0], q[1]+1
		k := q[2]
		var nodd int
		for j := range pre[r] {
			nodd += (pre[r][j] - pre[l][j]) & 1
		}
		nodd -= (r - l) & 1 // skip middle odd
		res[i] = k*2 >= nodd
	}
	return res
}
