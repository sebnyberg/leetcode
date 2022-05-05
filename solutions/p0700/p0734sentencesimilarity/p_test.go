package p0734sentencesimilarity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_areSentencesSimilar(t *testing.T) {
	for _, tc := range []struct {
		sentence1    []string
		sentence2    []string
		similarPairs [][]string
		want         bool
	}{
		{[]string{"great", "acting", "skills"}, []string{"fine", "drama", "talent"}, [][]string{{"great", "fine"}, {"drama", "acting"}, {"skills", "talent"}}, true},
		{[]string{"great"}, []string{"great"}, [][]string{}, true},
		{[]string{"great"}, []string{"doubleplus", "good"}, [][]string{}, false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.sentence1), func(t *testing.T) {
			require.Equal(t, tc.want, areSentencesSimilar(tc.sentence1, tc.sentence2, tc.similarPairs))
		})
	}
}

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	similar := make(map[string][]string, 2*len(similarPairs))
	for _, pair := range similarPairs {
		p1, p2 := pair[0], pair[1]
		similar[p1] = append(similar[p1], p2)
		similar[p2] = append(similar[p2], p1)
	}
	for i, s1 := range sentence1 {
		s2 := sentence2[i]
		if s1 == s2 {
			continue
		}
		var found bool
		for _, v := range similar[s1] {
			if v == s2 {
				found = true
			}
		}
		if !found {
			return false
		}
	}
	return true
}
