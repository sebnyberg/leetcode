package p1482minumumnumberofdayfstomakembouquets

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDays(t *testing.T) {
	for _, tc := range []struct {
		bloomDay []int
		m        int
		k        int
		want     int
	}{
		{[]int{1, 10, 3, 10, 2}, 3, 1, 3},
		{[]int{1, 10, 3, 10, 2}, 3, 2, -1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.bloomDay), func(t *testing.T) {
			require.Equal(t, tc.want, minDays(tc.bloomDay, tc.m, tc.k))
		})
	}
}

func minDays(bloomDay []int, m int, k int) int {
	// Either we add flowers and consider the number of bouquets we can make from
	// a given sequence, for example with union-find. Or we can consider the case
	// when the entire flowerbed has bloomed and remove flowers, splitting
	// sequences and calculating the number of bouquets lost from each split. In
	// both cases, we need a way to represent sequences.
	//
	// Going backwards is likely more efficient, but requires a lot of logic,
	// whereas union-find should be quite straight-forward.

	n := len(bloomDay)
	parent := make([]int, n)
	counts := make([]int, n)
	for i := range parent {
		parent[i] = i
		counts[i] = 1
	}

	var find func(a int) int
	find = func(a int) int {
		if parent[a] != a {
			// path compression
			root := find(parent[a])
			parent[a] = root
		}
		return parent[a]
	}
	var res int

	union := func(a, b int) {
		rootA, rootB := find(a), find(b)
		res -= (counts[rootA]/k + counts[rootB]/k)
		res += (counts[rootA] + counts[rootB]) / k
		counts[rootB] += counts[rootA]
		parent[rootA] = rootB
	}

	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return bloomDay[idx[i]] < bloomDay[idx[j]]
	})
	for _, x := range idx {
		if k == 1 {
			res++
		}
		if x > 0 && bloomDay[x-1] == 0 {
			union(x, x-1)
		}
		if x < n-1 && bloomDay[x+1] == 0 {
			union(x, x+1)
		}
		if res >= m {
			return bloomDay[x]
		}
		bloomDay[x] = 0
	}
	return -1
}
