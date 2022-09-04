package p0987verticalordertrav

import (
	"container/list"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_verticalTraversal(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want [][]int
	}{
		// {[]int{0, 2, 1, 3, -1, -1, -1, 4, 5, -1, 7, 6, -1, 10, 8, 11, 9}, [][]int{{4, 10, 11}, {3, 7, 6}, {2, 5, 8, 9}, {0}, {1}}},
		{[]int{3, 9, 20, -1, -1, 15, 7}, [][]int{{9}, {3, 15}, {20}, {7}}},
	} {
		t.Run(fmt.Sprintf("%v", tc.in), func(t *testing.T) {
			root := NewTreeFromList(tc.in)
			t.Log(root)
			got := verticalTraversal(root)
			require.Equal(t, tc.want, got)
		})
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nilNode *TreeNode

func NewTreeFromList(nodesList []int) *TreeNode {
	if len(nodesList) == 0 {
		return nil
	}
	root := &TreeNode{
		Val: nodesList[0],
	}
	levelNodes := list.New()
	levelNodes.PushBack(root)
	idx := 1
	for {
		for size := levelNodes.Len(); size > 0; size-- {
			if idx >= len(nodesList) {
				return root
			}
			node := levelNodes.Remove(levelNodes.Front()).(*TreeNode)
			if node == nil {
				levelNodes.PushBack(nilNode)
				levelNodes.PushBack(nilNode)
				idx += 2
				continue
			}
			// Left
			if nodesList[idx] != -1 {
				node.Left = &TreeNode{Val: nodesList[idx]}
				levelNodes.PushBack(node.Left)
			} else {
				levelNodes.PushBack(nilNode)
			}
			idx++
			if idx >= len(nodesList) {
				return root
			}
			// Right
			if nodesList[idx] != -1 {
				node.Right = &TreeNode{Val: nodesList[idx]}
				levelNodes.PushBack(node.Right)
			} else {
				levelNodes.PushBack(nilNode)
			}
			idx++
		}
	}
}

type nodepos struct {
	i   int
	j   int
	val int
}

func verticalTraversal(root *TreeNode) [][]int {
	l := make([]nodepos, 0)
	fill(root, 0, 0, &l)
	sort.Slice(l, func(i, j int) bool {
		if l[i].j == l[j].j {
			if l[i].i == l[j].i {
				return l[i].val < l[j].val
			}
			return l[i].i > l[j].i
		}
		return l[i].j < l[j].j
	})
	res := [][]int{{l[0].val}}
	j := 0
	for i := 1; i < len(l); i++ {
		if l[i].j != l[i-1].j {
			j++
			res = append(res, make([]int, 0))
		}
		res[j] = append(res[j], l[i].val)
	}
	return res
}

func fill(n *TreeNode, x int, y int, l *[]nodepos) {
	*l = append(*l, nodepos{x, y, n.Val})
	if n.Left != nil {
		fill(n.Left, x-1, y-1, l)
	}
	if n.Right != nil {
		fill(n.Right, x+1, y-1, l)
	}
}
