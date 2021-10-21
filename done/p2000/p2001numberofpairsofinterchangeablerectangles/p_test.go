package p2001numberofpairsofinterchangeablerectangles

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_interchangeableRectangles(t *testing.T) {
	for _, tc := range []struct {
		rectangles [][]int
		want       int64
	}{
		{[][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}, 6},
		{[][]int{{4, 2}, {1, 3}, {4, 1}, {4, 2}, {2, 4}, {1, 1}, {1, 1}}, 2},
		{[][]int{{4, 5}, {7, 8}}, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.rectangles), func(t *testing.T) {
			require.Equal(t, tc.want, interchangeableRectangles(tc.rectangles))
		})
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func interchangeableRectangles(rectangles [][]int) int64 {
	// Divide each entry by its gcd
	for i, rect := range rectangles {
		div := gcd(rect[0], rect[1])
		rectangles[i][0] /= div
		rectangles[i][1] /= div
	}

	// Iterate over rectangles from end to finish, counting pairs
	n := len(rectangles)
	pairCount := make(map[[2]int]int)
	var res int
	for i := n - 1; i >= 0; i-- {
		w, h := rectangles[i][0], rectangles[i][1]
		res += pairCount[[2]int{w, h}]
		pairCount[[2]int{w, h}]++
	}

	return int64(res)
}
