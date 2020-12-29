package d29_test

import (
	"container/list"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nilNode *TreeNode

func parseTree(nodes []int) *TreeNode {
	if len(nodes) == 0 {
		log.Fatalln("more than one node is required to parse a tree")
	}
	root := &TreeNode{Val: nodes[0]}
	q := list.New()
	n := len(nodes)
	q.PushBack(root)
	idx := 1

	for {
		if q.Len() == 0 {
			return root
		}
		for size := q.Len(); size > 0; size-- {
			el := q.Remove(q.Front()).(*TreeNode)
			if el == nilNode {
				idx += 2
				continue
			}
			// Left side
			if idx >= n {
				return root
			}
			if nodes[idx] == -1 {
				q.PushBack(nilNode)
			} else {
				el.Left = &TreeNode{Val: nodes[idx]}
				q.PushBack(el.Left)
			}
			idx++

			// Right side
			if idx >= n {
				return root
			}
			if nodes[idx] == -1 {
				q.PushBack(nilNode)
			} else {
				el.Right = &TreeNode{Val: nodes[idx]}
				q.PushBack(el.Right)
			}
			idx++
		}
	}
}

func Test_pseudoPalindromicPaths(t *testing.T) {
	for _, tc := range []struct {
		nodes []int
		want  int
	}{
		{[]int{2, 3, 1, 3, 1, -1, 1}, 2},
		{[]int{2, 1, 1, 1, 3, -1, -1, -1, -1, -1, 1}, 1},
		{[]int{9}, 1},
	} {
		t.Run(fmt.Sprintf("%v", tc.nodes), func(t *testing.T) {
			require.Equal(t, tc.want, pseudoPalindromicPaths(parseTree(tc.nodes)))
		})
	}
}

type nodeStack []int

func (s *nodeStack) Pop() int {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}

func (s *nodeStack) Push(x int) {
	old := *s
	old = append(old, x)
	*s = old
}

type palindromeCounter struct {
	odd   [10]bool
	count int
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pseudoPalindromicPaths(root *TreeNode) int {
	// Perform DFS of the tree, adding numbers to the stack
	explorer := palindromeCounter{}
	explorer.Explore(root)
	return explorer.count
}

// Explore traverses the tree until neither left or right node is found
func (e *palindromeCounter) Explore(node *TreeNode) {
	e.odd[node.Val] = !e.odd[node.Val]

	defer func(val int) {
		e.odd[val] = !e.odd[val]
	}(node.Val)

	// If at a leaf and the path is a palindrome, increase the counter
	if node.Left == nil && node.Right == nil {
		var foundodd bool
		for _, isodd := range e.odd {
			if isodd {
				if foundodd {
					return
				}
				foundodd = true
			}
		}
		e.count++
	}

	if node.Left != nil {
		e.Explore(node.Left)
	}
	if node.Right != nil {
		e.Explore(node.Right)
	}
}
