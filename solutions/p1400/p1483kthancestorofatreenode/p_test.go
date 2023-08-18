package p1483kthancestorofatreenode

import (
	"fmt"
	"testing"
)

func Test_TreeAncestor(t *testing.T) {
	for i, _ := range []struct {
		n      int
		parent []int
	}{
		{11, []int{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		// {7, []int{-1, 0, 0, 1, 1, 2, 2}},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			// ta := Constructor(tc.n, tc.parent)
			// a := ta.GetKthAncestor(8, 1)
			// a = ta.GetKthAncestor(8, 8)
			// _ = a
		})
	}
}

// TreeAncestor is a jump-tree specialized in quickly finding the ancestor of a
// node. Each node has a jump-list with its 2^kth parent.
type TreeAncestor struct {
	parent []int
	jump   [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	children := make([][]int, n)
	for i := range parent {
		if parent[i] != -1 {
			children[parent[i]] = append(children[parent[i]], i)
		}
	}
	jump := make([][]int, n)
	curr := []int{0}
	next := []int{}
	for k := 0; len(curr) > 0; k++ {
		next = next[:0]
		for _, x := range curr {
			for _, j := range children[x] {
				// Add parent
				jump[j] = append(jump[j], x)
				next = append(next, j)

				// For each time the 2^kth parent has a 2^kth parent, add it as
				// the 2^(k+1)th parent of the current node, and continue
				for k := 1; ; k++ {
					ancestor := jump[jump[j][k-1]]
					if len(ancestor) >= k {
						jump[j] = append(jump[j], ancestor[k-1])
					} else {
						break
					}
				}
			}
		}
		curr, next = next, curr
	}

	return TreeAncestor{
		parent: parent,
		jump:   jump,
	}
}
func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	for k > 0 && node != 0 {
		// Find largest factor of 2 in k which is supported by the jump table
		// given by the node
		i := len(this.jump[node]) - 1
		for (1 << i) > k {
			i--
		}
		node = this.jump[node][i]
		k -= (1 << i)
	}
	if node == 0 && k > 0 {
		return -1
	}
	return node
}
