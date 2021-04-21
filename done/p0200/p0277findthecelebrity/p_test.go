package p0277findthecelebrity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Solution(t *testing.T) {
	for _, tc := range []struct {
		graph [][]int
		n     int
		want  int
	}{
		{[][]int{{1, 0}, {0, 1}}, 2, -1},
		{[][]int{{1, 1, 0}, {0, 1, 0}, {1, 1, 1}}, 3, 1},
		{[][]int{{1, 0, 1}, {1, 1, 0}, {0, 1, 1}}, 3, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.graph), func(t *testing.T) {
			r := Relationships{
				knows: tc.graph,
			}
			f := solution(r.getFunc())
			res := f(tc.n)
			require.Equal(t, tc.want, res)
		})
	}
}

type Relationships struct {
	knows [][]int
}

func (r Relationships) getFunc() func(a int, b int) bool {
	return func(a, b int) bool {
		return r.knows[a][b] == 1
	}
}

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	// When A knows B and B knows A, neither is the celebrity
	// When A knows B and B does not know A, A could be the celebrity (B is not)
	// When A does not know B and B does not know A, neither could be the celebrity
	return func(n int) int {
		normalPerson := make([]bool, n)
		for a := 0; a < n; a++ {
			// If there is a person that does not know A, then A cannot
			// be the celebrity
			// If person B does know A, then B cannot be the celebrity
			for b := a + 1; b < n; b++ {
				if normalPerson[a] && normalPerson[b] {
					continue
				}
				if knows(a, b) {
					normalPerson[a] = true
				} else {
					normalPerson[b] = true
				}
				if knows(b, a) {
					normalPerson[b] = true
				} else {
					normalPerson[a] = true
				}
			}
		}
		res := -1
		for i, normal := range normalPerson {
			if !normal {
				res = i
			}
		}
		return res
	}
}
