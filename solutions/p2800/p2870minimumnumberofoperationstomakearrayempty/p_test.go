package p2870minimumnumberofoperationstomakearrayempty

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minOperations(t *testing.T) {
	for i, tc := range []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 3, 2, 2, 4, 2, 3, 4}, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, minOperations(tc.nums))
		})
	}
}

func minOperations(nums []int) int {
	count := make(map[int]int)
	for _, x := range nums {
		count[x]++
	}
	var res int
	for x, cnt := range count {
		_ = x
		if cnt == 1 {
			return -1
		}
		res += (cnt / 3)
		cnt -= (cnt / 3) * 3
		if cnt%2 != 0 {
			res--
			cnt += 3
		}
		res += cnt / 2
		if cnt%2 != 0 {
			panic("wut")
		}
	}
	return res
}
