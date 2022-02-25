package p0526beautifularrangement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countArrangement(t *testing.T) {
	for _, tc := range []struct {
		n    int
		want int
	}{
		{2, 2},
		{1, 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, countArrangement(tc.n))
		})
	}
}

func countArrangement(n int) int {
	res := dfs(make(map[int]int, n), 0, 1, n)
	return res
}

func dfs(mem map[int]int, bm, i, n int) int {
	if i > n {
		return 1
	}
	if v, exists := mem[bm]; exists {
		return v
	}

	var res int
	for j := 1; j <= n; j++ {
		if bm&(1<<j) > 0 {
			continue
		}
		if i%j == 0 || j%i == 0 {
			res += dfs(mem, bm|(1<<j), i+1, n)
		}
	}

	mem[bm] = res
	return mem[bm]
}
