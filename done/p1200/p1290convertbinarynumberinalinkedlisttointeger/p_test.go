package p1290convertbinarynumberinalinkedlisttointeger

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_getDecimalValue(t *testing.T) {
	for _, tc := range []struct {
		head string
		want int
	}{
		{"[1,0,1]", 5},
		{"[0]", 0},
		{"[1]", 1},
		{"[1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]", 18880},
		{"[0,0]", 0},
	} {
		t.Run(fmt.Sprintf("%+v", tc.head), func(t *testing.T) {
			node := ParseList(tc.head)
			require.Equal(t, tc.want, getDecimalValue(node))
		})
	}
}

func getDecimalValue(head *ListNode) int {
	var val int
	for head != nil {
		val <<= 1
		val += head.Val
		head = head.Next
	}
	return val
}
