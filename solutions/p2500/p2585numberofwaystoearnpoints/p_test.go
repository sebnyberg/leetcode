package p2585numberofwaystoearnpoints

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_waysToReachTarget(t *testing.T) {
	for i, tc := range []struct {
		target int
		types  [][]int
		want   int
	}{
		{
			6,
			leetcode.ParseMatrix("[[6,1],[3,2],[2,3]]"),
			7,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, waysToReachTarget(tc.target, tc.types))
		})
	}
}

const mod = 1e9 + 7

func waysToReachTarget(target int, types [][]int) int {
	mem := make(map[[2]int]int)
	res := dfs(mem, types, target, 0)
	return res
}

func dfs(mem map[[2]int]int, types [][]int, target, i int) int {
	if target == 0 {
		return 1
	}
	if i >= len(types) {
		return 0
	}
	k := [2]int{target, i}
	if v, exists := mem[k]; exists {
		return v
	}
	var res int
	count := types[i][0]
	marks := types[i][1]
	for k := 0; k <= count && target-k*marks >= 0; k++ {
		a := dfs(mem, types, target-k*marks, i+1)
		res = (res + a) % mod
	}
	mem[k] = res
	return res
}
