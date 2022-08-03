package p0491increasingsubsequences

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"

	"github.com/stretchr/testify/require"
)

func Test_findSubsequences(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want [][]int
	}{
		{[]int{4, 6, 7, 7}, leetcode.ParseMatrix("[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]")},
		{[]int{4, 4, 3, 2, 1}, leetcode.ParseMatrix("[[4,4]]")},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.ElementsMatch(t, tc.want, findSubsequences(tc.nums))
		})
	}
}

func findSubsequences(nums []int) [][]int {
	n := len(nums)
	var i int
	var empty [15]int8
	for i := range empty {
		empty[i] = -128
	}
	var res [][15]int8
	res = append(res, empty)
	seen := make(map[[15]int8]struct{})
	for x := 3; x < 1<<n; x++ {
		minVal := -128
		var j int
		for b := 0; b < n; b++ {
			if x&(1<<b) > 0 {
				if nums[b] < minVal {
					res[i] = empty
					goto continueLoop
				}
				res[i][j] = int8(nums[b])
				minVal = nums[b]
				j++
			}
		}
		if _, exists := seen[res[i]]; exists || j < 2 {
			res[i] = empty
		} else {
			seen[res[i]] = struct{}{}
			res = append(res, empty)
			i++
		}
	continueLoop:
	}
	ret := make([][]int, len(res)-1)
	for i := range res {
		for j := range res[i] {
			if res[i][j] == -128 {
				break
			}
			ret[i] = append(ret[i], int(res[i][j]))
		}
	}
	return ret
}
