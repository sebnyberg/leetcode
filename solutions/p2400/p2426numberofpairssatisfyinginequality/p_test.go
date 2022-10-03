package p2426numberofpairssatisfyinginequality

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOfPairs(t *testing.T) {
	for _, tc := range []struct {
		nums1 []int
		nums2 []int
		diff  int
		want  int64
	}{
		{[]int{5, 3}, []int{1, 2}, 2, 0},
		{[]int{3, 2, 5}, []int{2, 2, 1}, 1, 3},
		{[]int{3, -1}, []int{-2, 2}, -1, 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums1), func(t *testing.T) {
			require.Equal(t, tc.want, numberOfPairs(tc.nums1, tc.nums2, tc.diff))
		})
	}
}

func numberOfPairs(nums1 []int, nums2 []int, diff int) int64 {
	n := len(nums1)
	d := make([]int, n)
	for i := range nums1 {
		d[i] = nums1[i] - nums2[i]
	}
	var off int = 3*1e4 + 4
	t := NewFenwick(6*1e4 + 8)
	var res int64
	for _, x := range d {
		want := x + diff
		got := t.Sum(want + off)
		res += int64(got)
		t.Add(x+off, 1)
	}
	return res
}

type Fenwick struct {
	tree []int32
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int32, n+1),
	}
}

func (f *Fenwick) Init(vals []int32) {
	n := len(vals)
	copy(f.tree, vals)
	for i := 1; i < n; i++ {
		p := i + (i & -i)
		if p < n {
			f.tree[p] += f.tree[i]
		}
	}
}

func (f *Fenwick) Add(i int, k int32) {
	i++
	for i < len(f.tree) {
		f.tree[i] += k
		i += i & -i
	}
}

func (f *Fenwick) Sum(i int) int64 {
	i++
	var res int64
	for i > 0 {
		res += int64(f.tree[i])
		i -= i & -i
	}
	return res
}
