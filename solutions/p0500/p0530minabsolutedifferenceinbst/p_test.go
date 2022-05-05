package p0530minabsolutedifferenceinbst

import (
	"container/list"
	"fmt"
	"log"
	"math"
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

func (n *TreeNode) FindVal(val int) *TreeNode {
	if n == nil {
		return nil
	}
	if n.Val == val {
		return n
	}
	if l := n.Left.FindVal(val); l != nil {
		return l
	}
	return n.Right.FindVal(val)
}

func (n *TreeNode) Equals(other *TreeNode) bool {
	if n == nil || other == nil {
		return n == other
	}
	if n.Val != other.Val {
		return false
	}
	return n.Left.Equals(other.Left) && n.Right.Equals(other.Right)
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

func Test_getMinimumDifference(t *testing.T) {
	for _, tc := range []struct {
		root string
		want int
	}{
		{"[4,2,6,1,3]", 1},
		{"[1,0,48,null,null,12,49]", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			root := ParseTree(tc.root)
			require.Equal(t, tc.want, getMinimumDifference(root))
		})
	}
}

func getMinimumDifference(root *TreeNode) int {
	minDiff := math.MaxInt32
	// When you go right, set min value,
	// when you go left, set max value
	findMin(root, math.MinInt32, math.MaxInt32, &minDiff)
	return minDiff
}

func findMin(cur *TreeNode, min, max int, minDiff *int) {
	if cur == nil {
		return
	}
	if d := cur.Val - min; d < *minDiff {
		*minDiff = d
	}
	if d := max - cur.Val; d < *minDiff {
		*minDiff = d
	}
	findMin(cur.Left, min, cur.Val, minDiff)
	findMin(cur.Right, cur.Val, max, minDiff)
}
