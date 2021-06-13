package tree

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

// var nilNode *TreeNode

// func NewTreeFromList(nodesList []int) *TreeNode {
// 	if len(nodesList) == 0 {
// 		return nil
// 	}
// 	root := &TreeNode{
// 		Val: nodesList[0],
// 	}
// 	levelNodes := list.New()
// 	levelNodes.PushBack(root)
// 	idx := 1
// 	for {
// 		for size := levelNodes.Len(); size > 0; size-- {
// 			if idx >= len(nodesList) {
// 				return root
// 			}
// 			node := levelNodes.Remove(levelNodes.Front()).(*TreeNode)
// 			if node == nil {
// 				levelNodes.PushBack(nilNode)
// 				levelNodes.PushBack(nilNode)
// 				idx += 2
// 				continue
// 			}
// 			// Left
// 			if nodesList[idx] != -1 {
// 				node.Left = &TreeNode{Val: nodesList[idx]}
// 				levelNodes.PushBack(node.Left)
// 			} else {
// 				levelNodes.PushBack(nilNode)
// 			}
// 			idx++
// 			if idx >= len(nodesList) {
// 				return root
// 			}
// 			// Right
// 			if nodesList[idx] != -1 {
// 				node.Right = &TreeNode{Val: nodesList[idx]}
// 				levelNodes.PushBack(node.Right)
// 			} else {
// 				levelNodes.PushBack(nilNode)
// 			}
// 			idx++
// 		}
// 	}
// }
