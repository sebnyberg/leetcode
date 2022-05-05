package p2121intervalsbetweenidenticalelements

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getDistances(t *testing.T) {
	for _, tc := range []struct {
		arr  []int
		want []int64
	}{
		{[]int{2, 1, 3, 1, 2, 3, 3}, []int64{4, 2, 7, 2, 4, 4, 5}},
		{[]int{10, 5, 10, 10}, []int64{5, 0, 3, 4}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr), func(t *testing.T) {
			require.Equal(t, tc.want, getDistances(tc.arr))
		})
	}
}

func getDistances(arr []int) []int64 {
	// It does not matter exactly where the indices are, only how many elements
	// came prior to the current one and the indext of the last seen element
	type idxAndCount struct {
		idx   int
		count int
		val   int
	}
	n := len(arr)
	left := make(map[int]idxAndCount)
	res := make([]int, n)
	// First go from left to right
	for i, num := range arr {
		if _, exists := left[num]; !exists {
			left[num] = idxAndCount{i, 1, 0}
			continue
		}
		prev := left[num]
		prev.val += prev.count * (i - prev.idx)
		prev.count++
		prev.idx = i
		res[i] = prev.val
		left[num] = prev
	}

	right := make(map[int]idxAndCount)
	for i := len(arr) - 1; i >= 0; i-- {
		num := arr[i]
		if _, exists := right[num]; !exists {
			right[num] = idxAndCount{i, 1, 0}
			continue
		}
		prev := right[num]
		prev.val += prev.count * (prev.idx - i)
		prev.count++
		prev.idx = i
		res[i] += prev.val
		right[num] = prev
	}

	res64 := make([]int64, len(res))
	for i := range res {
		res64[i] = int64(res[i])
	}
	return res64
}
