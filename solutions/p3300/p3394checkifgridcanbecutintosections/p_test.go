package p3394checkifgridcanbecutintosections

import (
	"fmt"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_checkValidCuts(t *testing.T) {
	for _, tc := range []struct {
		n          int
		rectangles [][]int
		want       bool
	}{
		{5, leetcode.ParseMatrix("[[1,0,5,2],[0,2,2,4],[3,2,5,3],[0,4,4,5]]"), true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, checkValidCuts(tc.n, tc.rectangles))
		})
	}
}

func checkValidCuts(n int, rectangles [][]int) bool {
	m := len(rectangles)
	type change struct {
		coord int
		delta int
	}
	xs := make([]change, 0, m)
	ys := make([]change, 0, m)
	for _, r := range rectangles {
		xs = append(xs, change{r[0], 1}, change{r[2], -1})
		ys = append(ys, change{r[1], 1}, change{r[3], -1})
	}
	sortFunc := func(ls []change) func(i, j int) bool {
		return func(i, j int) bool {
			if ls[i].coord == ls[j].coord {
				return ls[i].delta < ls[j].delta
			}
			return ls[i].coord < ls[j].coord
		}
	}
	sort.Slice(xs, sortFunc(xs))
	sort.Slice(ys, sortFunc(ys))
	canCut := func(ls []change) bool {
		var delta int
		var cuts int
		var i int
		for i < len(ls)-1 {
			delta += ls[i].delta
			i++
			if delta != 0 {
				continue
			}
			// no overlapping rectangles, CUT!
			cuts++
			if cuts == 2 {
				return true
			}
			for i < len(ls)-1 && ls[i+1].coord == ls[i].coord {
				delta += ls[i].delta
				i++
			}
		}
		return false
	}
	return canCut(xs) || canCut(ys)
}
