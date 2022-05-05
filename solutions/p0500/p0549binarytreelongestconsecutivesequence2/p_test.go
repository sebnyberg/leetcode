package p0549binarytreelongestconsecutivesequence2

import (
	"container/list"
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ParseTree(input string) *TreeNode {
	// Trim start/end []
	input = input[1 : len(input)-1]

	// Split by comma
	inputParts := strings.Split(input, ",")
	n := len(inputParts)

	if n == 0 || inputParts[0] == "" {
		return nil
	}

	// Create one node per element in the array
	nodes := make([]*TreeNode, n)
	for i, part := range inputParts {
		if part != "null" {
			val, err := strconv.Atoi(part)
			if err != nil {
				log.Fatalln(err)
			}
			nodes[i] = &TreeNode{Val: val}
		}
	}

	q := list.New()
	q.PushBack(nodes[0])

	i := 1
	for q.Len() > 0 && i < n {
		el := q.Remove(q.Front()).(*TreeNode)
		if nodes[i] != nil {
			el.Left = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
		if i >= n {
			break
		}
		if nodes[i] != nil {
			el.Right = nodes[i]
			q.PushBack(nodes[i])
		}
		i++
	}

	return nodes[0]
}

func Test_longestConsecutive(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[1,2,3]", 2},
		{"[2,1,3]", 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.Equal(t, tc.want, longestConsecutive(tree))
		})
	}
}

func longestConsecutive(root *TreeNode) int {
	var f pathFinder
	f.visitArm(root)
	f.visitArm(root)
	return f.maxLen
}

type pathFinder struct {
	maxLen int
}

// visitArm returns the longest ascending and descending arm length starting in
// the current node. It also updates the pathFinder maxLen based on the arm
// and arrow length given the current node.
func (f *pathFinder) visitArm(cur *TreeNode) (int, int) {
	if cur == nil {
		return 0, 0
	}

	leftAscLen, leftDescLen := f.visitArm(cur.Left)
	rightAscLen, rightDescLen := f.visitArm(cur.Right)

	// If there is a descending arm starting in this node on the left side
	maxDescLen, maxAscLen := 1, 1
	leftDescending := cur.Left != nil && cur.Left.Val == cur.Val-1
	leftAscending := cur.Left != nil && cur.Left.Val == cur.Val+1
	rightDescending := cur.Right != nil && cur.Right.Val == cur.Val-1
	rightAscending := cur.Right != nil && cur.Right.Val == cur.Val+1
	if leftDescending {
		maxDescLen = max(maxDescLen, 1+leftDescLen)
		if rightAscending {
			f.maxLen = max(f.maxLen, 1+leftDescLen+rightAscLen)
		}
	}
	if rightDescending {
		maxDescLen = max(maxDescLen, 1+rightDescLen)
		if leftAscending {
			f.maxLen = max(f.maxLen, 1+leftAscLen+rightDescLen)
		}
	}
	if rightAscending {
		maxAscLen = max(maxAscLen, 1+rightAscLen)
	}
	if leftAscending {
		maxAscLen = max(maxAscLen, 1+leftAscLen)
	}
	f.maxLen = max(f.maxLen, max(maxAscLen, maxDescLen))
	return maxAscLen, maxDescLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
