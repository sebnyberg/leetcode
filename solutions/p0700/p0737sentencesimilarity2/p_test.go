package p0737sentencesimilarity2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_areSentencesSimilarTwo(t *testing.T) {
	for _, tc := range []struct {
		words1 []string
		words2 []string
		pairs  [][]string
		want   bool
	}{
		{
			[]string{"great", "acting", "skills"},
			[]string{"fine", "drama", "talent"},
			[][]string{{"great", "good"}, {"find", "good"}, {"acting", "drama"}, {"skills", "talent"}},
			true,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words1), func(t *testing.T) {
			require.Equal(t, tc.want, areSentencesSimilarTwo(tc.words1, tc.words2, tc.pairs))
		})
	}
}

func areSentencesSimilarTwo(words1 []string, words2 []string, pairs [][]string) bool {
	if len(words1) != len(words2) {
		return false
	}
	parent := make(map[string]string)
	find := func(s string) string {
		for parent[s] != s {
			s = parent[s]
		}
		return s
	}

	for _, pair := range pairs {
		a, b := pair[0], pair[1]
		if _, exists := parent[a]; !exists {
			parent[a] = a
		}
		if _, exists := parent[b]; !exists {
			parent[b] = b
		}
		if find(a) != find(b) {
			parent[find(a)] = find(b)
		}
	}

	for i := range words1 {
		if find(words1[i]) != find(words2[i]) {
			return false
		}
	}
	return true
}
