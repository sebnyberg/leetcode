package p0159longestsubstringwithatmosttwodistinctchars

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lengthOfLongestSubstringTwoDistinct(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"ccaabbb", 5},
		{"eceba", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, lengthOfLongestSubstringTwoDistinct(tc.s))
		})
	}
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	findMin := func(cc map[byte]int) byte {
		minVal := math.MaxInt32
		var res byte
		for k, v := range cc {
			if v < minVal {
				res = k
				minVal = v
			}
		}
		return res
	}

	windowChars := make(map[byte]int)
	k := 2
	var maxLen int
	var startIdx int
	for i := range s {
		r := s[i]
		if _, exists := windowChars[r]; !exists {
			if len(windowChars) == k {
				chToRemove := findMin(windowChars)
				maxLen = max(maxLen, i-startIdx)
				startIdx = windowChars[chToRemove] + 1
				delete(windowChars, chToRemove)
			}
		}
		windowChars[r] = i
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
