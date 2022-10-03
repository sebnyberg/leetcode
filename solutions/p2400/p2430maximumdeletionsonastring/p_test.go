package p2430maximumdeletionsonastring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_deleteString(t *testing.T) {
	for i, tc := range []struct {
		s    string
		want int
	}{
		{"aaaaa", 5},
		{"aaabaab", 4},
		{"abcabcdabc", 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, deleteString(tc.s))
		})
	}
}

func deleteString(s string) int {
	// We always delete from left to right, so it's always going to be a case of
	// "find the maximum number of operations needed starting with position j"
	//
	// Or in other words, we can memoize prior results.
	//
	// Because Leetcode is often not great at setting limits for Go, in
	// particular with relation to strings in Go, I think a suboptimal route
	// will still pass.
	//
	// Approach: for each index i, for each valid removal ending in
	// i+1,...,i+w/2 where w is the remaining length, add max number of removals
	// for remainder and add as possible result for index i.
	n := len(s)
	mem := make([]int, n+1)
	res := dp(mem, s, 0, n)
	return res
}

func dp(mem []int, s string, i, n int) int {
	if mem[i] != 0 {
		return mem[i]
	}

	if i == n {
		return 0
	}
	if i == n-1 {
		return 1
	}

	// Find all valid removals
	res := 1 // fallback
	// Check for matching pairs
	for j := 1; i+2*j <= n; j++ {
		if s[i:i+j] == s[i+j:i+(j*2)] {
			res = max(res, 1+dp(mem, s, i+j, n))
		}
	}

	mem[i] = res
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
