package p2064minimizedmaximumofproductsdistributedtoanystore

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minimizedMaximum(t *testing.T) {
	for _, tc := range []struct {
		n          int
		quantities []int
		want       int
	}{
		{1, []int{1}, 1},
		{4, []int{2, 2, 8, 7}, 8},
		{2, []int{5, 7}, 7},
		{7, []int{17, 13, 10}, 7},
		{6, []int{11, 6}, 3},
		{7, []int{15, 10, 10}, 5},
		{1, []int{1000000}, 1000000},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, minimizedMaximum(tc.n, tc.quantities))
		})
	}
}

func minimizedMaximum(n int, quantities []int) int {
	canDivide := func(part int) bool {
		var count int
		for _, q := range quantities {
			count += q / part
			if q%part != 0 {
				count++
			}
			if count > n {
				return false
			}
		}
		return true
	}

	lo, hi := 1, 10000001
	for lo < hi {
		mid := (lo + hi) / 2
		if !canDivide(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
