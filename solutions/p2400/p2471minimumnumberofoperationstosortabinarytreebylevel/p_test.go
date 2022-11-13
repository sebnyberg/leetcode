package p2471minimumnumberofoperationstosortabinarytreebylevel

import (
	"container/list"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Test_minimumOperations(t *testing.T) {
	for i, tc := range []struct {
		root string
		want int
	}{
		{"[1,4,3,7,6,8,5,null,null,null,null,9,null,10]", 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			tr := ParseTree(tc.root)
			res := minimumOperations(tr)
			require.Equal(t, tc.want, res)
		})
	}
}

type sorter struct {
	indices []int
	vals    []int
}

func (s sorter) Len() int {
	return len(s.indices)
}

func (s sorter) Less(i, j int) bool {
	return s.vals[i] < s.vals[j]
}

func (s *sorter) Swap(i, j int) {
	s.indices[i], s.indices[j] = s.indices[j], s.indices[i]
	s.vals[i], s.vals[j] = s.vals[j], s.vals[i]
}

func (s *sorter) Set(nodes []*TreeNode) {
	s.indices = s.indices[:0]
	s.vals = s.vals[:0]
	for i := range nodes {
		s.indices = append(s.indices, i)
		s.vals = append(s.vals, nodes[i].Val)
	}
}

func minimumOperations(root *TreeNode) int {
	curr := []*TreeNode{root}
	next := []*TreeNode{}
	var res int
	var s sorter
	for len(curr) > 0 {
		next = next[:0]
		s.Set(curr)
		sort.Sort(&s)
		for i := range s.indices {
			for s.indices[i] != i {
				j := s.indices[i]
				s.indices[i], s.indices[j] = s.indices[j], s.indices[i]
				res++
			}
		}
		for _, x := range curr {
			if x.Left != nil {
				next = append(next, x.Left)
			}
			if x.Right != nil {
				next = append(next, x.Right)
			}
		}
		curr, next = next, curr
	}
	return res
}
