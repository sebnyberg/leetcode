package p1302deepestleavessum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	todos := []*TreeNode{root}
	var sum int
	for len(todos) > 0 {
		sum = 0
		newTodos := make([]*TreeNode, 0)
		for _, n := range todos {
			if n.Left != nil {
				newTodos = append(newTodos, n.Left)
			}
			if n.Right != nil {
				newTodos = append(newTodos, n.Right)
			}
			sum += n.Val
		}
		todos = newTodos
	}
	return sum
}
