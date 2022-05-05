package p2007findoriginalarrayfromdoubledarray

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findOriginalArray(t *testing.T) {
	for _, tc := range []struct {
		changed []int
		want    []int
	}{
		{[]int{2, 1, 2, 4, 2, 4}, []int{1, 2, 2}},
		{[]int{0}, []int{}},
		{[]int{0, 0}, []int{0}},
		{[]int{1, 3, 4, 2, 6, 8}, []int{1, 3, 4}},
		{[]int{6, 3, 0, 1}, []int{}},
		{[]int{1}, []int{}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.changed), func(t *testing.T) {
			res := findOriginalArray(tc.changed)
			require.EqualValues(t, tc.want, res)
		})
	}
}

func findOriginalArray(changed []int) []int {
	numCounts := make(map[int]int)
	for _, num := range changed {
		numCounts[num]++
	}
	var res []int
	sortedNums := make([]int, 0)
	for num := range numCounts {
		sortedNums = append(sortedNums, num)
	}
	sort.Ints(sortedNums)
	for _, num := range sortedNums {
		count := numCounts[num]
		if num == 0 {
			if count%2 != 0 {
				return []int{}
			}
			for i := 0; i < count/2; i++ {
				res = append(res, 0)
			}
			continue
		}
		if count <= 0 {
			continue
		}
		if numCounts[num*2] < count {
			return []int{}
		}
		numCounts[num*2] -= count
		for i := 0; i < count; i++ {
			res = append(res, num)
		}
	}
	return res
}
