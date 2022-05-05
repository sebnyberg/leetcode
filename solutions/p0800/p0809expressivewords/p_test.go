package p0809expressivewords

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_expressiveWords(t *testing.T) {
	for _, tc := range []struct {
		S     string
		words []string
		want  int
	}{
		{"heeellooo", []string{"hello", "hi", "helo"}, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.S), func(t *testing.T) {
			require.Equal(t, tc.want, expressiveWords(tc.S, tc.words))
		})
	}
}

func expressiveWords(S string, words []string) int {
	chs, counts := rle(S)

	var ok int
	for _, word := range words {
		wordChs, wordCounts := rle(word)
		if len(wordChs) != len(chs) {
			continue
		}
		for i := range wordChs {
			if wordChs[i] != chs[i] || counts[i] < wordCounts[i] || counts[i] < 3 && counts[i] != wordCounts[i] {
				goto ContinueLoop
			}
		}
		ok++
	ContinueLoop:
	}
	return ok
}

func rle(s string) ([]byte, []int) {
	chs := []byte{}
	counts := []int{}
	for len(s) > 0 {
		ch := s[0]
		count := 1
		for i := 1; i < len(s); i++ {
			if s[i] != ch {
				break
			}
			count++
		}
		chs = append(chs, ch)
		counts = append(counts, count)
		s = s[count:]
	}
	return chs, counts
}
