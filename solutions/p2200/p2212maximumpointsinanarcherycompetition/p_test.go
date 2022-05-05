package p2212maximumpointsinanarcherycompetition

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maximumBobPoints(t *testing.T) {
	for _, tc := range []struct {
		numArrows   int
		aliceArrows []int
		want        []int
	}{
		{89, []int{3, 2, 28, 1, 7, 1, 16, 7, 3, 13, 3, 5}, []int{21, 3, 0, 2, 8, 2, 17, 8, 4, 14, 4, 6}},
		{9, []int{1, 1, 0, 1, 0, 0, 2, 1, 0, 1, 2}, []int{0, 0, 0, 0, 1, 1, 0, 0, 1, 2, 3, 1}},
		{3, []int{0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 2}, []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.numArrows), func(t *testing.T) {
			require.Equal(t, tc.want, maximumBobPoints(tc.numArrows, tc.aliceArrows))
		})
	}
}

func maximumBobPoints(numArrows int, aliceArrows []int) []int {
	// Classic DP exercise
	// Bob can either choose to shoot to win, shoot to go on par, or do nothing.

	seen := make(map[[2]int]struct{})
	var f maxFinder
	f.pre = make([]int, 12)
	f.res = make([]int, 12)
	f.dp(seen, 0, numArrows, 0, aliceArrows)
	return f.res
}

type maxFinder struct {
	maxPts int
	res    []int
	pre    []int
}

func (f *maxFinder) dp(seen map[[2]int]struct{}, i, arrowsLeft, pts int, aliceArrows []int) {
	if i == 12 {
		if pts > f.maxPts {
			if arrowsLeft > 0 {
				f.pre[11] += arrowsLeft
			}
			f.maxPts = pts
			copy(f.res, f.pre)
		}
		return
	}
	// Don't shoot
	f.pre[i] = 0
	f.dp(seen, i+1, arrowsLeft, pts, aliceArrows)

	// Or shoot
	var a int
	if i < len(aliceArrows) {
		a = aliceArrows[i]
	}
	if arrowsLeft > a {
		// Shoot more than alice
		f.pre[i] = a + 1
		f.dp(seen, i+1, arrowsLeft-a-1, pts+i, aliceArrows)
	}
	// seen[k] = struct{}{}
}
