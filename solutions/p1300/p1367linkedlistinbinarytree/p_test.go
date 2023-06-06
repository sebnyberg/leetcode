package p1367linkedlistinbinarytree

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return f(head, root) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func f(curr *ListNode, root *TreeNode) bool {
	if curr == nil {
		return true
	}
	if root == nil {
		return false
	}
	if root.Val != curr.Val {
		return false
	}
	return f(curr.Next, root.Left) || f(curr.Next, root.Right)
}
