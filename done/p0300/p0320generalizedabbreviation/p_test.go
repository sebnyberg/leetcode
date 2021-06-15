package p0320generalizedabbreviation

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateAbbreviations(t *testing.T) {
	for _, tc := range []struct {
		word string
		want []string
	}{
		{"word", []string{"4", "3d", "2r1", "2rd", "1o2", "1o1d", "1or1", "1ord", "w3", "w2d", "w1r1", "w1rd", "wo2", "wo1d", "wor1", "word"}},
		{"a", []string{"1", "a"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.word), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, generateAbbreviations(tc.word))
		})
	}
}

func generateAbbreviations(word string) []string {
	bs := []byte(word)
	n := len(word)
	a := &AbbrCollector{
		origWord: word,
		abbrs:    make([]string, 0, n*2),
	}
	prefix := make([]byte, 0, 15)
	a.collect(bs, prefix, 0, n, true)
	a.collect(bs, prefix, 0, n, false)
	return a.abbrs
}

type AbbrCollector struct {
	origWord string
	abbrs    []string
}

func (a *AbbrCollector) collect(word []byte, prefix []byte, pos, n int, doAbbr bool) {
	if pos == n {
		a.abbrs = append(a.abbrs, string(prefix))
		return
	}
	m := len(prefix)
	for i := 1; pos+i <= n; i++ {
		if doAbbr {
			// abbreviate segment
			prefix = append(prefix, strconv.Itoa(i)...)
			a.collect(word, prefix, pos+i, n, false)
			prefix = prefix[:m]
		} else {
			// pick segment as-is
			prefix = append(prefix, word[pos:pos+i]...)
			a.collect(word, prefix, pos+i, n, true)
			prefix = prefix[:m]
		}
	}
}
