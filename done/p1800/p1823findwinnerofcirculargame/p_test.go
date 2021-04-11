package p1823findwinnerofcirculargame

import (
	"container/list"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findTheWinner(t *testing.T) {
	for _, tc := range []struct {
		n    int
		k    int
		want int
	}{
		{8, 8, 4},
		{6, 5, 1},
		{5, 2, 3},
	} {
		t.Run(fmt.Sprintf("%+v", tc.n), func(t *testing.T) {
			require.Equal(t, tc.want, findTheWinner(tc.n, tc.k))
		})
	}
}

func findTheWinner(n int, k int) int {
	l := list.New()
	l.Init()
	for i := 1; i <= n; i++ {
		l.PushBack(i)
	}
	cur := l.Front()
	var next *list.Element
	for l.Len() > 1 {
		for i := 0; i < (k-1)%n; i++ {
			if cur.Next() == nil {
				cur = l.Front()
				continue
			}
			cur = cur.Next()
		}
		if cur.Next() == nil {
			next = l.Front()
		} else {
			next = cur.Next()
		}
		l.Remove(cur)
		cur = next
	}
	return l.Front().Value.(int)
}
