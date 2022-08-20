package p0833findandreplaceinstring

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findReplaceString(t *testing.T) {
	for _, tc := range []struct {
		s       string
		indices []int
		sources []string
		targets []string
		want    string
	}{
		{"abcd", []int{0, 2}, []string{"a", "cd"}, []string{"eee", "ffff"}, "eeebffff"},
		{"abcde", []int{2, 2}, []string{"cdef", "bc"}, []string{"f", "fe"}, "abcde"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, findReplaceString(tc.s, tc.indices, tc.sources, tc.targets))
		})
	}
}

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	type replacement struct {
		idx    int
		source string
		target string
	}
	n := len(indices)
	rs := make([]replacement, n)
	for i := range indices {
		rs[i] = replacement{indices[i], sources[i], targets[i]}
	}
	// rs = append(rs, replacement{len(s), "", ""})
	sort.Slice(rs, func(i, j int) bool {
		return rs[i].idx < rs[j].idx
	})
	res := make([]byte, 0, len(s))
	var i int
	for _, r := range rs {
		res = append(res, s[i:r.idx]...)
		j := r.idx + len(r.source)
		i = r.idx
		if j > len(s) || s[r.idx:j] != r.source {
			continue
		}
		i = r.idx + len(r.source)
		res = append(res, r.target...)
	}
	if i != len(s) {
		res = append(res, s[i:]...)
	}
	return string(res)
}
