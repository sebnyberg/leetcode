package p2470numberofsubarrayswithlcmequaltok

func subarrayLCM(nums []int, k int) int {
	var res int
	for i := range nums {
		lcm := 1
		for j := i; j < len(nums); j++ {
			lcm = nums[j] * lcm / gcd(lcm, nums[j])
			if lcm == k {
				res++
			}
		}
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
