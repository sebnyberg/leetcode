package p0369plusonelinkedlist

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func ParseList(s string) *ListNode {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")
	items := strings.Split(s, ",")
	dummy := &ListNode{}
	cur := dummy
	for _, item := range items {
		num, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("failed to convert string %v to integer", num)
		}
		cur.Next = &ListNode{
			Val: num,
		}
		cur = cur.Next
	}
	return dummy.Next
}

func (n *ListNode) String() string {
	cur := n
	if cur == nil {
		return "NIL"
	}
	i := 0
	var sb strings.Builder
	for cur != nil {
		if i != 0 {
			sb.WriteString("->")
		}
		sb.WriteString(strconv.Itoa(cur.Val))
		i++
		cur = cur.Next
	}
	return sb.String()
}

func (n *ListNode) Equals(other *ListNode) bool {
	if n == nil {
		return other == nil
	}
	if other == nil {
		return false
	}
	return n.Val == other.Val && n.Next.Equals(other.Next)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_plusOne(t *testing.T) {
	for _, tc := range []struct {
		head string
		want string
	}{
		{"[1,2,3]", "[1,2,4]"},
		{"[0]", "[1]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			head := ParseList(tc.head)
			res := plusOne(head)
			want := ParseList(tc.want)
			if !res.Equals(want) {
				t.Fatalf("list mismatch\nwant:\t%v\ngot:\t%v", want, res)
			}
		})
	}
}

func plusOne(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	carry := addOne(head)
	if carry > 0 {
		dummy.Val = 1
		return dummy
	}
	return head
}

// addOne adds one to the provided node and returns the overflow carry.
func addOne(head *ListNode) int {
	if head == nil {
		return 1
	}
	carry := addOne(head.Next)
	head.Val += carry
	if head.Val >= 10 {
		head.Val = 0
		return 1
	}
	return 0
}
