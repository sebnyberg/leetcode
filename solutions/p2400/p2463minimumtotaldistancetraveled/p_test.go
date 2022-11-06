package p2463minimumtotaldistancetraveled

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_minimumTotalDistance(t *testing.T) {
	for i, tc := range []struct {
		robot   []int
		factory [][]int
		want    int64
	}{
		{
			[]int{1, -1},
			leetcode.ParseMatrix("[[-2,1],[2,1]]"),
			2,
		},
		{
			[]int{0, 4, 6},
			leetcode.ParseMatrix("[[2,2],[6,2]]"),
			4,
		},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minimumTotalDistance(tc.robot, tc.factory))
		})
	}
}

func minimumTotalDistance(robot []int, factory [][]int) int64 {
	sort.Slice(factory, func(i, j int) bool {
		return factory[i][0] < factory[j][0]
	})
	var mem [101][101][101]int
	for i := range mem {
		for j := range mem {
			for k := range mem {
				mem[i][j][k] = -1
			}
		}
	}
	sort.Ints(robot)
	res := dp(&mem, robot, factory, 0, 0, 0)
	return int64(res)
}

func dp(mem *[101][101][101]int, robot []int, factory [][]int, i, j, k int) int {
	if i == len(robot) {
		return 0
	}
	if j == len(factory) {
		return math.MaxInt32 << 10
	}
	if mem[i][j][k] != -1 {
		return mem[i][j][k]
	}
	res := dp(mem, robot, factory, i, j+1, 0)
	if factory[j][1] > k {
		res = min(res, abs(robot[i]-factory[j][0])+dp(mem, robot, factory, i+1, j, k+1))
	}
	mem[i][j][k] = res
	return mem[i][j][k]
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
