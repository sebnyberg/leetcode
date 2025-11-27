package p3381maximumsubarraysumwithlengthdivisiblebyk

func maxSubarraySum(nums []int, k int) int64 {
	// For any subarray of nums that is divisible by k,
	// we can add any ajacent subarray of nums that is also divisible by k.
	// If we, for any index in nums, have calculated the maximum prior
	// subarray sum that is divisible by k, then we can add that to the
	// current subarray of length k to get the maximum subarray sum ending
	// in the current index.
	n := len(nums)
	maxArr := make([]int, n)
	for i := range maxArr {
		maxArr[i] = int(-1e15)
	}
	var sum int
	res := int(-1e15)
	for i := range nums {
		sum += nums[i]
		if i >= k {
			sum -= nums[i-k]
			maxArr[i] = max(maxArr[i], sum+maxArr[i-k])
		}
		if i >= k-1 {
			maxArr[i] = max(maxArr[i], sum)
		}
		res = max(res, maxArr[i])
	}
	return int64(res)
}
