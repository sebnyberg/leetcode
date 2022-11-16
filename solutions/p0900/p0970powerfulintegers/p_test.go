package p0970powerfulintegers

import (
	"fmt"
	"math"
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
	m := make(map[int]struct{})
	a := math.MaxInt32
	var res []int
	for xx := 1; xx <= bound; xx *= x {
		for yy := 1; yy+xx <= bound; yy *= y {
			if _, exists := m[xx+yy]; exists {
				continue
			}
			res = append(res, xx+yy)
			m[xx+yy] = struct{}{}
			if y == 1 {
				break
			}
		}
		if x == 1 {
			break
		}
	}
	return res
}
