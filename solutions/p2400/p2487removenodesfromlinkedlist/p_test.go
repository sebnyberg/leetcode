package p2487removenodesfromlinkedlist

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
		Val:  math.MaxInt32,
	}
	stack := []*ListNode{dummy}
	cur := head
	for cur != nil {
		for stack[len(stack)-1].Val < cur.Val {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, cur)
		cur = cur.Next
	}
	ret := stack[1]
	for i := 1; i < len(stack); i++ {
		if i < len(stack)-1 {
			stack[i].Next = stack[i+1]
		} else {
			stack[i].Next = nil
		}
	}
	return ret
}
