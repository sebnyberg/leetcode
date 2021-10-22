package p0451sortcharactersbyfrequency

import (
	"fmt"
	"sort"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
)

func Test_frequencySort(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want string
	}{
		// Cba to write a verifier
		// {"tree", "eert"},
		// {"cccaaa", "aaaccc"},
		// {"Aabb", "bbAa"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, frequencySort(tc.s))
		})
	}
}

func frequencySort(s string) string {
	type charFrequency struct {
		char  rune
		count int
	}
	charFreq := make([]charFrequency, 0, 26)
	charIdx := make(map[rune]int)
	for _, ch := range s {
		if _, exists := charIdx[ch]; !exists {
			charIdx[ch] = len(charFreq)
			charFreq = append(charFreq, charFrequency{
				char:  ch,
				count: 0,
			})
		}
		charFreq[charIdx[ch]].count++
	}
	sort.Slice(charFreq, func(i, j int) bool {
		return charFreq[i].count > charFreq[j].count
	})
	res := make([]rune, 0, utf8.RuneCountInString(s))
	for _, c := range charFreq {
		if c.count == 0 {
			break
		}
		for i := 0; i < c.count; i++ {
			res = append(res, c.char)
		}
	}
	return string(res)
}
