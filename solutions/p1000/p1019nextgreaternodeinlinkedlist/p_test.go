package p1019nextgreaternodeinlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	// Add values to a stack. This stack will be non-decreasing. While the
	// current value is larger than the value in the stack, pop from the stack
	// and add the result to its target position.
	targets := []int{}
	stack := []int{}
	res := []int{}
	curr := head
	var i int
	for curr != nil {
		for len(stack) > 0 && curr.Val > stack[len(stack)-1] {
			res[targets[len(targets)-1]] = curr.Val
			targets = targets[:len(targets)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, curr.Val)
		targets = append(targets, i)
		res = append(res, 0)
		i++
		curr = curr.Next
	}
	return res
}
