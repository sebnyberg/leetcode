package p2541minimumoperationstomakearrayequalii

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for i, tc := range []struct {
		nums1 []int
		nums2 []int
		k     int
		want  int64
	}{
		{[]int{10, 5, 15, 20}, []int{20, 10, 15, 5}, 0, 0},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums1, tc.nums2, tc.k))
		})
	}
}

func minOperations(nums1 []int, nums2 []int, k int) int64 {
	var extraAdd int
	var extraSub int
	var res int64
	for i := range nums1 {
		if nums1[i] == nums2[i] {
			continue
		}
		if k == 0 {
			return -1
		}
		d := nums2[i] - nums1[i]
		if abs(d)%k != 0 {
			return -1
		}
		nk := abs(d) / k
		if d < 0 {
			if extraSub > 0 {
				a := min(extraSub, nk)
				extraSub -= a
				nk -= a
			}
			res += int64(nk)
			extraAdd += nk
		} else {
			if extraAdd > 0 {
				a := min(extraAdd, nk)
				extraAdd -= a
				nk -= a
			}
			res += int64(nk)
			extraSub += nk
		}
	}
	if extraSub != 0 || extraAdd != 0 {
		return -1
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
