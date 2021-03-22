package p0966vowelspellchecker

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_spellchecker(t *testing.T) {
	for _, tc := range []struct {
		wordlist []string
		queries  []string
		want     []string
	}{
		{
			//        0        -       2       -
			[]string{"KiTe", "kite", "hare", "Hare"},
			[]string{"kite", "Kite", "KiTe", "Hare", "HARE", "Hear", "hear", "keti", "keet", "keto"},
			[]string{"kite", "KiTe", "KiTe", "Hare", "hare", "", "", "KiTe", "", "KiTe"},
		},
	} {
		t.Run(fmt.Sprintf("%+v/%+v", tc.wordlist, tc.queries), func(t *testing.T) {
			require.Equal(t, tc.want, spellchecker(tc.wordlist, tc.queries))
		})
	}
}

const zeroRune = rune(0)

func spellchecker(wordlist []string, queries []string) []string {
	wordMap := make(map[string]bool)
	lowercaseWordIndex := make(map[string]int)
	devowlizedWordIndex := make(map[string]int)
	replacer := strings.NewReplacer(
		"a", string(zeroRune),
		"e", string(zeroRune),
		"i", string(zeroRune),
		"o", string(zeroRune),
		"u", string(zeroRune),
	)
	for i, word := range wordlist {
		wordMap[word] = true
		lowercaseWord := strings.ToLower(word)
		if _, exists := lowercaseWordIndex[lowercaseWord]; !exists {
			lowercaseWordIndex[lowercaseWord] = i
		}

		devowelizedWord := replacer.Replace(lowercaseWord)
		tmp := []byte(devowelizedWord)
		_ = tmp
		if _, exists := devowlizedWordIndex[devowelizedWord]; !exists {
			devowlizedWordIndex[devowelizedWord] = i
		}
	}
	results := make([]string, 0, len(queries))
	for _, query := range queries {
		// First rule: perfect match
		if wordMap[query] {
			results = append(results, query)
			continue
		}

		// Second rule: case-insensitive match should return
		// first value found in the original list
		lowercaseWord := strings.ToLower(query)
		if idx, exists := lowercaseWordIndex[lowercaseWord]; exists {
			results = append(results, wordlist[idx])
			continue
		}

		devowelizedWord := replacer.Replace(lowercaseWord)
		if idx, exists := devowlizedWordIndex[devowelizedWord]; exists {
			results = append(results, wordlist[idx])
			continue
		}

		results = append(results, "")
	}
	return results
}
