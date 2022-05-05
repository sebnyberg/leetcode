package p0493reversepairs

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reversePairs(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{233, 2000000001, 234, 2000000006, 235, 2000000003, 236, 2000000007, 237, 2000000002, 2000000005, 233, 233, 233, 233, 233, 2000000004}, 40},
		{[]int{2, 2, -2, -2, -2, 2}, 9},
		{[]int{2147483647, 2147483647, -2147483647, -2147483647, -2147483647, 2147483647}, 9},
		{[]int{-5, -5}, 1},
		{[]int{1, 3, 2, 3, 1}, 2},
		{[]int{2, 4, 3, 5, 1}, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, reversePairs(tc.nums))
		})
	}
}

func reversePairs(nums []int) int {
	// Idea: have a fenwick tree with indices of numbers in sorted order + have
	// sorted numbers from nums.
	// Then for each element from the end to start, find the index of a number
	// that is greater than or equal to 2x the start, and do a range query of
	// number of elements above that index.
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	dedup := func(a []int) []int {
		var j int
		for i := 1; i < len(a) && j < len(a); i++ {
			if a[i] == a[j] {
				continue
			}
			j++
			a[j] = a[i]
		}
		a = a[:j+1]
		return a
	}
	// De-duplicate sorted list
	sort.Ints(sorted)
	sorted = append(sorted, 0, 0)
	copy(sorted[1:], sorted)
	sorted[0] = math.MinInt64
	sorted[len(sorted)-1] = math.MaxInt64
	sorted = dedup(sorted)
	numIdx := make(map[int]int, n)
	for i, n := range sorted {
		numIdx[n] = i
	}
	m := len(sorted)
	f := NewFenwick(m + 1)
	var res int
	for i := 0; i < n; i++ {
		num := nums[i]
		// Find smallest element greater than current
		want := num * 2
		idx := sort.SearchInts(sorted, want+1)
		above := f.Sum(m)
		below := f.Sum(idx - 1)
		res += above - below
		f.Add(numIdx[num], 1)
	}
	return res
}

type Fenwick struct {
	tree []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int, n+1),
	}
}

func (f *Fenwick) Init(vals []int) {
	n := len(vals)
	copy(f.tree, vals)
	for i := 1; i < n; i++ {
		p := i + (i & -i)
		if p < n {
			f.tree[p] += f.tree[i]
		}
	}
}

func (f *Fenwick) Add(i int, k int) {
	i++
	for i < len(f.tree) {
		f.tree[i] += k
		i += i & -i
	}
}

func (f *Fenwick) Sum(i int) int {
	i++
	res := 0
	for i > 0 {
		res += f.tree[i]
		i -= i & -i
	}
	return res
}
