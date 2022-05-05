package p0170intersectlinkedlists

import "testing"

func Test_getIntersectionNode(t *testing.T) {
	headA := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 6,
						},
					},
				},
			},
		},
	}

	headB := &ListNode{
		Val: 11,
		Next: &ListNode{
			Val:  12,
			Next: headA.Next.Next.Next.Next,
		},
	}

	res := getIntersectionNode(headA, headB)
	_ = res
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	var na int
	for curA != nil {
		curA = curA.Next
		na++
	}

	var nb int
	for curB != nil {
		curB = curB.Next
		nb++
	}

	if curA != curB {
		return nil
	}

	curA, curB = headA, headB
	for i := 0; i < na-nb; i++ {
		curA = curA.Next
	}
	for i := 0; i < nb-na; i++ {
		curB = curB.Next
	}

	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}

	return curA
}
