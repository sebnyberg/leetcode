package p0857mincosttohirekworkers

import (
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

func mincostToHireWorkers(quality []int, wage []int, K int) float64 {
	n := len(quality)
	workers := make([]worker, n)
	for i := range quality {
		workers[i] = worker{wage[i], quality[i], float64(quality[i]) / float64(wage[i])}
	}
	sort.Slice(workers, func(i, j int) bool {
		return workers[i].ratio < workers[j].ratio
	})
	minCost := math.MaxFloat32
	for i := 0; i < n-K+1; i++ {
		if diff := calculateCost(workers[i : i+K]); diff < minCost {
			minCost = diff
		}
	}
	return minCost
}

func calculateCost(workers []worker) float64 {
	minRatio := math.MaxFloat64
	var minRatioIdx int
	for i, w := range workers {
		if w.ratio < minRatio {
			minRatio = w.ratio
			minRatioIdx = i
		}
	}
	// minRatioIdx gets paid based on their minimum wage
	cost := float64(workers[minRatioIdx].wage)
	minWage := float64(workers[minRatioIdx].wage)
	minQuality := float64(workers[minRatioIdx].quality)
	for i, w := range workers {
		if i == minRatioIdx {
			continue
		}
		cost += (float64(w.quality) / minQuality) * minWage
	}
	return cost
}
