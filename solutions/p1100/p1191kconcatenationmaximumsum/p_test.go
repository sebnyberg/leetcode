package p1191kconcatenationmaximumsum

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_kConcatenationMaxSum(t *testing.T) {
	for i, tc := range []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{1, -9, 2, 6, 4}, 3, 20},
		{[]int{1, 2}, 3, 9},
		{[]int{1, -2, 1}, 5, 2},
		{[]int{-1, -2}, 7, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, kConcatenationMaxSum(tc.arr, tc.k))
		})
	}
}

func kConcatenationMaxSum(arr []int, k int) int {
	// Intuition:
	//
	// 1. A valid subarray sum(arr[l:i+1]) has a total sum > 0
	// 2. A valid subarray that is longer than len(arr) can increase in total
	// sum through repetition iff its last sum(arr[i-n:i]) > 0 and k > 2
	//
	n := len(arr)
	if k > 1 {
		arr = append(arr, arr...)
	}
	presum := make([]int, n*2+1)
	for i := range arr {
		presum[i+1] = presum[i] + arr[i]
	}
	var l int
	var sum int
	var res int
	var singleRes int
	const mod = 1e9 + 7
	for i := range arr {
		sum += arr[i]
		if sum <= 0 {
			l = i + 1
			sum = 0
			continue
		}
		if singleRes < sum {
			singleRes = sum
			if i-l >= n-1 && presum[i+1]-presum[i-n+1] > 0 {
				ss := presum[i+1] - presum[i-n+1]
				tt := presum[i-n+1] - presum[l]
				kk := k - 1
				if i == n-1 {
					kk = k
				}
				res = ((ss % mod) * (kk) % mod) + tt
				res %= mod
			} else {
				res = singleRes
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
