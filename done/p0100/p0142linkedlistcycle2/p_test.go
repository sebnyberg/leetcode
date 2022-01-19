package p0142linkedlistcycle2

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

func Test_detectCycle(t *testing.T) {
	for _, tc := range []struct {
		ll      string
		idx     int
		wantVal int
	}{
		{"3,2,0,-4", 1, 2},
	} {
		t.Run(fmt.Sprintf("%+v", tc.ll), func(t *testing.T) {
			l := ParseList(tc.ll)
			end := l
			for end.Next != nil {
				end = end.Next
			}
			cycleStart := l
			for i := 0; i < tc.idx; i++ {
				cycleStart = cycleStart.Next
			}
			end.Next = cycleStart

			res := detectCycle(l)
			require.Equal(t, tc.wantVal, res.Val)
		})
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	next := func(n *ListNode) *ListNode {
		if n == nil {
			return nil
		}
		return n.Next
	}

	tortoise := next(head)
	hare := next(next(head))

	// Floyd's Tortoise and Hare
	// How this works is easier to prove with pen and paper.
	// Introduce a variable for a cycle's length and the distance to the start of
	// the cycle. Then provide that when the 'hare' and the 'tortoise' coincide,
	// then it's possible to find the start of the cycle by resetting the position
	// of the tortoise.
	// See https://en.wikipedia.org/wiki/Cycle_detection#Floyd's_tortoise_and_hare
	for tortoise != hare {
		hare = next(next(hare))
		tortoise = next(tortoise)
		if hare == nil {
			return nil
		}
	}

	tortoise = head
	for tortoise != hare {
		tortoise = next(tortoise)
		hare = next(hare)
	}
	return tortoise
}
