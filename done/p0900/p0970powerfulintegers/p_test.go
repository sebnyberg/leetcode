package p0970powerfulintegers

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_powerfulIntegers(t *testing.T) {
	for _, tc := range []struct {
		x     int
		y     int
		bound int
		want  []int
	}{
		{1, 1, 0, []int{}},
		{2, 1, 10, []int{2, 3, 5, 9}},
		{2, 3, 10, []int{2, 3, 4, 5, 7, 9, 10}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.x), func(t *testing.T) {
			require.Equal(t, tc.want, powerfulIntegers(tc.x, tc.y, tc.bound))
		})
	}
}

func powerfulIntegers(x int, y int, bound int) []int {
	if y < x {
		x, y = y, x
	}
	if x == 1 && y == 1 {
		if bound >= 2 {
			return []int{2}
		}
		return []int{}
	}
	curX := 1
	res := make(map[int]struct{})
	for curX < bound {
		curY := 1
		for curX+curY <= bound {
			res[curX+curY] = struct{}{}
			curY *= y
		}
		if x == 1 {
			break
		}
		curX *= x
	}
	reslist := make([]int, 0, len(res))
	for k := range res {
		reslist = append(reslist, k)
	}
	sort.Ints(reslist)
	return reslist
}
