package p0111mindepthbintree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nodes := []*TreeNode{root}
	level := 1
	for len(nodes) > 0 {
		newNodes := make([]*TreeNode, 0)
		for _, node := range nodes {
			if node.Left == nil && node.Right == nil {
				return level
			}
			if node.Left != nil {
				newNodes = append(newNodes, node.Left)
			}
			if node.Right != nil {
				newNodes = append(newNodes, node.Right)
			}
		}
		level++
		nodes = newNodes
	}
	return level
}
