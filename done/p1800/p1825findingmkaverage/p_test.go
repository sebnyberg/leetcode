package p1825findingmkaverage

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type action func(t *testing.T, a *MKAverage)

func Add(num int) action {
	return func(t *testing.T, a *MKAverage) {
		a.AddElement(num)
	}
}

func Avg(want int) action {
	return func(t *testing.T, a *MKAverage) {
		res := a.CalculateMKAverage()
		require.Equal(t, want, res)
	}
}

func TestMKAverage(t *testing.T) {
	for _, tc := range []struct {
		m, n    int
		actions []action
	}{
		{
			6, 1,
			[]action{Add(3), Add(1), Add(12), Add(5), Add(3), Add(4), Avg(3)},
		},
		{
			3, 1,
			[]action{
				Add(3), Add(1),
				Avg(-1),
				Add(10),
				Avg(3),
				Add(5), Add(5), Add(5), Avg(5),
			},
		},
	} {
		t.Run(fmt.Sprintf("%v/%v/%v", tc.m, tc.n, len(tc.actions)), func(t *testing.T) {
			avg := Constructor(tc.m, tc.n)
			for _, a := range tc.actions {
				a(t, &avg)
			}
		})
	}
}

type MKAverage struct {
	m, k int
	// Elements is a ring-buffer of size m
	buf     *RingBuffer
	indices *Fenwick
	sums    *Fenwick
}

// Create a massive Fenwick tree which contains all values
// When adding a number, increment that value in the fenwick tree
// When removing, decrement that value in the tree
// When calculating the sum, use a deque to figure out where to start / end
func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:       m,
		k:       k,
		buf:     NewRingBuffer(m),
		indices: NewFenwick(1e5 + 1),
		sums:    NewFenwick(1e5 + 1),
	}
}

func (this *MKAverage) AddElement(num int) {
	overflow, prev := this.buf.Insert(num)
	if overflow {
		this.indices.Add(prev, -1)
		this.sums.Add(prev, -prev)
	}
	this.indices.Add(num, 1)
	this.sums.Add(num, num)
}

// Find index of k in the fenwick tree
func (this *MKAverage) getIndex(k int) int {
	return sort.Search(int(1e5), func(i int) bool {
		return this.indices.Sum(i) > k
	})
}

func (this *MKAverage) CalculateMKAverage() int {
	if this.buf.Len() < this.m {
		return -1
	}
	lo, hi := this.getIndex(this.k), this.getIndex(this.m-this.k)
	sum := this.sums.SumRange(lo, hi)
	sum += (this.indices.Sum(lo) - this.k) * lo
	sum -= (this.indices.Sum(hi) - (this.m - this.k)) * hi
	return sum / (this.m - 2*this.k)
}

type RingBuffer struct {
	curLen    int
	maxLen    int
	items     []int
	insertPos int
}

func NewRingBuffer(n int) *RingBuffer {
	return &RingBuffer{
		curLen:    0,
		maxLen:    n,
		items:     make([]int, n),
		insertPos: 0,
	}
}

func (b *RingBuffer) Len() int { return b.curLen }

// Inserts a value into the ring buffer. If an existing element was replaced,
// replaced is set to true, and prev contains the replaced element
func (b *RingBuffer) Insert(val int) (overflow bool, prev int) {
	if b.curLen == b.maxLen {
		overflow = true
		prev = b.items[b.insertPos]
	} else if b.curLen < b.maxLen {
		b.curLen++
	}
	b.items[b.insertPos] = val
	b.insertPos++
	b.insertPos %= b.maxLen
	return overflow, prev
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
