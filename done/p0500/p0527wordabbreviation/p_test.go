package p0527wordabbreviation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordsAbbreviation(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  []string
	}{
		{
			[]string{"like", "god", "internal", "me", "internet", "interval", "intension", "face", "intrusion"},
			[]string{"l2e", "god", "internal", "me", "i6t", "interval", "inte4n", "f2e", "intr4n"},
		},
		{
			[]string{"aa", "aaa"},
			[]string{"aa", "aaa"},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, wordsAbbreviation(tc.words))
		})
	}
}

func wordsAbbreviation(words []string) []string {
	// First off, two words cannot have the same abbreviation unless they have the
	// same number of letters. So we may partition the strings by size..
	var bySize [401][]string
	for _, w := range words {
		bySize[len(w)] = append(bySize[len(w)], w)
	}

	// Then, abbreviate strings one by one
	res := make(map[string]string, len(words))
	abbrs := make(map[[2]string]int, len(words))
	for size, words := range bySize {
		if len(words) == 0 {
			continue
		}
		for k := range abbrs {
			delete(abbrs, k)
		}
		var done [401]bool
		for x := size - 2; x > 1; x-- {
			for i, w := range words {
				if done[i] {
					continue
				}
				k := [2]string{w[:size-x-1], w[size-1:]}
				abbrs[k]++
			}
			for i, w := range words {
				k := [2]string{w[:size-x-1], w[size-1:]}
				if abbrs[k] == 1 {
					done[i] = true
					res[words[i]] = w[:size-x-1] + fmt.Sprint(x) + w[size-1:]
				}
			}
		}
	}
	resList := make([]string, len(words))
	for i, w := range words {
		if v, exists := res[w]; exists {
			resList[i] = v
		} else {
			resList[i] = w
		}
	}
	return resList
}
