package p3024typeoftriangle

import "sort"

func triangleType(nums []int) string {
	sort.Ints(nums)
	a := nums[0]
	b := nums[1]
	c := nums[2]
	switch {
	case a+b <= c:
		return "none"
	case a == b && b == c:
		return "equilateral"
	case a == b || b == c:
		return "isosceles"
	default:
		return "scalene"
	}
}
