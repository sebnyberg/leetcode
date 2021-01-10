package tree

// import (
// 	"container/list"
// 	"log"
// )

// type IntTreeNode struct {
// 	Val   int
// 	Left  *IntTreeNode
// 	Right *IntTreeNode
// }

// var nilIntNode *IntTreeNode

// func NewIntTreeFromList(nodes []int) *IntTreeNode {
// 	if len(nodes) == 0 {
// 		log.Fatalln("more than one node is required to parse a tree")
// 	}
// 	root := &IntTreeNode{Val: nodes[0]}
// 	q := list.New()
// 	n := len(nodes)
// 	q.PushBack(root)
// 	idx := 1

// 	for {
// 		if q.Len() == 0 {
// 			return root
// 		}
// 		for size := q.Len(); size > 0; size-- {
// 			el := q.Remove(q.Front()).(*IntTreeNode)
// 			if el == nilIntNode {
// 				idx += 2
// 				continue
// 			}
// 			// Left side
// 			if idx >= n {
// 				return root
// 			}
// 			if nodes[idx] == -1 {
// 				q.PushBack(nilNode)
// 			} else {
// 				el.Left = &IntTreeNode{Val: nodes[idx]}
// 				q.PushBack(el.Left)
// 			}
// 			idx++

// 			// Left side
// 			if idx >= n {
// 				return root
// 			}
// 			if nodes[idx] == -1 {
// 				q.PushBack(nilNode)
// 			} else {
// 				el.Right = &IntTreeNode{Val: nodes[idx]}
// 				q.PushBack(el.Left)
// 			}
// 			idx++
// 		}
// 	}
// }
