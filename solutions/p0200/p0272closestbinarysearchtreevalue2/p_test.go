package p0272closestbinarysearchtreevalue2

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClosestKValues(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 5},
	}

	res := closestKValues(tree, 3.714286, 2)
	require.Equal(t, []int{4, 3}, res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestKValues(root *TreeNode, target float64, k int) []int {
	// Perform inorder traversal of the tree until either there are no more nodes,
	// or the current value is more than k values below the target
	s := &search{maxDist: make(MaxHeap, 0)}
	s.Search(root, target, k)
	res := make([]int, k)
	for i, v := range s.maxDist {
		res[i] = v.nodeVal
	}
	return res
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

type search struct {
	maxDist MaxHeap
}

func (s *search) Search(cur *TreeNode, target float64, k int) {
	if cur == nil {
		return
	}
	d := abs(float64(cur.Val) - target)
	if len(s.maxDist) < k {
		heap.Push(&s.maxDist, nodeDist{d, cur.Val})
	} else {
		if d < s.maxDist[0].d {
			heap.Pop(&s.maxDist)
			heap.Push(&s.maxDist, nodeDist{d, cur.Val})
		}
	}
	s.Search(cur.Left, target, k)
	s.Search(cur.Right, target, k)
}

type nodeDist struct {
	d       float64
	nodeVal int
}

type MaxHeap []nodeDist

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].d > h[j].d
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(nodeDist))
}
func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
