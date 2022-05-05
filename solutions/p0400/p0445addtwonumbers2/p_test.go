package p0445addtwonumbers2

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
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
	var sb strings.Builder
	sb.WriteRune('[')
	i := 0
	for cur != nil {
		if i != 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(strconv.Itoa(cur.Val))
		i++
		cur = cur.Next
	}
	sb.WriteRune(']')
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

func Test_addTwoNumbers(t *testing.T) {
	for _, tc := range []struct {
		l1   string
		l2   string
		want string
	}{
		{"[7,2,4,3]", "[5,6,4]", "[7,8,0,7]"},
		{"[2,4,3]", "[5,6,4]", "[8,0,7]"},
		{"[0]", "[0]", "[0]"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.l1), func(t *testing.T) {
			l1 := ParseList(tc.l1)
			l2 := ParseList(tc.l2)
			res := addTwoNumbers(l1, l2)
			require.Equal(t, tc.want, res.String())
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var emit func(*ListNode, chan<- int)
	emit = func(cur *ListNode, ch chan<- int) {
		if cur == nil {
			return
		}
		emit(cur.Next, ch)
		ch <- cur.Val
	}
	emitAndClose := func(cur *ListNode, ch chan<- int) {
		emit(cur, ch)
		close(ch)
	}
	l1Vals := make(chan int)
	l2Vals := make(chan int)
	go emitAndClose(l1, l1Vals)
	go emitAndClose(l2, l2Vals)
	var carry int
	var cur, next *ListNode
	for {
		v1, ok1 := <-l1Vals
		v2, ok2 := <-l2Vals
		if !ok1 && !ok2 {
			break
		}
		v := (v1 + v2 + carry)
		cur = &ListNode{
			Val:  v % 10,
			Next: next,
		}
		carry = v / 10
		next = cur
	}
	if carry > 0 {
		cur = &ListNode{
			Val:  carry,
			Next: next,
		}
		next = cur
	}
	return next
}
