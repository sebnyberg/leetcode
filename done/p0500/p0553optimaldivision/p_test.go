package p0553optimaldivision

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_optimalDivision(t *testing.T) {
	for _, tc := range []struct {
		nums []int
		want string
	}{
		{[]int{3, 2}, "3/2"},
		{[]int{1000, 100, 10, 2}, "1000/(100/10/2)"},
	} {
		t.Run(fmt.Sprintf("%+v", tc.nums), func(t *testing.T) {
			require.Equal(t, tc.want, optimalDivision(tc.nums))
		})
	}
}

func optimalDivision(nums []int) string {
	var res []byte
	res = append(res, fmt.Sprint(nums[0])...)
	if len(nums) == 1 {
		return string(res)
	}
	if len(nums) == 2 {
		res = append(res, "/"...)
		res = append(res, fmt.Sprint(nums[1])...)
		return string(res)
	}
	res = append(res, "/("...)
	for i := 1; i < len(nums); i++ {
		res = append(res, fmt.Sprint(nums[i])...)
		if i != len(nums)-1 {
			res = append(res, "/"...)
		}
	}
	res = append(res, ")"...)
	return string(res)
}
