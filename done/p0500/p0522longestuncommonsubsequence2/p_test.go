package p0522longestuncommonsubsequence2

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLUSlength(t *testing.T) {
	for _, tc := range []struct {
		strs []string
		want int
	}{
		{[]string{"aabbcc", "aabbcc", "c"}, -1},
		{[]string{"aba", "cdc", "eae"}, 3},
		{[]string{"aaa", "aaa", "aa"}, -1},
		{[]string{"aaa", "aaa", "a"}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strs), func(t *testing.T) {
			require.Equal(t, tc.want, findLUSlength(tc.strs))
		})
	}
}

func findLUSlength(strs []string) int {
	// Sort strings by length
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})

	// While there are strings in strs
	seen := make(map[string]int)
	for width := len(strs[0]); width > 0 && len(strs) > 0; width-- {
		var i int
		for i = 0; i < len(strs) && len(strs[i]) == width; i++ {
			seen[strs[i]]++
		}
		// Any unique string is the max length
		for _, count := range seen {
			if count == 1 {
				return len(strs[0])
			}
		}
		// Add all possible substrings of all seen strings to next iteration
		newSeen := make(map[string]int)
		for s, count := range seen {
			for _, p := range getPerms(s) {
				newSeen[p] += count
			}
		}
		seen = newSeen
		strs = strs[i:]
	}
	return -1
}

func getPerms(s string) []string {
	res := make([]string, 0, len(s))
	for i := 0; i < len(s); i++ {
		res = append(res, s[:i]+s[i+1:])
	}
	return res
}
