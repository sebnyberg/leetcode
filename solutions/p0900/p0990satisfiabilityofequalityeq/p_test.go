package p0990satisfiabilityofequalityeq

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_equationsPossible(t *testing.T) {
	for _, tc := range []struct {
		equations []string
		want      bool
	}{
		{[]string{"a!=i", "g==k", "k==j", "k!=i", "c!=e", "a!=e", "k!=a", "a!=g", "g!=c"}, true},
		{[]string{"c==c", "b==d", "x!=z"}, true},
		{[]string{"a==b", "b!=c", "c==a"}, false},
		{[]string{"a==b", "b!=a"}, false},
		{[]string{"a==b", "b==a"}, true},
		{[]string{"a==b", "b==c", "a==c"}, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.equations), func(t *testing.T) {
			require.Equal(t, tc.want, equationsPossible(tc.equations))
		})
	}
}

func equationsPossible(equations []string) bool {
	var parent [26]byte
	for i := range parent {
		parent[i] = byte(i)
	}
	var find func(byte) byte
	find = func(a byte) byte {
		if parent[a] == a {
			return a
		}
		return find(parent[a])
	}
	union := func(a, b byte) {
		parent[find(a)] = find(b)
	}
	for _, eq := range equations {
		if eq[1] == '=' {
			a := eq[0] - 'a'
			b := eq[3] - 'a'
			union(a, b)
		}
	}
	for _, eq := range equations {
		if eq[1] == '!' {
			a := eq[0] - 'a'
			b := eq[3] - 'a'
			if find(a) == find(b) {
				return false
			}
		}
	}
	return true
}
