package p1993operationsontree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLockingTree(t *testing.T) {
	tree := Constructor([]int{-1, 0, 0, 1, 1, 2, 2})
	// lockingTree.lock(2, 2);    // return true because node 2 is unlocked.
	// 													 // Node 2 will now be locked by user 2.
	res := tree.Lock(2, 2)
	require.Equal(t, true, res)
	// lockingTree.unlock(2, 3);  // return false because user 3 cannot unlock a node locked by user 2.
	res = tree.Unlock(2, 3)
	require.Equal(t, false, res)
	// lockingTree.unlock(2, 2);  // return true because node 2 was previously locked by user 2.
	// 													 // Node 2 will now be unlocked.
	res = tree.Unlock(2, 2)
	require.Equal(t, true, res)
	// lockingTree.lock(4, 5);    // return true because node 4 is unlocked.
	// 													 // Node 4 will now be locked by user 5.
	res = tree.Lock(4, 5)
	require.Equal(t, true, res)
	// lockingTree.upgrade(0, 1); // return true because node 0 is unlocked and has at least one locked descendant (node 4).
	// 													 // Node 0 will now be locked by user 1 and node 4 will now be unlocked.
	res = tree.Upgrade(0, 1)
	require.Equal(t, true, res)
	// lockingTree.lock(0, 1);    // return false because node 0 is already locked.
	res = tree.Lock(0, 1)
	require.Equal(t, false, res)
}

type LockingTree struct {
	// lock  map[int]int
	nodes   []*LockTreeNode
	bfsCur  []*LockTreeNode
	bfsNext []*LockTreeNode
}

type LockTreeNode struct {
	id           int
	children     []*LockTreeNode
	parent       *LockTreeNode
	lockedByUser int
}

func Constructor(parent []int) LockingTree {
	tree := LockingTree{
		// lock:  make(map[int]int),
		bfsCur:  make([]*LockTreeNode, len(parent)),
		bfsNext: make([]*LockTreeNode, len(parent)),
		nodes:   make([]*LockTreeNode, len(parent)),
	}
	for i := range tree.nodes {
		tree.nodes[i] = &LockTreeNode{
			id:       i,
			children: []*LockTreeNode{},
			parent:   nil,
		}
	}
	for nodeIdx, parentIdx := range parent {
		if parentIdx != -1 {
			tree.nodes[nodeIdx].parent = tree.nodes[parentIdx]
			tree.nodes[parentIdx].children = append(tree.nodes[parentIdx].children,
				tree.nodes[nodeIdx],
			)
		}
	}
	return tree
}

func (this *LockingTree) Lock(num int, user int) bool {
	node := this.nodes[num]
	if node.lockedByUser == 0 {
		node.lockedByUser = user
		return true
	}
	return false
}

func (this *LockingTree) Unlock(num int, user int) bool {
	node := this.nodes[num]
	if node.lockedByUser == user {
		node.lockedByUser = 0
		return true
	}
	return false
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	node := this.nodes[num]
	// 1. Node must be unlocked (will become locked)
	if node.lockedByUser != 0 {
		return false
	}
	// 2. It has at least one locked descendant (by any user)
	// These descendants will become unlocked
	lockedDescendants := make([]int, 0)
	this.bfsCur = this.bfsCur[:0]
	this.bfsNext = this.bfsNext[:0]
	this.bfsCur = append(this.bfsCur, node.children...)
	for len(this.bfsCur) > 0 {
		this.bfsNext = this.bfsNext[:0]
		for _, child := range this.bfsCur {
			this.bfsNext = append(this.bfsNext, child.children...)
			if child.lockedByUser > 0 {
				lockedDescendants = append(lockedDescendants, child.id)
			}
		}
		this.bfsNext, this.bfsCur = this.bfsCur, this.bfsNext
	}
	if len(lockedDescendants) == 0 {
		return false
	}

	// No ancestor may be lockedj
	cur := node.parent
	for cur != nil {
		if cur.lockedByUser > 0 {
			return false
		}
		cur = cur.parent
	}

	// Unlock all locked descendants
	for _, desc := range lockedDescendants {
		this.nodes[desc].lockedByUser = 0
	}
	// Lock current node
	node.lockedByUser = user

	return true
}
