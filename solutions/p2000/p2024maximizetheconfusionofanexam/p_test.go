package p2024maximizetheconfusionofanexam

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxConsecutiveAnswers(t *testing.T) {
	for _, tc := range []struct {
		answerKey string
		k         int
		want      int
	}{
		{"TTFF", 2, 4},
		{"TFFT", 1, 3},
		{"TTFTTFTT", 1, 5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.answerKey), func(t *testing.T) {
			require.Equal(t, tc.want, maxConsecutiveAnswers(tc.answerKey, tc.k))
		})
	}
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	// Move a right cursor until it reaches the end or covers k 'T's
	// Update max window length
	// Move left cursor until window covers less than k 'T's
	// Move right cursor until end or covers k 'T's
	// Update max window length
	// etc.
	// Invert T's and F's, repeat procedure.
	// Total max is the answer.
	maxLen := maxWindowLen(answerKey, k, 'T')
	maxLen = max(maxLen, maxWindowLen(answerKey, k, 'F'))
	return maxLen
}

func maxWindowLen(answerKey string, k int, convertRune byte) int {
	var left, right, convCount, maxLen int
	n := len(answerKey)
	for right != n {
		// Right points to next convert rune
		for right != n {
			if answerKey[right] == convertRune {
				if convCount == k {
					break
				}
				convCount++
			}
			right++
		}
		maxLen = max(maxLen, right-left)
		for left != n && convCount == k {
			if answerKey[left] == convertRune {
				convCount--
			}
			left++
		}
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
