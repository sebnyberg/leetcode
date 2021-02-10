package p0138copylistrandom

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func Test_copyRandomList(t *testing.T) {
	nodes := []*Node{
		{Val: 7},
		{Val: 13},
		{Val: 11},
		{Val: 10},
		{Val: 1},
	}
	for i := 0; i < len(nodes)-2; i++ {
		nodes[i].Next = nodes[i+1]
	}
	nodes[0].Random = nil
	nodes[1].Random = nodes[0]
	nodes[2].Random = nodes[4]
	nodes[3].Random = nodes[2]
	nodes[4].Random = nodes[0]
	res := copyRandomList(nodes[0])
	printList(nodes[0])
	printList(res)
	require.NotNil(t, res)
}

func printList(head *Node) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// Pass 1. Tie together the original list with a copy with
	// "smart pointers"
	// 1. new = copy(cur) -> store first copy as newHead
	// 2. new.random = cur
	// 3. new.next = cur.next
	// 4. cur.next = new
	// 5. cur = new.next
	// 6. repeat
	var nextHead *Node
	cur := head
	for cur != nil {
		newNode := &Node{
			Val:    cur.Val,
			Next:   cur.Next,
			Random: cur,
		}
		if nextHead == nil {
			nextHead = newNode
		}
		cur.Next = newNode
		cur = newNode.Next
	}

	// Pass 2 - fix random pointers
	// 1. Consider the current and new (cur.next)
	// 2. cur.next.random = cur.random.next
	// 3. cur = cur.Next.Next
	cur = head
	for cur != nil {
		newNode := cur.Next
		if cur.Random != nil {
			newNode.Random = cur.Random.Next
		} else {
			newNode.Random = nil
		}
		cur = newNode.Next
	}

	// Pass 3
	cur = head
	for cur != nil {
		if cur.Next.Next == nil {
			cur.Next = nil
			break
		}
		cur.Next, cur.Next.Next = cur.Next.Next, cur.Next.Next.Next
		cur = cur.Next
	}
	return nextHead
}
