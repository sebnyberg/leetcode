package p1411numberofwaystopaintnx3grid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numOfWays(t *testing.T) {
	for i, tc := range []struct {
		n    int
		want int
	}{
		{1, 12},
		{2, 54},
		{5000, 30228214},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, numOfWays(tc.n))
		})
	}
}

var invalid = []int{
	1<<0 | 1<<3,
	1<<1 | 1<<4,
	1<<2 | 1<<5,
	1<<3 | 1<<6,
	1<<4 | 1<<7,
	1<<5 | 1<<8,
}

const mod = 1e9 + 7

func numOfWays(n int) int {
	alternatives := []int{}
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			for z := 0; z < 3; z++ {
				col := (1 << x) | (1 << (y + 3)) | (1 << (z + 6))
				for _, v := range invalid {
					if v&col == v {
						goto notValid
					}
				}
				alternatives = append(alternatives, col)
			notValid:
			}
		}
	}
	curr := map[int]int{}
	for _, x := range alternatives {
		curr[x]++
	}
	next := map[int]int{}
	for k := 2; k <= n; k++ {
		for k := range next {
			delete(next, k)
		}
		for x, count := range curr {
			for _, y := range alternatives {
				if x&y == 0 {
					next[y] = (next[y] + count) % mod
				}
			}
		}
		curr, next = next, curr
	}
	var res int
	for _, c := range curr {
		res = (res + c) % mod
	}
	return res
}
