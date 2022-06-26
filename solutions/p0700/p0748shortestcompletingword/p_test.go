package p0748shortestcompletingword

import (
	"fmt"
	"sort"
	"testing"
	"unicode"

	"github.com/stretchr/testify/require"
)

func Test_shortestCompletingWord(t *testing.T) {
	for _, tc := range []struct {
		licensePlate string
		words        []string
		want         string
	}{
		{
			"re69865",
			[]string{"population", "crime", "kid", "pressure", "store", "any", "relate", "will", "death", "when"},
			"crime",
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.licensePlate), func(t *testing.T) {
			require.Equal(t, tc.want, shortestCompletingWord(tc.licensePlate, tc.words))
		})
	}
}

func shortestCompletingWord(licensePlate string, words []string) string {
	freqCount := func(s string) [26]int {
		var freq [26]int
		for _, ch := range s {
			if !unicode.IsLetter(ch) {
				continue
			}
			if unicode.IsUpper(ch) {
				ch = unicode.ToLower(ch)
			}
			freq[ch-'a']++
		}
		return freq
	}

	want := freqCount(licensePlate)
	sort.SliceStable(words, func(i, j int) bool {
		if len(words[i]) == len(words[j]) {
			return i < j
		}
		return len(words[i]) < len(words[j])
	})

	for _, w := range words {
		got := freqCount(w)
		for ch, count := range want {
			if got[ch] < count {
				goto continueSearch
			}
		}
		return w

	continueSearch:
	}
	return ""
}
