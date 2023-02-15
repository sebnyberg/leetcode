package p1172dinnerplatestacks

import (
	"container/heap"
	"fmt"
	"testing"
)

func Test_DinnerPlates(t *testing.T) {
	c := Constructor(2)
	c.Push(1)
	c.Push(2)
	c.Push(3)
	c.Push(4)
	c.Push(5)
	fmt.Println(c.PopAtStack(0))
	c.Push(20)
	c.Push(21)
	fmt.Println(c.PopAtStack(0))
	fmt.Println(c.PopAtStack(2))
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
}

type DinnerPlates struct {
	cap    int
	push   *pushHeap
	pop    *popHeap
	stacks []*stack
}

func Constructor(capacity int) DinnerPlates {
	var dp DinnerPlates
	dp.cap = capacity
	first := &stack{}
	dp.pop = &popHeap{
		cap: capacity,
		vs:  []*stack{first},
	}
	dp.push = &pushHeap{
		cap: capacity,
		vs:  []*stack{first},
	}
	dp.stacks = append(dp.stacks, first)
	return dp
}

func (this *DinnerPlates) Push(val int) {
	n := len(this.push.vs[0].items)
	if n == this.cap {
		st := &stack{
			idx: len(this.push.vs),
		}
		heap.Push(this.push, st)
		heap.Push(this.pop, st)
		this.stacks = append(this.stacks, st)
	}

	this.push.vs[0].items = append(this.push.vs[0].items, val)
	heap.Fix(this.pop, this.push.vs[0].popIdx)
	heap.Fix(this.push, 0)
}

func (this *DinnerPlates) Pop() int {
	n := len(this.pop.vs[0].items)
	if n == 0 {
		return -1
	}
	val := this.pop.vs[0].items[n-1]
	this.pop.vs[0].items = this.pop.vs[0].items[:n-1]
	heap.Fix(this.push, this.pop.vs[0].pushIdx)
	heap.Fix(this.pop, 0)
	return val
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index >= len(this.stacks) {
		return -1
	}
	n := len(this.stacks[index].items)
	if n == 0 {
		return -1
	}
	x := this.stacks[index].items[n-1]
	this.stacks[index].items = this.stacks[index].items[:n-1]
	heap.Fix(this.push, this.stacks[index].pushIdx)
	heap.Fix(this.pop, this.stacks[index].popIdx)
	return x
}

type stack struct {
	idx     int
	pushIdx int
	popIdx  int
	items   []int
}

type pushHeap struct {
	vs  []*stack
	cap int
}

func (h pushHeap) Len() int { return len(h.vs) }
func (h pushHeap) Swap(i, j int) {
	h.vs[i], h.vs[j] = h.vs[j], h.vs[i]
	h.vs[i].pushIdx = i
	h.vs[j].pushIdx = j
}
func (h pushHeap) Less(i, j int) bool {
	if len(h.vs[i].items) >= h.cap {
		return false
	}
	if len(h.vs[j].items) >= h.cap {
		return true
	}
	return h.vs[i].idx < h.vs[j].idx
}
func (h *pushHeap) Push(x interface{}) {
	el := x.(*stack)
	el.pushIdx = len(h.vs)
	h.vs = append(h.vs, el)
}
func (h *pushHeap) Pop() interface{} {
	n := len(h.vs)
	el := h.vs[n-1]
	h.vs = h.vs[:n-1]
	return el
}

type popHeap struct {
	vs  []*stack
	cap int
}

func (h popHeap) Len() int { return len(h.vs) }
func (h popHeap) Swap(i, j int) {
	h.vs[i], h.vs[j] = h.vs[j], h.vs[i]
	h.vs[i].popIdx = i
	h.vs[j].popIdx = j
}
func (h popHeap) Less(i, j int) bool {
	if len(h.vs[i].items) == 0 {
		return false
	}
	if len(h.vs[j].items) == 0 {
		return true
	}
	return h.vs[i].idx > h.vs[j].idx
}
func (h *popHeap) Push(x interface{}) {
	el := x.(*stack)
	el.popIdx = len(h.vs)
	h.vs = append(h.vs, el)
}
func (h *popHeap) Pop() interface{} {
	n := len(h.vs)
	el := h.vs[n-1]
	h.vs = h.vs[:n-1]
	return el
}
