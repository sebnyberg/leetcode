package p1956minimumtimeforkvirusvariantstospread

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDayskVariants(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		k      int
		want   int
	}{
		{[][]int{{1, 1}, {6, 1}}, 2, 3},
		{[][]int{{3, 3}, {1, 2}, {9, 2}}, 2, 2},
		{[][]int{{3, 3}, {1, 2}, {9, 2}}, 3, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.points), func(t *testing.T) {
			require.Equal(t, tc.want, minDayskVariants(tc.points, tc.k))
		})
	}
}

func minDayskVariants(points [][]int, k int) int {
	check := func(day int) int {
		var grid [101][101]int
		var maxVal int
		for _, point := range points {
			x, y := int(point[0]), int(point[1])
			for i := -day; i <= day; i++ {
				if y+i < 0 {
					continue
				} else if y+i > 100 {
					break
				}
				for j := -day + abs(i); j <= day-abs(i); j++ {
					if x+j < 0 {
						continue
					} else if x+j > 100 {
						break
					}
					grid[y+i][x+j]++
					maxVal = max(maxVal, grid[y+i][x+j])
				}
			}
		}
		return maxVal
	}
	res := sort.Search(100, func(days int) bool {
		return check(int(days)) >= int(k)
	})
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
