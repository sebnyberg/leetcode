package p3423maximumdifferencebetweenadjacentelementsinacirculararray

func maxAdjacentDistance(nums []int) int {
	var res int
	for i := range nums {
		res = max(res, abs(nums[i]-nums[(i+1)%len(nums)]))
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
