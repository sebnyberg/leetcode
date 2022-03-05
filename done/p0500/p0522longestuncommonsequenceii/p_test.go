package p0522longestuncommonsequenceii

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
		{[]string{"aabbcc", "aabbcc", "cbbcc"}, 5},
		{[]string{"aabbcc", "aabbcc", "b"}, -1},
		{[]string{"aabbcc", "aabbcc", "c"}, -1},
		{[]string{"aaa", "abc"}, 3},
		{[]string{"aba", "cdc", "eae"}, 3},
		{[]string{"aaa", "aaa", "aa"}, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.strs), func(t *testing.T) {
			require.Equal(t, tc.want, findLUSlength(tc.strs))
		})
	}
}

func findLUSlength(strs []string) int {
	// Returns true if a is a substring of b
	isSubSeq := func(b, a string) bool {
		var j int
		for i := 0; i < len(a) && j < len(b); i++ {
			if a[i] == b[j] {
				j++
			}
		}
		return j == len(b)
	}

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) > len(strs[j])
	})
	n := len(strs)
	for i := 0; i < n; i++ {
		var j int
		for ; j < n; j++ {
			if i != j && isSubSeq(strs[i], strs[j]) {
				break
			}
		}
		if j == n {
			return len(strs[i])
		}
	}
	return -1
}
