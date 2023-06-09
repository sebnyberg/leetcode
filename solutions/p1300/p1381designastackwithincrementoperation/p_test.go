package p1381designastackwithincrementoperation

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDesignStack(t *testing.T) {
	a := Constructor(3)
	a.Push(1)
	a.Push(2)
	require.Equal(t, 2, a.Pop())
	a.Push(2)
	a.Push(3)
	a.Push(4)
	a.Increment(5, 100)
	a.Increment(2, 100)
	require.Equal(t, 103, a.Pop())
}

type CustomStack struct {
	items   []int
	changes []int
	n       int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		changes: make([]int, maxSize),
		items:   make([]int, 0, maxSize),
		n:       maxSize,
	}
}

func (this *CustomStack) Push(x int) {
	m := len(this.items)
	if m >= this.n {
		return
	}
	this.items = append(this.items, x)
}

func (this *CustomStack) Pop() int {
	m := len(this.items)
	if m == 0 {
		return -1
	}
	// There is an item. If the item has a pending change, add it to the value
	// and push the change further down the stack.
	d := this.changes[m-1]
	x := this.items[m-1] + d
	this.changes[m-1] = 0
	if m != 1 {
		this.changes[m-2] += d
	}
	this.items = this.items[:m-1]
	return x
}

func (this *CustomStack) Increment(k int, val int) {
	m := len(this.items)
	if m == 0 {
		return
	}
	this.changes[min(m-1, k-1)] += val
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
