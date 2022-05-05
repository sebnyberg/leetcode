package p0340longestsubstringwithatmostk

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	for _, tc := range []struct {
		s    string
		k    int
		want int
	}{
		{"ccaabbb", 2, 5},
		{"eceba", 2, 3},
		{"aa", 1, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, lengthOfLongestSubstringKDistinct(tc.s, tc.k))
		})
	}
}

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	// Note: this can be improved by keeping track of counts
	// rather than indices. When there are too many distinct chars in the window
	// move the left pointer and decrement from the window until one char
	// reaches zero, at which point the right pointer can start moving again.
	if k == 0 {
		return 0
	}
	findMin := func(cc map[byte]int) byte {
		minVal := math.MaxInt32
		var res byte
		for k, v := range cc {
			if int(v) < minVal {
				res = k
				minVal = int(v)
			}
		}
		return res
	}

	maxIndexForCh := make(map[byte]int)
	var maxLen int
	var startIdx int
	for i := range s {
		r := s[i]
		if _, exists := maxIndexForCh[r]; !exists {
			if len(maxIndexForCh) == k {
				chToRemove := findMin(maxIndexForCh)
				maxLen = max(maxLen, i-startIdx)
				startIdx = maxIndexForCh[chToRemove] + 1
				delete(maxIndexForCh, chToRemove)
			}
		}
		maxIndexForCh[r] = i
	}
	maxLen = max(maxLen, len(s)-startIdx)

	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
