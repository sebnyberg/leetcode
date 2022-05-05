package p1446consecutivecharacters

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxPower(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"leetcode", 2},
		{"abbcccddddeeeeedcba", 5},
		{"triplepillooooow", 5},
		{"hooraaaaaaaaaaay", 11},
		{"tourist", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, maxPower(tc.s))
		})
	}
}

func maxPower(s string) int {
	count := 1
	maxCount := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		maxCount = max(maxCount, count)
	}
	maxCount = max(maxCount, count)
	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
