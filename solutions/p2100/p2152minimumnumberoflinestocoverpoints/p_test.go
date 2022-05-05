package p2152minimumnumberoflinestocoverpoints

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func ParseMatrix(s string) [][]int {
	s = s[2 : len(s)-2]
	if s == "" {
		return nil
	}
	parts := strings.Split(s, "],[")
	res := make([][]int, len(parts))
	for i, part := range parts {
		if part == "" {
			continue
		}
		for _, numStr := range strings.Split(part, ",") {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatalf("failed to parse number, %v, %v\n", numStr, err)
			}
			res[i] = append(res[i], num)
		}
	}
	return res
}

func Test_minimumLines(t *testing.T) {
	for _, tc := range []struct {
		points string
		want   int
	}{
		{"[[-4,0],[0,4],[-1,-5],[2,3],[-4,1],[-4,3],[-1,5]]", 3},
		{"[[0,0],[2,-3],[-2,3]]", 1},
		{"[[0,0]]", 1},
		{"[[0,1],[2,3],[4,5],[4,3]]", 2},
		{"[[0,2],[-2,-2],[1,4]]", 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			points := ParseMatrix(tc.points)
			require.Equal(t, tc.want, minimumLines(points))
		})
	}
}

func minimumLines(points [][]int) int {
	var seen int
	mem := make(map[int]int)
	return dfs(mem, points, seen, 0, len(points))
}

func dfs(mem map[int]int, points [][]int, seen, nlines, n int) int {
	if seen == (1<<n)-1 {
		return nlines
	}
	if _, exists := mem[seen]; exists {
		return mem[seen]
	}
	res := math.MaxInt32
	// For each possible place to start a line from
	for i := 0; i < n; i++ {
		if seen&(1<<i) > 0 {
			continue
		}
		seen |= 1 << i
		if seen == (1<<n)-1 {
			res = min(res, nlines+1)
			seen &^= 1 << i
			continue
		}
		// For each pair (i, j) where j is not seen so far, draw a line.
		for j := 0; j < n; j++ {
			if seen&(1<<j) > 0 {
				continue
			}
			dy := points[j][1] - points[i][1]
			dx := points[j][0] - points[i][0]
			dydx := float64(dy) / float64(dx)
			seen |= 1 << j
			covered := 0
			// Then mark all points which fall on this line as seen
			for k := 0; k < n; k++ {
				if seen&(1<<k) > 0 || !coveredByLine(points[i], dydx, points[k]) {
					continue
				}
				covered |= 1 << k
			}
			res = min(res, dfs(mem, points, seen|covered, nlines+1, n))
			seen &^= 1 << j
		}
		seen &^= 1 << i
	}
	mem[seen] = res
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coveredByLine(start []int, dydx float64, pt []int) bool {
	dx := pt[0] - start[0]
	if dx == 0 {
		return math.IsInf(dydx, 0)
	}
	delta := pt[0] - start[0]
	return float64(start[1])+float64(delta)*dydx == float64(pt[1])
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
