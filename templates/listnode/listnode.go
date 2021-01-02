package listnode

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	root := &ListNode{
		Val: values[0],
	}
	prev := root
	for i := 1; i < len(values); i++ {
		prev.Next = &ListNode{
			Val: values[i],
		}
		prev = prev.Next
	}
	return root
}

func (l *ListNode) String() string {
	if l == nil {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")
	cur := l
	for cur.Next != nil {
		sb.WriteString(strconv.Itoa(cur.Val) + ",")
		cur = cur.Next
	}
	sb.WriteString(strconv.Itoa(cur.Val))
	sb.WriteString("]")
	return sb.String()
}
