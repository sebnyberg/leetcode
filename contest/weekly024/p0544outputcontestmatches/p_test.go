package p0544outputcontestmatches

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findContestMatch(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want string
	}{
		{4, "((1,4),(2,3))"},
		{8, "(((1,8),(4,5)),((2,7),(3,6)))"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findContestMatch(tc.n))
		})
	}
}

func findContestMatch(n int) string {
	var levels int
	for m := n; m > 1; m >>= 1 {
		levels++
	}
	var e expr
	e.res = make([]byte, 0, n*2)
	e.dfs(1, levels, 1)
	return string(e.res)
}

type expr struct {
	res []byte
}

func (e *expr) dfs(curLevel, nLevels, parent int) {
	if curLevel > nLevels {
		e.res = append(e.res, []byte(fmt.Sprint(parent))...)
		return
	}

	e.res = append(e.res, '(')

	// Create pair
	want := (1 << curLevel) + 1
	missing := want - parent
	e.dfs(curLevel+1, nLevels, parent)
	e.res = append(e.res, ',')
	e.dfs(curLevel+1, nLevels, missing)
	e.res = append(e.res, ')')
}
