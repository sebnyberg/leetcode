package p1120maxaveragesubtree

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

func Test_maximumAverageSubtree(t *testing.T) {
	for _, tc := range []struct {
		root string
		want float64
	}{
		// {"[5,6,1]", 6},
		{"[0,null,1]", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.InEpsilon(t, tc.want, maximumAverageSubtree(tree), 1e-5)
		})
	}
}

func maximumAverageSubtree(root *TreeNode) float64 {
	// Starting with leaf nodes, find the max average by returning the sum
	// and counts of each subtree
	var f maxAvgFinder
	f.visit(root)
	return f.maxAvg
}

type maxAvgFinder struct {
	maxAvg float64
}

func (f *maxAvgFinder) visit(cur *TreeNode) (sum int, count int) {
	var leftSum, leftCount, rightSum, rightCount int
	if cur.Left != nil {
		leftSum, leftCount = f.visit(cur.Left)
	}
	if cur.Right != nil {
		rightSum, rightCount = f.visit(cur.Right)
	}
	sum = leftSum + rightSum + cur.Val
	count = leftCount + rightCount + 1
	f.maxAvg = math.Max(f.maxAvg, float64(sum)/float64(count))
	return sum, count
}
