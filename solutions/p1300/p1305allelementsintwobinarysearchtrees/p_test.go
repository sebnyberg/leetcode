package p1305allelementsintwobinarysearchtrees

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

func Test_getAllElements(t *testing.T) {
	for _, tc := range []struct {
		root1 string
		root2 string
		want  []int
	}{
		{"[2,1,4]", "[1,0,3]", []int{0, 1, 1, 2, 3, 4}},
		{"[1,null,8]", "[8,1]", []int{1, 1, 8, 8}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.root1), func(t *testing.T) {
			t1 := ParseTree(tc.root1)
			t2 := ParseTree(tc.root2)
			require.Equal(t, tc.want, getAllElements(t1, t2))
		})
	}
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	// To make the exercise fun, let's use channels to emit the smallest number
	// from both trees, then merge them together in this function.
	first, second := make(chan int), make(chan int)
	go func() {
		emitInOrder(root1, first)
		close(first)
	}()
	go func() {
		emitInOrder(root2, second)
		close(second)
	}()
	firstOK, secondOK := true, true
	firstVal, secondVal := math.MinInt32, math.MinInt32
	res := make([]int, 0)
	for {
		// Load new value if needed
		if firstOK && firstVal == math.MinInt32 {
			firstVal, firstOK = <-first
		}
		if secondOK && secondVal == math.MinInt32 {
			secondVal, secondOK = <-second
		}
		if !firstOK && !secondOK {
			break
		}
		switch {
		case !firstOK:
			res = append(res, secondVal)
			secondVal = math.MinInt32
		case !secondOK:
			res = append(res, firstVal)
			firstVal = math.MinInt32
		case firstVal <= secondVal:
			res = append(res, firstVal)
			firstVal = math.MinInt32
		default:
			res = append(res, secondVal)
			secondVal = math.MinInt32
		}
	}
	return res
}

func emitInOrder(cur *TreeNode, ch chan<- int) {
	if cur == nil {
		return
	}
	emitInOrder(cur.Left, ch)
	ch <- cur.Val
	emitInOrder(cur.Right, ch)
}
