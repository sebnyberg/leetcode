package p0813largestsumofaverages

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_largestSumOfAverages(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		k    int
		want float64
	}{
		{[]int{4, 1, 7, 5, 6, 2, 3}, 4, 18.16667},
		{[]int{9, 1, 2, 3, 9}, 3, 20},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, 20.5},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.InEpsilon(t, tc.want, largestSumOfAverages(tc.nums, tc.k), 1e-5)
		})
	}
}

func largestSumOfAverages(nums []int, k int) float64 {
	var mem [101][101]float64
	for i := range mem {
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	res := dp(&mem, nums, 0, len(nums), k)
	return res
}

// dp returns the maximum possible average sum given that you're at
// position i in nums, and there are k partitions left
func dp(mem *[101][101]float64, nums []int, i, n, k int) float64 {
	if k > n-i {
		return -1
	}
	if k == 0 {
		if i < n {
			return -1
		}
		return 0
	}
	if mem[i][k] != -1 {
		return mem[i][k]
	}
	var sum int
	var res float64
	for j := i; j < len(nums)-k+1; j++ {
		sum += nums[j]
		avg := float64(sum) / float64(j-i+1)
		miniRes := dp(mem, nums, j+1, n, k-1)
		if miniRes < 0 {
			continue
		}
		res = math.Max(res, avg+miniRes)
	}
	mem[i][k] = res
	return res
}
