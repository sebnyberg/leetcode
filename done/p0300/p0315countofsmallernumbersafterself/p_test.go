package p0315countofsmallernumbersafterself

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countSmaller(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 6, 1}, []int{2, 1, 1, 0}},
		{[]int{-1}, []int{0}},
		{[]int{-1, -1}, []int{0, 0}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, countSmaller(tc.nums))
		})
	}
}

func countSmaller(nums []int) []int {
	// Idea: shift each number in nums by the smallest possible number
	// Then create a BIT where the index is the number, and the value at each
	// index is the number of occurrences at that index
	f := NewFenwick(20002)
	offset := 10000
	n := len(nums)
	res := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		res[i] = f.Sum(nums[i] + offset - 1)
		f.Add(nums[i]+offset, 1)
	}
	return res
}

type Fenwick struct {
	tree []int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{
		tree: make([]int, n),
	}
}

// Init initializes the Fenwick tree, overwriting old content with
// the provided list
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

// Add k to index i
func (f *Fenwick) Add(i int, k int) {
	for i < len(f.tree) {
		f.tree[i] += k
		i += i & -i
	}
}

func (f *Fenwick) Sum(i int) int {
	res := 0
	for i > 0 {
		res += f.tree[i]
		// flip last set bit to zero (use parent in tree)
		i -= i & -i
	}
	return res
}

func (f *Fenwick) SumRange(from int, to int) int {
	return f.Sum(to) - f.Sum(from)
}
