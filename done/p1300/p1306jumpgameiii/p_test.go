package p1306jumpgameiii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canReach(t *testing.T) {
	for _, tc := range []struct {
		arr   []int
		start int
		want  bool
	}{
		{[]int{3, 0, 2, 1, 2}, 2, false},
		{[]int{4, 2, 3, 0, 3, 1, 2}, 3, true},
		{[]int{4, 2, 3, 0, 3, 1, 2}, 0, true},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, canReach(tc.arr, tc.start))
		})
	}
}

func canReach(arr []int, start int) bool {
	n := len(arr)
	visited := make([]bool, n)
	toVisit := []int{start}
	visited[start] = true
	next := []int{}
	// BFS
	for len(toVisit) > 0 {
		next = next[:0]
		for _, idx := range toVisit {
			if arr[idx] == 0 {
				return true
			}
			width := arr[idx]
			l, r := idx-width, idx+width
			if l >= 0 && !visited[l] {
				visited[l] = true
				next = append(next, l)
			}
			if r < n && !visited[r] {
				visited[r] = true
				next = append(next, r)
			}
		}
		toVisit, next = next, toVisit
	}
	return false
}
