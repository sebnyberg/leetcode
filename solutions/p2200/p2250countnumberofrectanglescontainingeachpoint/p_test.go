package p2250countnumberofrectanglescontainingeachpoint

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countRectangles(t *testing.T) {
	for _, tc := range []struct {
		rectangles [][]int
		points     [][]int
		want       []int
	}{
		{
			leetcode.ParseMatrix("[[1,2],[2,3],[2,5]]"),
			leetcode.ParseMatrix("[[2,1],[1,4]]"),
			[]int{2, 1},
		},
		{
			leetcode.ParseMatrix("[[1,1],[2,2],[3,3]]"),
			leetcode.ParseMatrix("[[1,3],[1,1]]"),
			[]int{1, 3},
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rectangles), func(t *testing.T) {
			require.Equal(t, tc.want, countRectangles(tc.rectangles, tc.points))
		})
	}
}

func countRectangles(rectangles [][]int, points [][]int) []int {
	cover := make([][]int, 102)
	for _, r := range rectangles {
		for y := 0; y <= r[1]; y++ {
			cover[y] = append(cover[y], r[0])
		}
	}
	for i := range cover {
		sort.Ints(cover[i])
	}
	res := make([]int, len(points))
	for i, p := range points {
		a := cover[p[1]]
		j := sort.SearchInts(a, p[0])
		if j == len(a) {
			res[i] = 0
			continue
		}
		res[i] = len(a) - j
	}
	return res
}
