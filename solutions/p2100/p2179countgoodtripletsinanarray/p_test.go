package p2179

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_goodTriplets(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		want  int64
	}{
		{[]int{0, 5, 1, 4, 2, 3}, []int{5, 1, 0, 4, 3, 2}, 9},
		// {[]int{2, 0, 1, 3}, []int{0, 1, 2, 3}, 1},
		// {[]int{4, 0, 1, 3, 2}, []int{4, 1, 0, 2, 3}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, goodTriplets(tc.nums1, tc.nums2))
		})
	}
}

func goodTriplets(nums1 []int, nums2 []int) int64 {
	n := len(nums1)

	idx := make([]int, n)
	for i := range nums1 {
		idx[nums1[i]] = i
	}

	n1Pos := make([]int, n)
	for i, n2 := range nums2 {
		n1Pos[i] = idx[n2]
	}

	right := NewFenwick(n)
	left := NewFenwick(n)

	var count int64
	for i := 2; i < n; i++ {
		right.Add(n1Pos[i], 1)
	}
	left.Add(n1Pos[0], 1)

	for i := 1; i < n-1; i++ {
		below := left.Sum(n1Pos[i])
		above := right.Sum(n-1) - right.Sum(n1Pos[i])
		count += int64(below * above)
		left.Add(n1Pos[i], 1)
		right.Add(n1Pos[i+1], -1)
	}

	return int64(count)
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
