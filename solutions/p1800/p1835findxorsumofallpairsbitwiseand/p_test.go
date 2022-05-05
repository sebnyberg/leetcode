package p1835findxorsumofallpairsbitwiseand

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getXORSum(t *testing.T) {
	for _, tc := range []struct {
		arr1 []int
		arr2 []int
		want int
	}{
		{[]int{1, 2, 3}, []int{6, 5}, 0},
		{[]int{12}, []int{4}, 4},
	} {
		t.Run(fmt.Sprintf("%+v", tc.arr1), func(t *testing.T) {
			require.Equal(t, tc.want, getXORSum(tc.arr1, tc.arr2))
		})
	}
}

func getXORSum(arr1 []int, arr2 []int) int {
	// Remove duplicates
	arr1Count := make(map[int]int)
	arr2Count := make(map[int]int)
	for _, n := range arr1 {
		if arr1Count[n] == 1 {
			delete(arr1Count, n)
		} else {
			arr1Count[n] = 1
		}
	}
	for _, n := range arr2 {
		if arr2Count[n] == 1 {
			delete(arr2Count, n)
		} else {
			arr2Count[n] = 1
		}
	}
	// enumerate maps
	res := 0
	for n1 := range arr1Count {
		for n2 := range arr2Count {
			res ^= n1 & n2
		}
	}
	return res
}
