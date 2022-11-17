package p0995minimumnumberofkconsecutivebitflips

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minKBitFlips(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{0, 1, 0}, 1, 2},
		{[]int{1, 1, 0}, 2, -1},
		{[]int{0, 0, 0, 1, 0, 1, 1, 0}, 3, 3},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minKBitFlips(tc.nums, tc.k))
		})
	}
}

func minKBitFlips(nums []int, k int) int {
	// My intuition is that we can greedily flip so that the leftmost digit
	// is ensured to be a 1. Once the stamp won't fit, ensure that the remainder
	// has only ones.
	//
	n := len(nums)
	want := 1
	var res int
	unflip := []int{}
	for i := range nums {
		if n-i < k {
			break
		}
		if len(unflip) > 0 && unflip[0] == i {
			want = 1 - want
			unflip = unflip[1:]
		}
		if nums[i] == want {
			continue
		}
		// "flip" by changing the wanted bit
		want = 1 - want
		unflip = append(unflip, i+k)
		res++
	}
	for i := n - k + 1; i < n; i++ {
		if len(unflip) > 0 && unflip[0] == i {
			want = 1 - want
			unflip = unflip[1:]
		}
		if nums[i] != want {
			return -1
		}
	}
	return res
}
