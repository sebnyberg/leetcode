package p2063vowelsofallsubstrings

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countVowels(t *testing.T) {
	for _, tc := range []struct {
		word string
		want int64
	}{
		{"aba", 6},
		{"abc", 3},
		{"ltcd", 0},
		{"noosabasboosa", 237},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, countVowels(tc.word))
		})
	}
}

func countVowels(word string) int64 {
	var sumDelta int
	var res int
	for i, ch := range word {
		// If the current position is a consonant, and the sum addition of the
		// previous position was, let's say, 18. Then the sum addition from the current
		// position will be the same.
		if !strings.ContainsRune("aeiou", ch) {
			res += sumDelta
		} else {
			// If the current position is a vowel, then the value of all prior substrings
			// will increase by 1.
			sumDelta += i + 1
			res += sumDelta
		}
	}
	return int64(res)
}
