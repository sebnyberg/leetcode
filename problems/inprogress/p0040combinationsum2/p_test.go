package p0040combinationsum2

// import (
// 	"fmt"
// 	"sort"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func Test_combinationSum2(t *testing.T) {
// 	for _, tc := range []struct {
// 		candidates []int
// 		target     int
// 		want       [][]int
// 	}{
// 		{[]int{1, 1, 2, 5, 6, 7, 10}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
// 		{[]int{2, 5, 2, 1, 2}, 8, [][]int{{1, 2, 2}, {5}}},
// 	} {
// 		t.Run(fmt.Sprintf("%+v/%v", tc.candidates, tc.target), func(t *testing.T) {
// 			require.Equal(t, tc.want, combinationSum2(tc.candidates, tc.target))
// 		})
// 	}
// }

// func combinationSum2(candidates []int, target int) [][]int {
// 	sort.Ints(candidates)
// 	nums := make([]int, 0, len(candidates))
// 	res := make([][]int, 0)
// 	findCombinationSums(candidates, target, nums, &res)
// 	return res
// }

// func findCombinationSums(candidates []int, target int, nums []int, res *[][]int) {
// 	switch {
// 	case target < 0:
// 		return
// 	case target == 0:
// 		a := make([]int, len(nums))
// 		copy(a, nums)
// 		*res = append(*res, a)
// 		return
// 	}

// 	for i := 0; i < len(candidates); i++ {
// 		n := candidates[i]
// 		var j int
// 		for j = i + 1; candidates[j] == n; j++ {
// 		}
// 		nums = append(nums, n)
// 		findCombinationSums(candidates[i+1:], target-n, nums, res)
// 		nums = nums[:len(nums)-1]
// 	}
// }
