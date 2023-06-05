package p1739buildingboxes

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimumBoxes(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{10, 6},
		{4, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumBoxes(tc.n))
		})
	}
}

func minimumBoxes(n int) int {
	// With a piece of paper, I figured out that each round (1, 2, 3), we end up
	// with a highest height of the round, and the same number of boxes as the
	// round. E.g.
	//
	// [1]
	//
	// Then
	//
	// [2][1]
	// [1]
	//
	// Then
	//
	// [3][2][1]
	// [2][1]
	// [1]
	//
	// There are a couple of ways to add boxes to the new round, but the one i
	// found to be most useful is to fill row-by-row:
	//
	// [3][2][1]
	// [2][1]
	// [1]
	// [1] <
	//
	// [3][2][1]
	// [2][1]
	// [1][1] <
	// [1]
	//
	// [3][2][1]
	// [2][1]
	//>[2][1]
	// [1]
	//
	// [3][2][1]
	// [2][1][1] <
	// [2][1]
	// [1]
	//
	// etc.
	//
	// To count the maximum number of boxes for a certain height, just run the
	// series height*1 + (height-1)*2 + (height-2)*3 + ... + 1*height
	//
	// Once we know the total number of boxes, we can find where to run the
	// box-by-box approach.

	// count(dim) counts the number of boxes in a tower of dimension dim. The
	// dim, it turns out, is the same both in terms of height, width and height.
	count := func(dim int) int {
		var res int
		for x := 1; x <= dim; x++ {
			res += x * (dim - x + 1)
		}
		return res
	}
	var height int
	for height = 1; ; height++ {
		y := count(height + 1)
		if y > n {
			break
		}
	}
	n -= count(height)
	bottomBoxes := (height * (height + 1)) / 2
	for x := 1; n > 0 && x <= height+1; x++ {
		n -= x
		bottomBoxes++
	}
	return bottomBoxes
}
