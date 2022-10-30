package p2457minimumadditiontomakeintegerbeautiful

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_makeIntegerBeautiful(t *testing.T) {
	for i, tc := range []struct {
		n      int64
		target int
		want   int64
	}{
		{467, 6, 33},
		{1, 1, 0},
		{16, 6, 4},
	} {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			require.Equal(t, tc.want, makeIntegerBeautiful(tc.n, tc.target))
		})
	}
}

func makeIntegerBeautiful(n int64, target int) int64 {
	var sum int
	nums := []int{}
	for _, ch := range fmt.Sprint(n) {
		nums = append(nums, int(ch-'0'))
	}
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
	calc := func() {
		sum = 0
		for i := range nums {
			if nums[i] >= 10 {
				if i == len(nums)-1 {
					nums = append(nums, 1)
				} else {
					nums[i+1]++
				}
				nums[i] = 0
			}
			sum += nums[i]
		}
	}
	fac := 1
	calc()
	var res int
	for i := 0; i < len(nums); i++ {
		if sum <= target {
			break
		}
		res += fac * (10 - nums[i])
		fac *= 10
		nums[i] = 10
		calc()
	}
	return int64(res)
}
