package p0243shortestworddistance

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shortestDistance(t *testing.T) {
	for _, tc := range []struct {
		wordsDict []string
		word1     string
		word2     string
		want      int
	}{
		{[]string{"practice", "makes", "perfect", "coding", "makes"}, "coding", "practice", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.wordsDict), func(t *testing.T) {
			require.Equal(t, tc.want, shortestDistance(tc.wordsDict, tc.word1, tc.word2))
		})
	}
}

func shortestDistance(wordsDict []string, word1 string, word2 string) int {
	word1Idx := make([]int, 0)
	word2Idx := make([]int, 0)
	for i, word := range wordsDict {
		if word == word1 {
			word1Idx = append(word1Idx, i)
		}
		if word == word2 {
			word2Idx = append(word2Idx, i)
		}
	}
	minDist := math.MaxInt32
	for _, i := range word1Idx {
		for _, j := range word2Idx {
			if d := abs(i - j); d < minDist {
				minDist = d
			}
		}
	}
	return minDist
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
