package p1206designskiplist

import (
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
	res = a.Search(1)
	require.True(t, res)
	res = a.Search(2)
	require.True(t, res)
	res = a.Search(3)
	require.True(t, res)
}

const (
	maxLevel = 32
)

type node struct {
	val  int
	next []*node
}

type Skiplist struct {
	head  *node
	level int
}

func Constructor() Skiplist {
	head := &node{
		next: make([]*node, maxLevel),
	}
	return Skiplist{
		head: head,
	}
}

func (s *Skiplist) Search(target int) bool {
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < target {
			cur = cur.next[i]
		}
	}
	return cur.next[0] != nil && cur.next[0].val == target
}

func (s *Skiplist) Add(num int) {
	// update contains a list of nodes that come prior to the new node to insert
	update := make([]*node, maxLevel)
	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}
	level := randomLevel()
	if level > s.level {
		for i := s.level; i < level; i++ {
			update[i] = s.head
		}
		s.level = level
	}
	newNode := &node{
		val:  num,
		next: make([]*node, level),
	}
	for i := 0; i < level; i++ {
		newNode.next[i] = update[i].next[i]
		update[i].next[i] = newNode
	}
}

func randomLevel() int {
	level := 1
	for rand.Intn(2) == 1 && level < maxLevel {
		level++
	}
	return level
}

func (s *Skiplist) Erase(num int) bool {
	update := make([]*node, maxLevel)

	cur := s.head
	for i := s.level - 1; i >= 0; i-- {
		for cur.next[i] != nil && cur.next[i].val < num {
			cur = cur.next[i]
		}
		update[i] = cur
	}

	if cur.next[0] == nil || cur.next[0].val != num {
		return false
	}

	// erase
	nodeToRemove := cur.next[0]
	for i := 0; i < s.level; i++ {
		if update[i].next[i] != nodeToRemove {
			// all levels with this node have already been removed
			break
		}
		update[i].next[i] = nodeToRemove.next[i]
	}
	// prune empty levels (if any)
	for s.level > 1 && s.head.next[s.level-1] == nil {
		s.level--
	}
	return true
}
