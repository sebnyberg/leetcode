package p1206designskiplist

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSkiplist(t *testing.T) {
	a := Constructor()
	a.Add(1)
	a.Add(2)
	a.Add(3)
	res := a.Search(0)
	require.False(t, res)
}

const layers = 4

type node struct {
	val  int
	next [layers]*node
}

type Skiplist struct {
	dummy *node
}

func Constructor() Skiplist {
	return Skiplist{
		dummy: &node{val: -1},
	}
}

func next(n *node) int {
	if n != nil {
		return n.val
	}
	return math.MaxInt32
}

func (this *Skiplist) Search(target int) bool {
	curr := this.dummy
	for l := layers - 1; l >= 0; l-- {
		for target > next(curr.next[l]) {
			curr = curr.next[l]
		}
	}
	return curr.val == target
}

func (this *Skiplist) Add(num int) {
	curr := this.dummy
	// pre[i] is the node prior to the node that is being inserted on each level
	pre := make([]*node, layers)
	for l := layers - 1; l >= 0; l-- {
		for num > next(curr.next[l]) {
			curr = curr.next[l]
		}
		pre[l] = curr
	}
	level := randomLevel()
	node := &node{val: num}
	for i := 0; i < level; i++ {
		pre[i].next[i], node.next[i] = node, pre[i].next[i]
	}
}

func randomLevel() int {
	level := 1
	for rand.Intn(2) == 0 && level < layers {
		level++
	}
	return level
}

func (this *Skiplist) Erase(num int) bool {
	return true
}
