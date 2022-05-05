package p2013detectsquares

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDetectSquares(t *testing.T) {
	ds := Constructor()
	ds.Add([]int{3, 10})
	ds.Add([]int{11, 2})
	ds.Add([]int{3, 2})
	res := ds.Count([]int{11, 10})
	require.Equal(t, 1, res)
	res = ds.Count([]int{14, 8})
	require.Equal(t, 0, res)
	ds.Add([]int{11, 2})
	res = ds.Count([]int{11, 10})
	require.Equal(t, 2, res)
}

func Test2(t *testing.T) {
	ds := Constructor()
	for idx, tc := range []struct {
		op      string
		point   []int
		wantVal int
	}{
		{"add", []int{5, 10}, 0},
		{"add", []int{10, 5}, 0},
		{"add", []int{10, 10}, 0},
		{"count", []int{5, 5}, 1},
		{"add", []int{3, 0}, 0},
		{"add", []int{8, 0}, 0},
		{"add", []int{8, 5}, 0},
		{"count", []int{3, 5}, 1},
		{"add", []int{9, 0}, 0},
		{"add", []int{9, 8}, 0},
		{"add", []int{1, 8}, 0},
		{"count", []int{1, 0}, 1},
		{"add", []int{0, 0}, 0},
		{"add", []int{8, 0}, 0},
		{"add", []int{8, 8}, 0},
		{"count", []int{0, 8}, 2},
	} {
		t.Run(fmt.Sprintf("%v, %v, %v", idx, tc.op, tc.wantVal), func(t *testing.T) {
			switch tc.op {
			case "add":
				ds.Add(tc.point)
			case "count":
				res := ds.Count(tc.point)
				require.Equal(t, tc.wantVal, res)
			}
		})
	}
}

type Point struct {
	x, y int16
}

type DetectSquares struct {
	points      []Point
	pointsCount map[Point]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		points:      make([]Point, 0, 1000),
		pointsCount: make(map[Point]int, 100),
	}
}

func (this *DetectSquares) Add(point []int) {
	p := Point{int16(point[0]), int16(point[1])}
	this.points = append(this.points, p)
	this.pointsCount[p]++
}

func (this *DetectSquares) Count(point []int) int {
	p := Point{int16(point[0]), int16(point[1])}
	var res int
	// Find points that could be on the diagonal from this point
	for _, q := range this.points {
		if q.x == p.x || q.y == p.y || abs(q.y-p.y) != abs(q.x-p.x) {
			continue
		}
		res += this.pointsCount[Point{q.x, p.y}] * this.pointsCount[Point{p.x, q.y}]
	}
	return res
}

func abs(a int16) int16 {
	if a < 0 {
		return -a
	}
	return a
}
