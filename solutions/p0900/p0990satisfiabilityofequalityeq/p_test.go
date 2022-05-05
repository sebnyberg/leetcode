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
	eq := NewDSU(26)
	for _, equation := range equations {
		if equation[1] != '=' {
			continue
		}
		eq.union(int(equation[0]-'a'), int(equation[3]-'a'))
	}

	for _, equation := range equations {
		if equation[1] != '!' {
			continue
		}
		a, b := int(equation[0]-'a'), int(equation[3]-'a')
		if a == b || eq.find(a) == eq.find(b) {
			return false
		}
	}
	return true
}

type DSU struct {
	parent []int
}

func NewDSU(n int) *DSU {
	dsu := &DSU{
		parent: make([]int, n),
	}
	for i := 0; i < n; i++ {
		dsu.parent[i] = i
	}
	return dsu
}

func (d *DSU) find(a int) int {
	if p := d.parent[a]; p != a {
		return d.find(d.parent[a])
	}
	return a
}

func (d *DSU) union(a, b int) {
	d.parent[d.find(a)] = d.find(b)
}
