package p0508mostfrequentsubtreesum

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

func Test_findFrequentTreeSum(t *testing.T) {
	for _, tc := range []struct {
		root string
		want []int
	}{
		{"[5,2,-3]", []int{2, -3, 4}},
		{"[5,2,-5]", []int{2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root), func(t *testing.T) {
			tree := ParseTree(tc.root)
			require.ElementsMatch(t, tc.want, findFrequentTreeSum(tree))

		})
	}
}

func findFrequentTreeSum(root *TreeNode) []int {
	sums := make(map[int]int)
	visit(sums, root)
	var maxCount int
	var res []int
	for sum, count := range sums {
		if count > maxCount {
			res = res[:0]
			res = append(res, sum)
			maxCount = count
		} else if count == maxCount {
			res = append(res, sum)
		}
	}
	return res
}

func visit(sums map[int]int, cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	sum := cur.Val + visit(sums, cur.Left) + visit(sums, cur.Right)
	sums[sum]++
	return sum
}
