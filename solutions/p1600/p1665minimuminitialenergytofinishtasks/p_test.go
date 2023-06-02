package p1665minimuminitialenergytofinishtasks

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumEffort(t *testing.T) {
	for i, tc := range []struct {
		tasks [][]int
		want  int
	}{
		{
			leetcode.ParseMatrix("[[1,3],[2,4],[10,11],[10,12],[8,9]]"),
			32,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumEffort(tc.tasks))
		})
	}
}

func minimumEffort(tasks [][]int) int {
	var t int
	sort.Slice(tasks, func(i, j int) bool {
		a := tasks[i]
		b := tasks[j]
		da := a[1] - a[0]
		db := b[1] - b[0]
		return da < db
	})
	for _, x := range tasks {
		t = max(t+x[0], x[1])
	}
	return t
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
