package p2067numberofequalcountsubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_equalCountSubstrings(t *testing.T) {
	for _, tc := range []struct {
		s     string
		count int
		want  int
	}{
		{"aaaaaaaa", 4, 5},
		{"aaabcbbcc", 3, 3},
		{"abcd", 2, 0},
		{"a", 5, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, equalCountSubstrings(tc.s, tc.count))
		})
	}
}

func equalCountSubstrings(s string, count int) int {
	// Find number of unique characters in s
	var hasCh [26]bool
	for _, ch := range s {
		hasCh[ch-'a'] = true
	}
	var uniqueCount int
	for _, has := range hasCh {
		if has {
			uniqueCount++
		}
	}

	var res int
	for nunique := 1; nunique <= uniqueCount; nunique++ {
		// A window containing i unique characters must be of length count*i
		var charFreq [26]int
		windowLen := nunique * count
		if windowLen > len(s) {
			break
		}
		var uniqueCount int
		for i := 0; i < len(s); i++ {
			end := int(s[i] - 'a')
			charFreq[end]++
			if charFreq[end] == count {
				uniqueCount++
			} else if charFreq[end] == count+1 {
				uniqueCount--
			}

			if i >= windowLen {
				start := int(s[i-windowLen] - 'a')
				charFreq[start]--
				if charFreq[start] == count {
					uniqueCount++
				} else if charFreq[start] == count-1 {
					uniqueCount--
				}
			}

			if uniqueCount == nunique {
				res++
			}
		}
	}
	return res
}
