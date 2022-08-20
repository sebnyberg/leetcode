package p0835imageoverlap

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_largestOverlap(t *testing.T) {
	for _, tc := range []struct {
		img1 [][]int
		img2 [][]int
		want int
	}{
		{
			leetcode.ParseMatrix("[[0,1],[1,1]]"),
			leetcode.ParseMatrix("[[1,1],[1,0]]"),
			2,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.img1), func(t *testing.T) {
			require.Equal(t, tc.want, largestOverlap(tc.img1, tc.img2))
		})
	}
}

func largestOverlap(img1 [][]int, img2 [][]int) int {
	// There are only 30*30 different offsets, so we can just try all options
	m := len(img1)
	n := len(img1[0])
	shifted := make([][]int, m)
	for i := range shifted {
		shifted[i] = make([]int, n)
	}
	reset := func() {
		for i := range shifted {
			for j := range shifted[i] {
				shifted[i][j] = 0
			}
		}
	}
	shift := func(dx, dy int) {
		reset()
		for i := range img1 {
			for j, v := range img1[i] {
				ii := i + dy
				jj := j + dx
				if ii < 0 || jj < 0 || ii >= m || jj >= n {
					continue
				}
				shifted[ii][jj] = v
			}
		}
	}
	countOverlap := func() int {
		var res int
		for i := range img2 {
			for j, v := range img2[i] {
				res += v & shifted[i][j]
			}
		}
		return res
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var res int
	for dy := -m; dy <= m; dy++ {
		for dx := -n; dx <= n; dx++ {
			shift(dx, dy)
			res = max(res, countOverlap())
		}
	}
	return res
}
