package p0833findandreplaceinstring

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findReplaceString(t *testing.T) {
	for _, tc := range []struct {
		S       string
		indexes []int
		sources []string
		target  []string
		want    string
	}{
		{"vmokgggqzp", []int{3, 5, 1}, []string{"kg", "ggq", "mo"}, []string{"s", "so", "bfr"}, "vbfrssozp"},
		{"abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}, "eeebffff"},
		{"abcd", []int{0, 2}, []string{"ab", "ec"}, []string{"eee", "ffff"}, "eeecd"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.S), func(t *testing.T) {
			require.Equal(t, tc.want, findReplaceString(tc.S, tc.indexes, tc.sources, tc.target))
		})
	}
}

type wordReplacer struct {
	idx       int
	source    string
	sourceLen int
	target    string
}

func findReplaceString(S string, indexes []int, sources []string, target []string) string {
	if len(indexes) == 0 {
		return S
	}
	replacers := make([]wordReplacer, len(indexes))
	for i := range indexes {
		replacers[i] = wordReplacer{indexes[i], sources[i], len(sources[i]), target[i]}
	}
	sort.Slice(replacers, func(i, j int) bool {
		return replacers[i].idx < replacers[j].idx
	})
	m := len(replacers)
	var matchIdx int
	res := make([]byte, 0)
	for i := 0; i < len(S); i++ {
		if matchIdx >= m || i < replacers[matchIdx].idx {
			res = append(res, S[i])
			continue
		}
		r := replacers[matchIdx]
		// Attempt to match
		if len(S[i:]) < r.sourceLen || S[i:i+r.sourceLen] != r.source {
			// fail
			matchIdx++
			res = append(res, S[i])
			continue
		}
		// success
		i += r.sourceLen - 1
		res = append(res, []byte(r.target)...)
		matchIdx++
	}
	return string(res)
}
