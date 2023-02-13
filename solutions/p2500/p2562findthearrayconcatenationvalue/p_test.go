package p2562findthearrayconcatenationvalue

import (
	"fmt"
	"strconv"
)

func findTheArrayConcVal(nums []int) int64 {
	var res int
	for l, r := 0, len(nums)-1; l < r; l, r = l+1, r-1 {
		a := fmt.Sprint(nums[l]) + fmt.Sprint(nums[r])
		x, _ := strconv.Atoi(a)
		res += x
	}
	if len(nums)&1 == 1 {
		res += nums[len(nums)/2]
	}
	return int64(res)
}
