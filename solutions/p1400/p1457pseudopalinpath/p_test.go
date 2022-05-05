package p1457pseudopalinpath

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
		{[]int{2, 1, 1, 1, 3, -1, -1, -1, -1, -1, 1}, 1},
		{[]int{9}, 1},
	} {
		t.Run(fmt.Sprintf("%v", tc.nodes), func(t *testing.T) {
			tree := parseTree(tc.nodes)
			require.Equal(t, tc.want, pseudoPalindromicPaths(tree))
		})
	}
}

func init() {
	for i := 1; i < 10; i++ {
		bitMasks[i] = 1 << (i - 1)
	}
}

var bitMasks [10]int16

func pseudoPalindromicPaths(root *TreeNode) int {
	return explore(root, 0)
}

func explore(node *TreeNode, oddBits int16) int {
	if node == nil {
		return 0
	}
	if oddBits&bitMasks[node.Val] == 0 {
		oddBits |= bitMasks[node.Val]
	} else {
		oddBits -= bitMasks[node.Val]
	}

	if node.Left == nil && node.Right == nil {
		var foundodd bool
		for i := 1; i < 10; i++ {
			if oddBits&bitMasks[i] == bitMasks[i] {
				if foundodd {
					return 0
				}
				foundodd = true
			}
		}
		return 1
	}

	return explore(node.Left, oddBits) + explore(node.Right, oddBits)
}
