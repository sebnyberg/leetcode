package p0861scoreafterflippingmatrix

import (
	"fmt"
	"testing"

	"github.com/sebnyberg/leetcode"
	"github.com/stretchr/testify/require"
)

func Test_matrixScore(t *testing.T) {
	for _, tc := range []struct {
		grid [][]int
		want int
	}{
		{leetcode.ParseMatrix("[[0,0,1,1],[1,0,1,0],[1,1,0,0]]"), 39},
		{leetcode.ParseMatrix("[[0]]"), 1},
	} {
		t.Run(fmt.Sprintf("%+v", tc.grid), func(t *testing.T) {
			require.Equal(t, tc.want, matrixScore(tc.grid))
		})
	}
}

func matrixScore(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	mask := (1 << n) - 1
	nums1 := make([]int, m)
	for i := range grid {
		for _, v := range grid[i] {
			nums1[i] <<= 1
			nums1[i] += v
		}
	}
	nums2 := make([]int, len(nums1))
	copy(nums2, nums1)
	leftmostBit := (1 << (n - 1))
	for i := range nums1 {
		if nums1[i]&leftmostBit == 0 {
			nums1[i] = mask & ^nums1[i]
		} else {
			nums2[i] = mask & ^nums2[i]
		}
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var calcResult func(nums []int, bit int) int
	calcResult = func(nums []int, bit int) int {
		if bit == 0 {
			return 0
		}
		var res [2]int
		for _, x := range nums {
			res[0] += x & bit
			res[1] += ^x & bit
		}
		return max(res[0], res[1]) + calcResult(nums, bit>>1)
	}

	res1 := calcResult(nums1, leftmostBit)
	res2 := calcResult(nums2, leftmostBit)
	return max(res1, res2)
}
