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
		{"[5,10,10,null,null,2,3]", true},
		{"[1,2,10,null,null,2,20]", false},
	} {
		t.Run(fmt.Sprintf("%+v", tc.tree), func(t *testing.T) {
			root := ParseTree(tc.tree)
			require.Equal(t, tc.want, checkEqualTree(root))
		})
	}
}

func checkEqualTree(root *TreeNode) bool {
	// Collect total sum O(n)/O(n)
	var sum int
	tovisit := []*TreeNode{root}
	for i := 0; i < len(tovisit); i++ {
		if tovisit[i] == nil {
			continue
		}
		sum += tovisit[i].Val
		tovisit = append(tovisit, tovisit[i].Left)
		tovisit = append(tovisit, tovisit[i].Right)
	}

	// Check if there is a branch for which total sum - branch == branch
	_, res := checkHasPartition(root, sum)
	return res
}

func checkHasPartition(cur *TreeNode, total int) (int, bool) {
	var leftSum int
	if cur.Left != nil {
		var leftHasPart bool
		leftSum, leftHasPart = checkHasPartition(cur.Left, total)
		if leftHasPart || total-leftSum == leftSum {
			return -1, true
		}
	}
	var rightSum int
	if cur.Right != nil {
		var rightHasPart bool
		rightSum, rightHasPart = checkHasPartition(cur.Right, total)
		if rightHasPart || total-rightSum == rightSum {
			return -1, true
		}
	}
	return leftSum + rightSum + cur.Val, false
}
