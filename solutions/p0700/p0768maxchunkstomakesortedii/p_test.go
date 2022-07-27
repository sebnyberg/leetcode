package p0768maxchunkstomakesortedii

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxChunksToSorted(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want int
	}{
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{2, 1, 3, 4, 4}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, maxChunksToSorted(tc.arr))
		})
	}
}

func maxChunksToSorted(arr []int) int {
	n := len(arr)
	sortedArr := make([]int, n)
	copy(sortedArr, arr)
	sort.Ints(sortedArr)
	valToIdx := make(map[int]int, n)
	for i, v := range sortedArr {
		if _, exists := valToIdx[v]; !exists {
			valToIdx[v] = i
		}
	}
	adjusted := make([]int, n)
	for i, v := range arr {
		adjusted[i] = valToIdx[v]
		valToIdx[v]++
	}

	seen := make([]bool, n)
	var j int
	var res int
	for i, v := range adjusted {
		seen[v] = true
		for j < n && seen[j] {
			j++
		}
		if j == i+1 {
			res++
		}
	}
	return res
}
