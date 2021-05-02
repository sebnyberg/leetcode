package p0288uniquewordabbrev

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidWordAbbr(t *testing.T) {
	dict := []string{"deer", "door", "cake", "card"}
	a := Constructor(dict)
	res := a.IsUnique("dear")
	require.Equal(t, false, res)
	res = a.IsUnique("cart")
	require.Equal(t, true, res)
	res = a.IsUnique("cane")
	require.Equal(t, false, res)
	res = a.IsUnique("make")
	require.Equal(t, true, res)
}

type ValidWordAbbr struct {
	words   map[string]bool
	abbrevs map[string]int
}

func Constructor(dictionary []string) ValidWordAbbr {
	a := ValidWordAbbr{
		words:   make(map[string]bool),
		abbrevs: make(map[string]int),
	}
	for _, w := range dictionary {
		if !a.words[w] {
			a.words[w] = true
			a.abbrevs[abbr(w)]++
		}
	}
	return a
}

func abbr(word string) string {
	n := len(word)
	if n <= 2 {
		return word
	}
	return string(word[0]) + string(rune(n-2+'0')) + string(word[n-1])
}

func (this *ValidWordAbbr) IsUnique(word string) bool {
	a := abbr(word)
	switch {
	case this.abbrevs[a] == 0:
		return true
	case this.abbrevs[a] == 1 && this.words[word]:
		return true
	default:
		return false
	}
}
