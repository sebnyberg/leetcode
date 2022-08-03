package p0638shoppingoffers

import (
	"fmt"
	"github.com/sebnyberg/leetcode"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_shoppingOffers(t *testing.T) {
	for _, tc := range []struct {
		price   []int
		special [][]int
		needs   []int
		want    int
	}{
		{
			[]int{2, 5},
			leetcode.ParseMatrix("[[3,0,5],[1,2,10]]"),
			[]int{3, 2},
			14,
		},
		{
			[]int{2, 3, 4},
			leetcode.ParseMatrix("[[1,1,0,4],[2,2,1,9]]"),
			[]int{1, 2, 1},
			11,
		},
	} {
		t.Run(fmt.Sprintf("%+v", tc.price), func(t *testing.T) {
			require.Equal(t, tc.want, shoppingOffers(tc.price, tc.special, tc.needs))
		})
	}
}

func shoppingOffers(price []int, special [][]int, needs []int) int {
	mem := make(map[string]int)
	res := dfs(mem, price, special, needs, len(price))
	return res
}

func done(needs []int) bool {
	for _, n := range needs {
		if n > 0 {
			return false
		}
	}
	return true
}

func dfs(mem map[string]int, price []int, special [][]int, needs []int, n int) int {
	if done(needs) {
		return 0
	}
	k := key(needs)
	if v, exists := mem[k]; exists {
		return v
	}
	res := dot(needs, price)
	pp := make([]int, len(needs))
	for i := range special {
		copy(pp, needs)
		var j int
		for j = 0; j < n; j++ {
			pp[j] -= special[i][j]
			if pp[j] < 0 {
				break
			}
		}
		if j == n {
			res = min(res, special[i][j]+dfs(mem, price, special, pp, n))
		}
	}
	mem[k] = res
	return mem[k]
}

func dot(needs []int, prices []int) int {
	var sum int
	for i := range needs {
		sum += needs[i] * prices[i]
	}
	return sum
}

func key(needs []int) string {
	ss := make([]string, len(needs))
	for i := range needs {
		ss[i] = strconv.Itoa(needs[i])
	}
	return strings.Join(ss, ",")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
