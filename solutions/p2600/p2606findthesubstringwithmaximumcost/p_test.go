package p2606findthesubstringwithmaximumcost

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumCostSubstring(t *testing.T) {
	for i, tc := range []struct {
		s     string
		chars string
		vals  []int
		want  int
	}{
		{"adaa", "d", []int{-1000}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maximumCostSubstring(tc.s, tc.chars, tc.vals))
		})
	}
}

func maximumCostSubstring(s string, chars string, vals []int) int {
	// Greedy with stack
	var cost [26]int
	for i := range cost {
		cost[i] = i + 1
	}
	for i := range chars {
		cost[chars[i]-'a'] = vals[i]
	}
	var subres int
	var l int
	var res int
	for i := range s {
		subres += cost[s[i]-'a']
		res = max(res, subres)
		for l < i && (cost[s[l]-'a'] < 0 || subres < 0) {
			subres -= cost[s[l]-'a']
			l++
		}
		res = max(res, subres)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
