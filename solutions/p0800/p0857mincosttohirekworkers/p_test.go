package p0857mincosttohirekworkers

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mincostToHireWorkers(t *testing.T) {
	for _, tc := range []struct {
		quality []int
		wage    []int
		K       int
		want    float64
	}{
		{[]int{10, 20, 5}, []int{70, 50, 30}, 2, 105},
		{[]int{3, 1, 10, 10, 1}, []int{4, 8, 2, 2, 7}, 3, 30.66667},
	} {
		t.Run(fmt.Sprintf("%+v/%v/%v", tc.quality, tc.wage, tc.K), func(t *testing.T) {
			require.InEpsilon(t, tc.want, mincostToHireWorkers(tc.quality, tc.wage, tc.K), 0.1)
		})
	}
}

type worker struct {
	wage    int
	quality int
	ratio   float64
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	n := len(quality)

	workers := make([]worker, n)
	for i := range quality {
		workers[i] = worker{wage[i], quality[i], float64(wage[i]) / float64(quality[i])}
	}
	// It is very likely that the best worker is one which has low wage / quality ratio
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].ratio < workers[j].ratio
	})
	minCost := math.MaxFloat32
	groupQuality := 0
	h := IntHeap{}
	heap.Init(&h)
	for i := 0; i < n; i++ {
		heap.Push(&h, workers[i].quality)
		groupQuality += workers[i].quality
		if len(h) > K {
			groupQuality -= heap.Pop(&h).(int)
		}
		if len(h) == K {
			minCost = math.Min(minCost, float64(groupQuality)*workers[i].ratio)
		}
	}
	return minCost
}
