package p0593validsquare

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_validSquare(t *testing.T) {
	for i, tc := range []struct {
		p1, p2, p3, p4 []int
		want           bool
	}{
		{[]int{0, 0}, []int{1, 1}, []int{0, 0}, []int{1, 1}, false},
		{[]int{1, 1}, []int{5, 3}, []int{3, 5}, []int{7, 7}, false},
		{[]int{0, 1}, []int{1, 0}, []int{-1, 0}, []int{0, -1}, true},
		{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}, true},
		{[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}, true},
		{[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 12}, false},
	} {
		t.Run(fmt.Sprintf("%+v", i), func(t *testing.T) {
			require.Equal(t, tc.want, validSquare(tc.p1, tc.p2, tc.p3, tc.p4))
		})
	}
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	p := [][]int{p1, p2, p3, p4}
	return dfs(p, make([]int, 4), 0, 0)
}

func dfs(p [][]int, ord []int, i, bm int) bool {
	if bm == (1<<4)-1 {
		res := check(p, ord)
		return res
	}
	for j := 0; j < 4; j++ {
		if bm&(1<<j) > 0 {
			continue
		}
		ord[i] = j
		if dfs(p, ord, i+1, bm|(1<<j)) {
			return true
		}
	}
	return false
}

func check(p [][]int, ord []int) bool {
	p1 := p[ord[0]]
	p2 := p[ord[1]]
	p3 := p[ord[2]]
	p4 := p[ord[3]]
	return dist(p1, p2) > 0 && dist(p1, p3) > 0 &&
		dist(p1, p2) == dist(p2, p3) &&
		dist(p2, p3) == dist(p3, p4) &&
		dist(p3, p4) == dist(p4, p1) &&
		dist(p1, p3) == dist(p2, p4)
}

func dist(p1, p2 []int) int {
	dx := p1[0] - p2[0]
	dy := p1[1] - p2[1]
	return dx*dx + dy*dy
}
