package p0678validparenthesisstring

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkValidString(t *testing.T) {
	for _, tc := range []struct {
		s    string
		want bool
	}{
		{"(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())", false},
		{"()", true},
		{"(*)", true},
		{"(*))", true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.s), func(t *testing.T) {
			require.Equal(t, tc.want, checkValidString(tc.s))
		})
	}
}

type state struct {
	open  int
	index int
}

func checkValidString(s string) bool {
	m := make(map[state]bool)
	return dfs(m, s, 0, 0)
}

func dfs(m map[state]bool, s string, i, open int) bool {
	if open < 0 {
		return false
	}
	if i == len(s) {
		return open == 0
	}
	k := state{i, open}
	if v, exists := m[k]; exists {
		return v
	}
	var res bool
	switch s[i] {
	case '(':
		res = dfs(m, s, i+1, open+1)
	case ')':
		res = dfs(m, s, i+1, open-1)
	case '*':
		res = dfs(m, s, i+1, open) ||
			dfs(m, s, i+1, open+1) ||
			dfs(m, s, i+1, open-1)
	}
	m[k] = res
	return m[k]
}
