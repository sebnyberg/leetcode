package p2207

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximizeSubsequenceCount(t *testing.T) {
	for _, tc := range []struct {
		text    string
		pattern string
		want    int64
	}{
		{"fwymvreuftzgrcrxczjacqovduqaiig", "yy", 1},
		{"k", "jk", 1},
		{"abdcdbc", "ac", 4},
		{"aabb", "ab", 6},
	} {
		t.Run(fmt.Sprintf("%+v", tc.text), func(t *testing.T) {
			require.Equal(t, tc.want, maximumSubsequenceCount(tc.text, tc.pattern))
		})
	}
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	n := len(text)
	rightCount := make([]int, n+1)
	leftCount := make([]int, n+1)
	for i := range text {
		if text[i] == pattern[0] {
			leftCount[i+1]++
		}
		leftCount[i+1] += leftCount[i]
	}
	for i := n; i >= 1; i-- {
		if text[i-1] == pattern[1] {
			rightCount[i-1]++
		}
		rightCount[i-1] += rightCount[i]
	}
	var res int
	// Count patterns currently
	for i := range text {
		if text[i] == pattern[0] {
			res += rightCount[i+1]
		}
	}
	// By adding pattern[0] to first we get rightCount[0]
	addToStart := rightCount[0]
	addToEnd := leftCount[n]
	maxRes := res + max(addToEnd, addToStart)
	return int64(maxRes)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
