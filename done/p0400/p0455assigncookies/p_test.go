package p0445assigncookies

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findContentChildren(t *testing.T) {
	for _, tc := range []struct {
		g    []int
		s    []int
		want int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.g), func(t *testing.T) {
			require.Equal(t, tc.want, findContentChildren(tc.g, tc.s))
		})
	}
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	var res int
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if s[j] >= g[i] {
			res++
			i++
		}
	}
	return res
}
