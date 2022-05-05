package p1087braceexpansion

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_expand(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want []string
	}{
		{"{a,b}c{d,e}f", []string{"acdf", "acef", "bcdf", "bcef"}},
		{"abcd", []string{"abcd"}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, expand(tc.s))
		})
	}
}

func expand(s string) []string {
	// Cool exercise!
	//
	// Perform DFS, when finding a {, parse options until finding a },
	// Sort options in lexicographical order, then pick one by one and add to
	// the prefix.

	prefix := make([]byte, 50)
	var res []string
	dfs(prefix, s, 0, 0, &res)
	return res
}

func dfs(prefix []byte, s string, i, j int, res *[]string) {
	if i == len(s) {
		*res = append(*res, string(prefix[:j]))
		return
	}

	if s[i] != '{' {
		prefix[j] = s[i]
		dfs(prefix, s, i+1, j+1, res)
		return
	}

	// Parse options
	var opts []byte
	for i++; s[i] != '}'; i++ {
		if s[i] == ',' {
			continue
		}
		opts = append(opts, s[i])
	}
	sort.Slice(opts, func(i, j int) bool {
		return opts[i] < opts[j]
	})

	// Try each option in lexicographical order
	for _, o := range opts {
		prefix[j] = o
		dfs(prefix, s, i+1, j+1, res)
	}
}
