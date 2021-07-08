package p0382linkedlistrandomnode

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

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

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestSolution(t *testing.T) {
	head := ParseList("[1,2,3]")
	s := Constructor(head)
	count := make([]int, 4)
	for i := 0; i < 1e4; i++ {
		count[s.GetRandom()]++
	}
	for i := 1; i <= 3; i++ {
		require.InEpsilon(t, 1.0/3.0, float64(count[i])/1e4, 0.04)
	}
}

type Solution struct {
	rnd  *rand.Rand
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
		rnd:  rand.New(rand.NewSource(time.Now().Unix())),
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	cur := this.head
	var picked int
	i := 1
	for cur != nil {
		if rand.Intn(i) == 0 {
			picked = cur.Val
		}
		cur = cur.Next
		i++
	}
	return picked
}
