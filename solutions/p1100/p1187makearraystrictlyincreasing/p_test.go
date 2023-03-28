package p1187makearraystrictlyincreasing

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeArrayIncreasing(t *testing.T) {
	for i, tc := range []struct {
		arr1 []int
		arr2 []int
		want int
	}{
		{[]int{1, 5, 3, 6, 7}, []int{1, 3, 2, 4}, 1},
		{[]int{1, 5, 3, 6, 7}, []int{4, 3, 1}, 2},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, makeArrayIncreasing(tc.arr1, tc.arr2))
		})
	}
}

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	// If the number of distinct values in arr1 and arr2 is >= len(arr1) then
	// there exists a solution.
	m := make(map[int]struct{})
	for _, x := range arr1 {
		m[x] = struct{}{}
	}
	for _, x := range arr2 {
		m[x] = struct{}{}
	}
	if len(m) < len(arr1) {
		return -1
	}

	// mem[{i,prev}] = minimum number of changes needed to sort values of
	// arr[i:] given that the previous values was prev.
	//
	// There is a stop-early criterion, which may be incorporated if this turns
	// out to be too slow.
	mem := make(map[[2]int]int)
	sort.Ints(arr2)
	res := dp(mem, arr1, arr2, 0, 0, -1)
	if res != math.MaxInt32 {
		return res
	}
	return -1
}

func dp(mem map[[2]int]int, arr1, arr2 []int, i, j, prev int) int {
	if i == len(arr1) {
		return 0
	}

	k := [2]int{i, prev}
	if v, exists := mem[k]; exists {
		return v
	}

	res := math.MaxInt32
	if arr1[i] > prev {
		// continue without doing anything
		res = min(res, dp(mem, arr1, arr2, i+1, j, arr1[i]))
	}
	// Try to insert a value into this position
	for j < len(arr2) && arr2[j] <= prev {
		j++
	}
	if j != len(arr2) {
		res = min(res, 1+dp(mem, arr1, arr2, i+1, j+1, arr2[j]))
	}

	mem[k] = res
	return mem[k]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
