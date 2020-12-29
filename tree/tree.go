package tree

import (
	"container/list"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

var nilNode *TreeNode

// NewTreeFromList parses a tree from a list of node values.
// Nodes with a value of nil and their descendants are not added to the tree.
// If the list is empty, the result is nil.
//
// Example: [1,2,3,4]
//
//          1
//        /   \
//       2     3
//     /
//    4
//
// Example: [1,nil,3,nil,nil,4]
//
//         1
//           \
//            3
//              \
//               4

func NewTreeFromList(nodesList []interface{}) *TreeNode {
	if len(nodesList) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nodesList[0],
	}
	levelNodes := list.New()
	levelNodes.PushBack(root)
	idx := 1
	for {
		for size := levelNodes.Len(); size > 0; size-- {
			if idx >= len(nodesList) {
				return root
			}
			node := levelNodes.Remove(levelNodes.Front()).(*TreeNode)
			if node == nil {
				levelNodes.PushBack(nilNode)
				levelNodes.PushBack(nilNode)
				idx += 2
				continue
			}
			// Left
			if nodesList[idx] != nil {
				node.Left = &TreeNode{Val: nodesList[idx]}
				levelNodes.PushBack(node.Left)
			} else {
				levelNodes.PushBack(nilNode)
			}
			idx++
			if idx >= len(nodesList) {
				return root
			}
			// Right
			if nodesList[idx] != nil {
				node.Right = &TreeNode{Val: nodesList[idx]}
				levelNodes.PushBack(node.Right)
			} else {
				levelNodes.PushBack(nilNode)
			}
			idx++
		}
	}
}

func (n *TreeNode) Equals(other *TreeNode) bool {
	if n == nil && other == nil {
		return true
	}
	if n == nil && other != nil {
		return false
	}
	if n != nil && other == nil {
		return false
	}
	return n.Val == other.Val && n.Left.Equals(other.Left) && n.Right.Equals(other.Right)
}
