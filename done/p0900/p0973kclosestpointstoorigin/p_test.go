package p0973kclosestpointstoorigin

import (
	"container/heap"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kClosest(t *testing.T) {
	for _, tc := range []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
	} {
		t.Run(fmt.Sprintf("%+v/%v", tc.points, tc.k), func(t *testing.T) {
			require.Equal(t, tc.want, kClosest(tc.points, tc.k))
		})
	}
}

type pointDistance struct {
	p    point
	dist float64
	idx  int
}

type point struct {
	x, y int
}

func (p point) origoDistance() float64 {
	return math.Sqrt(math.Pow(float64(p.x), 2) + math.Pow(float64(p.y), 2))
}

type MinDistHeap []*pointDistance

func (h *MinDistHeap) Len() int { return len(*h) }
func (h *MinDistHeap) Less(i int, j int) bool {
	return (*h)[i].dist < (*h)[j].dist
}

func (h MinDistHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].idx = i
	h[j].idx = j
}

func (h *MinDistHeap) Push(x interface{}) {
	(*h) = append((*h), x.(*pointDistance))
}

func (h *MinDistHeap) Pop() interface{} {
	n := len(*h)
	el := (*h)[n-1]
	(*h)[n-1] = nil
	(*h) = (*h)[:n-1]
	return el
}

func kClosest(points [][]int, k int) [][]int {
	distHeap := make(MinDistHeap, len(points))
	for i, p := range points {
		point := point{p[0], p[1]}
		distHeap[i] = &pointDistance{
			point,
			point.origoDistance(),
			0,
		}
	}
	heap.Init(&distHeap)
	res := make([][]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(&distHeap).(*pointDistance)
		res[i] = []int{p.p.x, p.p.y}
	}
	return res
}
