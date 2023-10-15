package p2901longestunequaladjacentgroupssubsequenceii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	for i, tc := range []struct {
		n      int
		words  []string
		groups []int
		want   []string
	}{
		{4, []string{"a", "b", "c", "d"}, []int{1, 2, 3, 4}, []string{"a", "b", "c", "d"}},
		{3, []string{"bab", "dab", "cab"}, []int{1, 2, 2}, []string{"bab", "cab"}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, getWordsInLongestSubsequence(tc.n, tc.words, tc.groups))
		})
	}
}

func getWordsInLongestSubsequence(n int, words []string, groups []int) []string {
	// This is a DP-type problem.
	//
	// Given that a certain index is chosen, return the maximum list such that
	// the prior element is the given index.
	//
	m := len(words)
	mem := make([][]string, m)
	for i := range mem {
		mem[i] = []string{}
	}
	res := dp(mem, words, groups, -1, 0)
	return res
}

func dp(mem [][]string, words []string, groups []int, prevIdx, i int) []string {
	if i == len(mem) {
		return []string{}
	}

	// First do nothing
	res := dp(mem, words, groups, prevIdx, i+1)

	// Then, if possible, start a new subsequence from this index
	if prevIdx == -1 ||
		(hammingIsOne(words[prevIdx], words[i]) && groups[prevIdx] != groups[i]) {
		if len(mem[i]) == 0 {
			// We can select the current index.
			mem[i] = append(mem[i], words[i])
			mem[i] = append(mem[i], dp(mem, words, groups, i, i+1)...)
		}
		if len(mem[i]) > len(res) {
			res = mem[i]
		}
	}

	return res
}

func hammingIsOne(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var count int
	for i := range s {
		if s[i] != t[i] {
			count++
		}
	}
	return count == 1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
