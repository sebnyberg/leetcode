package p1593splitastringintothemaxnumberofuniquesubstrings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxUniqueSplit(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want int
	}{
		{"ababccc", 5},
		{"aba", 2},
		{"aa", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, maxUniqueSplit(tc.s))
		})
	}
}

func maxUniqueSplit(s string) int {
	seen := make(map[string]struct{})
	var res int
	for j := 1; j <= len(s); j++ {
		explore(s, j, seen, &res)
	}
	return res
}

func explore(s string, j int, seen map[string]struct{}, res *int) {
	if _, exists := seen[s[:j]]; exists {
		return
	}
	seen[s[:j]] = struct{}{}
	defer func(s string) {
		delete(seen, s[:j])
	}(s)
	s = s[j:]
	if len(s) == 0 {
		*res = max(*res, len(seen))
		return
	}
	for k := 1; k <= len(s); k++ {
		explore(s, k, seen, res)
	}
}
