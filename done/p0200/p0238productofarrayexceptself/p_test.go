package p0238productofarrayexceptself

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_productExceptSelf(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want []int
	}{
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, productExceptSelf(tc.nums))
		})
	}
}

func productExceptSelf(nums []int) []int {
	var zeroCount int
	prod := 1
	for _, n := range nums {
		if n == 0 {
			zeroCount++
			continue
		}
		prod *= n
	}

	res := make([]int, len(nums))
	if zeroCount >= 2 {
		return res
	}
	if zeroCount == 0 {
		copy(res, nums)
	}
	for i, n := range nums {
		if zeroCount == 1 {
			if n == 0 {
				res[i] = prod
				return res
			}
			continue
		}
		res[i] = prod / n
	}
	return res
}
