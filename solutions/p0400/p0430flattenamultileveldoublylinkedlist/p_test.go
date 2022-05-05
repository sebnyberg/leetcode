package p0430flattenamultileveldoublylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func makeList(nums []int) []*Node {
	res := make([]*Node, len(nums))
	res[0] = &Node{Val: nums[0]}
	for i := 1; i < len(nums); i++ {
		res[i] = &Node{Val: nums[i]}
		res[i-1].Next = res[i]
		res[i].Prev = res[i-1]
	}
	return res
}

func collectList(root *Node) []int {
	res := make([]int, 0)
	cur := root
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

func TestFlatten(t *testing.T) {
	first := makeList([]int{1, 2, 3, 4, 5, 6})
	second := makeList([]int{7, 8, 9, 10})
	third := makeList([]int{11, 12})
	first[2].Child = second[0]
	second[1].Child = third[0]
	res := flatten(first[0])
	require.Equal(t, collectList(res), []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6})

	// hehe
	a := res
	_ = a
}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	cur := root
	parents := make([]*Node, 0)
	for cur != nil {
		next := cur.Next
		if cur.Child != nil {
			parents = append(parents, next)
			next = cur.Child
			cur.Child = nil
		}
		for next == nil && len(parents) > 0 {
			next = parents[len(parents)-1]
			parents = parents[:len(parents)-1]
		}
		if next == nil {
			break
		}
		cur.Next, next.Prev = next, cur
		cur = next
	}
	return root
}
