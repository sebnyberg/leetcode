package p2062countvowelsofastring

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countVowelSubstrings(t *testing.T) {
	for _, tc := range []struct {
		word string
		want int
	}{
		{"aeiouu", 2},
		{"unicornarihan", 0},
		{"cuaieuouac", 7},
		{"bbaeixoubb", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.Equal(t, tc.want, countVowelSubstrings(tc.word))
		})
	}
}

func countVowelSubstrings(word string) int {
	bs := []byte(word)
	bs = append(bs, 'l')
	n := len(bs)
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}
	var ans int
	for i := range bs {
		if !bytes.Contains(vowels, bs[i:i+1]) {
			continue
		}
		hasVowel := make(map[byte]bool, 5)
		hasVowel[bs[i]] = true

		for j := i + 1; j < n; j++ {
			if !bytes.Contains(vowels, bs[j:j+1]) {
				break
			}
			hasVowel[bs[j]] = true
			if len(hasVowel) >= 5 {
				ans++
			}
		}
	}
	return ans
}
