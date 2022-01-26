package p1924erectthefence2

import (
	"fmt"
	"log"
	"math"
	"math/rand"
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

func Test_outerTrees(t *testing.T) {
	for _, tc := range []struct {
		trees string
		want  []float64
	}{
		{"[[1,2],[2,2],[4,2]]", []float64{2.5, 2, 1.5}},
		{"[[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]", []float64{2, 2, 2}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.trees), func(t *testing.T) {
			trees := ParseMatrix(tc.trees)
			require.Equal(t, tc.want, outerTrees(trees))
		})
	}
}

func outerTrees(trees [][]int) []float64 {
	boundaryPts := make([]point, 0, 3)
	rand.Shuffle(len(trees), func(i, j int) {
		trees[i], trees[j] = trees[j], trees[i]
	})
	res := welzl(trees, 0, boundaryPts)
	return []float64{res.x, res.y, res.radius}
}

func welzl(trees [][]int, idx int, boundaryPts []point) circle {
	if idx == len(trees) || len(boundaryPts) == 3 {
		res := circleFromPoints(boundaryPts)
		return res
	}
	c := welzl(trees, idx+1, boundaryPts)
	p := point{float64(trees[idx][0]), float64(trees[idx][1])}
	if p.dist(c.center()) <= c.radius {
		return c
	}
	boundaryPts = append(boundaryPts, p)
	return welzl(trees, idx+1, boundaryPts)
}

type point struct {
	x, y float64
}

func (p point) dist(to point) float64 {
	dx := p.x - to.x
	dy := p.y - to.y
	s := math.Sqrt(dx*dx + dy*dy)
	return s
}

type circle struct {
	x, y, radius float64
}

func (c circle) center() point {
	return point{c.x, c.y}
}

func circleFrom3Points(a, b, c point) circle {
	bx := b.x - a.x
	by := b.y - a.y
	cx := c.x - a.x
	cy := c.y - a.y
	bb := bx*bx + by*by
	cc := cx*cx + cy*cy
	dd := bx*cy - by*cx
	center := point{
		x: (cy*bb-by*cc)/(2*dd) + a.x,
		y: (bx*cc-cx*bb)/(2*dd) + a.y,
	}
	return circle{
		x:      center.x,
		y:      center.y,
		radius: center.dist(a),
	}
}

func circleFromPoints(ps []point) circle {
	switch len(ps) {
	case 0:
		return circle{0, 0, 0}
	case 1:
		return circle{ps[0].x, ps[0].y, 0}
	case 2:
		return circle{
			x:      (ps[0].x + ps[1].x) / 2,
			y:      (ps[0].y + ps[1].y) / 2,
			radius: ps[0].dist(ps[1]) / 2,
		}
	default:
		return circleFrom3Points(ps[0], ps[1], ps[2])
	}
}
