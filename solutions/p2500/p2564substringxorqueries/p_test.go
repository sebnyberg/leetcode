package p2564substringxorqueries

import (
	"fmt"
	"math"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_substringXorQueries(t *testing.T) {
	for i, tc := range []struct {
		s       string
		queries [][]int
		want    [][]int
	}{
		{
			"1",
			leetcode.ParseMatrix("[[4,5]]"),
			leetcode.ParseMatrix("[[0,0]]"),
		},
		{
			"101101",
			leetcode.ParseMatrix("[[0,5],[1,2]]"),
			leetcode.ParseMatrix("[[0,2],[2,3]]"),
		},
		{
			"0101",
			leetcode.ParseMatrix("[[12,8]]"),
			leetcode.ParseMatrix("[[-1,-1]]"),
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, substringXorQueries(tc.s, tc.queries))
		})
	}
}

func substringXorQueries(s string, queries [][]int) [][]int {
	// Shortest just means that we should remove leading zeroes.
	//
	// Either a match exists, or it doesn't.
	//
	// We could partition queries by string length, then run over S with that
	// window, matching as needed to avoid using too much storage.
	//
	n := len(queries)
	want := make(map[string][]int)
	var maxlen int
	minlen := math.MaxInt32
	for i, q := range queries {
		b := fmt.Sprintf("%b", q[0]^q[1])
		want[b] = append(want[b], i)
		maxlen = max(maxlen, len(b))
		minlen = min(minlen, len(b))
	}
	res := make([][]int, n)
	for i := range res {
		res[i] = []int{-1, -1}
	}
	for k := minlen; k <= maxlen; k++ {
		for i := k; i <= len(s); i++ {
			key := s[i-k : i]
			for _, j := range want[key] {
				if res[j][0] == -1 {
					res[j][0] = i - k
					res[j][1] = i - 1
				}
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
