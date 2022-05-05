package p1799maxscoreafternoperations

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxScore(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 14},
		{[]int{1, 2}, 1},
		{[]int{3, 4, 6, 8}, 11},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, maxScore(tc.nums))
		})
	}
}

func maxScore(nums []int) int {
	n := len(nums)
	var mem [8][16384]int
	gcdCache := &cachedGCD{
		mem: make(map[[2]int]int, 1000),
	}
	res := dfs(&mem, gcdCache, nums, n, 1, 0)
	return res
}

func dfs(mem *[8][16384]int, gcd *cachedGCD, nums []int, n int, i int, mask int) int {
	if i > (n / 2) {
		return 0
	}
	if mem[i][mask] != 0 {
		return mem[i][mask]
	}
	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			b := (1 << j) + (1 << k)
			if mask&b == 0 {
				mem[i][mask] = max(mem[i][mask], i*gcd.gcd(nums[j], nums[k])+dfs(mem, gcd, nums, n, i+1, mask+b))
			}
		}
	}
	return mem[i][mask]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type cachedGCD struct {
	mem map[[2]int]int
}

func (g *cachedGCD) gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	k := [2]int{a, b}
	if _, exists := g.mem[k]; !exists {
		for b != 0 {
			a, b = b, a%b
		}
		g.mem[k] = a
	}
	return g.mem[k]
}
