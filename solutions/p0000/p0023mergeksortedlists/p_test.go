package p0023mergeksortedlists

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_mergeKLists(t *testing.T) {
	for _, tc := range []struct {
		lists [][]int
		want  []int
	}{
		{[][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{[][]int{}, []int{}},
		{[][]int{{}}, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.lists), func(t *testing.T) {
			lists := make([]*ListNode, len(tc.lists))
			for i := range lists {
				lists[i] = createList(tc.lists[i])
			}
			require.Equal(t, createList(tc.want).String(), mergeKLists(lists).String())
		})
	}
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

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Todo: There is a more efficient version using divide and conquer
func mergeKLists(lists []*ListNode) *ListNode {
	h := &NodeHeap{}
	for _, l := range lists {
		if l != nil {
			*h = append(*h, l)
		}
	}
	heap.Init(h)
	if h.Len() == 0 {
		return nil
	}
	if h.Len() == 1 {
		return heap.Pop(h).(*ListNode)
	}

	// Two or more lists, pop two smallest nodes
	cur := heap.Pop(h).(*ListNode)
	other := heap.Pop(h).(*ListNode)
	head := cur
	for {
		for cur.Next != nil && cur.Next.Val <= other.Val {
			cur = cur.Next
		}
		// Switch to other
		if cur.Next != nil {
			heap.Push(h, cur.Next)
		}
		if h.Len() > 0 {
			cur, cur.Next, other = other, other, heap.Pop(h).(*ListNode)
			continue
		}
		// h is empty, make the last link and exit
		cur.Next = other
		return head
	}
}
