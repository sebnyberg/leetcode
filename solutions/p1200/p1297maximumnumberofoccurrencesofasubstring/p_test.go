package p1297maximumnumberofoccurrencesofasubstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxFreq(t *testing.T) {
	for i, tc := range []struct {
		s          string
		maxLetters int
		minSize    int
		maxSize    int
		want       int
	}{
		{"aababcaab", 2, 3, 4, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, maxFreq(tc.s, tc.maxLetters, tc.minSize, tc.maxSize))
		})
	}
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	// Slide a window of length minSize over the string, counting occurrences.
	// Note that any string larger than minSize is irrelevant because it would
	// contain the smaller string anyway.
	var freq [26]int
	var nunique int
	var res int
	count := make(map[string]int)
	for i := range s {
		freq[s[i]-'a']++
		if freq[s[i]-'a'] == 1 {
			nunique++
		}
		if i < minSize-1 {
			continue
		}
		if i >= minSize {
			freq[s[i-minSize]-'a']--
			if freq[s[i-minSize]-'a'] == 0 {
				nunique--
			}
		}
		if nunique <= maxLetters {
			t := s[i-minSize+1 : i+1]
			count[t]++
			res = max(res, count[t])
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
