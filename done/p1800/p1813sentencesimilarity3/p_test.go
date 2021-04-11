package p1813sentencesimilarity3

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_areSentencesSimilar(t *testing.T) {
	for _, tc := range []struct {
		sentence1 string
		sentence2 string
		want      bool
	}{
		{"c h p Ny", "c BDQ r h p Ny", true},
		{"B", "ByI BMyQIqce b bARkkMaABi vlR RLHhqjNzCN oXvyK zRXR q ff B yHS OD KkvJA P JdWksnH", false},
		{"My name is Haley", "My Haley", true},
		{"Luky", "Lucccky", false},
		{"of", "A lot of words", false},
		{"Eating right now", "Eating", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sentence1), func(t *testing.T) {
			require.Equal(t, tc.want, areSentencesSimilar(tc.sentence1, tc.sentence2))
		})
	}
}

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	if sentence1 == sentence2 {
		return true
	}
	if len(sentence1) < len(sentence2) {
		sentence1, sentence2 = sentence2, sentence1
	}
	s1 := strings.Split(sentence1, " ")
	s2 := strings.Split(sentence2, " ")

	n1, n2 := len(s1), len(s2)
	matched := 0
	for l := 0; l < n2; l++ {
		if s1[l] != s2[l] {
			break
		}
		matched++
	}
	d := n1 - n2
	for r := n2 - 1; r >= 0; r-- {
		if s1[d+r] != s2[r] {
			break
		}
		matched++
	}
	return matched >= n2
}
