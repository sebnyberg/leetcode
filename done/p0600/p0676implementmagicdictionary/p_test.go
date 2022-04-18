package p0676implementmagicdictionary

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMagicDictionary(t *testing.T) {
	c := Constructor()
	c.BuildDict([]string{"hello", "leetcode"})
	res := c.Search("hello")
	require.Equal(t, false, res)
}

// Notes:
// Removal is easy - just concat

type MagicDictionary struct {
	words map[string]struct{}
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		words: make(map[string]struct{}, 100),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	// Build masked version of each word in the dictionary
	for _, w := range dictionary {
		b := []byte(w)
		for i := 0; i < len(b); i++ {
			for j := 'a'; j <= 'z'; j++ {
				if byte(j) == w[i] {
					continue
				}
				this.words[w[:i]+string(j)+w[i+1:]] = struct{}{}
			}
		}
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	if _, exists := this.words[searchWord]; exists {
		return true
	}
	return false
}
