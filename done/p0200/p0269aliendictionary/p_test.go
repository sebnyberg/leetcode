package p0269aliendictionary

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_alienOrder(t *testing.T) {
	for _, tc := range []struct {
		words []string
		want  string
	}{
		{[]string{"wrt", "wrf", "er", "ett", "rftt", "te"}, "wertf"},
		{[]string{"abc", "ab"}, ""},
		{[]string{"ac", "ab", "zc", "zb"}, "aczb"},
		{[]string{"wrt", "wrf", "er", "ett", "rftt"}, "wertf"},
		{[]string{"z", "x"}, "zx"},
		{[]string{"z", "x", "z"}, ""},
	} {
		t.Run(fmt.Sprintf("%+v", tc.words), func(t *testing.T) {
			require.Equal(t, tc.want, alienOrder(tc.words))
		})
	}
}

func alienOrder(words []string) string {
	n := len(words)
	adj := make(map[byte]map[byte]struct{})
	deg := make(map[byte]int)
	for _, word := range words {
		for i := range word {
			adj[word[i]] = make(map[byte]struct{})
			deg[word[i]] = 0
		}
	}
	for i := 0; i < n-1; i++ {
		n1, n2 := len(words[i]), len(words[i+1])
		for j, maxlen := 0, min(n1, n2); j < maxlen; j++ {
			a, b := words[i][j], words[i+1][j]
			if a == b {
				// Idiotic that this is a testcase.. is this about solving the problem
				// or validating invalid input? What's next? Symbols? Emojis?
				if j == maxlen-1 && n1 > n2 {
					return ""
				}
				continue
			}
			if _, exists := adj[a][b]; !exists {
				deg[b]++
				adj[a][b] = struct{}{}
			}
			break
		}
	}

	current := make([]byte, 0)
	for ch, indegree := range deg {
		if indegree == 0 {
			current = append(current, ch)
		}
	}

	res := make([]byte, 0, len(adj))
	for len(current) > 0 {
		new := make([]byte, 0)
		for _, ch := range current {
			res = append(res, ch)
			for near := range adj[ch] {
				deg[near]--
				if deg[near] == 0 {
					new = append(new, near)
				}
			}
		}
		current = new
	}

	if len(res) != len(adj) {
		return ""
	}

	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
