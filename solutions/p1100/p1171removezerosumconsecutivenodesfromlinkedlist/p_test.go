package p1171removezerosumconsecutivenodesfromlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	// Keep a running sum and check whether that sum can be removed (sum up to
	// zero). Then, splice the linked list and reboot the algo.
	m := make(map[int]*ListNode)
	dummy := &ListNode{
		Next: head,
	}
	var sum int
outer:
	for {
		for k := range m {
			delete(m, k)
		}
		m[0] = dummy
		sum = 0
		curr := dummy.Next
		for curr != nil {
			sum += curr.Val
			if n, exists := m[sum]; exists {
				n.Next = curr.Next
				goto outer
			}
			m[sum] = curr
			curr = curr.Next
		}
		break
	}
	return dummy.Next
}
