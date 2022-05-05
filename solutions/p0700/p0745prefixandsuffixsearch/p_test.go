package p0745prefixandsuffixsearch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordFilter(t *testing.T) {
	wf := Constructor([]string{"apple"})
	res := wf.F("a", "e")
	require.Equal(t, 0, res)
}

type WordFilter struct {
	wordIndex map[string]int
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{
		wordIndex: make(map[string]int),
	}
	for wordIdx, w := range words {
		n := len(w)
		for i := 0; i <= n; i++ {
			for j := 0; j <= n; j++ {
				k := w[j:] + "#" + w[:n-i]
				wf.wordIndex[k] = wordIdx
			}
		}
	}
	return wf
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if v, exists := this.wordIndex[suffix+"#"+prefix]; exists {
		return v
	}
	return -1
}
