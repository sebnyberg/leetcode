package p0637avglevelbintree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) (levelAvgs []float64) {
	if root == nil {
		return
	}
	levelNodes := []*TreeNode{root}
	for len(levelNodes) > 0 {
		n := len(levelNodes)
		var levelSum int
		var newLevelNodes []*TreeNode
		for _, node := range levelNodes {
			levelSum += node.Val
			for _, n := range []*TreeNode{node.Left, node.Right} {
				if n != nil {
					newLevelNodes = append(newLevelNodes, n)
				}
			}
		}
		levelAvgs = append(levelAvgs, float64(levelSum)/float64(n))
		levelNodes = newLevelNodes
	}
	return levelAvgs
}
