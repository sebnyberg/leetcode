package p0497randompointinnonoverlappingrectangles

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	s := Constructor(
		[][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}},
	)
	var a []int
	for i := 0; i < 1000; i++ {
		a = s.Pick()
	}
	fmt.Println(a)
}

type Solution struct {
	rects [][]int
	areas []float64
	sum   float64
}

func Constructor(rects [][]int) Solution {
	n := len(rects)
	areas := make([]float64, n)
	var sum float64
	for i, r := range rects {
		dx := float64(r[2] - r[0] + 1)
		dy := float64(r[3] - r[1] + 1)
		sum += dx * dy
		areas[i] = sum
	}

	return Solution{
		rects: rects,
		areas: areas,
		sum:   sum,
	}
}

func (this *Solution) Pick() []int {
	// First, pick rectangle.
	randSum := rand.Float64() * this.sum
	idx := sort.SearchFloat64s(this.areas, randSum)
	// Then pick x/y
	r := this.rects[idx]
	dx := r[2] - r[0]
	dy := r[3] - r[1]
	x := rand.Intn(dx + 1)
	y := rand.Intn(dy + 1)
	return []int{r[0] + x, r[1] + y}
}
