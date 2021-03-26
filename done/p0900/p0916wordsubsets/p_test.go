package p0916wordsubsets

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_wordSubsets(t *testing.T) {
	for _, tc := range []struct {
		A    []string
		B    []string
		want []string
	}{
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"e", "o"},
			[]string{"facebook", "google", "leetcode"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"l", "e"},
			[]string{"apple", "google", "leetcode"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"e", "oo"},
			[]string{"facebook", "google"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"lo", "eo"},
			[]string{"google", "leetcode"},
		},
		{
			[]string{"amazon", "apple", "facebook", "google", "leetcode"},
			[]string{"ec", "oc", "ceo"},
			[]string{"facebook", "leetcode"},
		},
	} {
		t.Run(fmt.Sprintf("%+v/+%v", tc.A, tc.B), func(t *testing.T) {
			require.Subset(t, tc.want, wordSubsets(tc.A, tc.B))
		})
	}
}

func wordSubsets(A []string, B []string) []string {
	type charCount [26]int
	var mergedCount charCount
	for _, b := range B {
		var bCount charCount
		for _, ch := range b {
			bCount[ch-'a']++
		}
		for i, count := range bCount {
			mergedCount[i] = max(mergedCount[i], count)
		}
	}

	matches := make([]string, 0)
	for _, a := range A {
		required := mergedCount
		for _, ch := range a {
			required[ch-'a']--
		}
		for _, count := range required {
			if count > 0 {
				goto Continue
			}
		}
		matches = append(matches, a)
	Continue:
	}

	return matches
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
