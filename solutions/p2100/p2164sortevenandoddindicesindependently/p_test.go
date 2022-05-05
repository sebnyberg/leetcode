package p2164sortevenandoddindicesindependently

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortEvenOdd(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{4, 1, 2, 3}, []int{2, 3, 4, 1}},
		{[]int{2, 1}, []int{2, 1}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, sortEvenOdd(tc.nums))
		})
	}
}

func sortEvenOdd(nums []int) []int {
	n := len(nums)
	odd := make([]int, 0, n/2)
	even := make([]int, 0, n/2)
	for i, n := range nums {
		if i%2 == 0 {
			even = append(even, n)
		} else {
			odd = append(odd, n)
		}
	}
	sort.Ints(even)
	sort.Slice(odd, func(i, j int) bool {
		return odd[i] > odd[j]
	})
	for i := range nums {
		if i%2 == 0 {
			nums[i] = even[i/2]
		} else {
			nums[i] = odd[i/2]
		}
	}
	return nums
}
