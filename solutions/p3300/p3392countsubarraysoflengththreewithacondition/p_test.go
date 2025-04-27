package p3392countsubarraysoflengththreewithacondition

func countSubarrays(nums []int) int {
	var res int
	for i := range len(nums) - 2 {
		a := nums[i]
		b := nums[i+1]
		c := nums[i+2]
		if (a+c)*2 == b {
			res++
		}
	}
	return res
}
