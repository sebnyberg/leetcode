package p3306countofsubstringscontainingeveryvowelandkconsonantsii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countOfSubstrings(t *testing.T) {
	for _, tc := range []struct {
		word string
		k    int
		want int64
	}{
		// {"aeioqq", 1, 0},
		// {"aeiou", 0, 1},
		{"ieaouqqieaouqq", 1, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, countOfSubstrings(tc.word, tc.k))
		})
	}
}

func countOfSubstrings(word string, k int) int64 {
	isvowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}
	isconsonant := func(ch byte) bool {
		return !isvowel(ch)
	}

	var idxs []int
	for i := range word {
		if isconsonant(word[i]) {
			idxs = append(idxs, i)
		}
	}
	idxs = append(idxs, len(word)) // sentinel value

	var vowelCount [26]int
	var uniqueVowels int
	var res int64
	var consonantCount int
	j := 0 // nextConsonantIndex

	// The range [l,r] is valid IFF numConsonants == k
	// If so, then any vowel index to the right of r also forms a valid range.
	var l int
	r := -1
outer:
	for {
		switch {
		case consonantCount < k || uniqueVowels < 5: // move r
			r++
			if r == len(word) {
				break outer
			}
			if isvowel(word[r]) {
				vowelCount[word[r]-'a']++
				if vowelCount[word[r]-'a'] == 1 {
					uniqueVowels++
				}
			} else {
				consonantCount++
				j++
			}
		default:
			if consonantCount == k {
				res += int64(idxs[j] - r)
			}
			if isvowel(word[l]) {
				vowelCount[word[l]-'a']--
				if vowelCount[word[l]-'a'] == 0 {
					uniqueVowels--
				}
			} else {
				consonantCount--
			}
			l++
		}
	}

	return res
}
