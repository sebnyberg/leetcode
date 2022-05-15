package p0499themazeiii

import (
	"container/heap"
	"fmt"
	"leetcode"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findShortestWay(t *testing.T) {
	for _, tc := range []struct {
		maze [][]int
		ball []int
		hole []int
		want string
	}{
		{leetcode.ParseMatrix("[[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]]"), []int{4, 3}, []int{0, 1}, "lul"},
		{leetcode.ParseMatrix("[[0,0,0,0,0],[1,1,0,0,1],[0,0,0,0,0],[0,1,0,0,1],[0,1,0,0,0]]"), []int{4, 3}, []int{3, 0}, "impossible"},
		{leetcode.ParseMatrix("[[0,0,0,0,0,0,0],[0,0,1,0,0,1,0],[0,0,0,0,1,0,0],[0,0,0,0,0,0,1]]"), []int{0, 4}, []int{3, 5}, "dldr"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.maze), func(t *testing.T) {
			require.Equal(t, tc.want, findShortestWay(tc.maze, tc.ball, tc.hole))
		})
	}
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	// Idea is to perform a kind of BFS with memoization and stopping early

	m := len(maze)
	n := len(maze[0])
	dirs := [][2]int{{1, 0}, {0, -1}, {0, 1}, {-1, 0}}
	dirNames := []string{"d", "l", "r", "u"}
	ok := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n && maze[i][j] != 1
	}
	goal := func(i, j int) bool {
		return i == hole[0] && j == hole[1]
	}

	var minDist [101][101]int
	for i := range minDist {
		for j := range minDist[i] {
			minDist[i][j] = math.MaxInt32
		}
	}
	h := MinHeap{{"", 0, ball[0], ball[1]}}
	for len(h) > 0 {
		x := heap.Pop(&h).(state)
		if goal(x.i, x.j) {
			return x.moves
		}
		for dir, delta := range dirs {
			// Move in direction until hitting the hole, or the wall
			s := x
			for {
				s.i += delta[0]
				s.j += delta[1]
				if !ok(s.i, s.j) {
					s.i -= delta[0]
					s.j -= delta[1]
					break
				}
				s.dist++
				if goal(s.i, s.j) {
					break
				}
			}
			if s == x {
				continue
			}
			s.moves += dirNames[dir]
			if minDist[s.i][s.j] < s.dist {
				continue
			}
			minDist[s.i][s.j] = s.dist
			heap.Push(&h, s)
		}
	}
	return "impossible"
}

type state struct {
	moves string
	dist  int
	i, j  int
}

type MinHeap []state

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	// h[i].idx = i
	// h[j].idx = j
}
func (h MinHeap) Less(i, j int) bool {
	if h[i].dist == h[j].dist {
		return h[i].moves < h[j].moves
	}
	return h[i].dist < h[j].dist
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(state))
}
func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	it := (*h)[n-1]
	*h = (*h)[:n-1]
	return it
}
