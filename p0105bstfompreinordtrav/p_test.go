package p0105bstfompreinordtrav

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) != len(inorder) {
		return nil
	}
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		preorder[0],
	}
}
