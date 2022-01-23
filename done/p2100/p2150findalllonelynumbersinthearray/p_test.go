package p2150findalllonelynumbersinthearray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLonely(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{10, 6, 5, 8}, []int{10, 8}},
		{[]int{1, 3, 5, 3}, []int{1, 5}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findLonely(tc.nums))
		})
	}
}

func findLonely(nums []int) []int {
	seen := make(map[int]int)
	for _, x := range nums {
		seen[x]++
	}
	var res []int
	for num, count := range seen {
		if count > 1 {
			continue
		}
		if seen[num-1] == 0 && seen[num+1] == 0 {
			res = append(res, num)
		}
	}
	return res
}
