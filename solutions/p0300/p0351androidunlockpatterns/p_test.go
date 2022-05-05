package p0351androidunlockpatterns

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfPatterns(t *testing.T) {
	for _, tc := range []struct {
		m    int
		n    int
		want int
	}{
		{1, 1, 9},
		{1, 2, 65},
		{1, 4, 2009},
	} {
		t.Run(fmt.Sprintf("%+v", tc.m), func(t *testing.T) {
			res := numberOfPatterns(tc.m, tc.n)
			require.Equal(t, tc.want, res)
		})
	}
}

func numberOfPatterns(m int, n int) int {
	// avoid[i][j] contains the index to avoid when drawing a line from i to j
	var avoid [10][10]int
	avoid[1][3] = 2
	avoid[1][7] = 4
	avoid[3][9] = 6
	avoid[7][9] = 8
	avoid[1][9] = 5
	avoid[2][8] = 5
	avoid[3][7] = 5
	avoid[4][6] = 5
	for i := range avoid {
		for j := range avoid[i] {
			if j > i {
				avoid[j][i] = avoid[i][j]
			}
		}
	}

	var res int
	for i := m; i <= n; i++ {
		res += combs(1<<1, &avoid, 1, i-1) * 4
		res += combs(1<<2, &avoid, 2, i-1) * 4
		res += combs(1<<5, &avoid, 5, i-1)
	}
	return res
}

func combs(bm int, avoid *[10][10]int, prev, remain int) int {
	if remain == 0 {
		return 1
	}
	var res int
	bit := 1
	for i := 1; i <= 9; i++ {
		bit <<= 1
		if bm&bit == 0 && (avoid[prev][i] == 0 || bm&(1<<avoid[prev][i]) > 0) {
			res += combs(bm+bit, avoid, i, remain-1)
		}
	}
	return res
}
