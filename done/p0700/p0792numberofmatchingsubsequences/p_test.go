package p0792numberofmatchingsubsequences

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numMatchingSubseq(t *testing.T) {
	for _, tc := range []struct {
		s     string
		words []string
		want  int
	}{
		{"abcde", []string{"a", "bb", "acd", "ace"}, 3},
		{"dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, numMatchingSubseq(tc.s, tc.words))
		})
	}
}

func numMatchingSubseq(s string, words []string) int {
	bs := []byte(s)
	var next [26][][]byte
	for i := 0; i < 26; i++ {
		next[i] = make([][]byte, 0)
	}
	for _, word := range words {
		next[word[0]-'a'] = append(next[word[0]-'a'], []byte(word[1:]))
	}

	var count int
	for _, b := range bs {
		n := len(next[b-'a'])
		for _, w := range next[b-'a'] {
			if len(w) == 0 {
				count++
				continue
			}
			next[w[0]-'a'] = append(next[w[0]-'a'], w[1:])
		}
		next[b-'a'] = next[b-'a'][n:]
	}
	return count
}
