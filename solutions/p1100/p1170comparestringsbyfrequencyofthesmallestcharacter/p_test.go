package p1170comparestringsbyfrequencyofthesmallestcharacter

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numSmallerByFrequency(t *testing.T) {
	for i, tc := range []struct {
		queries []string
		words   []string
		want    []int
	}{
		{[]string{"bbb", "cc"}, []string{"aaaa", "aaa", "aa", "a"}, []int{1, 2}},
		{[]string{"cbd"}, []string{"zaaaz"}, []int{1}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numSmallerByFrequency(tc.queries, tc.words))
		})
	}
}

func numSmallerByFrequency(queries []string, words []string) []int {
	f := func(w string) int {
		var count [26]int
		for _, ch := range w {
			count[ch-'a']++
		}
		var res int
		for _, c := range count {
			if c > 0 {
				return c
			}
		}
		return res
	}
	n := len(words)
	fs := make([]int, n)
	for i := range words {
		fs[i] = f(words[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(fs)))
	m := len(queries)
	qfs := make([]int, m)
	for i := range qfs {
		qfs[i] = f(queries[i])
	}
	idx := make([]int, m)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return qfs[idx[i]] > qfs[idx[j]]
	})
	res := make([]int, m)
	var l, r int
	for r < len(qfs) {
		for l < len(fs) && qfs[idx[r]] < fs[l] {
			l++
		}
		res[idx[r]] = l
		r++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
