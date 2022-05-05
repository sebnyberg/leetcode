package p1902depthofbstgiveninsertionorder2

import (
	"container/heap"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxDepthBST(t *testing.T) {
	for _, tc := range []struct {
		order []int
		want  int
	}{
		{[]int{1}, 1},
		{[]int{2, 1, 4, 3}, 3},
		{[]int{2, 1, 3, 4}, 3},
		{[]int{1, 2, 3, 4}, 4},
		// {[]int{8, 14, 6, 10, 11, 13, 3, 2, 9, 1, 5, 12, 4, 7}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.order), func(t *testing.T) {
			require.Equal(t, tc.want, maxDepthBST(tc.order))
		})
	}
}

type treeNode struct {
	left, right *treeNode
	val         int
	depth       int
}

// func (n *treeNode) insert(x, depth int) int {
// 	if x < n.val {
// 		if n.left != nil {
// 			return n.left.insert(x, depth+1)
// 		} else {
// 			n.left = &treeNode{val: x}
// 			return depth + 1
// 		}
// 	} else {
// 		if n.right != nil {
// 			return n.right.insert(x, depth+1)
// 		} else {
// 			n.right = &treeNode{val: x}
// 			return depth + 1
// 		}
// 	}
// }

func maxDepth(n *treeNode) int {
	if n == nil {
		return 0
	}
	return 1 + max(maxDepth(n.left), maxDepth(n.right))
}

func maxDepthBST(order []int) int {
	n := len(order)

	// The idea is to insert steadily increasing or steadily decreasing
	// subsequences from order such that the values stay within a pre-determined
	// interval.

	// Insert steadily increasing / decreasing numbers into the BST before
	// doing a regular insert.
	root := &treeNode{val: order[0], depth: 1}

	todo := make(nodeHeap, 0)
	heap.Push(&todo, &interval{dirRight, root, 0, 10001, order[0], 10001})
	heap.Push(&todo, &interval{dirLeft, root, 0, 10001, 0, order[0]})

	seen := make([]bool, n)
	seen[0] = true
	// maxDepth := 1

	for len(todo) > 0 {
		next := heap.Pop(&todo).(*interval)
		curNode := next.root
		if next.dir == dirRight {
			// Add increasing integers from idx and forward which have not yet
			// been seen and fall within the interval. Also add all new intervals
			// to heap.
			prev := next.lo
			for i := next.idx + 1; i < n; i++ {
				if !seen[i] && order[i] > prev && order[i] < next.hi {
					curNode.right = &treeNode{val: order[i], depth: curNode.depth + 1}
					// maxDepth = max(maxDepth, curNode.depth+1)
					if order[i]-prev > 1 {
						heap.Push(&todo, &interval{dirLeft, curNode.right, i, order[i] - prev, prev, order[i]})
					}
					curNode = curNode.right
					prev = order[i]
					seen[i] = true
				}
			}
		} else { // next.dir == dirLeft
			prev := next.hi
			for i := next.idx + 1; i < n; i++ {
				if !seen[i] && order[i] < prev && order[i] > next.lo {
					curNode.left = &treeNode{val: order[i], depth: curNode.depth + 1}
					// maxDepth = max(maxDepth, curNode.depth+1)
					if prev-order[i] > 1 {
						heap.Push(&todo, &interval{dirRight, curNode.left, i, prev - order[i], order[i], prev})
					}
					curNode = curNode.left
					prev = order[i]
					seen[i] = true
				}
			}
		}
	}
	return maxDepth(root)

	// // maxDepth := max(depthRight, depthLeft)

	// // // Finally, do a regular BST insert of remaining values
	// // for i, num := range order {
	// // 	if seen[i] {
	// // 		continue
	// // 	}
	// // 	maxDepth = max(maxDepth, root.insert(num, 1))
	// // }

	// return 0
}

type direction uint8

const (
	dirLeft  direction = 0
	dirRight direction = 1
)

type interval struct {
	dir    direction
	root   *treeNode
	idx    int
	d      int
	lo, hi int
}

type nodeHeap []*interval

func (h nodeHeap) Len() int { return len(h) }
func (h nodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h nodeHeap) Less(i, j int) bool {
	return h[i].idx < h[j].idx
}
func (h *nodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*interval))
}
func (h *nodeHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
