package p0160intersectionoftwolinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	// The idea is to invert edges in the linked list.
	// Then the problem is a matter of finding the last node in common between the
	// two, if any.

	// First, lets try a hacky solution - pointers are comparable, so we can
	// put them in a map and check for a match.
	m := make(map[*ListNode]struct{}, 1000)
	cur := headA
	for cur != nil {
		m[cur] = struct{}{}
		cur = cur.Next
	}

	cur = headB
	for cur != nil {
		if _, exists := m[cur]; exists {
			return cur
		}
		cur = cur.Next
	}
	return nil
}
