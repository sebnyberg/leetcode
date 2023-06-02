package p1655distributerepeatingintegers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canDistribute(t *testing.T) {
	for i, tc := range []struct {
		nums     []int
		quantity []int
		want     bool
	}{
		{
			[]int{1, 2, 3, 3},
			[]int{2},
			true,
		},
		// {
		// 	[]int{1, 1, 2, 2, 1},
		// 	[]int{2, 3},
		// 	false,
		// },
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, canDistribute(tc.nums, tc.quantity))
		})
	}
}

func canDistribute(nums []int, quantity []int) bool {
	// Obviously, only counts matter
	idx := make(map[int]int)
	counts := []int{}
	for _, x := range nums {
		if _, exists := idx[x]; !exists {
			idx[x] = len(counts)
			counts = append(counts, 0)
		}
		counts[idx[x]]++
	}
	// This looks like a DP / knapsack problem. We could do: given the i'th
	// number and bitmask of current satisfied customers, store whether that
	// state does not have a solution (for early exits).
	mem := make(map[[2]int]bool)
	possible := dfs(mem, counts, quantity, 0, 0)
	return possible
}

func dfs(impossible map[[2]int]bool, counts, quantity []int, i, bm int) bool {
	n := len(quantity)
	if bm == (1<<n)-1 {
		return true
	}
	if i == len(counts) {
		return false
	}
	key := [2]int{i, bm}
	if impossible[key] {
		return false
	}
	for x := 0; x < (1 << n); x++ {
		if x&bm > 0 {
			continue
		}
		var amt int
		for i := range quantity {
			if x&(1<<i) > 0 {
				amt += quantity[i]
			}
		}
		if amt > counts[i] {
			continue
		}
		ok := dfs(impossible, counts, quantity, i+1, bm|x)
		if ok {
			return true
		}
	}
	impossible[key] = true
	return false
}
