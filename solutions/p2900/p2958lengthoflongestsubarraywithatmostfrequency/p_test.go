package p2958lengthoflongestsubarraywithatmostfrequency

func maxSubarrayLength(nums []int, k int) int {
	// This is a greedy problem. Keep a left pointer and iterate over each
	// element. If the count of frequencies is >k for all seen elements, then
	// remove the element pointed to by the left pointer until the subarray is
	// valid again. Then conditionally update the maximum subarray length.

	var l int
	count := make(map[int]int)
	var res int
	for r, x := range nums {
		count[x]++
		for count[x] > k {
			count[nums[l]]--
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
