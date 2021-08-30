package p0663equaltreepartition

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

func Test_checkEqualTree(t *testing.T) {
	for _, tc := range []struct {
		tree string
		want bool
	}{
		{"[1,-1]", false},
		// {"[5,10,10,null,null,2,3]", true},
		// {"[1,2,10,null,null,2,20]", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			root := ParseTree(tc.tree)
			require.Equal(t, tc.want, checkEqualTree(root))
		})
	}
}

func checkEqualTree(root *TreeNode) bool {
	// O(n)/O(n) two-pass:
	var p partitioner
	// Calculate total sum
	p.sums = make(map[*TreeNode]int)
	p.collectSum(root)
	res := p.checkHasPartition(root, 0)
	return res
}

type partitioner struct {
	sums map[*TreeNode]int
	// hasSum bool
}

func (p *partitioner) collectSum(cur *TreeNode) int {
	if cur == nil {
		return 0
	}
	if _, exists := p.sums[cur]; !exists {
		leftSum := p.collectSum(cur.Left)
		rightSum := p.collectSum(cur.Right)
		p.sums[cur] = leftSum + rightSum + cur.Val
	}
	return p.sums[cur]
}

func (p *partitioner) checkHasPartition(cur *TreeNode, parentVal int) bool {
	if cur == nil {
		return false
	}
	leftSum := p.collectSum(cur.Left)
	rightSum := p.collectSum(cur.Right)
	if leftSum == parentVal-leftSum+cur.Val+rightSum ||
		rightSum == parentVal-rightSum+cur.Val+leftSum {
		return true
	}
	return p.checkHasPartition(cur.Left, leftSum+rightSum+cur.Val) ||
		p.checkHasPartition(cur.Right, leftSum+rightSum+cur.Val)
}
